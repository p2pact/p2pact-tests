package bench

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"sync"
	"sync/atomic"
	"time"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// --- HTTP/2: the same API as HTTPS+JSON, multiplexed on one connection ---

// RunHTTP2 measures POST /chats/<id>/messages over HTTP/2 and TLS 1.3.
func RunHTTP2(sc *Scenario, p Profile) (*Result, error) {
	res := newResult("HTTP/2", p.Name)
	srv := &http.Server{Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.Copy(io.Discard, r.Body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}), TLSConfig: &tls.Config{Certificates: []tls.Certificate{selfSigned()}}}
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, err
	}
	go srv.ServeTLS(ln, "", "")
	defer srv.Close()
	addr, stop, err := linkProxy(ln.Addr().String(), p)
	if err != nil {
		return nil, err
	}
	defer stop()
	var in, out atomic.Int64
	client := func() *http.Client {
		return &http.Client{Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, a string) (net.Conn, error) {
				c, err := (&net.Dialer{}).DialContext(ctx, network, a)
				if err != nil {
					return nil, err
				}
				return counting{c, &in, &out}, nil
			},
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			ForceAttemptHTTP2: true,
		}}
	}
	send := func(c *http.Client, text string) error {
		body, _ := json.Marshal(map[string]string{"text": text})
		req, _ := http.NewRequest("POST", "https://"+addr+"/chats/"+chatID+"/messages", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+bearer)
		resp, err := c.Do(req)
		if err != nil {
			return err
		}
		_, _ = io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		if resp.ProtoMajor != 2 {
			return fmt.Errorf("negotiated %s, not HTTP/2", resp.Proto)
		}
		return nil
	}
	return measureTCP(sc, p, res, &in, &out, func() (func(string) error, func(), error) {
		c := client()
		return func(t string) error { return send(c, t) }, c.CloseIdleConnections, nil
	})
}

// --- gRPC: a unary call with a protobuf message {chat, text} over HTTP/2 and TLS 1.3 ---

type rawCodec struct{}

func (rawCodec) Marshal(v any) ([]byte, error) {
	switch x := v.(type) {
	case *[]byte:
		return *x, nil
	case []byte:
		return x, nil
	}
	return nil, fmt.Errorf("rawCodec: %T", v)
}

func (rawCodec) Unmarshal(data []byte, v any) error {
	*(v.(*[]byte)) = append([]byte(nil), data...)
	return nil
}

func (rawCodec) Name() string { return "proto" }

// protoMessage encodes {1: chat, 2: text} as protobuf by hand (what generated code would emit).
func protoMessage(chat, text string) []byte {
	b := []byte{0x0a}
	b = binary.AppendUvarint(b, uint64(len(chat)))
	b = append(b, chat...)
	b = append(b, 0x12)
	b = binary.AppendUvarint(b, uint64(len(text)))
	return append(b, text...)
}

var chatService = grpc.ServiceDesc{
	ServiceName: "p2pact.bench.Chat",
	HandlerType: (*any)(nil),
	Methods: []grpc.MethodDesc{{MethodName: "Send", Handler: func(_ any, _ context.Context, dec func(any) error, _ grpc.UnaryServerInterceptor) (any, error) {
		var in []byte
		if err := dec(&in); err != nil {
			return nil, err
		}
		out := []byte{}
		return &out, nil
	}}},
}

// RunGRPC measures a unary gRPC call Chat/Send.
func RunGRPC(sc *Scenario, p Profile) (*Result, error) {
	res := newResult("gRPC", p.Name)
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, err
	}
	gs := grpc.NewServer(grpc.Creds(credentials.NewTLS(&tls.Config{Certificates: []tls.Certificate{selfSigned()}})),
		grpc.ForceServerCodec(rawCodec{}), grpc.MaxRecvMsgSize(64<<20))
	gs.RegisterService(&chatService, struct{}{})
	go gs.Serve(ln)
	defer gs.Stop()
	addr, stop, err := linkProxy(ln.Addr().String(), p)
	if err != nil {
		return nil, err
	}
	defer stop()
	var in, out atomic.Int64
	return measureTCP(sc, p, res, &in, &out, func() (func(string) error, func(), error) {
		cc, err := grpc.NewClient("passthrough:///"+addr,
			grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})),
			grpc.WithContextDialer(func(ctx context.Context, a string) (net.Conn, error) {
				c, err := (&net.Dialer{}).DialContext(ctx, "tcp", a)
				if err != nil {
					return nil, err
				}
				return counting{c, &in, &out}, nil
			}),
			grpc.WithDefaultCallOptions(grpc.ForceCodec(rawCodec{}), grpc.MaxCallSendMsgSize(64<<20)))
		if err != nil {
			return nil, nil, err
		}
		send := func(text string) error {
			req, resp := protoMessage(chatID, text), []byte{}
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
			defer cancel()
			return cc.Invoke(ctx, "/p2pact.bench.Chat/Send", &req, &resp)
		}
		return send, func() { cc.Close() }, nil
	})
}

// measureTCP runs the cold start and the sizes for a TCP-based protocol; connect returns a
// send function on a fresh connection and its close.
func measureTCP(sc *Scenario, p Profile, res *Result, in, out *atomic.Int64,
	connect func() (func(string) error, func(), error)) (*Result, error) {
	send, closeCold, err := connect()
	if err != nil {
		return nil, err
	}
	start := time.Now()
	if err := send(sc.message(sc.Sizes[0], 1)); err != nil {
		return nil, err
	}
	res.ColdMS = ms(time.Since(start))
	res.ColdBytes = tcpHandshake + tcpWire(out.Load(), in.Load()) + 3*tcpHeader
	closeCold()

	send, closeWarm, err := connect()
	if err != nil {
		return nil, err
	}
	defer closeWarm()
	if err := send("warm up"); err != nil {
		return nil, err
	}
	for _, s := range sc.Sizes {
		var rtts []time.Duration
		in0, out0 := in.Load(), out.Load()
		cnt := sc.count(s)
		for i := 0; i < cnt; i++ {
			t := time.Now()
			if err := send(sc.message(s, i)); err != nil {
				return nil, err
			}
			rtts = append(rtts, time.Since(t))
		}
		up, down := (out.Load()-out0)/int64(cnt), (in.Load()-in0)/int64(cnt)
		res.Payload[s.Name] = up + down
		res.Wire[s.Name] = tcpWire(up, down)
		res.RTTms[s.Name] = ms(median(rtts))
	}
	if p.loss() > 0 {
		res.Note = "no packet loss emulated for TCP (it would need root); real numbers would be worse"
	}
	return res, nil
}

// --- HTTP/3 over QUIC: UDP, so the link (loss included) is emulated exactly as for P2PACT ---

// udpProxy forwards datagrams between one client and a server with the profile applied in
// both directions, and counts them.
type udpProxy struct {
	front, back *net.UDPConn
	mu          sync.Mutex
	client      *net.UDPAddr
	bytes, pkts atomic.Int64
	free        [2]time.Time
	rng         *rand.Rand
}

func newUDPProxy(target string, p Profile, seed int64) (*udpProxy, error) {
	front, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		return nil, err
	}
	ta, _ := net.ResolveUDPAddr("udp", target)
	back, err := net.DialUDP("udp", nil, ta)
	if err != nil {
		front.Close()
		return nil, err
	}
	_ = front.SetReadBuffer(4 << 20)
	_ = back.SetReadBuffer(4 << 20)
	x := &udpProxy{front: front, back: back, rng: rand.New(rand.NewSource(seed))}
	deliver := func(dir int, b []byte, write func([]byte)) {
		x.mu.Lock()
		x.bytes.Add(int64(len(b)))
		x.pkts.Add(1)
		drop := p.loss() > 0 && x.rng.Float64() < p.loss()
		now := time.Now()
		due := now
		if p.rate() > 0 {
			if x.free[dir].Before(now) {
				x.free[dir] = now
			}
			x.free[dir] = x.free[dir].Add(time.Duration(len(b)+28) * time.Second / time.Duration(p.rate()))
			due = x.free[dir]
		}
		x.mu.Unlock()
		if drop {
			return
		}
		cp := append([]byte(nil), b...)
		time.AfterFunc(time.Until(due)+p.delay(), func() { write(cp) })
	}
	go func() {
		buf := make([]byte, 65536)
		for {
			n, from, err := front.ReadFromUDP(buf)
			if err != nil {
				return
			}
			x.mu.Lock()
			x.client = from
			x.mu.Unlock()
			deliver(0, buf[:n], func(b []byte) { _, _ = back.Write(b) })
		}
	}()
	go func() {
		buf := make([]byte, 65536)
		for {
			n, err := back.Read(buf)
			if err != nil {
				return
			}
			x.mu.Lock()
			c := x.client
			x.mu.Unlock()
			if c != nil {
				deliver(1, buf[:n], func(b []byte) { _, _ = front.WriteToUDP(b, c) })
			}
		}
	}()
	return x, nil
}

func (x *udpProxy) addr() string { return x.front.LocalAddr().String() }
func (x *udpProxy) close()       { x.front.Close(); x.back.Close() }
func (x *udpProxy) wire() int64  { return x.bytes.Load() + 28*x.pkts.Load() }

// RunHTTP3 measures POST /chats/<id>/messages over HTTP/3 (QUIC, TLS 1.3 inside).
func RunHTTP3(sc *Scenario, p Profile) (*Result, error) {
	res := newResult("HTTP/3", p.Name)
	pc, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		return nil, err
	}
	srv := &http3.Server{Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.Copy(io.Discard, r.Body)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"ok":true}`))
	}), TLSConfig: http3.ConfigureTLSConfig(&tls.Config{Certificates: []tls.Certificate{selfSigned()}})}
	go srv.Serve(pc)
	defer srv.Close()

	newClient := func(px *udpProxy) (*http.Client, func()) {
		tr := &http3.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			QUICConfig: &quic.Config{MaxIdleTimeout: time.Minute}}
		return &http.Client{Transport: tr}, func() { tr.Close() }
	}
	send := func(c *http.Client, px *udpProxy, text string) error {
		body, _ := json.Marshal(map[string]string{"text": text})
		req, _ := http.NewRequest("POST", "https://"+px.addr()+"/chats/"+chatID+"/messages", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+bearer)
		resp, err := c.Do(req)
		if err != nil {
			return err
		}
		_, _ = io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		return nil
	}

	cold, err := newUDPProxy(pc.LocalAddr().String(), p, sc.Seed)
	if err != nil {
		return nil, err
	}
	c, closeC := newClient(cold)
	start := time.Now()
	if err := send(c, cold, sc.message(sc.Sizes[0], 1)); err != nil {
		return nil, err
	}
	res.ColdMS = ms(time.Since(start))
	time.Sleep(3*p.delay() + 100*time.Millisecond)
	res.ColdBytes = cold.wire()
	closeC()
	cold.close()

	px, err := newUDPProxy(pc.LocalAddr().String(), p, sc.Seed+1)
	if err != nil {
		return nil, err
	}
	defer px.close()
	c, closeC = newClient(px)
	defer closeC()
	if err := send(c, px, "warm up"); err != nil {
		return nil, err
	}
	for _, s := range sc.Sizes {
		var rtts []time.Duration
		cnt := sc.count(s)
		time.Sleep(3*p.delay() + 50*time.Millisecond)
		b0, p0 := px.bytes.Load(), px.pkts.Load()
		for i := 0; i < cnt; i++ {
			t := time.Now()
			if err := send(c, px, sc.message(s, i)); err != nil {
				return nil, err
			}
			rtts = append(rtts, time.Since(t))
		}
		time.Sleep(3*p.delay() + 100*time.Millisecond)
		b, pk := px.bytes.Load()-b0, px.pkts.Load()-p0
		res.Payload[s.Name] = b / int64(cnt)
		res.Wire[s.Name] = (b + 28*pk) / int64(cnt)
		res.RTTms[s.Name] = ms(median(rtts))
	}
	return res, nil
}

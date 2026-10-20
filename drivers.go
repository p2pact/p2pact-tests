package bench

import (
	"bufio"
	"bytes"
	"context"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/fxamacker/cbor/v2"
	p2pact "github.com/p2pact/p2pact-go"
	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// The sending side. A client sends a chat message, waits for "accepted" and checks the
// fingerprint; bytes() is what crossed its link so far, both directions.

var errMismatch = errors.New("the receiver got something else")

type client interface {
	send(text string) error
	bytes() (up, down, packets int64)
	close()
}

// driver opens clients of one protocol over one emulated link.
type driver struct {
	name string
	udp  bool // bytes are datagrams (exact headers) rather than a TCP stream (estimated)
	dial func() (client, error)
	stop func()
}

func bufioReader(c net.Conn) *bufio.Reader { return bufio.NewReaderSize(c, 64<<10) }

func checkHex(got string, text string) error {
	if got != hex.EncodeToString(fingerprint(text)) {
		return errMismatch
	}
	return nil
}

// --- P2PACT ---

// p2pactSetup pairs a client device with the server once; later clients reuse its keys, so
// each new client is a cold start between paired devices (a new session, no pairing).
type p2pactSetup struct {
	dir    string
	server string // the server's address
	client string // ours
}

func newP2PACTSetup(invite string) (*p2pactSetup, error) {
	dir, err := os.MkdirTemp("", "p2pact-bench-client")
	if err != nil {
		return nil, err
	}
	n, err := p2pact.NewNode(dir, "bench-client")
	if err != nil {
		return nil, err
	}
	defer n.Close()
	if err := n.ListenUDP(":0"); err != nil {
		return nil, err
	}
	if err := n.ListenTCP(""); err != nil {
		return nil, err
	}
	n.Start()
	inv, err := p2pact.ParseInvite(invite)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	p, err := n.Join(ctx, inv)
	if err != nil {
		return nil, fmt.Errorf("pairing with the server: %w", err)
	}
	return &p2pactSetup{dir: dir, server: p.Addr, client: n.ID.Addr()}, nil
}

type p2pactClient struct {
	n      *p2pact.Node
	road   string
	server string
	calls  atomic.Uint64
}

// dial starts a fresh client device with the paired keys; addr is where the server is
// reached on this link (a proxy or the server itself).
func (s *p2pactSetup) dial(road, addr string) (client, error) {
	n, err := p2pact.NewNode(s.dir, "bench-client")
	if err != nil {
		return nil, err
	}
	if road == "udp" {
		err = n.ListenUDP(":0")
	} else {
		err = n.ListenTCP("")
	}
	if err != nil {
		n.Close()
		return nil, err
	}
	h := p2pact.Hints{}
	if road == "udp" {
		h.UDP = []string{addr}
	} else {
		h.TCP = []string{addr}
	}
	if p := n.Trust.Get(s.server); p != nil {
		p.Hints = h
		n.Trust.Put(*p)
	}
	// keep the server on this link even after its hello lists other addresses
	n.OnHello = func(peer string) {
		if p := n.Trust.Get(peer); p != nil {
			p.Hints = h
			n.Trust.Put(*p)
		}
	}
	// announce no addresses of our own: the server must answer over the link under test
	// (the emulating proxy), not find a direct way around it
	n.SetAdvertise(&p2pact.Hints{})
	n.Start()
	return &p2pactClient{n: n, road: road, server: s.server}, nil
}

func (c *p2pactClient) send(text string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	// one stream per chat: messages of different chats do not wait for each other's losses
	stream := 1 + c.calls.Add(1)%64
	res, err := c.n.CallOn(ctx, c.server, stream, 1, 2, stream, text)
	if err != nil {
		return err
	}
	var got []byte
	if len(res) != 1 || cbor.Unmarshal(res[0], &got) != nil || !bytes.Equal(got, fingerprint(text)) {
		return errMismatch
	}
	return nil
}

func (c *p2pactClient) bytes() (int64, int64, int64) {
	var st p2pact.RoadStats
	if c.road == "udp" {
		st = c.n.UDP().Stats()
	} else {
		st = c.n.TCP().Stats()
	}
	return st.SentBytes, st.RecvBytes, st.SentPackets + st.RecvPackets
}

func (c *p2pactClient) close() { c.n.Close() }

// --- HTTP: HTTPS (HTTP/1.1), HTTP/2 ---

type httpClient struct {
	c       *http.Client
	url     string
	in, out atomic.Int64
	proto   int
}

func dialHTTP(addr string, h2 bool) client {
	hc := &httpClient{url: "https://" + addr + "/chats/" + chatID + "/messages", proto: 1}
	if h2 {
		hc.proto = 2
	}
	tr := &http.Transport{
		DialContext: func(ctx context.Context, network, a string) (net.Conn, error) {
			c, err := (&net.Dialer{}).DialContext(ctx, network, a)
			if err != nil {
				return nil, err
			}
			return counting{c, &hc.in, &hc.out}, nil
		},
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		ForceAttemptHTTP2:   h2,
		MaxIdleConnsPerHost: 256,
	}
	if !h2 {
		tr.TLSNextProto = map[string]func(string, *tls.Conn) http.RoundTripper{}
	}
	hc.c = &http.Client{Transport: tr, Timeout: 2 * time.Minute}
	return hc
}

func (h *httpClient) send(text string) error {
	body, _ := json.Marshal(map[string]string{"text": text})
	req, _ := http.NewRequest("POST", h.url, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+bearer)
	resp, err := h.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.ProtoMajor != h.proto {
		return fmt.Errorf("negotiated %s", resp.Proto)
	}
	var m struct{ H string }
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return err
	}
	return checkHex(m.H, text)
}

func (h *httpClient) bytes() (int64, int64, int64) { return h.out.Load(), h.in.Load(), 0 }
func (h *httpClient) close()                       { h.c.CloseIdleConnections() }

// --- HTTP/3 ---

type h3Client struct {
	c   *http.Client
	tr  *http3.Transport
	url string
	px  *udpProxy
}

func (h *h3Client) send(text string) error {
	body, _ := json.Marshal(map[string]string{"text": text})
	req, _ := http.NewRequest("POST", h.url, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+bearer)
	resp, err := h.c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var m struct{ H string }
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return err
	}
	return checkHex(m.H, text)
}

func (h *h3Client) bytes() (int64, int64, int64) {
	return h.px.up.Load(), h.px.down.Load(), h.px.pkts.Load()
}
func (h *h3Client) close() { h.tr.Close() }

// --- gRPC ---

type grpcClient struct {
	cc      *grpc.ClientConn
	in, out atomic.Int64
}

func dialGRPC(addr string) (client, error) {
	g := &grpcClient{}
	cc, err := grpc.NewClient("passthrough:///"+addr,
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})),
		grpc.WithContextDialer(func(ctx context.Context, a string) (net.Conn, error) {
			c, err := (&net.Dialer{}).DialContext(ctx, "tcp", a)
			if err != nil {
				return nil, err
			}
			return counting{c, &g.in, &g.out}, nil
		}),
		grpc.WithDefaultCallOptions(grpc.ForceCodec(rawCodec{}), grpc.MaxCallSendMsgSize(64<<20), grpc.MaxCallRecvMsgSize(64<<20)))
	if err != nil {
		return nil, err
	}
	g.cc = cc
	return g, nil
}

func (g *grpcClient) send(text string) error {
	req, resp := protoMessage(chatID, text), []byte{}
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	if err := g.cc.Invoke(ctx, "/p2pact.bench.Chat/Send", &req, &resp); err != nil {
		return err
	}
	if !bytes.Equal(resp, fingerprint(text)) {
		return errMismatch
	}
	return nil
}

func (g *grpcClient) bytes() (int64, int64, int64) { return g.out.Load(), g.in.Load(), 0 }
func (g *grpcClient) close()                       { g.cc.Close() }

// --- WebSocket: messages pipelined on one connection, replies in order ---

type wsClient struct {
	c       net.Conn
	r       *bufio.Reader
	in, out atomic.Int64
	mu      sync.Mutex
	waiting []chan string
	wmu     sync.Mutex
	err     error
}

func dialWS(addr string) (client, error) {
	w := &wsClient{}
	raw, err := net.DialTimeout("tcp", addr, 10*time.Second)
	if err != nil {
		return nil, err
	}
	c := tls.Client(counting{raw, &w.in, &w.out}, &tls.Config{InsecureSkipVerify: true})
	if _, err := io.WriteString(c, "GET /chat HTTP/1.1\r\nHost: "+addr+"\r\nUpgrade: websocket\r\nConnection: Upgrade\r\n"+
		"Sec-WebSocket-Key: dGhlIHNhbXBsZSBub25jZQ==\r\nSec-WebSocket-Version: 13\r\nAuthorization: Bearer "+bearer+"\r\n\r\n"); err != nil {
		return nil, err
	}
	w.c, w.r = c, bufioReader(c)
	for {
		line, err := w.r.ReadString('\n')
		if err != nil {
			return nil, err
		}
		if line == "\r\n" {
			break
		}
	}
	go func() {
		for {
			b, err := readFrame(w.r)
			w.mu.Lock()
			if err != nil {
				w.err = err
				for _, ch := range w.waiting {
					close(ch)
				}
				w.waiting = nil
				w.mu.Unlock()
				return
			}
			var m struct{ H string }
			_ = json.Unmarshal(b, &m)
			if len(w.waiting) > 0 {
				w.waiting[0] <- m.H
				w.waiting = w.waiting[1:]
			}
			w.mu.Unlock()
		}
	}()
	return w, nil
}

func (w *wsClient) send(text string) error {
	msg, _ := json.Marshal(map[string]string{"type": "send", "chat": chatID, "text": text})
	ch := make(chan string, 1)
	w.wmu.Lock()
	w.mu.Lock()
	if w.err != nil {
		w.mu.Unlock()
		w.wmu.Unlock()
		return w.err
	}
	w.waiting = append(w.waiting, ch)
	w.mu.Unlock()
	_, err := w.c.Write(frame(msg, true))
	w.wmu.Unlock()
	if err != nil {
		return err
	}
	select {
	case h, ok := <-ch:
		if !ok {
			return errors.New("connection closed")
		}
		return checkHex(h, text)
	case <-time.After(2 * time.Minute):
		return errors.New("no reply")
	}
}

func (w *wsClient) bytes() (int64, int64, int64) { return w.out.Load(), w.in.Load(), 0 }
func (w *wsClient) close()                       { w.c.Close() }

// --- drivers over one link ---

// Target is where the receivers are (local or the remote server).
type Target struct {
	Servers                                       *Servers // local only
	Invite                                        func() (string, error)
	P2PUDP, P2PTCP, WS, HTTPS, HTTP2, GRPC, HTTP3 string
}

// makeDriver puts the profile's emulated link in front of the target of one protocol.
func makeDriver(sc *Scenario, p Profile, t *Target, setup *p2pactSetup, proto string) (*driver, error) {
	tcp := func(target string, dial func(addr string) (client, error), name string) (*driver, error) {
		addr, stop, err := linkProxy(target, p)
		if err != nil {
			return nil, err
		}
		return &driver{name: name, dial: func() (client, error) { return dial(addr) }, stop: stop}, nil
	}
	switch proto {
	case "p2pact-udp":
		px, err := newUDPProxy(t.P2PUDP, p, sc.Seed)
		if err != nil {
			return nil, err
		}
		if t.Servers != nil { // local: the server announces the proxy
			t.Servers.Node.SetAdvertise(&p2pact.Hints{UDP: []string{px.addr()}})
		}
		return &driver{name: "P2PACT UDP", udp: true, dial: func() (client, error) { return setup.dial("udp", px.addr()) },
			stop: func() {
				px.close()
				if t.Servers != nil {
					t.Servers.Node.SetAdvertise(nil)
				}
			}}, nil
	case "p2pact-tcp":
		addr, stop, err := linkProxy(t.P2PTCP, p)
		if err != nil {
			return nil, err
		}
		if t.Servers != nil {
			t.Servers.Node.SetAdvertise(&p2pact.Hints{TCP: []string{addr}})
		}
		return &driver{name: "P2PACT TCP", dial: func() (client, error) { return setup.dial("tcp", addr) },
			stop: func() {
				stop()
				if t.Servers != nil {
					t.Servers.Node.SetAdvertise(nil)
				}
			}}, nil
	case "websocket":
		return tcp(t.WS, dialWS, "WebSocket")
	case "https":
		return tcp(t.HTTPS, func(a string) (client, error) { return dialHTTP(a, false), nil }, "HTTPS+JSON")
	case "http2":
		return tcp(t.HTTP2, func(a string) (client, error) { return dialHTTP(a, true), nil }, "HTTP/2")
	case "grpc":
		return tcp(t.GRPC, dialGRPC, "gRPC")
	case "http3":
		px, err := newUDPProxy(t.HTTP3, p, sc.Seed+1)
		if err != nil {
			return nil, err
		}
		return &driver{name: "HTTP/3", udp: true, dial: func() (client, error) {
			tr := &http3.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, QUICConfig: &quic.Config{MaxIdleTimeout: time.Minute}}
			return &h3Client{c: &http.Client{Transport: tr, Timeout: 2 * time.Minute}, tr: tr,
				url: "https://" + px.addr() + "/chats/" + chatID + "/messages", px: px}, nil
		}, stop: px.close}, nil
	}
	return nil, fmt.Errorf("unknown protocol %q (want one of %v)", proto, AllProtocols)
}

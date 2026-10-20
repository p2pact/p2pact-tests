package bench

import (
	"context"
	"crypto/sha256"
	"crypto/tls"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/fxamacker/cbor/v2"
	p2pact "github.com/p2pact/p2pact-go"
	"github.com/quic-go/quic-go/http3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// The receiving side of every protocol. It answers each chat message with a fingerprint of
// what it received (the first 8 bytes of SHA-256 of the text), so the sender can check that
// the message arrived intact.

// fingerprint is what the receiver sends back for a message.
func fingerprint(text string) []byte {
	h := sha256.Sum256([]byte(text))
	return h[:8]
}

// Ports are the fixed ports of a remote server (p2pact-bench serve).
const BasePort = 48731

// Servers are the receivers of all protocols.
type Servers struct {
	Node                                          *p2pact.Node
	P2PUDP, P2PTCP, WS, HTTPS, HTTP2, GRPC, HTTP3 string
	closers                                       []func()
}

func (s *Servers) Close() {
	for _, c := range s.closers {
		c()
	}
}

// StartServers starts every receiver on host; base 0 — free ports, else base, base+1, …
func StartServers(dir, host string, base int) (*Servers, error) {
	port := func(i int) string {
		if base == 0 {
			return net.JoinHostPort(host, "0")
		}
		return net.JoinHostPort(host, strconv.Itoa(base+i))
	}
	s := &Servers{}
	fail := func(err error) (*Servers, error) { s.Close(); return nil, err }
	cert := selfSigned()

	// P2PACT: UDP and TCP on the same port number
	n, err := p2pact.NewNode(dir, "bench-server")
	if err != nil {
		return fail(err)
	}
	s.Node = n
	if err := n.ListenUDP(port(0)); err != nil {
		return fail(err)
	}
	tcpAddr := port(0)
	if base == 0 {
		_, p, _ := net.SplitHostPort(n.UDP().LocalAddr())
		tcpAddr = net.JoinHostPort(host, p)
	}
	if err := n.ListenTCP(tcpAddr); err != nil {
		return fail(err)
	}
	n.Handle(1, 2, func(_ string, args []cbor.RawMessage) ([]any, error) {
		var text string
		if len(args) < 2 || cbor.Unmarshal(args[1], &text) != nil {
			return nil, &p2pact.Error{Code: p2pact.ErrBadArgs, Text: "want [chat, text]"}
		}
		return []any{fingerprint(text)}, nil
	})
	n.Start()
	s.closers = append(s.closers, n.Close)
	s.P2PUDP, s.P2PTCP = n.UDP().LocalAddr(), n.TCP().LocalAddr()

	jsonReply := func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/invite" {
			_, _ = io.WriteString(w, n.Invite(10*time.Minute).String())
			return
		}
		var m struct{ Text string }
		_ = json.NewDecoder(r.Body).Decode(&m)
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{"ok":true,"h":"%s"}`, hex.EncodeToString(fingerprint(m.Text)))
	}
	listenTCP := func(i int) (net.Listener, error) { return net.Listen("tcp", port(i)) }

	// WebSocket: upgrade, then one reply frame per message frame, in order
	ln, err := tls.Listen("tcp", port(1), &tls.Config{Certificates: []tls.Certificate{cert}})
	if err != nil {
		return fail(err)
	}
	s.WS = ln.Addr().String()
	s.closers = append(s.closers, func() { ln.Close() })
	go func() {
		for {
			c, err := ln.Accept()
			if err != nil {
				return
			}
			go serveWS(c)
		}
	}()

	// HTTPS (HTTP/1.1 only) and HTTP/2
	for i, h2 := range []bool{false, true} {
		l, err := listenTCP(2 + i)
		if err != nil {
			return fail(err)
		}
		srv := &http.Server{Handler: http.HandlerFunc(jsonReply), TLSConfig: &tls.Config{Certificates: []tls.Certificate{cert}}}
		if !h2 {
			srv.TLSNextProto = map[string]func(*http.Server, *tls.Conn, http.Handler){}
			s.HTTPS = l.Addr().String()
		} else {
			s.HTTP2 = l.Addr().String()
		}
		go srv.ServeTLS(l, "", "")
		s.closers = append(s.closers, func() { srv.Close() })
	}

	// gRPC
	gl, err := listenTCP(4)
	if err != nil {
		return fail(err)
	}
	gs := grpc.NewServer(grpc.Creds(credentials.NewTLS(&tls.Config{Certificates: []tls.Certificate{cert}})),
		grpc.ForceServerCodec(rawCodec{}), grpc.MaxRecvMsgSize(64<<20))
	gs.RegisterService(&chatService, struct{}{})
	go gs.Serve(gl)
	s.GRPC = gl.Addr().String()
	s.closers = append(s.closers, gs.Stop)

	// HTTP/3
	pc, err := net.ListenPacket("udp", port(5))
	if err != nil {
		return fail(err)
	}
	h3 := &http3.Server{Handler: http.HandlerFunc(jsonReply), TLSConfig: http3.ConfigureTLSConfig(&tls.Config{Certificates: []tls.Certificate{cert}})}
	go h3.Serve(pc)
	s.HTTP3 = pc.LocalAddr().String()
	s.closers = append(s.closers, func() { h3.Close(); pc.Close() })
	return s, nil
}

func serveWS(c net.Conn) {
	defer c.Close()
	r := bufioReader(c)
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			return
		}
		if line == "\r\n" {
			break
		}
	}
	_, _ = io.WriteString(c, "HTTP/1.1 101 Switching Protocols\r\nUpgrade: websocket\r\nConnection: Upgrade\r\n"+
		"Sec-WebSocket-Accept: s3pPLMBiTxaQ9kYGzzhZRbK+xOo=\r\n\r\n")
	for {
		b, err := readFrame(r)
		if err != nil {
			return
		}
		var m struct{ Text string }
		_ = json.Unmarshal(b, &m)
		if _, err := c.Write(frame([]byte(`{"type":"ok","h":"`+hex.EncodeToString(fingerprint(m.Text))+`"}`), false)); err != nil {
			return
		}
	}
}

var chatService = grpc.ServiceDesc{
	ServiceName: "p2pact.bench.Chat",
	HandlerType: (*any)(nil),
	Methods: []grpc.MethodDesc{{MethodName: "Send", Handler: func(_ any, _ context.Context, dec func(any) error, _ grpc.UnaryServerInterceptor) (any, error) {
		var in []byte
		if err := dec(&in); err != nil {
			return nil, err
		}
		out := fingerprint(protoText(in))
		return &out, nil
	}}},
}

// protoMessage encodes {1: chat, 2: text} as protobuf by hand (what generated code would emit).
func protoMessage(chat, text string) []byte {
	b := []byte{0x0a}
	b = binary.AppendUvarint(b, uint64(len(chat)))
	b = append(b, chat...)
	b = append(b, 0x12)
	b = binary.AppendUvarint(b, uint64(len(text)))
	return append(b, text...)
}

// protoText reads field 2 (text) of that message.
func protoText(b []byte) string {
	for len(b) > 0 {
		tag := b[0]
		n, k := binary.Uvarint(b[1:])
		if k <= 0 || int(n) > len(b)-1-k {
			return ""
		}
		v := b[1+k : 1+k+int(n)]
		if tag == 0x12 {
			return string(v)
		}
		b = b[1+k+int(n):]
	}
	return ""
}

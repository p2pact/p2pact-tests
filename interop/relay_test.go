package interop

import (
	"crypto/tls"
	"net"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/p2pact/p2pact-go/relay"
)

// dropListener remembers what it accepted, so a test can cut every connection to the relay.
type dropListener struct {
	net.Listener
	mu    sync.Mutex
	conns []net.Conn
}

func (l *dropListener) Accept() (net.Conn, error) {
	c, err := l.Listener.Accept()
	if err == nil {
		l.mu.Lock()
		l.conns = append(l.conns, c)
		l.mu.Unlock()
	}
	return c, err
}

func (l *dropListener) drop() {
	l.mu.Lock()
	defer l.mu.Unlock()
	for _, c := range l.conns {
		_ = c.Close()
	}
	l.conns = nil
}

// startRelay runs the Go relay in this process: TLS with a self-signed certificate (the devices
// pin its fingerprint) and a plain WebSocket listener on loopback.
func startRelay(t *testing.T) (addr, fp, ws string, l *dropListener) {
	t.Helper()
	dir := t.TempDir()
	cert, err := relay.LoadOrMakeCert(filepath.Join(dir, "cert.pem"), filepath.Join(dir, "key.pem"))
	if err != nil {
		t.Fatal(err)
	}
	srv := relay.NewServer()
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	l = &dropListener{Listener: ln}
	go func() { _ = srv.Serve(tls.NewListener(l, &tls.Config{Certificates: []tls.Certificate{cert}, MinVersion: tls.VersionTLS12})) }()
	wl, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	go func() { _ = srv.ServeWS(wl) }()
	t.Cleanup(func() { _ = ln.Close(); _ = wl.Close(); l.drop() })
	return ln.Addr().String(), relay.Fingerprint(cert), "ws://" + wl.Addr().String() + "/", l
}

// TestInteropRelay: every pair talks over the relay road (spec §3) and nothing else — neither
// device listens on UDP or TCP, so the invite carries only the relay (and its WebSocket URL,
// which a native device just keeps). They pair through the relay, exchange messages both ways,
// and keep talking after the relay cuts every connection (both sides reconnect).
func TestInteropRelay(t *testing.T) {
	list := impls(t)
	for _, a := range list {
		for _, b := range list {
			t.Run(a.name+"-"+b.name, func(t *testing.T) {
				addr, fp, ws, rl := startRelay(t)
				owner, joiner := start(t, a, a.name+"-owner"), start(t, b, b.name+"-joiner")
				oa := owner.do(map[string]any{"cmd": "listen"})["addr"].(string)
				ja := joiner.do(map[string]any{"cmd": "listen"})["addr"].(string)
				hint := map[string]any{"cmd": "relay", "addr": addr, "fp": fp, "ws": []string{ws}}
				for _, p := range []*peer{owner, joiner} {
					if _, e := p.try(hint); e != "" {
						t.Skipf("%s has no relay command: %s", p.name, e)
					}
				}
				inv := owner.do(map[string]any{"cmd": "invite"})["invite"].(string)
				if got := joiner.do(map[string]any{"cmd": "join", "invite": inv})["peer"]; got != oa {
					t.Fatalf("joined %v, want %s", got, oa)
				}
				eventually(t, "the owner knows the joiner", func() bool {
					return owner.do(map[string]any{"cmd": "known", "peer": ja})["known"] == true
				})
				stream := 0
				for name, data := range messages() {
					stream = (stream + 1) % 3
					call(t, name, joiner, oa, owner, ja, data, stream)
				}
				answerLater(t, joiner, oa, owner, ja, 2)
				for i := range 50 {
					call(t, "conversation", joiner, oa, owner, ja, chatLine(i), i%3)
				}
				// the relay drops everyone: the devices reconnect and the talk goes on
				rl.drop()
				time.Sleep(200 * time.Millisecond)
				for i := 50; i < 80; i++ {
					call(t, "after the relay dropped", joiner, oa, owner, ja, chatLine(i), i%3)
				}
			})
		}
	}
}

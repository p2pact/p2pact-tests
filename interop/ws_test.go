package interop

// The relay over WebSocket (core spec §3): the Go relay's WebSocket listener on loopback, with a
// plain ws:// URL. A device with a WebSocket road (today the TypeScript one) meets a device of
// any language that dials the relay over TLS and tells its WebSocket URL in the invite hints.

import (
	"context"
	"crypto/rand"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/coder/websocket"
	"github.com/p2pact/p2pact-go/relay"
)

// startWSRelay runs the relay in this process: TLS (pinned) and a WebSocket listener on
// loopback. The WebSocket address is dropListener-wrapped so a test can cut the WebSocket
// connections only. With proxy the relay trusts X-Forwarded-For and the devices reach the
// WebSocket listener through a reverse proxy of ours that sets it (one proxy per device).
func startWSRelay(t *testing.T, proxy bool) (addr, fp string, ws func(ip string) string, drop func()) {
	t.Helper()
	dir := t.TempDir()
	cert, err := relay.LoadOrMakeCert(filepath.Join(dir, "cert.pem"), filepath.Join(dir, "key.pem"))
	if err != nil {
		t.Fatal(err)
	}
	srv := relay.NewServer()
	if proxy {
		srv.TrustProxy()
	}
	tl, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	tlsL := &dropListener{Listener: tl}
	go func() { _ = srv.Serve(tlsCfg(tlsL, cert)) }()
	wl, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	wsL := &dropListener{Listener: wl}
	go func() { _ = srv.ServeWS(wsL) }()
	drops := []func(){wsL.drop}
	t.Cleanup(func() { _ = tl.Close(); _ = wl.Close(); tlsL.drop(); wsL.drop() })

	ws = func(ip string) string {
		if !proxy {
			return "ws://" + wl.Addr().String() + "/"
		}
		return startXFFProxy(t, wl.Addr().String(), ip, &drops)
	}
	drop = func() {
		for _, d := range drops {
			d()
		}
	}
	return tl.Addr().String(), relay.Fingerprint(cert), ws, drop
}

// startXFFProxy is a tiny reverse proxy in front of the relay's WebSocket listener: it sets
// X-Forwarded-For to ip, as nginx would. Its client connections can be cut (the returned list of
// drops grows by one).
func startXFFProxy(t *testing.T, upstream, ip string, drops *[]func()) string {
	t.Helper()
	rp := &httputil.ReverseProxy{
		Rewrite: func(r *httputil.ProxyRequest) {
			r.SetURL(&url.URL{Scheme: "http", Host: upstream})
			r.Out.Header.Set("X-Forwarded-For", ip)
		},
	}
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}
	dl := &dropListener{Listener: ln}
	srv := &http.Server{Handler: rp, ReadHeaderTimeout: 15 * time.Second}
	go func() { _ = srv.Serve(dl) }()
	*drops = append(*drops, dl.drop)
	t.Cleanup(func() { _ = srv.Close(); dl.drop() })
	return "ws://" + ln.Addr().String() + "/"
}

func tlsCfg(l net.Listener, cert tls.Certificate) net.Listener {
	return tls.NewListener(l, &tls.Config{Certificates: []tls.Certificate{cert}, MinVersion: tls.VersionTLS12})
}

// wsPair says whether the pair can meet over the relay with ts on the WebSocket road: the
// WebSocket-only device must be the joiner (an invite from it has no TLS relay for the other
// to dial), the owner must tell its WebSocket URL in the hints (Go, Swift and TS do).
func wsPair(t *testing.T, a, b impl) {
	t.Helper()
	if b.name != "ts" {
		t.Skip("neither side is the joiner with a WebSocket road")
	}
	switch a.name {
	case "go", "swift", "ts":
	default:
		t.Skipf("%s does not tell a relay WebSocket URL in the invite", a.name)
	}
}

// TestInteropWSRelay: a device on the WebSocket road (ts) and a device on the TLS road meet
// through the relay's WebSocket listener, call both ways, exchange events, and go on after the
// WebSocket connections are cut. Once with the devices connected straight to the listener, once
// behind a reverse proxy that sets X-Forwarded-For (the relay trusts it).
func TestInteropWSRelay(t *testing.T) {
	list := impls(t)
	hasTS := false
	for _, im := range list {
		hasTS = hasTS || im.name == "ts"
	}
	if !hasTS {
		t.Skip("set P2PACT_TS: the WebSocket road is a TypeScript one")
	}
	for _, mode := range []string{"direct", "proxy"} {
		for _, a := range list {
			for _, b := range list {
				t.Run(mode+"/"+a.name+"-"+b.name, func(t *testing.T) {
					wsPair(t, a, b)
					addr, fp, ws, drop := startWSRelay(t, mode == "proxy")
					owner, joiner := start(t, a, a.name+"-owner"), start(t, b, b.name+"-joiner")
					oa := owner.do(map[string]any{"cmd": "listen"})["addr"].(string)
					ja := joiner.do(map[string]any{"cmd": "listen"})["addr"].(string)
					for i, p := range []*peer{owner, joiner} {
						hint := map[string]any{"cmd": "relay", "addr": addr, "fp": fp, "ws": []string{ws(fmt.Sprintf("10.0.0.%d", i+1))}}
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
					msgs := messages()
					for _, name := range []string{"short", "russian", "random 300 KB"} {
						call(t, name, joiner, oa, owner, ja, msgs[name], 1)
					}
					answerLater(t, joiner, oa, owner, ja, 2)
					for i := range 20 {
						call(t, "conversation", joiner, oa, owner, ja, chatLine(i), i%3)
					}
					// the WebSocket connections drop: the device reconnects and the talk goes on
					drop()
					time.Sleep(200 * time.Millisecond)
					for i := 20; i < 40; i++ {
						call(t, "after the WebSocket dropped", joiner, oa, owner, ja, chatLine(i), i%3)
					}
					// events both ways (a device with read on a tool gets them; this narrows the grants, so last)
					wsEvents(t, owner, ja, joiner, oa, 0)
					wsEvents(t, joiner, oa, owner, ja, 1)
				})
			}
		}
	}
}

// wsEvents: from registers a named tool (one per round), grants to read, emits an event, and
// to receives it with the tool name.
func wsEvents(t *testing.T, from *peer, toAddr string, to *peer, fromAddr string, round int) {
	t.Helper()
	tool := fmt.Sprintf("com.example.ws%d", round)
	from.do(map[string]any{"cmd": "tool", "name": tool})
	from.do(map[string]any{"cmd": "set-grants", "peer": toAddr, "grants": []any{gl(tool, 1)}})
	waitGranted(t, to, fromAddr, []any{gl(tool, 1)})
	before := len(events(to))
	from.do(map[string]any{"cmd": "emit", "peer": toAddr, "name": tool, "event": 1, "data": b64(tool)})
	eventually(t, from.name+" → "+to.name+": the event", func() bool { return len(events(to)) > before })
	ev := events(to)[before].(map[string]any)
	if ev["tool"] != tool || ev["event"] != float64(1) || ev["data"] != b64(tool) {
		t.Fatalf("event %v", ev)
	}
}

// relayBinary builds p2pact-relay and starts it with the flags; it returns the TLS address and
// the WebSocket URL the relay logged.
func relayBinary(t *testing.T, flags ...string) (tlsAddr, wsURL string) {
	t.Helper()
	dir := t.TempDir()
	bin := filepath.Join(dir, "relay")
	if out, err := exec.Command("go", "build", "-o", bin, "github.com/p2pact/p2pact-go/cmd/p2pact-relay").CombinedOutput(); err != nil {
		t.Fatalf("build the relay: %v\n%s", err, out)
	}
	cmd := exec.Command(bin, append([]string{"-data", dir, "-addr", "127.0.0.1:0"}, flags...)...)
	pr, pw, _ := os.Pipe()
	cmd.Stderr, cmd.Stdout = pw, pw
	if err := cmd.Start(); err != nil {
		t.Fatal(err)
	}
	_ = pw.Close()
	t.Cleanup(func() { _ = cmd.Process.Kill(); _ = cmd.Wait(); _ = pr.Close() })
	re := regexp.MustCompile(`url=(\S+)|addr=(\S+)`)
	found := make(chan struct{})
	var mu sync.Mutex
	go func() {
		buf := make([]byte, 4096)
		var text string
		for {
			n, err := pr.Read(buf)
			text += string(buf[:n])
			for _, m := range re.FindAllStringSubmatch(text, -1) {
				mu.Lock()
				if m[1] != "" {
					wsURL = m[1]
				}
				if m[2] != "" {
					tlsAddr = m[2]
				}
				ok := wsURL != "" && tlsAddr != ""
				mu.Unlock()
				if ok {
					close(found)
					return
				}
			}
			if err != nil {
				return
			}
		}
	}()
	select {
	case <-found:
	case <-time.After(60 * time.Second):
		t.Fatal("the relay did not start")
	}
	mu.Lock()
	defer mu.Unlock()
	return tlsAddr, wsURL
}

// wsJoin opens a raw WebSocket to the relay on a fresh channel and returns the relay's first
// answer ({"ok":true} or an error line) with the connection to hold.
func wsJoin(t *testing.T, url string) (*websocket.Conn, string) {
	t.Helper()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	c, _, err := websocket.Dial(ctx, url, nil)
	if err != nil {
		t.Fatalf("dial %s: %v", url, err)
	}
	t.Cleanup(func() { _ = c.CloseNow() })
	var ch [16]byte
	_, _ = rand.Read(ch[:])
	if err := c.Write(ctx, websocket.MessageText, []byte(`{"relay":1,"channel":"`+hex.EncodeToString(ch[:])+`","role":"a"}`)); err != nil {
		// the relay may already have refused and closed
		return c, "write: " + err.Error()
	}
	_, msg, err := c.Read(ctx)
	if err != nil {
		return c, "read: " + err.Error()
	}
	return c, strings.TrimSpace(string(msg))
}

// TestWSRelayPerIP runs the relay binary with -ws-addr and -ws-trust-proxy: connections are
// counted per client address, and behind a trusted proxy that address is X-Forwarded-For — one
// address cannot take more than 16 connections, another is not held up by it. Without
// -ws-trust-proxy the proxy's own address (loopback) is what counts, whatever the header says.
func TestWSRelayPerIP(t *testing.T) {
	const limit = 16
	for _, trust := range []bool{true, false} {
		t.Run(fmt.Sprintf("trust-proxy=%v", trust), func(t *testing.T) {
			flags := []string{"-ws-addr", "127.0.0.1:0"}
			if trust {
				flags = append(flags, "-ws-trust-proxy")
			}
			_, direct := relayBinary(t, flags...)
			u, _ := url.Parse(direct)
			var drops []func()
			a := startXFFProxy(t, u.Host, "198.51.100.1", &drops)
			b := startXFFProxy(t, u.Host, "198.51.100.2", &drops)
			for i := range limit {
				if _, ans := wsJoin(t, a); ans != `{"ok":true}` {
					t.Fatalf("connection %d of %d via a: %s", i+1, limit, ans)
				}
			}
			if _, ans := wsJoin(t, a); !strings.Contains(ans, "too many connections") {
				t.Fatalf("connection %d via a: %s, want too many connections", limit+1, ans)
			}
			_, ans := wsJoin(t, b)
			if trust && ans != `{"ok":true}` {
				t.Fatalf("another address behind the proxy: %s, want ok", ans)
			}
			if !trust && !strings.Contains(ans, "too many connections") {
				t.Fatalf("without -ws-trust-proxy the proxy's address counts: %s, want too many connections", ans)
			}
		})
	}
}

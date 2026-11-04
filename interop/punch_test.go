package interop

import (
	"crypto/rand"
	"net"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/p2pact/p2pact-go/relay"
)

// pathsOf asks p for its live roads to the device at addr.
func pathsOf(t *testing.T, p *peer, addr string) []string {
	t.Helper()
	var out []string
	for _, v := range p.do(map[string]any{"cmd": "paths", "peer": addr})["paths"].([]any) {
		out = append(out, v.(string))
	}
	return out
}

// TestInteropPunch: direct across the internet (core spec §7, observe §4.1). Both devices listen
// on UDP but advertise no address of their own, so the only road they share at first is the
// relay: they pair and talk over it. Then the relay starts answering observe on UDP (the port
// number of its TLS listener) and the devices are told their relay again, which makes them ask
// it at once where they are seen from. Each puts that address into its hints, the hello carries
// it, and both punch: the session must move to UDP within a few seconds, and messages must still
// flow. On loopback the observed address is the socket's own, so a punch always gets through.
//
// A peer without the relay, advertise or paths command does not punch: the pair is skipped
// (TestInterop and TestInteropRelay keep checking it on its other roads).
func TestInteropPunch(t *testing.T) {
	list := impls(t)
	for _, a := range list {
		for _, b := range list {
			t.Run(a.name+"-"+b.name, func(t *testing.T) {
				addr, fp, _, _ := startRelay(t)
				owner, joiner := start(t, a, a.name+"-owner"), start(t, b, b.name+"-joiner")
				oa := owner.do(map[string]any{"cmd": "listen", "udp": "127.0.0.1:0"})["addr"].(string)
				ja := joiner.do(map[string]any{"cmd": "listen", "udp": "127.0.0.1:0"})["addr"].(string)
				hint := map[string]any{"cmd": "relay", "addr": addr, "fp": fp}
				noAddrs := map[string]any{"cmd": "advertise", "hints": map[string]any{"udp": []string{}, "tcp": []string{}}}
				for _, p := range []*peer{owner, joiner} {
					for _, req := range []map[string]any{noAddrs, hint} {
						if _, e := p.try(req); e != "" {
							t.Skipf("%s cannot take part (%v: %s): it does not punch", p.name, req["cmd"], e)
						}
					}
					if _, e := p.try(map[string]any{"cmd": "paths", "peer": "x"}); strings.Contains(e, "unknown") {
						t.Skipf("%s does not tell its roads: %s", p.name, e)
					}
				}
				inv := owner.do(map[string]any{"cmd": "invite"})["invite"].(string)
				if got := joiner.do(map[string]any{"cmd": "join", "invite": inv})["peer"]; got != oa {
					t.Fatalf("joined %v, want %s", got, oa)
				}
				eventually(t, "the owner knows the joiner", func() bool {
					return owner.do(map[string]any{"cmd": "known", "peer": ja})["known"] == true
				})
				call(t, "over the relay", joiner, oa, owner, ja, []byte("over the relay"), 1)
				for _, s := range []struct {
					p  *peer
					to string
				}{{owner, ja}, {joiner, oa}} {
					if ps := pathsOf(t, s.p, s.to); slices.Contains(ps, "udp") || len(ps) == 0 {
						t.Fatalf("%s before the punch: %v, want the relay only", s.p.name, ps)
					}
				}

				// the relay answers observe now; telling the devices their relay again makes them ask
				pc, err := net.ListenPacket("udp", addr)
				if err != nil {
					t.Fatal(err)
				}
				t.Cleanup(func() { _ = pc.Close() })
				go func() { _ = relay.NewServer().ServeUDP(pc) }() // observe needs no channels: a server of its own
				began := time.Now()
				owner.do(hint)
				joiner.do(hint)
				onUDP := func(p *peer, to string) bool { return slices.Contains(pathsOf(t, p, to), "udp") }
				for deadline := began.Add(10 * time.Second); !(onUDP(owner, ja) && onUDP(joiner, oa)); time.Sleep(50 * time.Millisecond) {
					if time.Now().After(deadline) {
						t.Fatalf("not on UDP 10 s after observe: %s %v, %s %v",
							owner.name, pathsOf(t, owner, ja), joiner.name, pathsOf(t, joiner, oa))
					}
				}
				t.Logf("on UDP both ways %v after the relay began to answer observe", time.Since(began).Round(10*time.Millisecond))

				// messages go on, the long one too
				big := make([]byte, 300<<10)
				_, _ = rand.Read(big)
				call(t, "random 300 KB after the punch", joiner, oa, owner, ja, big, 2)
				for i := range 30 {
					call(t, "after the punch", joiner, oa, owner, ja, chatLine(i), i%3)
				}
				for _, s := range []struct {
					p  *peer
					to string
				}{{owner, ja}, {joiner, oa}} {
					if !onUDP(s.p, s.to) {
						t.Fatalf("%s left UDP: %v", s.p.name, pathsOf(t, s.p, s.to))
					}
				}
			})
		}
	}
}

package interop

import (
	"slices"
	"strings"
	"testing"
)

// TestInteropForward: two devices that cannot reach each other talk through a third one both
// trust (spec §3, via a third device). A and B pair over UDP, then restart on new ports and
// announce no addresses, so neither can find the other; both reach X, which opts in as the
// forwarder for the pair. Their session runs over the road via X, messages go both ways, and
// X forwarded packets it could not read. Every triple of implementations; a peer without the
// forward commands is skipped.
func TestInteropForward(t *testing.T) {
	list := impls(t)
	for _, ia := range list {
		for _, ix := range list {
			for _, ib := range list {
				t.Run(ia.name+"-via-"+ix.name+"-"+ib.name, func(t *testing.T) { forwardScenario(t, ia, ix, ib) })
			}
		}
	}
}

func forwardScenario(t *testing.T, ia, ix, ib impl) {
	da, db := t.TempDir(), t.TempDir()
	a, x, b := startIn(t, ia, ia.name+"-a", da), start(t, ix, ix.name+"-x"), startIn(t, ib, ib.name+"-b", db)
	for _, p := range []*peer{a, x, b} {
		if _, e := p.try(map[string]any{"cmd": "forward-stats"}); e != "" {
			t.Skipf("%s has no forward commands: %s", p.name, e)
		}
	}
	listen := func(p *peer) string {
		return p.do(map[string]any{"cmd": "listen", "udp": "127.0.0.1:0"})["addr"].(string)
	}
	aa, xa, ba := listen(a), listen(x), listen(b)
	join := func(owner, joiner *peer, ja string) {
		inv := owner.do(map[string]any{"cmd": "invite"})["invite"].(string)
		joiner.do(map[string]any{"cmd": "join", "invite": inv})
		eventually(t, owner.name+" knows "+joiner.name, func() bool {
			return owner.do(map[string]any{"cmd": "known", "peer": ja})["known"] == true
		})
	}
	join(a, b, ba)
	join(x, a, aa)
	join(x, b, ba)

	// A and B come back on other ports and tell nobody where: the addresses they know of each
	// other are dead, and the hellos (through X too) carry none
	a.stop()
	b.stop()
	a, b = startIn(t, ia, ia.name+"-a-again", da), startIn(t, ib, ib.name+"-b-again", db)
	for _, p := range []*peer{a, b} {
		p.do(map[string]any{"cmd": "hide"})
		listen(p)
	}
	x.do(map[string]any{"cmd": "forward", "a": aa, "b": ba, "on": true})

	stream := 0
	for name, data := range messages() {
		stream = (stream + 1) % 3
		call(t, name, a, ba, b, aa, data, stream)
	}
	for i := range 30 {
		call(t, "conversation", a, ba, b, aa, chatLine(i), i%3)
	}
	via := "via:" + xa
	for _, d := range []struct {
		p    *peer
		peer string
	}{{a, ba}, {b, aa}} {
		res := d.p.do(map[string]any{"cmd": "paths", "peer": d.peer})
		var paths []string
		for _, v := range res["paths"].([]any) {
			paths = append(paths, v.(string))
		}
		if !slices.Equal(paths, []string{via}) {
			t.Fatalf("%s's paths: %v, want [%s]", d.p.name, paths, via)
		}
	}
	st := x.do(map[string]any{"cmd": "forward-stats"})
	if st["forwarded"].(float64) == 0 {
		t.Fatalf("%s forwarded nothing: %v", x.name, st)
	}
	t.Logf("%s forwarded %v packets, dropped %v", x.name, st["forwarded"], st["dropped"])

	// the forwarder steps down: both are told, and X refuses to forward to a device it does
	// not trust
	x.do(map[string]any{"cmd": "forward", "a": aa, "b": ba, "on": false})
	if _, e := x.try(map[string]any{"cmd": "forward", "a": aa, "b": strings.Repeat("a", 26), "on": true}); e == "" {
		t.Fatalf("%s agreed to forward to a device it does not know", x.name)
	}
}

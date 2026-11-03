package interop

// Interop of the platform features (core spec §1 kind, §2 wants and local programs, §6 0.2/0.6,
// §8 named tools, §8.1 grants, §8.2 hosted tools): every pair of implementations, one after
// the other as owner and joiner (for the hosted tool: host and program, with a Go client).

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

const (
	kindComputer = 1
	kindClient   = 2
	kindPlugin   = 3

	notes = "com.example.notes"
	tasks = "com.example.tasks"
)

func eventually(t *testing.T, what string, f func() bool) {
	t.Helper()
	for deadline := time.Now().Add(20 * time.Second); ; time.Sleep(50 * time.Millisecond) {
		if f() {
			return
		}
		if time.Now().After(deadline) {
			t.Fatalf("timeout: %s", what)
		}
	}
}

func b64(s string) string { return base64.StdEncoding.EncodeToString([]byte(s)) }

// pairOver pairs joiner with owner over UDP (both listen); extra commands go to the peers first.
func pairOver(t *testing.T, owner, joiner *peer, join map[string]any) (oa, ja string) {
	t.Helper()
	oa = owner.do(map[string]any{"cmd": "listen", "udp": "127.0.0.1:0"})["addr"].(string)
	ja = joiner.do(map[string]any{"cmd": "listen", "udp": "127.0.0.1:0"})["addr"].(string)
	join["cmd"] = "join"
	join["invite"] = owner.do(map[string]any{"cmd": "invite"})["invite"]
	if got := joiner.do(join)["peer"]; got != oa {
		t.Fatalf("joined %v, want %s", got, oa)
	}
	eventually(t, "the owner knows the joiner", func() bool { return owner.do(map[string]any{"cmd": "known", "peer": ja})["known"] == true })
	return
}

func pairs(t *testing.T, name string, f func(t *testing.T, a, b impl)) {
	list := impls(t)
	for _, a := range list {
		for _, b := range list {
			t.Run(name+"/"+a.name+"-"+b.name, func(t *testing.T) { f(t, a, b) })
		}
	}
}

func kindOf(p *peer, addr string) int {
	return int(p.do(map[string]any{"cmd": "peer-info", "peer": addr})["kind"].(float64))
}

func echoOK(t *testing.T, from, _ *peer, toAddr string) {
	t.Helper()
	from.do(map[string]any{"cmd": "call", "peer": toAddr, "data": []byte("hi")})
}

// toolCall calls a tool by name: the data, or the error code (0 — no code) and text.
func toolCall(p *peer, peerAddr, name string, method int, data string) (got string, code int, errText string) {
	p.t.Helper()
	res, e := p.try(map[string]any{"cmd": "call-tool", "peer": peerAddr, "name": name, "method": method, "data": b64(data)})
	if e != "" {
		if c, ok := p.last["code"].(float64); ok {
			code = int(c)
		}
		return "", code, e
	}
	raw, _ := base64.StdEncoding.DecodeString(res["data"].(string))
	return string(raw), 0, ""
}

func mustCall(t *testing.T, p *peer, peerAddr, name string, method int, data, want string) {
	t.Helper()
	if got, code, e := toolCall(p, peerAddr, name, method, data); e != "" || got != want {
		t.Fatalf("%s calls %s.%d: got %q, error %d %q, want %q", p.name, name, method, got, code, e, want)
	}
}

func wantForbidden(t *testing.T, p *peer, peerAddr, name string, method int) {
	t.Helper()
	if got, code, e := toolCall(p, peerAddr, name, method, "x"); code != 4 {
		t.Fatalf("%s calls %s.%d: got %q, error %d %q, want error 4", p.name, name, method, got, code, e)
	}
}

func grantsOf(p *peer, peerAddr, which string) any {
	return p.do(map[string]any{"cmd": "grants", "peer": peerAddr})[which]
}

func waitGranted(t *testing.T, p *peer, peerAddr string, want []any) {
	t.Helper()
	eventually(t, fmt.Sprintf("%s sees the grants %v (event 0.6)", p.name, want), func() bool {
		return reflect.DeepEqual(grantsOf(p, peerAddr, "granted"), want)
	})
}

func gl(tool string, level int) []any { return []any{tool, float64(level)} }

func events(p *peer) []any { return p.do(map[string]any{"cmd": "events"})["events"].([]any) }

// TestInteropKind: the kind of a device is told at pairing and in the hello, and kept; a device that
// says nothing (an older one) still pairs and has kind 0.
func TestInteropKind(t *testing.T) {
	pairs(t, "pairing", func(t *testing.T, a, b impl) {
		owner, joiner := start(t, a, a.name+"-owner"), start(t, b, b.name+"-joiner")
		owner.do(map[string]any{"cmd": "kind", "kind": kindComputer})
		joiner.do(map[string]any{"cmd": "kind", "kind": kindClient})
		oa, ja := pairOver(t, owner, joiner, map[string]any{})
		if k := kindOf(owner, ja); k != kindClient {
			t.Fatalf("the owner sees kind %d of the joiner, want %d", k, kindClient)
		}
		if k := kindOf(joiner, oa); k != kindComputer {
			t.Fatalf("the joiner sees kind %d of the owner, want %d", k, kindComputer)
		}
		echoOK(t, joiner, owner, oa) // a session, and the hello in it, keep the kinds
		echoOK(t, owner, joiner, ja)
		time.Sleep(300 * time.Millisecond)
		if kindOf(owner, ja) != kindClient || kindOf(joiner, oa) != kindComputer {
			t.Fatal("the hello changed a kind")
		}
	})
	pairs(t, "old-style", func(t *testing.T, a, b impl) {
		owner, joiner := start(t, a, a.name+"-owner"), start(t, b, b.name+"-joiner")
		oa, ja := pairOver(t, owner, joiner, map[string]any{}) // nobody says a kind: no field on the wire
		call(t, "echo", joiner, oa, owner, ja, []byte("old"), 0)
		if kindOf(owner, ja) != 0 || kindOf(joiner, oa) != 0 {
			t.Fatal("a kind appeared from nowhere")
		}
	})
	pairs(t, "hello", func(t *testing.T, a, b impl) {
		dir := t.TempDir()
		owner, joiner := startIn(t, a, a.name+"-owner", dir), start(t, b, b.name+"-joiner")
		joiner.do(map[string]any{"cmd": "kind", "kind": kindClient})
		oa, ja := pairOver(t, owner, joiner, map[string]any{})
		if kindOf(joiner, oa) != 0 {
			t.Fatal("the owner said no kind")
		}
		// the owner comes back saying what it is: the joiner learns it from the hello
		owner.stop()
		owner = startIn(t, a, a.name+"-owner-again", dir)
		owner.do(map[string]any{"cmd": "kind", "kind": kindComputer})
		owner.do(map[string]any{"cmd": "listen", "udp": "127.0.0.1:0"})
		echoOK(t, owner, joiner, ja)
		eventually(t, "the joiner learns the owner's kind from the hello", func() bool { return kindOf(joiner, oa) == kindComputer })
	})
}

// TestInteropNamedTools: a named tool gets a number of its device's own; the other side resolves it
// by name (the hello's named map) and calls it; the same names registered in another order get
// other numbers.
func TestInteropNamedTools(t *testing.T) {
	pairs(t, "names", func(t *testing.T, a, b impl) {
		owner, joiner := start(t, a, a.name+"-owner"), start(t, b, b.name+"-joiner")
		num := func(p *peer, name string) int {
			return int(p.do(map[string]any{"cmd": "tool", "name": name})["number"].(float64))
		}
		an, at := num(owner, notes), num(owner, tasks)
		bt, bn := num(joiner, tasks), num(joiner, notes)
		if an == bn || at == bt {
			t.Fatalf("numbers do not differ: owner %d,%d joiner %d,%d", an, at, bn, bt)
		}
		oa, ja := pairOver(t, owner, joiner, map[string]any{})
		echoOK(t, joiner, owner, oa)
		for _, c := range []struct {
			p    *peer
			addr string
			name string
			want int
		}{{joiner, oa, notes, an}, {joiner, oa, tasks, at}, {owner, ja, notes, bn}, {owner, ja, tasks, bt}} {
			eventually(t, c.p.name+" resolves "+c.name, func() bool {
				info := c.p.do(map[string]any{"cmd": "peer-info", "peer": c.addr, "tool": c.name})
				return info["number"] == float64(c.want)
			})
			mustCall(t, c.p, c.addr, c.name, 1, "ping", "ping")
			mustCall(t, c.p, c.addr, c.name, 3, "x", c.name) // the tool itself answers with its name
		}
		tools := joiner.do(map[string]any{"cmd": "peer-info", "peer": oa})["tools"]
		if !reflect.DeepEqual(tools, []any{notes, tasks}) {
			t.Fatalf("the joiner sees the owner's tools %v", tools)
		}
		if _, code, _ := toolCall(joiner, oa, "com.example.none", 1, "x"); code != 1 {
			t.Fatalf("an unknown tool: error %d, want 1", code)
		}
	})
}

// TestInteropGrants: the owner lets the joiner read a tool: a read method works, a write method
// is error 4, the joiner sees the list (event 0.6), and the tool's events reach only a device
// with read.
func TestInteropGrants(t *testing.T) {
	pairs(t, "grants", func(t *testing.T, a, b impl) {
		owner, joiner := start(t, a, a.name+"-owner"), start(t, b, b.name+"-joiner")
		owner.do(map[string]any{"cmd": "tool", "name": notes})
		oa, ja := pairOver(t, owner, joiner, map[string]any{})
		eventually(t, "the joiner resolves the tool", func() bool {
			echoOK(t, joiner, owner, oa)
			return joiner.do(map[string]any{"cmd": "peer-info", "peer": oa, "tool": notes})["number"] != nil
		})
		mustCall(t, joiner, oa, notes, 2, "w", "w") // a device of the same person: everything
		read := []any{gl(notes, 1)}
		owner.do(map[string]any{"cmd": "set-grants", "peer": ja, "grants": read})
		waitGranted(t, joiner, oa, read)
		if got := grantsOf(owner, ja, "allowed"); !reflect.DeepEqual(got, read) {
			t.Fatalf("the owner allows %v", got)
		}
		mustCall(t, joiner, oa, notes, 1, "r", "r")
		mustCall(t, joiner, oa, notes, 3, "r", notes)
		wantForbidden(t, joiner, oa, notes, 2)
		// events: only with read
		emit := func(n string) {
			owner.do(map[string]any{"cmd": "emit", "peer": ja, "name": notes, "event": 1, "data": b64(n)})
		}
		emit("e1")
		eventually(t, "the event with read", func() bool { return len(events(joiner)) == 1 })
		ev := events(joiner)[0].(map[string]any)
		if ev["tool"] != notes || ev["event"] != float64(1) || ev["data"] != b64("e1") {
			t.Fatalf("event %v", ev)
		}
		none := []any{gl(tasks, 2)} // write on another tool: nothing on notes
		owner.do(map[string]any{"cmd": "set-grants", "peer": ja, "grants": none})
		waitGranted(t, joiner, oa, none)
		emit("e2") // dropped: no read
		wantForbidden(t, joiner, oa, notes, 1)
		owner.do(map[string]any{"cmd": "set-grants", "peer": ja, "grants": []any{gl(notes, 2)}})
		waitGranted(t, joiner, oa, []any{gl(notes, 2)})
		mustCall(t, joiner, oa, notes, 2, "w", "w")
		emit("e3")
		eventually(t, "the event after read again", func() bool { return len(events(joiner)) >= 2 })
		time.Sleep(200 * time.Millisecond)
		evs := events(joiner)
		if len(evs) != 2 || evs[1].(map[string]any)["data"] != b64("e3") {
			t.Fatalf("the joiner got events %v, want e1 and e3 only", evs)
		}
	})
}

// TestInteropWants: a joiner that asks for permissions at pairing gets them, narrowed by the
// owner's onWants when it narrows, and tells the result with event 0.6.
func TestInteropWants(t *testing.T) {
	pairs(t, "narrowed", func(t *testing.T, a, b impl) {
		owner, joiner := start(t, a, a.name+"-owner"), start(t, b, b.name+"-joiner")
		owner.do(map[string]any{"cmd": "tool", "name": notes})
		owner.do(map[string]any{"cmd": "on-wants", "level": 1})
		oa, ja := pairOver(t, owner, joiner, map[string]any{"wants": []any{gl(notes, 2)}})
		waitGranted(t, joiner, oa, []any{gl(notes, 1)})
		eventually(t, "the owner keeps the narrowed grants", func() bool {
			return reflect.DeepEqual(grantsOf(owner, ja, "allowed"), []any{gl(notes, 1)})
		})
		eventually(t, "the joiner resolves the tool", func() bool {
			return joiner.do(map[string]any{"cmd": "peer-info", "peer": oa, "tool": notes})["number"] != nil
		})
		mustCall(t, joiner, oa, notes, 1, "r", "r")
		wantForbidden(t, joiner, oa, notes, 2)
	})
	pairs(t, "as-asked", func(t *testing.T, a, b impl) {
		owner, joiner := start(t, a, a.name+"-owner"), start(t, b, b.name+"-joiner")
		owner.do(map[string]any{"cmd": "tool", "name": notes})
		oa, ja := pairOver(t, owner, joiner, map[string]any{"wants": []any{gl(notes, 1)}}) // no onWants: as asked
		waitGranted(t, joiner, oa, []any{gl(notes, 1)})
		if got := grantsOf(owner, ja, "allowed"); !reflect.DeepEqual(got, []any{gl(notes, 1)}) {
			t.Fatalf("the owner allows %v", got)
		}
		eventually(t, "the joiner resolves the tool", func() bool {
			return joiner.do(map[string]any{"cmd": "peer-info", "peer": oa, "tool": notes})["number"] != nil
		})
		mustCall(t, joiner, oa, notes, 3, "r", notes)
		wantForbidden(t, joiner, oa, notes, 2)
	})
}

// program starts a program of language b on host (language a's peer): the host listens on a local
// socket and makes the plugin environment, the program reaches it over the socket and pairs as a
// plugin. It returns the program, its address and the host's.
func program(t *testing.T, host *peer, b impl, grants []any) (prog *peer, pa, ha string) {
	t.Helper()
	dir, err := os.MkdirTemp("/tmp", "p2pl")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.RemoveAll(dir) })
	ha = host.do(map[string]any{"cmd": "listen", "udp": "127.0.0.1:0"})["addr"].(string)
	host.do(map[string]any{"cmd": "local", "listen": dir + "/s"})
	env := host.do(map[string]any{"cmd": "plugin-env", "grants": grants})
	prog = start(t, b, b.name+"-plugin")
	prog.do(map[string]any{"cmd": "kind", "kind": kindPlugin})
	prog.do(map[string]any{"cmd": "local", "use": env["socket"]})
	pa = prog.do(map[string]any{"cmd": "listen"})["addr"].(string)
	if got := prog.do(map[string]any{"cmd": "join", "invite": env["invite"]})["peer"]; got != ha {
		t.Fatalf("the program joined %v, want %s", got, ha)
	}
	eventually(t, "the host knows the program", func() bool { return host.do(map[string]any{"cmd": "known", "peer": pa})["known"] == true })
	return
}

// TestInteropLocal: a program of one language pairs with a device of another over the local socket
// (the environment the device makes), as kind plugin, and they call each other's named tools.
func TestInteropLocal(t *testing.T) {
	pairs(t, "plugin", func(t *testing.T, a, b impl) {
		host := start(t, a, a.name+"-host")
		host.do(map[string]any{"cmd": "tool", "name": tasks})
		prog, pa, ha := program(t, host, b, []any{gl("*", 2)})
		prog.do(map[string]any{"cmd": "tool", "name": notes})
		if k := kindOf(host, pa); k != kindPlugin {
			t.Fatalf("the host sees kind %d of the program, want %d", k, kindPlugin)
		}
		eventually(t, "the host resolves the program's tool", func() bool {
			return host.do(map[string]any{"cmd": "peer-info", "peer": pa, "tool": notes})["number"] != nil
		})
		eventually(t, "the program resolves the host's tool", func() bool {
			return prog.do(map[string]any{"cmd": "peer-info", "peer": ha, "tool": tasks})["number"] != nil
		})
		mustCall(t, host, pa, notes, 3, "x", notes)
		mustCall(t, prog, ha, tasks, 3, "x", tasks)
		mustCall(t, prog, ha, tasks, 2, "w", "w")
	})
}

// TestInteropHosted: a program (language b) behind a host (language a) serves a named tool through
// the host's tool-wide handler (spec §8.2); a third device (Go) calls the host and is answered by
// the program, with the grants the host gave it: write methods are error 4 with read only.
func TestInteropHosted(t *testing.T) {
	pairs(t, "hosted", func(t *testing.T, a, b impl) {
		list := impls(t)
		host := start(t, a, a.name+"-host")
		prog, pa, ha := program(t, host, b, []any{gl("*", 2)})
		prog.do(map[string]any{"cmd": "tool", "name": notes})
		eventually(t, "the host resolves the program's tool", func() bool {
			return host.do(map[string]any{"cmd": "peer-info", "peer": pa, "tool": notes})["number"] != nil
		})
		host.do(map[string]any{"cmd": "proxy", "name": notes, "plugin": pa})
		client := start(t, list[0], "go-client")
		ca := client.do(map[string]any{"cmd": "listen", "udp": "127.0.0.1:0"})["addr"].(string)
		client.do(map[string]any{"cmd": "join", "invite": host.do(map[string]any{"cmd": "invite"})["invite"]})
		eventually(t, "the host knows the client", func() bool { return host.do(map[string]any{"cmd": "known", "peer": ca})["known"] == true })
		host.do(map[string]any{"cmd": "set-grants", "peer": ca, "grants": []any{gl(notes, 1)}})
		eventually(t, "the client resolves the hosted tool", func() bool {
			return client.do(map[string]any{"cmd": "peer-info", "peer": ha, "tool": notes})["number"] != nil
		})
		mustCall(t, client, ha, notes, 1, "ping", "ping")
		mustCall(t, client, ha, notes, 3, "x", notes) // answered by the program
		wantForbidden(t, client, ha, notes, 2)
		host.do(map[string]any{"cmd": "set-grants", "peer": ca, "grants": []any{gl(notes, 2)}})
		mustCall(t, client, ha, notes, 2, "w", "w")
	})
}

var _ = json.Marshal

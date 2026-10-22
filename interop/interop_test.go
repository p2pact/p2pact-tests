// Package interop checks that implementations of P2PACT understand each other: every pair
// (including an implementation with itself) pairs by invite and exchanges messages both ways,
// over UDP and over TCP, and over an application road that stands in for Bluetooth LE; a call
// answered later does not hold up the calls behind it on its stream.
//
// Implementations are interop peers (see p2pact-go/cmd/p2pact-interop) driven by JSON lines.
// The Go one is always built; others are added through environment variables:
//
//	P2PACT_TS=/path/to/p2pact-ts/dist/src/interop.js   (run with node)
//	P2PACT_PY=/path/to/python                            (with p2pact-py installed: python -m p2pact.interop)
//	P2PACT_KOTLIN=/path/to/p2pact-kotlin/build/install/p2pact-interop/bin/p2pact-interop
package interop

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

type impl struct {
	name string
	cmd  []string
}

func impls(t *testing.T) []impl {
	bin := filepath.Join(t.TempDir(), "go-peer")
	out, err := exec.Command("go", "build", "-o", bin, "github.com/p2pact/p2pact-go/cmd/p2pact-interop").CombinedOutput()
	if err != nil {
		t.Fatalf("build the Go peer: %v\n%s", err, out)
	}
	list := []impl{{"go", []string{bin}}}
	if js := os.Getenv("P2PACT_TS"); js != "" {
		list = append(list, impl{"ts", []string{"node", js}})
	}
	if py := os.Getenv("P2PACT_PY"); py != "" {
		list = append(list, impl{"py", []string{py, "-m", "p2pact.interop"}})
	}
	if kt := os.Getenv("P2PACT_KOTLIN"); kt != "" {
		list = append(list, impl{"kotlin", []string{kt}})
	}
	return list
}

type peer struct {
	t    *testing.T
	name string
	cmd  *exec.Cmd
	in   *json.Encoder
	out  *bufio.Scanner
}

func start(t *testing.T, im impl, name string) *peer {
	c := exec.Command(im.cmd[0], append(im.cmd[1:], t.TempDir(), name)...)
	c.Stderr = os.Stderr
	w, _ := c.StdinPipe()
	r, _ := c.StdoutPipe()
	if err := c.Start(); err != nil {
		t.Fatal(err)
	}
	p := &peer{t: t, name: name, cmd: c, in: json.NewEncoder(w), out: bufio.NewScanner(r)}
	p.out.Buffer(make([]byte, 1<<20), 64<<20)
	t.Cleanup(func() {
		_ = p.in.Encode(map[string]string{"cmd": "close"})
		done := make(chan struct{})
		go func() { _ = c.Wait(); close(done) }()
		select {
		case <-done:
		case <-time.After(3 * time.Second):
			_ = c.Process.Kill()
		}
	})
	return p
}

func (p *peer) do(req map[string]any) map[string]any {
	p.t.Helper()
	res, e := p.try(req)
	if e != "" {
		p.t.Fatalf("%s %v: %v", p.name, req["cmd"], e)
	}
	return res
}

// try is do that returns an error answer instead of failing the test.
func (p *peer) try(req map[string]any) (map[string]any, string) {
	p.t.Helper()
	if err := p.in.Encode(req); err != nil {
		p.t.Fatalf("%s: %v", p.name, err)
	}
	line := make(chan []byte, 1)
	go func() {
		if p.out.Scan() {
			line <- append([]byte(nil), p.out.Bytes()...)
		}
		close(line)
	}()
	select {
	case b, ok := <-line:
		if !ok {
			p.t.Fatalf("%s exited", p.name)
		}
		var res map[string]any
		if err := json.Unmarshal(b, &res); err != nil {
			p.t.Fatalf("%s: %s", p.name, b)
		}
		if e, ok := res["error"]; ok {
			return nil, fmt.Sprint(e)
		}
		return res, ""
	case <-time.After(90 * time.Second):
		p.t.Fatalf("%s %v: no answer", p.name, req["cmd"])
	}
	return nil, ""
}

// ownDictInUse waits until p has trained its own dictionary for the device at addr and uses it:
// that device took it (fetched it with 0.5 and announced it back). A peer without the
// dict-state command is skipped.
func ownDictInUse(t *testing.T, p *peer, addr string) {
	t.Helper()
	req := map[string]any{"cmd": "dict-state", "peer": addr}
	for deadline := time.Now().Add(60 * time.Second); ; time.Sleep(100 * time.Millisecond) {
		res, e := p.try(req)
		if strings.Contains(e, "unknown command") {
			t.Logf("%s does not report its dictionaries: not checked", p.name)
			return
		}
		if e != "" {
			t.Fatalf("%s dict-state: %s", p.name, e)
		}
		if res["own"] != nil && res["using"] == res["own"] && res["training"] == false {
			t.Logf("%s compresses with its own dictionary %.12s…", p.name, res["own"])
			return
		}
		if time.Now().After(deadline) {
			t.Fatalf("%s does not compress with its own dictionary: %v", p.name, res)
		}
	}
}

func messages() map[string][]byte {
	r := rand.New(rand.NewSource(1))
	random := make([]byte, 300<<10)
	r.Read(random)
	return map[string][]byte{
		"short":         []byte("run the tests"),
		"compressible":  []byte(strings.Repeat("the agent finished the task and pushed a commit. ", 200)),
		"random 300 KB": random,
		"russian":       []byte(strings.Repeat("агент закончил задачу, тесты прошли. ", 50)),
	}
}

// call sends data both ways (joiner → owner, owner → joiner) and checks the echo.
func call(t *testing.T, name string, joiner *peer, oa string, owner *peer, ja string, data []byte, stream int) {
	t.Helper()
	for _, dir := range []struct {
		from *peer
		to   string
	}{{joiner, oa}, {owner, ja}} {
		res := dir.from.do(map[string]any{"cmd": "call", "peer": dir.to, "data": data, "stream": stream})
		var got []byte
		b, _ := json.Marshal(res["data"])
		_ = json.Unmarshal(b, &got)
		if !bytes.Equal(got, data) {
			t.Fatalf("%s from %s: got %d bytes back, want %d", name, dir.from.name, len(got), len(data))
		}
	}
}

// answerLater checks a call answered later (spec §6), both ways: the callee holds a call until
// a second call, sent behind it on the same stream, has run. A peer without the later command
// is skipped.
func answerLater(t *testing.T, joiner *peer, oa string, owner *peer, ja string, stream int) {
	t.Helper()
	for _, dir := range []struct {
		from *peer
		to   string
	}{{joiner, oa}, {owner, ja}} {
		data := []byte("held by " + dir.from.name)
		res, e := dir.from.try(map[string]any{"cmd": "later", "peer": dir.to, "data": data, "stream": stream})
		if strings.Contains(e, "unknown command") {
			t.Logf("%s does not call to be answered later: not checked", dir.from.name)
			continue
		}
		if e != "" {
			t.Fatalf("answered later, from %s: %s", dir.from.name, e)
		}
		var got []byte
		b, _ := json.Marshal(res["data"])
		_ = json.Unmarshal(b, &got)
		if !bytes.Equal(got, data) {
			t.Fatalf("answered later, from %s: got %q", dir.from.name, got)
		}
	}
}

func chatLine(i int) []byte {
	words := []string{"the agent", "finished", "the task", "run the tests", "and commit", "сообщения не доходят",
		"до телефона", "why is it so slow", "checked the logs", "a delivery bug", "fixed it", "pushed to the branch"}
	var b strings.Builder
	fmt.Fprintf(&b, "step %d: ", i)
	for j := range 12 {
		b.WriteString(words[(i*7+j*j)%len(words)])
		b.WriteString(", ")
	}
	return []byte(b.String())
}

func TestInterop(t *testing.T) {
	list := impls(t)
	if len(list) < 2 {
		t.Log("only the Go peer: set P2PACT_TS to test against TypeScript")
	}
	for _, road := range []string{"udp", "tcp"} {
		for _, a := range list {
			for _, b := range list {
				t.Run(fmt.Sprintf("%s/%s-%s", road, a.name, b.name), func(t *testing.T) {
					owner, joiner := start(t, a, a.name+"-owner"), start(t, b, b.name+"-joiner")
					listen := map[string]any{"cmd": "listen"}
					dial := map[string]any{"cmd": "listen"}
					if road == "udp" {
						listen["udp"], dial["udp"] = "127.0.0.1:0", "127.0.0.1:0"
					} else {
						listen["tcp"], dial["tcp"] = "127.0.0.1:0", ""
					}
					oa := owner.do(listen)["addr"].(string)
					ja := joiner.do(dial)["addr"].(string)
					inv := owner.do(map[string]any{"cmd": "invite"})["invite"].(string)
					if got := joiner.do(map[string]any{"cmd": "join", "invite": inv})["peer"]; got != oa {
						t.Fatalf("joined %v, want %s", got, oa)
					}
					for i := 0; ; i++ {
						if owner.do(map[string]any{"cmd": "known", "peer": ja})["known"] == true {
							break
						}
						if i > 100 {
							t.Fatal("the owner does not know the joiner")
						}
						time.Sleep(20 * time.Millisecond)
					}
					stream := 0
					for name, data := range messages() {
						stream = (stream + 1) % 3 // stream 0 (in order) and streams 1, 2
						call(t, name, joiner, oa, owner, ja, data, stream)
					}
					answerLater(t, joiner, oa, owner, ja, 2)
					// a conversation long enough for a device to train its own dictionary (a Go
					// one does at 32 KB); talking goes on after it, with it where the other side
					// takes it
					for i := range 400 {
						call(t, "conversation", joiner, oa, owner, ja, chatLine(i), i%3)
					}
					for name, data := range messages() {
						call(t, name+" after training", joiner, oa, owner, ja, data, 0)
					}
					// both trained one for the other and the other took it: talk with them
					ownDictInUse(t, owner, ja)
					ownDictInUse(t, joiner, oa)
					for i := 400; i < 450; i++ {
						call(t, "conversation with own dictionaries", joiner, oa, owner, ja, chatLine(i), i%3)
					}
				})
			}
		}
	}
}

// TestInteropBluetooth: every pair whose peers have the road command talks over an application
// road (spec §3) and nothing else — a local TCP connection standing in for a Bluetooth LE GATT
// link: writes of 20 bytes, 40 KB/s, framed as a stream road. They pair over it (the invite has
// no hints; the road is bound by the handshake) and exchange messages both ways.
func TestInteropBluetooth(t *testing.T) {
	list := impls(t)
	for _, a := range list {
		for _, b := range list {
			t.Run(a.name+"-"+b.name, func(t *testing.T) {
				owner, joiner := start(t, a, a.name+"-owner"), start(t, b, b.name+"-joiner")
				oa := owner.do(map[string]any{"cmd": "listen"})["addr"].(string)
				ja := joiner.do(map[string]any{"cmd": "listen"})["addr"].(string)
				link := map[string]any{"cmd": "road", "mtu": 20, "rate": 40 << 10}
				link["listen"] = "127.0.0.1:0"
				res, e := owner.try(link)
				if strings.Contains(e, "unknown command") {
					t.Skipf("%s has no application roads", a.name)
				} else if e != "" {
					t.Fatalf("%s road: %s", a.name, e)
				}
				delete(link, "listen")
				link["dial"] = res["road"]
				if _, e := joiner.try(link); strings.Contains(e, "unknown command") {
					t.Skipf("%s has no application roads", b.name)
				} else if e != "" {
					t.Fatalf("%s road: %s", b.name, e)
				}
				inv := owner.do(map[string]any{"cmd": "invite"})["invite"].(string)
				if got := joiner.do(map[string]any{"cmd": "join", "invite": inv})["peer"]; got != oa {
					t.Fatalf("joined %v, want %s", got, oa)
				}
				for i := 0; owner.do(map[string]any{"cmd": "known", "peer": ja})["known"] != true; i++ {
					if i > 200 {
						t.Fatal("the owner does not know the joiner")
					}
					time.Sleep(20 * time.Millisecond)
				}
				msgs := messages()
				msgs["random 30 KB"] = msgs["random 300 KB"][:30<<10] // 40 KB/s: a shorter long message
				delete(msgs, "random 300 KB")
				stream := 0
				for name, data := range msgs {
					stream = (stream + 1) % 3
					call(t, name, joiner, oa, owner, ja, data, stream)
				}
				answerLater(t, joiner, oa, owner, ja, 2)
				for i := range 50 {
					call(t, "conversation", joiner, oa, owner, ja, chatLine(i), i%3)
				}
			})
		}
	}
}

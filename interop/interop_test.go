// Package interop checks that implementations of P2PACT understand each other: every pair
// (including an implementation with itself) pairs by invite and exchanges messages both ways,
// over UDP and over TCP.
//
// Implementations are interop peers (see p2pact-go/cmd/p2pact-interop) driven by JSON lines.
// The Go one is always built; others are added through environment variables:
//
//	P2PACT_TS=/path/to/p2pact-ts/dist/src/interop.js   (run with node)
//	P2PACT_PY=/path/to/python                            (with p2pact-py installed: python -m p2pact.interop)
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
			p.t.Fatalf("%s %v: %v", p.name, req["cmd"], e)
		}
		return res
	case <-time.After(90 * time.Second):
		p.t.Fatalf("%s %v: no answer", p.name, req["cmd"])
	}
	return nil
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
				})
			}
		}
	}
}

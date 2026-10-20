package bench

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/fxamacker/cbor/v2"
	p2pact "github.com/p2pact/p2pact-go"
)

// RunP2PACT measures a chat message as a P2PACT call [0, id, 1 chat, 2 send, chat, text] → [1, id]
// over one road ("udp" or "tcp") with the profile emulated on both sides.
func RunP2PACT(sc *Scenario, p Profile, road string) (*Result, error) {
	res := newResult("P2PACT "+map[string]string{"udp": "UDP", "tcp": "TCP"}[road], p.Name)
	dirA, _ := os.MkdirTemp("", "p2pact-bench-a")
	dirB, _ := os.MkdirTemp("", "p2pact-bench-b")
	defer os.RemoveAll(dirA)
	defer os.RemoveAll(dirB)

	start := func(dir, name, addr string, listen bool) (*p2pact.Node, error) {
		nd, err := p2pact.NewNode(dir, name)
		if err != nil {
			return nil, err
		}
		if road == "udp" {
			err = nd.ListenUDP(addr)
		} else {
			if !listen {
				addr = ""
			}
			err = nd.ListenTCP(addr)
		}
		if err != nil {
			return nil, err
		}
		nd.NoCompression = !*sc.Compression
		nd.Handle(1, 2, func(string, []cbor.RawMessage) ([]any, error) { return []any{}, nil })
		nd.Start()
		return nd, nil
	}
	local := func(nd *p2pact.Node) string {
		if road == "udp" {
			return nd.UDP().LocalAddr()
		}
		return nd.TCP().LocalAddr()
	}
	stats := func(nd *p2pact.Node) p2pact.RoadStats {
		if road == "udp" {
			return nd.UDP().Stats()
		}
		return nd.TCP().Stats()
	}
	var rngMu sync.Mutex
	rng := rand.New(rand.NewSource(sc.Seed))
	emulate := func(nd *p2pact.Node) {
		if road == "tcp" {
			t := nd.TCP()
			t.Delay, t.Rate = p.delay(), p.rate()
			return
		}
		u := nd.UDP()
		u.Delay, u.Rate = p.delay(), p.rate()
		if p.loss() > 0 {
			// Drop runs on both send and receive, so each side gets half the loss
			u.Drop = func() bool { rngMu.Lock(); defer rngMu.Unlock(); return rng.Float64() < p.loss()/2 }
		}
	}

	srv, err := start(dirA, "Laptop", "127.0.0.1:0", true)
	if err != nil {
		return nil, err
	}
	srvAddr := local(srv)
	cli, err := start(dirB, "Phone", "127.0.0.1:0", false)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	inv, _ := p2pact.ParseInvite(srv.Invite(time.Minute).String())
	peer, err := cli.Join(ctx, inv)
	if err != nil {
		return nil, err
	}
	for i := 0; i < 200 && srv.Trust.Get(cli.ID.Addr()) == nil; i++ {
		time.Sleep(10 * time.Millisecond)
	}
	// cold: both restart with the same keys — paired, but no session
	srv.Close()
	cli.Close()
	if srv, err = start(dirA, "Laptop", srvAddr, true); err != nil {
		return nil, err
	}
	defer srv.Close()
	if cli, err = start(dirB, "Phone", "127.0.0.1:0", false); err != nil {
		return nil, err
	}
	defer cli.Close()
	emulate(srv)
	emulate(cli)
	// up — client → server, down — server → client
	snap := func() (up, down, pkts int64) {
		c, s := stats(cli), stats(srv)
		return c.SentBytes, s.SentBytes, c.SentPackets + s.SentPackets
	}
	wire := func(up, down, pkts int64) int64 {
		if road == "udp" {
			return up + down + 28*pkts
		}
		return tcpWire(up, down)
	}
	send := func(text string) error {
		_, err := cli.Call(ctx, peer.Addr, 1, 2, 1, text)
		return err
	}
	settle := func() { time.Sleep(3*p.delay() + 100*time.Millisecond) } // let the last ack land

	u0, d0, p0 := snap()
	t := time.Now()
	if err := send(sc.message(sc.Sizes[0], 1)); err != nil {
		return nil, err
	}
	res.ColdMS = ms(time.Since(t))
	settle()
	u1, d1, p1 := snap()
	res.ColdBytes = wire(u1-u0, d1-d0, p1-p0)
	if road == "tcp" {
		res.ColdBytes += tcpHandshake
	}

	for _, s := range sc.Sizes {
		var rtts []time.Duration
		cnt := sc.count(s)
		u0, d0, p0 := snap()
		for i := 0; i < cnt; i++ {
			t := time.Now()
			if err := send(sc.message(s, i)); err != nil {
				return nil, err
			}
			rtts = append(rtts, time.Since(t))
		}
		settle()
		u1, d1, p1 := snap()
		c := int64(cnt)
		res.Payload[s.Name] = (u1 - u0 + d1 - d0) / c
		res.Wire[s.Name] = wire((u1-u0)/c, (d1-d0)/c, (p1-p0)/c)
		res.RTTms[s.Name] = ms(median(rtts))
	}
	if road == "tcp" && p.loss() > 0 {
		res.Note = "no packet loss emulated for TCP (it would need root); real numbers would be worse"
	}
	return res, nil
}

// Run measures the scenario's protocols over its profiles.
func Run(sc *Scenario, progress func(string)) ([]*Result, error) {
	var out []*Result
	for _, p := range sc.Profiles {
		for _, proto := range sc.Protocols {
			var r *Result
			var err error
			switch proto {
			case "p2pact-udp":
				r, err = RunP2PACT(sc, p, "udp")
			case "p2pact-tcp":
				r, err = RunP2PACT(sc, p, "tcp")
			case "websocket":
				r, err = RunWebSocket(sc, p)
			case "https":
				r, err = RunHTTPS(sc, p)
			case "http2":
				r, err = RunHTTP2(sc, p)
			case "grpc":
				r, err = RunGRPC(sc, p)
			case "http3":
				r, err = RunHTTP3(sc, p)
			default:
				err = fmt.Errorf("unknown protocol %q (want one of %v)", proto, AllProtocols)
			}
			if err != nil {
				return nil, fmt.Errorf("%s · %s: %w", proto, p.Name, err)
			}
			if progress != nil {
				progress(r.Protocol + " · " + r.Profile)
			}
			out = append(out, r)
		}
	}
	return out, nil
}

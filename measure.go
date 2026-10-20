package bench

import (
	"errors"
	"sort"

	p2pact "github.com/p2pact/p2pact-go"

	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

// Load is the throughput test: many messages of one size, several in flight at once (like
// many chats), on a live connection.
type Load struct {
	Messages    int    `json:"messages"`
	Concurrency int    `json:"concurrency"`
	Size        string `json:"size"` // name of a size bin
}

// LoadResult is how a protocol coped with the load.
type LoadResult struct {
	Messages    int     `json:"messages"`
	Concurrency int     `json:"concurrency"`
	MS          float64 `json:"ms"`
	PerSecond   float64 `json:"per_second"`
	WireBytes   int64   `json:"wire_bytes"`
	Errors      int     `json:"errors"`
}

// wire is bytes on the wire for traffic of up/down bytes in packets datagrams (UDP: exact
// headers; TCP: estimated per message by tcpWire).
func (d *driver) wire(up, down, packets int64, msgs int64) int64 {
	if d.udp {
		return up + down + 28*packets
	}
	if msgs <= 0 {
		msgs = 1
	}
	return tcpWire(up/msgs, down/msgs) * msgs
}

func settle(p Profile) { time.Sleep(3*p.delay() + 150*time.Millisecond) }

// measure runs one protocol over one link: a cold start, every size, then the load.
func measure(sc *Scenario, p Profile, d *driver) (*Result, error) {
	res := newResult(d.name, p.Name)
	// cold: a new connection (P2PACT: a new session between paired devices) and one message;
	// five times, the median — one lost first packet on the internet must not decide it
	var colds []time.Duration
	var coldBytes []int64
	for i := 0; i < 5; i++ {
		t := time.Now() // the connection itself counts: some protocols connect in dial
		c, err := d.dial()
		if err != nil {
			return nil, err
		}
		if err := c.send(sc.message(sc.Sizes[0], i)); err != nil {
			c.close()
			return nil, fmt.Errorf("first message: %w", err)
		}
		colds = append(colds, time.Since(t))
		settle(p)
		up, down, pk := c.bytes()
		b := d.wire(up, down, pk, 1)
		if !d.udp {
			b += tcpHandshake
		}
		coldBytes = append(coldBytes, b)
		c.close()
	}
	res.ColdMS = ms(median(colds))
	sort.Slice(coldBytes, func(i, j int) bool { return coldBytes[i] < coldBytes[j] })
	res.ColdBytes = coldBytes[len(coldBytes)/2]

	c, err := d.dial()
	if err != nil {
		return nil, err
	}
	defer c.close()
	if err := c.send("warm up"); err != nil {
		return nil, err
	}
	for _, s := range sc.Sizes {
		settle(p)
		var rtts []time.Duration
		cnt := sc.count(s)
		u0, d0, p0 := c.bytes()
		for i := 0; i < cnt; i++ {
			t := time.Now()
			err := c.send(sc.message(s, i))
			if err != nil {
				res.Errors++
				if res.Errors > 20 {
					return nil, fmt.Errorf("%s: too many errors, last: %w", s.Name, err)
				}
				continue
			}
			rtts = append(rtts, time.Since(t))
		}
		settle(p)
		u1, d1, p1 := c.bytes()
		n := int64(cnt)
		res.Payload[s.Name] = (u1 - u0 + d1 - d0) / n
		res.Wire[s.Name] = d.wire(u1-u0, d1-d0, p1-p0, n) / n
		res.RTTms[s.Name] = ms(median(rtts))
		res.Sent += cnt
	}
	if sc.Load != nil {
		lr, err := load(sc, p, d, c)
		if err != nil {
			return nil, err
		}
		res.Load = lr
		res.Errors += lr.Errors
	}
	return res, nil
}

func load(sc *Scenario, p Profile, d *driver, c client) (*LoadResult, error) {
	var size *Size
	for i := range sc.Sizes {
		if sc.Sizes[i].Name == sc.Load.Size {
			size = &sc.Sizes[i]
		}
	}
	if size == nil {
		size = &sc.Sizes[0]
	}
	settle(p)
	lr := &LoadResult{Messages: sc.Load.Messages, Concurrency: max(1, sc.Load.Concurrency)}
	u0, d0, p0 := c.bytes()
	var next, errs atomic.Int64
	var wg sync.WaitGroup
	t := time.Now()
	for w := 0; w < lr.Concurrency; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				i := next.Add(1) - 1
				if i >= int64(lr.Messages) {
					return
				}
				if err := c.send(sc.message(*size, int(i))); err != nil {
					errs.Add(1)
				}
			}
		}()
	}
	wg.Wait()
	took := time.Since(t)
	settle(p)
	u1, d1, p1 := c.bytes()
	lr.MS = ms(took)
	lr.PerSecond = float64(lr.Messages) / took.Seconds()
	lr.WireBytes = d.wire(u1-u0, d1-d0, p1-p0, int64(lr.Messages))
	lr.Errors = int(errs.Load())
	return lr, nil
}

// Run measures the scenario's protocols over its profiles against the target.
func Run(sc *Scenario, t *Target, log io.Writer) ([]*Result, error) {
	inv, err := t.Invite()
	if err != nil {
		return nil, fmt.Errorf("server invite: %w", err)
	}
	setup, err := newP2PACTSetup(inv)
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(setup.dir)
	if t.Servers != nil { // forget the address the client paired from: only the link under test reaches it
		if p := t.Servers.Node.Trust.Get(setup.client); p != nil {
			p.Hints = p2pact.Hints{}
			t.Servers.Node.Trust.Put(*p)
		}
	}
	var out []*Result
	for _, p := range sc.Profiles {
		for _, proto := range sc.Protocols {
			d, err := makeDriver(sc, p, t, setup, proto)
			if err != nil {
				return nil, err
			}
			r, err := measure(sc, p, d)
			d.stop()
			if err != nil {
				return nil, errors.Join(fmt.Errorf("%s · %s", proto, p.Name), err)
			}
			fmt.Fprintf(log, "done: %s · %s (errors: %d)\n", r.Protocol, r.Profile, r.Errors)
			out = append(out, r)
		}
	}
	return out, nil
}

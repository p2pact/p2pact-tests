// Package bench compares P2PACT with WebSocket (JSON over TLS) and HTTPS+JSON: bytes on the
// wire and time for chat messages of several sizes over emulated network profiles.
//
// A run is described by a scenario (scenarios/*.json): profiles, size bins, protocols, where
// the texts come from, repetitions and a seed. Change any field to get a new experiment.
// Everything runs on one machine with emulated links, so results are reproducible; bytes are
// counted on the sockets, not computed from formulas.
package bench

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"
	"unicode/utf8"
)

// Scenario is one experiment.
type Scenario struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	// Data is where texts come from: "synthetic", "oasst2:human", "oasst2:assistant",
	// "oasst2:human:ru" (any language code), "swe-gym:agent", "swe-gym:tool", or
	// "file:<path>" — a JSON array of strings.
	Data        string    `json:"data"`
	Seed        int64     `json:"seed"`
	Repeat      int       `json:"repeat"`       // messages per size
	RepeatLarge int       `json:"repeat_large"` // for sizes of 10 000+ characters
	Compression *bool     `json:"compression"`  // P2PACT zstd; default true
	Sizes       []Size    `json:"sizes"`
	Profiles    []Profile `json:"profiles"`
	Load        *Load     `json:"load,omitempty"`
	Protocols   []string  `json:"protocols"` // p2pact-udp, p2pact-tcp, websocket, https, http2, grpc, http3
}

// Profile is an emulated link: one-way latency, bandwidth per direction and packet loss.
type Profile struct {
	Name     string  `json:"name"`
	DelayMS  float64 `json:"delay_ms"`
	RateKBps int     `json:"rate_kbps"`    // kilobytes per second per direction, 0 — unlimited
	LossPct  float64 `json:"loss_percent"` // P2PACT over UDP only: TCP loss needs root to emulate
}

func (p Profile) delay() time.Duration { return time.Duration(p.DelayMS * float64(time.Millisecond)) }
func (p Profile) rate() int            { return p.RateKBps * 1024 }
func (p Profile) loss() float64        { return p.LossPct / 100 }

// Size is a bin of message lengths in characters, [Min, Max). Chars is the generated length
// for synthetic data; with real data the bin's messages are used and Chars becomes their
// median.
type Size struct {
	Name  string `json:"name"`
	Min   int    `json:"min"`
	Max   int    `json:"max"`
	Chars int    `json:"chars"`

	msgs []string
}

// DefaultSizes are the bins used when a scenario has none.
var DefaultSizes = []Size{
	{Name: "very short", Min: 1, Max: 40, Chars: 20},
	{Name: "medium", Min: 40, Max: 200, Chars: 100},
	{Name: "long", Min: 200, Max: 2000, Chars: 600},
	{Name: "very long", Min: 2000, Max: 20000, Chars: 5000},
	{Name: "huge", Min: 20000, Max: 1 << 30, Chars: 200000},
}

// DefaultProfiles are the network conditions used when a scenario has none.
var DefaultProfiles = []Profile{
	{Name: "local network", DelayMS: 1},
	{Name: "bluetooth", DelayMS: 30, RateKBps: 40},
	{Name: "internet", DelayMS: 40, LossPct: 1},
	{Name: "mobile", DelayMS: 100, RateKBps: 500, LossPct: 3},
}

// AllProtocols in report order.
var AllProtocols = []string{"p2pact-udp", "p2pact-tcp", "websocket", "https", "http2", "grpc", "http3"}

// LoadScenario reads a scenario file and fills defaults.
func LoadScenario(path string) (*Scenario, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var s Scenario
	if err := json.Unmarshal(b, &s); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	s.defaults()
	return &s, nil
}

func (s *Scenario) defaults() {
	if s.Name == "" {
		s.Name = "scenario"
	}
	if s.Data == "" {
		s.Data = "synthetic"
	}
	if s.Seed == 0 {
		s.Seed = 1
	}
	if s.Repeat <= 0 {
		s.Repeat = 20
	}
	if s.RepeatLarge <= 0 {
		s.RepeatLarge = 10
	}
	if s.Compression == nil {
		t := true
		s.Compression = &t
	}
	if len(s.Sizes) == 0 {
		s.Sizes = append([]Size(nil), DefaultSizes...)
	}
	if len(s.Profiles) == 0 {
		s.Profiles = append([]Profile(nil), DefaultProfiles...)
	}
	if len(s.Protocols) == 0 {
		s.Protocols = AllProtocols
	}
}

// count is how many messages of a size to send.
func (s *Scenario) count(sz Size) int {
	if sz.Chars >= 10000 {
		return s.RepeatLarge
	}
	return s.Repeat
}

// message is the i-th message of a size.
func (s *Scenario) message(sz Size, i int) string {
	if len(sz.msgs) > 0 {
		return sz.msgs[i%len(sz.msgs)]
	}
	return Text(sz.Chars, s.Seed*1000003+int64(i))
}

// fill sorts real messages into the size bins (a random but seeded order); a bin without
// messages is dropped from the run.
func (s *Scenario) fill(msgs []string) {
	r := rand.New(rand.NewSource(s.Seed))
	r.Shuffle(len(msgs), func(i, j int) { msgs[i], msgs[j] = msgs[j], msgs[i] })
	var kept []Size
	for _, sz := range s.Sizes {
		sz.msgs = nil
		var lens []int
		for _, m := range msgs {
			if n := utf8.RuneCountInString(m); n >= sz.Min && n < sz.Max {
				sz.msgs = append(sz.msgs, m)
				lens = append(lens, n)
			}
		}
		if len(lens) == 0 {
			continue
		}
		sort.Ints(lens)
		sz.Chars = lens[len(lens)/2]
		kept = append(kept, sz)
	}
	s.Sizes = kept
}

// Text returns a deterministic chat-like text of n characters: words drawn from a Zipf
// distribution over a generated vocabulary, so it compresses like natural text instead of
// repeating itself.
func Text(n int, seed int64) string {
	r := rand.New(rand.NewSource(seed))
	z := rand.NewZipf(r, 1.1, 2, uint64(len(vocab)-1))
	var b strings.Builder
	for b.Len() < n {
		if b.Len() > 0 {
			if r.Intn(12) == 0 {
				b.WriteString(". ")
			} else {
				b.WriteByte(' ')
			}
		}
		b.WriteString(vocab[z.Uint64()])
	}
	return b.String()[:n]
}

var vocab = func() []string {
	r := rand.New(rand.NewSource(42))
	syl := []string{"ka", "re", "to", "mi", "sa", "lo", "ne", "di", "pu", "va", "te", "ro", "chi", "an", "el", "ou",
		"ing", "er", "st", "pr", "co", "de", "fi", "ma", "ti", "on", "ex", "un", "go", "fu", "nc", "rt"}
	common := []string{"the", "a", "to", "and", "is", "of", "in", "it", "test", "file", "fix", "run", "error", "go",
		"func", "return", "chat", "agent", "ok", "done", "commit", "build", "we", "this", "that", "for", "not"}
	out := append([]string{}, common...)
	seen := map[string]bool{}
	for _, w := range common {
		seen[w] = true
	}
	for len(out) < 4000 {
		var w string
		for i := 0; i < 1+r.Intn(4); i++ {
			w += syl[r.Intn(len(syl))]
		}
		if !seen[w] {
			seen[w] = true
			out = append(out, w)
		}
	}
	return out
}()

// Result is one protocol over one profile.
type Result struct {
	Protocol string `json:"protocol"`
	Profile  string `json:"profile"`
	// per size name
	Payload map[string]int64   `json:"payload_bytes"` // above TCP/UDP per message, both directions
	Wire    map[string]int64   `json:"wire_bytes"`    // ≈ with IP and TCP/UDP headers
	RTTms   map[string]float64 `json:"rtt_ms"`        // median: send a message, get the ack
	// a new connection (P2PACT: a new session) plus the first very short message
	ColdMS    float64     `json:"cold_ms"`
	ColdBytes int64       `json:"cold_bytes"`
	Sent      int         `json:"sent"`   // messages sent in the size tests
	Errors    int         `json:"errors"` // lost or arrived not intact (fingerprint mismatch)
	Load      *LoadResult `json:"load,omitempty"`
	Note      string      `json:"note,omitempty"`
}

func newResult(proto, profile string) *Result {
	return &Result{Protocol: proto, Profile: profile, Payload: map[string]int64{}, Wire: map[string]int64{},
		RTTms: map[string]float64{}}
}

func ms(d time.Duration) float64 { return float64(d.Microseconds()) / 1000 }

func median(d []time.Duration) time.Duration {
	if len(d) == 0 {
		return 0
	}
	s := append([]time.Duration(nil), d...)
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	return s[len(s)/2]
}

package bench

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// LoadReport reads a report saved by p2pact-bench run.
func LoadReport(path string) (*Report, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var r Report
	if err := json.Unmarshal(b, &r); err != nil {
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	r.Scenario.defaults()
	return &r, nil
}

// summaryCol is one column of the summary: a case and how to read it from a result.
type summaryCol struct {
	title, scenario, profile string
	get                      func(*Result) (float64, bool)
	format                   func(float64) string
	higherBetter             bool
}

func rtt(size string) func(*Result) (float64, bool) {
	return func(r *Result) (float64, bool) { v, ok := r.RTTms[size]; return v, ok }
}

func cold(r *Result) (float64, bool) { return r.ColdMS, r.ColdMS > 0 }

func loadRate(r *Result) (float64, bool) {
	if r.Load == nil {
		return 0, false
	}
	return r.Load.PerSecond, true
}

func wireOf(size string) func(*Result) (float64, bool) {
	return func(r *Result) (float64, bool) { v, ok := r.Wire[size]; return float64(v), ok }
}

func perSecond(v float64) string { return fmt.Sprintf("%.0f/s", v) }
func bytesOf(v float64) string   { return human(int64(v)) }

var protoOrder = []string{"P2PACT UDP", "P2PACT TCP", "WebSocket", "HTTPS+JSON", "HTTP/2", "gRPC", "HTTP/3"}

// Summary renders the comparison tables: the real network and the emulated links, each cell
// the value and how many times worse the protocol is than P2PACT over UDP.
func Summary(realDir, emuDir string) (string, error) {
	reports := map[string]*Report{}
	for _, dir := range []string{realDir, emuDir} {
		files, _ := filepath.Glob(filepath.Join(dir, "*.json"))
		for _, f := range files {
			r, err := LoadReport(f)
			if err != nil {
				return "", err
			}
			reports[dir+"/"+strings.TrimSuffix(filepath.Base(f), ".json")] = r
		}
	}
	find := func(dir, scenario, profile, proto string) *Result {
		r := reports[dir+"/"+scenario]
		if r == nil {
			return nil
		}
		for _, x := range r.Results {
			if x.Profile == profile && x.Protocol == proto {
				return x
			}
		}
		return nil
	}
	table := func(dir string, cols []summaryCol) string {
		var b strings.Builder
		b.WriteString("| Protocol |")
		for _, c := range cols {
			b.WriteString(" " + c.title + " |")
		}
		b.WriteString("\n|---" + strings.Repeat("|---:", len(cols)) + "|\n")
		for _, p := range protoOrder {
			b.WriteString("| " + p + " |")
			for _, c := range cols {
				x, base := find(dir, c.scenario, c.profile, p), find(dir, c.scenario, c.profile, "P2PACT UDP")
				if x == nil || base == nil {
					b.WriteString(" — |")
					continue
				}
				v, ok := c.get(x)
				bv, bok := c.get(base)
				if !ok {
					b.WriteString(" — |")
					continue
				}
				cell := c.format(v)
				if p != "P2PACT UDP" && bok && v > 0 && bv > 0 {
					k := v / bv
					if c.higherBetter {
						k = bv / v
					}
					cell += " " + times(k)
				} else if p == "P2PACT UDP" {
					cell = "**" + cell + "**"
				}
				b.WriteString(" " + cell + " |")
			}
			b.WriteString("\n")
		}
		return b.String()
	}
	var b strings.Builder
	b.WriteString("# P2PACT — summary\n\nEach cell: the value, and in brackets how many times worse it is than P2PACT over UDP " +
		"(longer time, more bytes, fewer messages a second). Above 1× P2PACT is better, below 1× it is worse. " +
		"Every message was checked by a fingerprint of what the receiver got; errors are listed in the detailed results.\n\n")
	b.WriteString("## Real network\n\nA laptop on home Wi-Fi to a cloud server, nothing emulated. Details: " +
		"[results-real/RESULTS.md](results-real/RESULTS.md).\n\n")
	real := "real network"
	b.WriteString(table(realDir, []summaryCol{
		{title: "Medium message", scenario: "human-messages", profile: real, get: rtt("medium"), format: dur},
		{title: "Long message", scenario: "human-messages", profile: real, get: rtt("long"), format: dur},
		{title: "Huge message (200 000 chars)", scenario: "synthetic", profile: real, get: rtt("huge"), format: dur},
		{title: "New connection + first message", scenario: "human-messages", profile: real, get: cold, format: dur},
		{title: "Load: messages a second (50 at a time)", scenario: "human-messages", profile: real, get: loadRate, format: perSecond, higherBetter: true},
		{title: "Bytes per medium message", scenario: "human-messages", profile: real, get: wireOf("medium"), format: bytesOf},
	}))
	b.WriteString("\n## Emulated links\n\nOne machine, delay, bandwidth and loss emulated. Details: " +
		"[results/RESULTS.md](results/RESULTS.md).\n\n")
	b.WriteString(table(emuDir, []summaryCol{
		{title: "Long, Bluetooth", scenario: "coding-agent", profile: "bluetooth", get: rtt("very long"), format: dur},
		{title: "Huge, Bluetooth", scenario: "synthetic", profile: "bluetooth", get: rtt("huge"), format: dur},
		{title: "Huge, local network", scenario: "synthetic", profile: "local network", get: rtt("huge"), format: dur},
		{title: "Huge, internet 1% loss *", scenario: "synthetic", profile: "internet", get: rtt("huge"), format: dur},
		{title: "Huge, mobile 3% loss *", scenario: "synthetic", profile: "mobile", get: rtt("huge"), format: dur},
		{title: "New connection, mobile", scenario: "human-messages", profile: "mobile", get: cold, format: dur},
		{title: "Load, internet 1% loss *", scenario: "human-messages", profile: "internet", get: loadRate, format: perSecond, higherBetter: true},
		{title: "Load, local network", scenario: "human-messages", profile: "local network", get: loadRate, format: perSecond, higherBetter: true},
	}))
	b.WriteString("\n\\* TCP-based protocols (P2PACT TCP, WebSocket, HTTPS, HTTP/2, gRPC) run without loss in the " +
		"emulation, and the emulating proxy terminates TCP, so these columns are optimistic for them. On lossy links " +
		"the fair comparison is with HTTP/3, which goes through the same emulated UDP link as P2PACT.\n")
	return b.String(), nil
}

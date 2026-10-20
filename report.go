package bench

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
	"time"
)

// Report is everything needed to compare a run with someone else's.
type Report struct {
	Scenario *Scenario `json:"scenario"`
	Data     string    `json:"data_source"`
	Env      Env       `json:"environment"`
	Results  []*Result `json:"results"`
}

// Env describes the machine and the code that produced a report.
type Env struct {
	Date     string `json:"date"`
	Go       string `json:"go"`
	OS       string `json:"os"`
	CPUs     int    `json:"cpus"`
	Revision string `json:"revision,omitempty"`
}

// CurrentEnv describes this machine and build.
func CurrentEnv() Env {
	e := Env{Date: time.Now().UTC().Format("2006-01-02"), Go: runtime.Version(), OS: runtime.GOOS + "/" + runtime.GOARCH,
		CPUs: runtime.NumCPU()}
	if bi, ok := debug.ReadBuildInfo(); ok {
		for _, s := range bi.Settings {
			if s.Key == "vcs.revision" && len(s.Value) >= 12 {
				e.Revision = s.Value[:12]
			}
		}
	}
	return e
}

var protoTitle = map[string]string{"P2PACT UDP": "P2PACT UDP", "P2PACT TCP": "P2PACT TCP", "WebSocket": "WebSocket", "HTTPS+JSON": "HTTPS+JSON"}

// Markdown renders a report: who, what data, bytes per message, time per profile.
func (r *Report) Markdown() string {
	var b strings.Builder
	sc := r.Scenario
	fmt.Fprintf(&b, "## %s\n\n", sc.Name)
	if sc.Description != "" {
		b.WriteString(sc.Description + "\n\n")
	}
	fmt.Fprintf(&b, "Data: %s. Messages per size: %d (%d for 10 000+ characters). Compression: %v. Seed: %d.  \n",
		r.Data, sc.Repeat, sc.RepeatLarge, *sc.Compression, sc.Seed)
	fmt.Fprintf(&b, "Machine: %s, %d CPUs, %s, %s", r.Env.OS, r.Env.CPUs, r.Env.Go, r.Env.Date)
	if r.Env.Revision != "" {
		fmt.Fprintf(&b, ", commit %s", r.Env.Revision)
	}
	b.WriteString(".\n\n")

	byProfile := map[string][]*Result{}
	for _, x := range r.Results {
		byProfile[x.Profile] = append(byProfile[x.Profile], x)
	}
	first := byProfile[sc.Profiles[0].Name]
	if len(first) == 0 {
		return b.String()
	}
	header := func(rs []*Result) string {
		var cols []string
		for _, x := range rs {
			cols = append(cols, x.Protocol)
		}
		return "| Message | " + strings.Join(cols, " | ") + " |\n|---" + strings.Repeat("|---:", len(rs)) + "|\n"
	}
	label := func(s Size) string { return fmt.Sprintf("%s (~%s chars)", s.Name, num(s.Chars)) }

	fmt.Fprintf(&b, "**Bytes on the wire per message, both directions** (%s; the number in brackets is how many "+
		"times more than P2PACT UDP):\n\n", sc.Profiles[0].Name)
	b.WriteString(header(first))
	base := first[0]
	cell := func(v, ref int64, approx bool) string {
		s := human(v)
		if approx {
			s = "≈" + s
		}
		if ref > 0 && v != ref {
			s += fmt.Sprintf(" (%.1f×)", float64(v)/float64(ref))
		}
		return s
	}
	for _, s := range sc.Sizes {
		var cols []string
		for _, x := range first {
			cols = append(cols, cell(x.Wire[s.Name], base.Wire[s.Name], x.Protocol != "P2PACT UDP"))
		}
		b.WriteString("| " + label(s) + " | " + strings.Join(cols, " | ") + " |\n")
	}
	var cols []string
	for _, x := range first {
		cols = append(cols, cell(x.ColdBytes, base.ColdBytes, x.Protocol != "P2PACT UDP"))
	}
	b.WriteString("| new connection + first message | " + strings.Join(cols, " | ") + " |\n\n")

	// every message is checked: the receiver answers with a fingerprint of what it got
	var lines []string
	for _, x := range r.Results {
		if x.Errors > 0 {
			lines = append(lines, fmt.Sprintf("%s · %s: %d", x.Protocol, x.Profile, x.Errors))
		}
	}
	if len(lines) == 0 {
		b.WriteString("**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).\n\n")
	} else {
		b.WriteString("**Delivery errors** (lost, or arrived not intact): " + strings.Join(lines, "; ") + ".\n\n")
	}

	b.WriteString("**Time from sending a message to its ack** (median):\n\n")
	for _, p := range sc.Profiles {
		rs := byProfile[p.Name]
		if len(rs) == 0 {
			continue
		}
		fmt.Fprintf(&b, "*%s — %s*\n\n", p.Name, describe(p))
		b.WriteString(header(rs))
		for _, s := range sc.Sizes {
			var cols []string
			for _, x := range rs {
				cols = append(cols, dur(x.RTTms[s.Name]))
			}
			b.WriteString("| " + label(s) + " | " + strings.Join(cols, " | ") + " |\n")
		}
		var cols []string
		for _, x := range rs {
			cols = append(cols, dur(x.ColdMS))
		}
		b.WriteString("| new connection + first message | " + strings.Join(cols, " | ") + " |\n\n")
		if rs[0].Load != nil {
			l := rs[0].Load
			fmt.Fprintf(&b, "Load: %d messages, %d at a time — messages per second (errors):\n\n", l.Messages, l.Concurrency)
			var hs, vs []string
			for _, x := range rs {
				if x.Load == nil {
					continue
				}
				hs = append(hs, x.Protocol)
				vs = append(vs, fmt.Sprintf("%.0f (%d)", x.Load.PerSecond, x.Load.Errors))
			}
			b.WriteString("| " + strings.Join(hs, " | ") + " |\n|" + strings.Repeat("---:|", len(hs)) + "\n| " + strings.Join(vs, " | ") + " |\n\n")
		}
		if p.LossPct > 0 {
			b.WriteString("TCP-based protocols run without loss here (emulating it needs root), so their times " +
				"are optimistic.\n\n")
		}
	}
	return b.String()
}

func describe(p Profile) string {
	if p.DelayMS == 0 && p.RateKBps == 0 && p.LossPct == 0 {
		return "a real network, nothing emulated"
	}
	parts := []string{fmt.Sprintf("%g ms one way", p.DelayMS)}
	if p.RateKBps > 0 {
		parts = append(parts, fmt.Sprintf("%d KB/s", p.RateKBps))
	}
	if p.LossPct > 0 {
		parts = append(parts, fmt.Sprintf("%g%% loss", p.LossPct))
	}
	return strings.Join(parts, ", ")
}

func human(n int64) string {
	switch {
	case n >= 10<<10:
		return fmt.Sprintf("%.0f KB", float64(n)/1024)
	case n >= 1<<10:
		return fmt.Sprintf("%.1f KB", float64(n)/1024)
	default:
		return fmt.Sprintf("%d B", n)
	}
}

func num(n int) string {
	s := fmt.Sprint(n)
	for i := len(s) - 3; i > 0; i -= 3 {
		s = s[:i] + " " + s[i:]
	}
	return s
}

func dur(msv float64) string {
	switch {
	case msv >= 1000:
		return fmt.Sprintf("%.2f s", msv/1000)
	case msv >= 10:
		return fmt.Sprintf("%.0f ms", msv)
	default:
		return fmt.Sprintf("%.1f ms", msv)
	}
}

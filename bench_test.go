package bench

import (
	"io"
	"testing"
)

// On a local network P2PACT over UDP must be smaller than WebSocket and HTTPS for every
// message size and for a new connection. `p2pact-bench run scenarios/*.json` gives the full
// picture.
func TestSmallerThanWebProtocols(t *testing.T) {
	if testing.Short() {
		t.Skip("benchmark")
	}
	sc := &Scenario{Name: "test", Repeat: 3, RepeatLarge: 1, Profiles: DefaultProfiles[:1],
		Protocols: []string{"p2pact-udp", "websocket", "https"}}
	tg, stop, err := LocalTarget()
	if err != nil {
		t.Fatal(err)
	}
	defer stop()
	rep, err := Execute(sc, tg, io.Discard)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("\n" + rep.Markdown())
	ours := rep.Results[0]
	for _, other := range rep.Results[1:] {
		for _, s := range sc.Sizes {
			if ours.Wire[s.Name] >= other.Wire[s.Name] {
				t.Errorf("%s: P2PACT %d B is not smaller than %s %d B", s.Name, ours.Wire[s.Name], other.Protocol, other.Wire[s.Name])
			}
		}
		if other.Errors > 0 || ours.Errors > 0 {
			t.Errorf("errors: P2PACT %d, %s %d", ours.Errors, other.Protocol, other.Errors)
		}
		if ours.ColdBytes >= other.ColdBytes {
			t.Errorf("new connection: P2PACT %d B is not smaller than %s %d B", ours.ColdBytes, other.Protocol, other.ColdBytes)
		}
	}
}

func TestTextDeterministic(t *testing.T) {
	for _, s := range DefaultSizes {
		a, b := Text(s.Chars, 3), Text(s.Chars, 3)
		if a != b || len(a) != s.Chars {
			t.Fatalf("%s: %d chars, deterministic %v", s.Name, len(a), a == b)
		}
	}
}

func TestFillBins(t *testing.T) {
	sc := &Scenario{}
	sc.defaults()
	sc.fill([]string{"ok", "да, давай", "a slightly longer message that goes past forty characters", "ещё одно сообщение подлиннее сорока символов точно"})
	if len(sc.Sizes) != 2 || sc.Sizes[0].Name != "very short" || len(sc.Sizes[1].msgs) != 2 {
		t.Fatalf("bins: %+v", sc.Sizes)
	}
}

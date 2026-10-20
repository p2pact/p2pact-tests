package bench

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Execute loads the scenario's data, runs it and returns the report.
func Execute(sc *Scenario, t *Target, log io.Writer) (*Report, error) {
	sc.defaults()
	msgs, credit, err := LoadData(sc.Data, log)
	if err != nil {
		return nil, err
	}
	if msgs != nil {
		sc.fill(msgs)
		if len(sc.Sizes) == 0 {
			return nil, fmt.Errorf("%s: no messages fall into the size bins", sc.Data)
		}
		for _, s := range sc.Sizes {
			fmt.Fprintf(log, "%s: %d messages, median %d chars\n", s.Name, len(s.msgs), s.Chars)
		}
	}
	res, err := Run(sc, t, log)
	if err != nil {
		return nil, err
	}
	return &Report{Scenario: sc, Data: credit, Env: CurrentEnv(), Results: res}, nil
}

// LocalTarget starts all receivers on this machine.
func LocalTarget() (*Target, func(), error) {
	dir, err := os.MkdirTemp("", "p2pact-bench-server")
	if err != nil {
		return nil, nil, err
	}
	s, err := StartServers(dir, "127.0.0.1", 0)
	if err != nil {
		return nil, nil, err
	}
	t := &Target{Servers: s, P2PUDP: s.P2PUDP, P2PTCP: s.P2PTCP, WS: s.WS, HTTPS: s.HTTPS, HTTP2: s.HTTP2,
		GRPC: s.GRPC, HTTP3: s.HTTP3,
		Invite: func() (string, error) { return s.Node.Invite(10 * time.Minute).String(), nil }}
	return t, func() { s.Close(); os.RemoveAll(dir) }, nil
}

// RemoteTarget is a server started with `p2pact-bench serve` on host.
func RemoteTarget(host string) *Target {
	at := func(i int) string { return net.JoinHostPort(host, strconv.Itoa(BasePort+i)) }
	t := &Target{P2PUDP: at(0), P2PTCP: at(0), WS: at(1), HTTPS: at(2), HTTP2: at(3), GRPC: at(4), HTTP3: at(5)}
	t.Invite = func() (string, error) {
		c := &http.Client{Timeout: 15 * time.Second, Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			TLSNextProto: map[string]func(string, *tls.Conn) http.RoundTripper{}}}
		resp, err := c.Get("https://" + t.HTTPS + "/invite")
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		b, err := io.ReadAll(resp.Body)
		return strings.TrimSpace(string(b)), err
	}
	return t
}

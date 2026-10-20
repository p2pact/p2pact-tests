package bench

// Network plumbing shared by all protocols: emulated links (a TCP proxy and a UDP proxy that
// delay, pace and drop), byte counters, certificates, WebSocket frames, a raw gRPC codec.

import (
	"bufio"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"math/big"
	mrand "math/rand"
	"net"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// TCP header estimate: IPv4 20 + TCP 20 + timestamps 12.
const (
	tcpHeader = 52
	tcpMSS    = 1448
)

// tcpWire estimates bytes on the wire for one exchange: data segments in both directions
// with headers, plus pure ACKs (one per two segments, delayed ACK).
func tcpWire(up, down int64) int64 {
	segs := func(n int64) int64 { return (n + tcpMSS - 1) / tcpMSS }
	su, sd := segs(up), segs(down)
	acks := (su+1)/2 + (sd+1)/2
	return up + down + tcpHeader*(su+sd+acks)
}

// tcpHandshake is SYN, SYN-ACK, ACK on the wire.
const tcpHandshake = 64 + 64 + 52

// emulated link between a client and a server: a TCP proxy that delays and paces every chunk
// per direction without serialising them (pipelined like a real link).
func linkProxy(target string, p Profile) (addr string, stop func(), err error) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return "", nil, err
	}
	pipe := func(dst, src net.Conn) {
		type chunk struct {
			b   []byte
			due time.Time
		}
		q := make(chan chunk, 4096)
		go func() {
			var free time.Time
			buf := make([]byte, 32<<10)
			for {
				n, err := src.Read(buf)
				if n > 0 {
					now := time.Now()
					due := now
					if p.rate() > 0 {
						if free.Before(now) {
							free = now
						}
						free = free.Add(time.Duration(n) * time.Second / time.Duration(p.rate()))
						due = free
					}
					q <- chunk{append([]byte(nil), buf[:n]...), due.Add(p.delay())}
				}
				if err != nil {
					close(q)
					return
				}
			}
		}()
		for c := range q {
			time.Sleep(time.Until(c.due))
			if _, err := dst.Write(c.b); err != nil {
				break
			}
		}
		dst.Close()
	}
	go func() {
		for {
			c, err := ln.Accept()
			if err != nil {
				return
			}
			s, err := net.Dial("tcp", target)
			if err != nil {
				c.Close()
				continue
			}
			go pipe(s, c)
			go pipe(c, s)
		}
	}()
	return ln.Addr().String(), func() { ln.Close() }, nil
}

// counting wraps a connection and counts bytes both ways (TLS payload above TCP).
type counting struct {
	net.Conn
	in, out *atomic.Int64
}

func (c counting) Read(b []byte) (int, error) {
	n, err := c.Conn.Read(b)
	c.in.Add(int64(n))
	return n, err
}

func (c counting) Write(b []byte) (int, error) {
	n, err := c.Conn.Write(b)
	c.out.Add(int64(n))
	return n, err
}

func selfSigned() tls.Certificate {
	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	tpl := &x509.Certificate{SerialNumber: big.NewInt(1), NotBefore: time.Now().Add(-time.Hour), NotAfter: time.Now().Add(time.Hour),
		IPAddresses: []net.IP{net.ParseIP("127.0.0.1")}}
	der, _ := x509.CreateCertificate(rand.Reader, tpl, tpl, &key.PublicKey, key)
	return tls.Certificate{Certificate: [][]byte{der}, PrivateKey: key}
}

const chatID = "62b7b92a-252b-43a2-971a-ec011b3c299c"

var bearer = strings.Repeat("3f9a1c7e", 4)

// frame is a WebSocket text frame (RFC 6455); clients mask.
func frame(payload []byte, masked bool) []byte {
	b := []byte{0x81}
	m := byte(0)
	if masked {
		m = 0x80
	}
	switch n := len(payload); {
	case n < 126:
		b = append(b, m|byte(n))
	case n < 1<<16:
		b = append(b, m|126)
		b = binary.BigEndian.AppendUint16(b, uint16(n))
	default:
		b = append(b, m|127)
		b = binary.BigEndian.AppendUint64(b, uint64(n))
	}
	if !masked {
		return append(b, payload...)
	}
	key := []byte{0x12, 0x34, 0x56, 0x78}
	b = append(b, key...)
	for i, x := range payload {
		b = append(b, x^key[i%4])
	}
	return b
}

func readFrame(r *bufio.Reader) ([]byte, error) {
	var h [2]byte
	if _, err := io.ReadFull(r, h[:]); err != nil {
		return nil, err
	}
	n := uint64(h[1] & 0x7f)
	switch n {
	case 126:
		var x [2]byte
		if _, err := io.ReadFull(r, x[:]); err != nil {
			return nil, err
		}
		n = uint64(binary.BigEndian.Uint16(x[:]))
	case 127:
		var x [8]byte
		if _, err := io.ReadFull(r, x[:]); err != nil {
			return nil, err
		}
		n = binary.BigEndian.Uint64(x[:])
	}
	if n > 64<<20 {
		return nil, errors.New("frame too big")
	}
	var key [4]byte
	if h[1]&0x80 != 0 {
		if _, err := io.ReadFull(r, key[:]); err != nil {
			return nil, err
		}
	}
	b := make([]byte, n)
	if _, err := io.ReadFull(r, b); err != nil {
		return nil, err
	}
	if h[1]&0x80 != 0 {
		for i := range b {
			b[i] ^= key[i%4]
		}
	}
	return b, nil
}

type rawCodec struct{}

func (rawCodec) Marshal(v any) ([]byte, error) {
	switch x := v.(type) {
	case *[]byte:
		return *x, nil
	case []byte:
		return x, nil
	}
	return nil, fmt.Errorf("rawCodec: %T", v)
}

func (rawCodec) Unmarshal(data []byte, v any) error {
	*(v.(*[]byte)) = append([]byte(nil), data...)
	return nil
}

func (rawCodec) Name() string { return "proto" }

// udpProxy forwards datagrams between one client and a server with the profile applied in
// both directions, and counts them.
type udpProxy struct {
	front, back *net.UDPConn
	mu          sync.Mutex
	client      *net.UDPAddr
	up, down    atomic.Int64 // bytes client → server, server → client
	pkts        atomic.Int64
	free        [2]time.Time
	rng         *mrand.Rand
}

func newUDPProxy(target string, p Profile, seed int64) (*udpProxy, error) {
	front, err := net.ListenUDP("udp", &net.UDPAddr{IP: net.IPv4(127, 0, 0, 1)})
	if err != nil {
		return nil, err
	}
	ta, _ := net.ResolveUDPAddr("udp", target)
	back, err := net.DialUDP("udp", nil, ta)
	if err != nil {
		front.Close()
		return nil, err
	}
	_ = front.SetReadBuffer(4 << 20)
	_ = back.SetReadBuffer(4 << 20)
	x := &udpProxy{front: front, back: back, rng: mrand.New(mrand.NewSource(seed))}
	deliver := func(dir int, b []byte, write func([]byte)) {
		x.mu.Lock()
		if dir == 0 {
			x.up.Add(int64(len(b)))
		} else {
			x.down.Add(int64(len(b)))
		}
		x.pkts.Add(1)
		drop := p.loss() > 0 && x.rng.Float64() < p.loss()
		now := time.Now()
		due := now
		if p.rate() > 0 {
			if x.free[dir].Before(now) {
				x.free[dir] = now
			}
			x.free[dir] = x.free[dir].Add(time.Duration(len(b)+28) * time.Second / time.Duration(p.rate()))
			due = x.free[dir]
		}
		x.mu.Unlock()
		if drop {
			return
		}
		cp := append([]byte(nil), b...)
		time.AfterFunc(time.Until(due)+p.delay(), func() { write(cp) })
	}
	go func() {
		buf := make([]byte, 65536)
		for {
			n, from, err := front.ReadFromUDP(buf)
			if err != nil {
				return
			}
			x.mu.Lock()
			x.client = from
			x.mu.Unlock()
			deliver(0, buf[:n], func(b []byte) { _, _ = back.Write(b) })
		}
	}()
	go func() {
		buf := make([]byte, 65536)
		for {
			n, err := back.Read(buf)
			if err != nil {
				return
			}
			x.mu.Lock()
			c := x.client
			x.mu.Unlock()
			if c != nil {
				deliver(1, buf[:n], func(b []byte) { _, _ = front.WriteToUDP(b, c) })
			}
		}
	}()
	return x, nil
}

func (x *udpProxy) addr() string { return x.front.LocalAddr().String() }
func (x *udpProxy) close()       { x.front.Close(); x.back.Close() }

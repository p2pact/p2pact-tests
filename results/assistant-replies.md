## Assistant replies

What an assistant answers: oasst2 assistant messages.

Data: OpenAssistant oasst2 (Apache-2.0), assistant. Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 169 B | ≈321 B (1.9×) | ≈391 B (2.3×) | ≈685 B (4.1×) | ≈416 B (2.5×) | ≈504 B (3.0×) | ≈475 B (2.8×) |
| medium (~103 chars) | 311 B | ≈463 B (1.5×) | ≈583 B (1.9×) | ≈876 B (2.8×) | ≈610 B (2.0×) | ≈709 B (2.3×) | ≈675 B (2.2×) |
| long (~697 chars) | 738 B | ≈896 B (1.2×) | ≈1.6 KB (2.2×) | ≈1.9 KB (2.7×) | ≈1.6 KB (2.3×) | ≈1.7 KB (2.3×) | ≈1.8 KB (2.4×) |
| very long (~2 491 chars) | 1.4 KB | ≈1.5 KB (1.1×) | ≈3.7 KB (2.6×) | ≈3.9 KB (2.9×) | ≈3.7 KB (2.7×) | ≈3.7 KB (2.7×) | ≈3.9 KB (2.8×) |
| new connection + first message | 639 B | ≈858 B (1.3×) | ≈4.2 KB (6.8×) | ≈4.1 KB (6.6×) | ≈4.3 KB (6.8×) | ≈4.3 KB (6.9×) | ≈50 KB (80.4×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 2.5 ms | 2.4 ms (1.0×) | 2.4 ms (1.0×) | 2.6 ms (1.0×) | 2.4 ms (1.0×) | 2.4 ms (1.0×) | 2.4 ms (1.0×) |
| medium (~103 chars) | 2.7 ms | 2.5 ms (0.9×) | 2.5 ms (0.9×) | 2.5 ms (0.9×) | 2.4 ms (0.9×) | 2.5 ms (0.9×) | 2.4 ms (0.9×) |
| long (~697 chars) | 2.6 ms | 2.6 ms (1.0×) | 2.4 ms (0.9×) | 2.6 ms (1.0×) | 2.4 ms (0.9×) | 2.4 ms (0.9×) | 2.4 ms (0.9×) |
| very long (~2 491 chars) | 2.6 ms | 2.7 ms (1.0×) | 2.4 ms (0.9×) | 2.6 ms (1.0×) | 2.8 ms (1.1×) | 2.4 ms (0.9×) | 2.4 ms (0.9×) |
| new connection + first message | 4.0 ms | 5.1 ms (1.3×) | 8.4 ms (2.1×) | 5.3 ms (1.3×) | 6.7 ms (1.7×) | 8.1 ms (2.0×) | 5.7 ms (1.4×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 20978 (0 errors) | 19342 (0 errors) (1.1×) | 20930 (0 errors) (1.0×) | 19745 (0 errors) (1.1×) | 21647 (0 errors) (1.0×) | 19768 (0 errors) (1.1×) | 21298 (0 errors) (1.0×) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 65 ms | 64 ms (1.0×) | 66 ms (1.0×) | 73 ms (1.1×) | 66 ms (1.0×) | 68 ms (1.0×) | 72 ms (1.1×) |
| medium (~103 chars) | 69 ms | 68 ms (1.0×) | 70 ms (1.0×) | 77 ms (1.1×) | 71 ms (1.0×) | 73 ms (1.1×) | 76 ms (1.1×) |
| long (~697 chars) | 81 ms | 82 ms (1.0×) | 100 ms (1.2×) | 108 ms (1.3×) | 101 ms (1.2×) | 102 ms (1.3×) | 109 ms (1.3×) |
| very long (~2 491 chars) | 99 ms | 97 ms (1.0×) | 137 ms (1.4×) | 144 ms (1.5×) | 137 ms (1.4×) | 138 ms (1.4×) | 148 ms (1.5×) |
| new connection + first message | 73 ms | 73 ms (1.0×) | 280 ms (3.8×) | 215 ms (2.9×) | 215 ms (2.9×) | 276 ms (3.8×) | 270 ms (3.7×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 213 (0 errors) | 246 (0 errors) (0.9×) | 167 (0 errors) (1.3×) | 631 (0 errors) (0.3×) | 177 (0 errors) (1.2×) | 174 (0 errors) (1.2×) | 104 (0 errors) (2.0×) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 81 ms | 81 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) |
| medium (~103 chars) | 81 ms | 81 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) |
| long (~697 chars) | 81 ms | 82 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) |
| very long (~2 491 chars) | 81 ms | 81 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) |
| new connection + first message | 83 ms | 82 ms (1.0×) | 244 ms (2.9×) | 163 ms (2.0×) | 163 ms (2.0×) | 246 ms (3.0×) | 163 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 535 (0 errors) | 614 (0 errors) (0.9×) | 617 (0 errors) (0.9×) | 600 (0 errors) (0.9×) | 614 (0 errors) (0.9×) | 612 (0 errors) (0.9×) | 582 (0 errors) (0.9×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 201 ms | 202 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) |
| medium (~103 chars) | 202 ms | 202 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) |
| long (~697 chars) | 203 ms | 203 ms (1.0×) | 204 ms (1.0×) | 205 ms (1.0×) | 205 ms (1.0×) | 204 ms (1.0×) | 205 ms (1.0×) |
| very long (~2 491 chars) | 204 ms | 204 ms (1.0×) | 207 ms (1.0×) | 209 ms (1.0×) | 208 ms (1.0×) | 207 ms (1.0×) | 208 ms (1.0×) |
| new connection + first message | 203 ms | 203 ms (1.0×) | 611 ms (3.0×) | 411 ms (2.0×) | 411 ms (2.0×) | 612 ms (3.0×) | 416 ms (2.1×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 143 (0 errors) | 246 (0 errors) (0.6×) | 245 (0 errors) (0.6×) | 240 (0 errors) (0.6×) | 246 (0 errors) (0.6×) | 241 (0 errors) (0.6×) | 195 (0 errors) (0.7×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


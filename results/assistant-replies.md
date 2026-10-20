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

**Time from sending a message to its ack** (median):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 2.5 ms | 2.4 ms | 2.4 ms | 2.6 ms | 2.4 ms | 2.4 ms | 2.4 ms |
| medium (~103 chars) | 2.7 ms | 2.5 ms | 2.5 ms | 2.5 ms | 2.4 ms | 2.5 ms | 2.4 ms |
| long (~697 chars) | 2.6 ms | 2.6 ms | 2.4 ms | 2.6 ms | 2.4 ms | 2.4 ms | 2.4 ms |
| very long (~2 491 chars) | 2.6 ms | 2.7 ms | 2.4 ms | 2.6 ms | 2.8 ms | 2.4 ms | 2.4 ms |
| new connection + first message | 4.0 ms | 5.1 ms | 8.4 ms | 5.3 ms | 6.7 ms | 8.1 ms | 5.7 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 20978 (0) | 19342 (0) | 20930 (0) | 19745 (0) | 21647 (0) | 19768 (0) | 21298 (0) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 65 ms | 64 ms | 66 ms | 73 ms | 66 ms | 68 ms | 72 ms |
| medium (~103 chars) | 69 ms | 68 ms | 70 ms | 77 ms | 71 ms | 73 ms | 76 ms |
| long (~697 chars) | 81 ms | 82 ms | 100 ms | 108 ms | 101 ms | 102 ms | 109 ms |
| very long (~2 491 chars) | 99 ms | 97 ms | 137 ms | 144 ms | 137 ms | 138 ms | 148 ms |
| new connection + first message | 73 ms | 73 ms | 280 ms | 215 ms | 215 ms | 276 ms | 270 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 213 (0) | 246 (0) | 167 (0) | 631 (0) | 177 (0) | 174 (0) | 104 (0) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 81 ms | 81 ms | 81 ms | 81 ms | 81 ms | 81 ms | 82 ms |
| medium (~103 chars) | 81 ms | 81 ms | 81 ms | 81 ms | 81 ms | 81 ms | 82 ms |
| long (~697 chars) | 81 ms | 82 ms | 81 ms | 81 ms | 81 ms | 82 ms | 82 ms |
| very long (~2 491 chars) | 81 ms | 81 ms | 81 ms | 81 ms | 81 ms | 81 ms | 81 ms |
| new connection + first message | 83 ms | 82 ms | 244 ms | 163 ms | 163 ms | 246 ms | 163 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 535 (0) | 614 (0) | 617 (0) | 600 (0) | 614 (0) | 612 (0) | 582 (0) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 201 ms | 202 ms | 202 ms | 202 ms | 202 ms | 202 ms | 203 ms |
| medium (~103 chars) | 202 ms | 202 ms | 203 ms | 203 ms | 202 ms | 202 ms | 202 ms |
| long (~697 chars) | 203 ms | 203 ms | 204 ms | 205 ms | 205 ms | 204 ms | 205 ms |
| very long (~2 491 chars) | 204 ms | 204 ms | 207 ms | 209 ms | 208 ms | 207 ms | 208 ms |
| new connection + first message | 203 ms | 203 ms | 611 ms | 411 ms | 411 ms | 612 ms | 416 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 143 (0) | 246 (0) | 245 (0) | 240 (0) | 246 (0) | 241 (0) | 195 (0) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


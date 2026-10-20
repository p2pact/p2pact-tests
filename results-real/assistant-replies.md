## Assistant replies

What an assistant answers: oasst2 assistant messages.

Data: OpenAssistant oasst2 (Apache-2.0), assistant. Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 172 B | ≈319 B (1.9×) | ≈391 B (2.3×) | ≈685 B (4.0×) | ≈418 B (2.4×) | ≈513 B (3.0×) | ≈1.1 KB (6.6×) |
| medium (~103 chars) | 374 B | ≈459 B (1.2×) | ≈583 B (1.6×) | ≈876 B (2.3×) | ≈613 B (1.6×) | ≈704 B (1.9×) | ≈727 B (1.9×) |
| long (~697 chars) | 774 B | ≈893 B (1.2×) | ≈1.6 KB (2.1×) | ≈1.9 KB (2.6×) | ≈1.6 KB (2.2×) | ≈1.7 KB (2.2×) | ≈1.9 KB (2.5×) |
| very long (~2 491 chars) | 1.4 KB | ≈1.5 KB (1.1×) | ≈3.7 KB (2.6×) | ≈3.9 KB (2.8×) | ≈3.7 KB (2.6×) | ≈3.7 KB (2.6×) | ≈3.9 KB (2.8×) |
| new connection + first message | 701 B | ≈844 B (1.2×) | ≈4.2 KB (6.2×) | ≈4.1 KB (6.0×) | ≈4.3 KB (6.2×) | ≈4.3 KB (6.2×) | ≈21 KB (30.8×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 121 ms | 122 ms (1.0×) | 125 ms (1.0×) | 121 ms (1.0×) | 122 ms (1.0×) | 121 ms (1.0×) | 122 ms (1.0×) |
| medium (~103 chars) | 126 ms | 125 ms (1.0×) | 120 ms (1.0×) | 123 ms (1.0×) | 131 ms (1.0×) | 122 ms (1.0×) | 127 ms (1.0×) |
| long (~697 chars) | 131 ms | 127 ms (1.0×) | 126 ms (1.0×) | 127 ms (1.0×) | 122 ms (0.9×) | 127 ms (1.0×) | 124 ms (1.0×) |
| very long (~2 491 chars) | 132 ms | 132 ms (1.0×) | 122 ms (0.9×) | 126 ms (1.0×) | 126 ms (1.0×) | 125 ms (0.9×) | 125 ms (1.0×) |
| new connection + first message | 130 ms | 256 ms (2.0×) | 531 ms (4.1×) | 406 ms (3.1×) | 405 ms (3.1×) | 539 ms (4.1×) | 250 ms (1.9×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 337 (0 errors) | 352 (0 errors) (1.0×) | 365 (0 errors) (0.9×) | 233 (0 errors) (1.4×) | 370 (0 errors) (0.9×) | 379 (0 errors) (0.9×) | 366 (0 errors) (0.9×) |


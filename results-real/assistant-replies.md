## Assistant replies

What an assistant answers: oasst2 assistant messages.

Data: OpenAssistant oasst2 (Apache-2.0), assistant. Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 299 B | ≈319 B (1.1×) | ≈391 B (1.3×) | ≈685 B (2.3×) | ≈419 B (1.4×) | ≈509 B (1.7×) | ≈1.1 KB (3.6×) |
| medium (~103 chars) | 433 B | ≈459 B (1.1×) | ≈583 B (1.3×) | ≈876 B (2.0×) | ≈613 B (1.4×) | ≈706 B (1.6×) | ≈710 B (1.6×) |
| long (~697 chars) | 879 B | ≈893 B (1.0×) | ≈1.6 KB (1.9×) | ≈1.9 KB (2.2×) | ≈1.6 KB (1.9×) | ≈1.7 KB (2.0×) | ≈1.9 KB (2.3×) |
| very long (~2 491 chars) | 1.5 KB | ≈1.5 KB (1.0×) | ≈3.7 KB (2.4×) | ≈3.9 KB (2.6×) | ≈3.7 KB (2.4×) | ≈3.7 KB (2.4×) | ≈4.0 KB (2.6×) |
| new connection + first message | 888 B | ≈951 B (1.1×) | ≈4.2 KB (4.9×) | ≈4.1 KB (4.8×) | ≈4.3 KB (4.9×) | ≈4.3 KB (4.9×) | ≈21 KB (24.2×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 136 ms | 159 ms (1.2×) | 136 ms (1.0×) | 145 ms (1.1×) | 127 ms (0.9×) | 142 ms (1.0×) | 136 ms (1.0×) |
| medium (~103 chars) | 130 ms | 145 ms (1.1×) | 135 ms (1.0×) | 129 ms (1.0×) | 125 ms (1.0×) | 155 ms (1.2×) | 134 ms (1.0×) |
| long (~697 chars) | 146 ms | 151 ms (1.0×) | 145 ms (1.0×) | 136 ms (0.9×) | 128 ms (0.9×) | 154 ms (1.1×) | 139 ms (1.0×) |
| very long (~2 491 chars) | 132 ms | 169 ms (1.3×) | 136 ms (1.0×) | 132 ms (1.0×) | 139 ms (1.1×) | 154 ms (1.2×) | 150 ms (1.1×) |
| new connection + first message | 132 ms | 280 ms (2.1×) | 543 ms (4.1×) | 461 ms (3.5×) | 439 ms (3.3×) | 633 ms (4.8×) | 297 ms (2.2×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 375 (0 errors) | 313 (0 errors) (1.2×) | 342 (0 errors) (1.1×) | 216 (0 errors) (1.7×) | 348 (0 errors) (1.1×) | 322 (0 errors) (1.2×) | 329 (0 errors) (1.1×) |


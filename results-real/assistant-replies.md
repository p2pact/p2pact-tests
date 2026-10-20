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

**Time from sending a message to its ack** (median):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 136 ms | 159 ms | 136 ms | 145 ms | 127 ms | 142 ms | 136 ms |
| medium (~103 chars) | 130 ms | 145 ms | 135 ms | 129 ms | 125 ms | 155 ms | 134 ms |
| long (~697 chars) | 146 ms | 151 ms | 145 ms | 136 ms | 128 ms | 154 ms | 139 ms |
| very long (~2 491 chars) | 132 ms | 169 ms | 136 ms | 132 ms | 139 ms | 154 ms | 150 ms |
| new connection + first message | 132 ms | 280 ms | 543 ms | 461 ms | 439 ms | 633 ms | 297 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 375 (0) | 313 (0) | 342 (0) | 216 (0) | 348 (0) | 322 (0) | 329 (0) |


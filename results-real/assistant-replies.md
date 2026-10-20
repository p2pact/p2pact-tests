## Assistant replies

What an assistant answers: oasst2 assistant messages.

Data: OpenAssistant oasst2 (Apache-2.0), assistant. Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 179 B | ≈319 B (1.8×) | ≈391 B (2.2×) | ≈685 B (3.8×) | ≈418 B (2.3×) | ≈507 B (2.8×) | ≈1.0 KB (5.9×) |
| medium (~103 chars) | 311 B | ≈459 B (1.5×) | ≈583 B (1.9×) | ≈876 B (2.8×) | ≈613 B (2.0×) | ≈696 B (2.2×) | ≈669 B (2.2×) |
| long (~697 chars) | 744 B | ≈928 B (1.2×) | ≈1.6 KB (2.2×) | ≈1.9 KB (2.7×) | ≈1.6 KB (2.2×) | ≈1.7 KB (2.3×) | ≈1.8 KB (2.5×) |
| very long (~2 491 chars) | 1.4 KB | ≈1.6 KB (1.1×) | ≈3.7 KB (2.6×) | ≈3.9 KB (2.9×) | ≈3.7 KB (2.7×) | ≈3.7 KB (2.7×) | ≈3.9 KB (2.8×) |
| new connection + first message | 675 B | ≈844 B (1.3×) | ≈4.2 KB (6.4×) | ≈4.1 KB (6.3×) | ≈4.3 KB (6.5×) | ≈4.3 KB (6.5×) | ≈21 KB (32.0×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 109 ms | 120 ms | 97 ms | 113 ms | 119 ms | 105 ms | 118 ms |
| medium (~103 chars) | 109 ms | 120 ms | 97 ms | 113 ms | 119 ms | 104 ms | 119 ms |
| long (~697 chars) | 109 ms | 120 ms | 97 ms | 113 ms | 118 ms | 106 ms | 118 ms |
| very long (~2 491 chars) | 110 ms | 120 ms | 98 ms | 114 ms | 120 ms | 106 ms | 120 ms |
| new connection + first message | 113 ms | 213 ms | 442 ms | 336 ms | 329 ms | 430 ms | 239 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 448 (0) | 404 (0) | 494 (0) | 276 (0) | 394 (0) | 465 (0) | 415 (0) |


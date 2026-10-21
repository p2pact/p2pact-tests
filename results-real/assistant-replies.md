## Assistant replies

What an assistant answers: oasst2 assistant messages.

Data: OpenAssistant oasst2 (Apache-2.0), assistant. Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 176 B | ≈461 B (2.6×) | ≈401 B (2.3×) | ≈695 B (3.9×) | ≈429 B (2.4×) | ≈532 B (3.0×) | ≈1.1 KB (6.2×) |
| medium (~104 chars) | 270 B | ≈553 B (2.0×) | ≈513 B (1.9×) | ≈806 B (3.0×) | ≈541 B (2.0×) | ≈642 B (2.4×) | ≈636 B (2.4×) |
| long (~694 chars) | 654 B | ≈943 B (1.4×) | ≈1.4 KB (2.2×) | ≈1.7 KB (2.6×) | ≈1.4 KB (2.2×) | ≈1.5 KB (2.4×) | ≈1.5 KB (2.4×) |
| very long (~2 476 chars) | 1.3 KB | ≈1.6 KB (1.2×) | ≈3.2 KB (2.4×) | ≈3.5 KB (2.6×) | ≈3.2 KB (2.4×) | ≈3.3 KB (2.5×) | ≈3.7 KB (2.7×) |
| new connection + first message | 990 B | ≈1.2 KB (1.2×) | ≈4.2 KB (4.4×) | ≈4.1 KB (4.3×) | ≈4.3 KB (4.4×) | ≈4.3 KB (4.4×) | ≈21 KB (21.6×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 118 ms | 126 ms (1.1×) | 120 ms (1.0×) | 132 ms (1.1×) | 120 ms (1.0×) | 115 ms (1.0×) | 114 ms (1.0×) |
| medium (~104 chars) | 127 ms | 132 ms (1.0×) | 121 ms (1.0×) | 134 ms (1.1×) | 128 ms (1.0×) | 127 ms (1.0×) | 116 ms (0.9×) |
| long (~694 chars) | 125 ms | 129 ms (1.0×) | 126 ms (1.0×) | 131 ms (1.1×) | 122 ms (1.0×) | 118 ms (0.9×) | 125 ms (1.0×) |
| very long (~2 476 chars) | 124 ms | 133 ms (1.1×) | 123 ms (1.0×) | 136 ms (1.1×) | 129 ms (1.0×) | 123 ms (1.0×) | 125 ms (1.0×) |
| new connection + first message | 131 ms | 252 ms (1.9×) | 516 ms (3.9×) | 383 ms (2.9×) | 386 ms (2.9×) | 517 ms (3.9×) | 246 ms (1.9×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 364 (0 errors) | 350 (0 errors) (1.0×) | 387 (0 errors) (0.9×) | 233 (0 errors) (1.6×) | 360 (0 errors) (1.0×) | 385 (0 errors) (0.9×) | 388 (0 errors) (0.9×) |


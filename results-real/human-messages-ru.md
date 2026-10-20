## Human messages, Russian

Russian-language prompts from oasst2: Cyrillic takes two bytes per character.

Data: OpenAssistant oasst2 (Apache-2.0), human, language ru. Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 209 B | ≈341 B (1.6×) | ≈412 B (2.0×) | ≈706 B (3.4×) | ≈440 B (2.1×) | ≈540 B (2.6×) | ≈1.0 KB (5.1×) |
| medium (~75 chars) | 273 B | ≈430 B (1.6×) | ≈529 B (1.9×) | ≈822 B (3.0×) | ≈558 B (2.0×) | ≈651 B (2.4×) | ≈603 B (2.2×) |
| long (~329 chars) | 562 B | ≈690 B (1.2×) | ≈1.1 KB (2.0×) | ≈1.4 KB (2.5×) | ≈1.1 KB (2.1×) | ≈1.2 KB (2.2×) | ≈1.2 KB (2.2×) |
| very long (~3 305 chars) | 2.6 KB | ≈2.7 KB (1.0×) | ≈6.5 KB (2.5×) | ≈6.8 KB (2.6×) | ≈6.6 KB (2.5×) | ≈6.6 KB (2.5×) | ≈7.1 KB (2.7×) |
| new connection + first message | 731 B | ≈844 B (1.2×) | ≈4.2 KB (6.0×) | ≈4.2 KB (5.8×) | ≈4.3 KB (6.0×) | ≈4.3 KB (6.0×) | ≈21 KB (29.4×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 111 ms | 120 ms (1.1×) | 135 ms (1.2×) | 117 ms (1.0×) | 121 ms (1.1×) | 121 ms (1.1×) | 126 ms (1.1×) |
| medium (~75 chars) | 116 ms | 121 ms (1.0×) | 143 ms (1.2×) | 119 ms (1.0×) | 119 ms (1.0×) | 123 ms (1.1×) | 132 ms (1.1×) |
| long (~329 chars) | 123 ms | 120 ms (1.0×) | 135 ms (1.1×) | 122 ms (1.0×) | 120 ms (1.0×) | 124 ms (1.0×) | 127 ms (1.0×) |
| very long (~3 305 chars) | 129 ms | 127 ms (1.0×) | 146 ms (1.1×) | 127 ms (1.0×) | 126 ms (1.0×) | 129 ms (1.0×) | 141 ms (1.1×) |
| new connection + first message | 117 ms | 251 ms (2.1×) | 557 ms (4.7×) | 397 ms (3.4×) | 390 ms (3.3×) | 487 ms (4.1×) | 269 ms (2.3×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 372 (0 errors) | 357 (0 errors) (1.0×) | 343 (0 errors) (1.1×) | 231 (0 errors) (1.6×) | 384 (0 errors) (1.0×) | 354 (0 errors) (1.1×) | 372 (0 errors) (1.0×) |


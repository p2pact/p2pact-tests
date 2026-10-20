## Human messages, Russian

Russian-language prompts from oasst2: Cyrillic takes two bytes per character.

Data: OpenAssistant oasst2 (Apache-2.0), human, language ru. Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 188 B | ≈341 B (1.8×) | ≈412 B (2.2×) | ≈706 B (3.8×) | ≈440 B (2.3×) | ≈540 B (2.9×) | ≈1.1 KB (5.8×) |
| medium (~75 chars) | 315 B | ≈421 B (1.3×) | ≈529 B (1.7×) | ≈822 B (2.6×) | ≈558 B (1.8×) | ≈653 B (2.1×) | ≈624 B (2.0×) |
| long (~329 chars) | 572 B | ≈690 B (1.2×) | ≈1.1 KB (2.0×) | ≈1.4 KB (2.5×) | ≈1.1 KB (2.0×) | ≈1.2 KB (2.2×) | ≈1.3 KB (2.4×) |
| very long (~3 305 chars) | 2.8 KB | ≈2.6 KB (0.9×) | ≈6.5 KB (2.3×) | ≈6.8 KB (2.4×) | ≈6.6 KB (2.4×) | ≈6.6 KB (2.4×) | ≈6.7 KB (2.4×) |
| new connection + first message | 731 B | ≈844 B (1.2×) | ≈4.2 KB (5.9×) | ≈4.2 KB (5.8×) | ≈4.3 KB (6.0×) | ≈4.3 KB (6.0×) | ≈21 KB (29.3×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 132 ms | 137 ms (1.0×) | 133 ms (1.0×) | 138 ms (1.0×) | 151 ms (1.2×) | 150 ms (1.1×) | 154 ms (1.2×) |
| medium (~75 chars) | 139 ms | 144 ms (1.0×) | 145 ms (1.0×) | 137 ms (1.0×) | 157 ms (1.1×) | 148 ms (1.1×) | 152 ms (1.1×) |
| long (~329 chars) | 137 ms | 150 ms (1.1×) | 144 ms (1.0×) | 141 ms (1.0×) | 149 ms (1.1×) | 140 ms (1.0×) | 161 ms (1.2×) |
| very long (~3 305 chars) | 144 ms | 152 ms (1.1×) | 158 ms (1.1×) | 149 ms (1.0×) | 154 ms (1.1×) | 144 ms (1.0×) | 179 ms (1.2×) |
| new connection + first message | 158 ms | 296 ms (1.9×) | 582 ms (3.7×) | 443 ms (2.8×) | 428 ms (2.7×) | 550 ms (3.5×) | 333 ms (2.1×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 333 (0 errors) | 313 (0 errors) (1.1×) | 328 (0 errors) (1.0×) | 203 (0 errors) (1.6×) | 286 (0 errors) (1.2×) | 324 (0 errors) (1.0×) | 279 (0 errors) (1.2×) |


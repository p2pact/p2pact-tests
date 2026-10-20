## Human messages, Russian

Russian-language prompts from oasst2: Cyrillic takes two bytes per character.

Data: OpenAssistant oasst2 (Apache-2.0), human, language ru. Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 225 B | ≈341 B (1.5×) | ≈412 B (1.8×) | ≈706 B (3.1×) | ≈440 B (2.0×) | ≈529 B (2.4×) | ≈1.1 KB (4.8×) |
| medium (~75 chars) | 336 B | ≈421 B (1.3×) | ≈529 B (1.6×) | ≈822 B (2.4×) | ≈557 B (1.7×) | ≈643 B (1.9×) | ≈613 B (1.8×) |
| long (~329 chars) | 696 B | ≈768 B (1.1×) | ≈1.1 KB (1.6×) | ≈1.4 KB (2.0×) | ≈1.1 KB (1.7×) | ≈1.2 KB (1.8×) | ≈1.3 KB (1.9×) |
| very long (~3 305 chars) | 3.1 KB | ≈2.6 KB (0.9×) | ≈6.5 KB (2.1×) | ≈6.8 KB (2.2×) | ≈6.6 KB (2.1×) | ≈6.6 KB (2.1×) | ≈6.9 KB (2.3×) |
| new connection + first message | 701 B | ≈844 B (1.2×) | ≈4.2 KB (6.2×) | ≈4.2 KB (6.1×) | ≈4.3 KB (6.3×) | ≈4.3 KB (6.3×) | ≈21 KB (30.7×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 105 ms | 112 ms | 103 ms | 109 ms | 116 ms | 100 ms | 102 ms |
| medium (~75 chars) | 105 ms | 112 ms | 104 ms | 109 ms | 117 ms | 98 ms | 101 ms |
| long (~329 chars) | 105 ms | 112 ms | 103 ms | 110 ms | 115 ms | 100 ms | 101 ms |
| very long (~3 305 chars) | 107 ms | 114 ms | 105 ms | 111 ms | 118 ms | 102 ms | 103 ms |
| new connection + first message | 109 ms | 231 ms | 418 ms | 337 ms | 320 ms | 454 ms | 207 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 460 (0) | 421 (0) | 468 (0) | 271 (0) | 402 (0) | 488 (0) | 471 (0) |


## Human messages

What people type to an assistant: OpenAssistant oasst2 prompts, all languages.

Data: OpenAssistant oasst2 (Apache-2.0), human. Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 180 B | ≈333 B (1.9×) | ≈405 B (2.2×) | ≈698 B (3.9×) | ≈433 B (2.4×) | ≈527 B (2.9×) | ≈1.1 KB (6.1×) |
| medium (~80 chars) | 232 B | ≈386 B (1.7×) | ≈462 B (2.0×) | ≈754 B (3.2×) | ≈489 B (2.1×) | ≈588 B (2.5×) | ≈610 B (2.6×) |
| long (~296 chars) | 539 B | ≈641 B (1.2×) | ≈858 B (1.6×) | ≈1.1 KB (2.1×) | ≈890 B (1.7×) | ≈979 B (1.8×) | ≈994 B (1.8×) |
| very long (~2 718 chars) | 3.1 KB | ≈2.0 KB (0.6×) | ≈4.4 KB (1.4×) | ≈4.7 KB (1.5×) | ≈4.5 KB (1.4×) | ≈4.5 KB (1.4×) | ≈5.2 KB (1.7×) |
| new connection + first message | 723 B | ≈861 B (1.2×) | ≈4.2 KB (6.0×) | ≈4.1 KB (5.9×) | ≈4.3 KB (6.1×) | ≈4.3 KB (6.1×) | ≈21 KB (29.8×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 124 ms | 135 ms (1.1×) | 129 ms (1.0×) | 138 ms (1.1×) | 121 ms (1.0×) | 117 ms (0.9×) | 114 ms (0.9×) |
| medium (~80 chars) | 120 ms | 134 ms (1.1×) | 135 ms (1.1×) | 137 ms (1.1×) | 125 ms (1.0×) | 121 ms (1.0×) | 119 ms (1.0×) |
| long (~296 chars) | 134 ms | 142 ms (1.1×) | 131 ms (1.0×) | 141 ms (1.1×) | 120 ms (0.9×) | 129 ms (1.0×) | 117 ms (0.9×) |
| very long (~2 718 chars) | 133 ms | 136 ms (1.0×) | 135 ms (1.0×) | 136 ms (1.0×) | 128 ms (1.0×) | 125 ms (0.9×) | 126 ms (0.9×) |
| new connection + first message | 121 ms | 242 ms (2.0×) | 478 ms (3.9×) | 401 ms (3.3×) | 397 ms (3.3×) | 517 ms (4.3×) | 244 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 355 (0 errors) | 343 (0 errors) (1.0×) | 357 (0 errors) (1.0×) | 230 (0 errors) (1.5×) | 368 (0 errors) (1.0×) | 398 (0 errors) (0.9×) | 392 (0 errors) (0.9×) |


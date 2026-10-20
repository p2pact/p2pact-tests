## Human messages

What people type to an assistant: OpenAssistant oasst2 prompts, all languages.

Data: OpenAssistant oasst2 (Apache-2.0), human. Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 187 B | ≈333 B (1.8×) | ≈405 B (2.2×) | ≈698 B (3.7×) | ≈432 B (2.3×) | ≈521 B (2.8×) | ≈1.1 KB (5.9×) |
| medium (~80 chars) | 232 B | ≈379 B (1.6×) | ≈462 B (2.0×) | ≈754 B (3.2×) | ≈489 B (2.1×) | ≈575 B (2.5×) | ≈627 B (2.7×) |
| long (~296 chars) | 423 B | ≈572 B (1.4×) | ≈858 B (2.0×) | ≈1.1 KB (2.7×) | ≈890 B (2.1×) | ≈973 B (2.3×) | ≈1.2 KB (2.9×) |
| very long (~2 718 chars) | 2.0 KB | ≈1.9 KB (1.0×) | ≈4.4 KB (2.2×) | ≈4.7 KB (2.4×) | ≈4.5 KB (2.2×) | ≈4.5 KB (2.2×) | ≈4.8 KB (2.4×) |
| new connection + first message | 671 B | ≈833 B (1.2×) | ≈4.2 KB (6.5×) | ≈4.1 KB (6.3×) | ≈4.3 KB (6.5×) | ≈4.3 KB (6.6×) | ≈21 KB (31.7×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 116 ms | 116 ms | 108 ms | 102 ms | 114 ms | 115 ms | 114 ms |
| medium (~80 chars) | 115 ms | 116 ms | 108 ms | 102 ms | 115 ms | 115 ms | 113 ms |
| long (~296 chars) | 115 ms | 116 ms | 108 ms | 101 ms | 115 ms | 115 ms | 115 ms |
| very long (~2 718 chars) | 116 ms | 117 ms | 110 ms | 103 ms | 117 ms | 115 ms | 114 ms |
| new connection + first message | 119 ms | 215 ms | 439 ms | 325 ms | 325 ms | 435 ms | 229 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 421 (0) | 402 (0) | 441 (0) | 280 (0) | 417 (0) | 431 (0) | 427 (0) |


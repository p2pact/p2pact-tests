## Human messages

What people type to an assistant: OpenAssistant oasst2 prompts, all languages.

Data: OpenAssistant oasst2 (Apache-2.0), human. Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 180 B | ≈333 B (1.9×) | ≈405 B (2.2×) | ≈698 B (3.9×) | ≈433 B (2.4×) | ≈534 B (3.0×) | ≈1.1 KB (6.2×) |
| medium (~80 chars) | 226 B | ≈388 B (1.7×) | ≈462 B (2.0×) | ≈754 B (3.3×) | ≈490 B (2.2×) | ≈578 B (2.6×) | ≈634 B (2.8×) |
| long (~296 chars) | 417 B | ≈572 B (1.4×) | ≈858 B (2.1×) | ≈1.1 KB (2.8×) | ≈890 B (2.1×) | ≈975 B (2.3×) | ≈971 B (2.3×) |
| very long (~2 718 chars) | 2.9 KB | ≈2.0 KB (0.7×) | ≈4.4 KB (1.5×) | ≈4.7 KB (1.6×) | ≈4.5 KB (1.5×) | ≈4.5 KB (1.5×) | ≈4.7 KB (1.6×) |
| new connection + first message | 723 B | ≈840 B (1.2×) | ≈4.2 KB (6.0×) | ≈4.1 KB (5.9×) | ≈4.3 KB (6.1×) | ≈4.3 KB (6.1×) | ≈21 KB (29.7×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 127 ms | 139 ms | 137 ms | 140 ms | 152 ms | 145 ms | 142 ms |
| medium (~80 chars) | 136 ms | 142 ms | 145 ms | 141 ms | 156 ms | 148 ms | 145 ms |
| long (~296 chars) | 138 ms | 151 ms | 139 ms | 154 ms | 153 ms | 154 ms | 155 ms |
| very long (~2 718 chars) | 137 ms | 144 ms | 143 ms | 148 ms | 163 ms | 159 ms | 146 ms |
| new connection + first message | 138 ms | 286 ms | 572 ms | 450 ms | 464 ms | 553 ms | 304 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 368 (0) | 311 (0) | 339 (0) | 209 (0) | 308 (0) | 297 (0) | 333 (0) |


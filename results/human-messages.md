## Human messages

What people type to an assistant: OpenAssistant oasst2 prompts, all languages.

Data: OpenAssistant oasst2 (Apache-2.0), human. Messages per size: 20 (2 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 163 B | ≈319 B (2.0×) | ≈382 B (2.3×) | ≈675 B (4.1×) | ≈407 B (2.5×) | ≈520 B (3.2×) | ≈919 B (5.6×) |
| medium (~80 chars) | 213 B | ≈369 B (1.7×) | ≈439 B (2.1×) | ≈731 B (3.4×) | ≈463 B (2.2×) | ≈578 B (2.7×) | ≈542 B (2.5×) |
| long (~296 chars) | 403 B | ≈557 B (1.4×) | ≈835 B (2.1×) | ≈1.1 KB (2.8×) | ≈864 B (2.1×) | ≈982 B (2.4×) | ≈952 B (2.4×) |
| very long (~2 718 chars) | 1.7 KB | ≈1.9 KB (1.1×) | ≈4.4 KB (2.5×) | ≈4.7 KB (2.7×) | ≈4.4 KB (2.6×) | ≈4.4 KB (2.6×) | ≈4.6 KB (2.7×) |
| new connection + first message | 771 B | ≈919 B (1.2×) | ≈4.4 KB (5.8×) | ≈4.3 KB (5.7×) | ≈4.4 KB (5.9×) | ≈4.4 KB (5.9×) | ≈15 KB (20.4×) |

**Time from sending a message to its ack** (median):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 2.4 ms | 2.4 ms | 2.4 ms | 2.5 ms | 2.5 ms | 2.4 ms | 2.5 ms |
| medium (~80 chars) | 2.4 ms | 2.5 ms | 2.4 ms | 2.5 ms | 2.5 ms | 2.4 ms | 2.4 ms |
| long (~296 chars) | 2.5 ms | 2.5 ms | 2.4 ms | 2.5 ms | 2.5 ms | 2.4 ms | 2.5 ms |
| very long (~2 718 chars) | 2.6 ms | 2.6 ms | 2.5 ms | 2.4 ms | 2.5 ms | 2.4 ms | 2.5 ms |
| new connection + first message | 3.2 ms | 3.1 ms | 8.1 ms | 5.7 ms | 5.5 ms | 7.8 ms | 5.5 ms |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 65 ms | 63 ms | 65 ms | 73 ms | 66 ms | 68 ms | 73 ms |
| medium (~80 chars) | 66 ms | 65 ms | 67 ms | 74 ms | 67 ms | 70 ms | 73 ms |
| long (~296 chars) | 70 ms | 68 ms | 72 ms | 79 ms | 72 ms | 76 ms | 79 ms |
| very long (~2 718 chars) | 98 ms | 96 ms | 139 ms | 145 ms | 140 ms | 140 ms | 149 ms |
| new connection + first message | 72 ms | 69 ms | 277 ms | 213 ms | 214 ms | 276 ms | 268 ms |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 81 ms | 81 ms | 81 ms | 81 ms | 81 ms | 81 ms | 81 ms |
| medium (~80 chars) | 81 ms | 81 ms | 81 ms | 81 ms | 81 ms | 81 ms | 81 ms |
| long (~296 chars) | 81 ms | 81 ms | 81 ms | 81 ms | 81 ms | 81 ms | 81 ms |
| very long (~2 718 chars) | 81 ms | 81 ms | 81 ms | 81 ms | 81 ms | 81 ms | 81 ms |
| new connection + first message | 81 ms | 81 ms | 245 ms | 163 ms | 163 ms | 243 ms | 163 ms |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 201 ms | 202 ms | 203 ms | 203 ms | 203 ms | 203 ms | 203 ms |
| medium (~80 chars) | 202 ms | 202 ms | 203 ms | 203 ms | 203 ms | 203 ms | 203 ms |
| long (~296 chars) | 202 ms | 202 ms | 203 ms | 204 ms | 203 ms | 203 ms | 203 ms |
| very long (~2 718 chars) | 204 ms | 204 ms | 208 ms | 209 ms | 209 ms | 207 ms | 208 ms |
| new connection + first message | 202 ms | 202 ms | 616 ms | 412 ms | 410 ms | 614 ms | 414 ms |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


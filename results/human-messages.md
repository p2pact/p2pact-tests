## Human messages

What people type to an assistant: OpenAssistant oasst2 prompts, all languages.

Data: OpenAssistant oasst2 (Apache-2.0), human. Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 184 B | ≈332 B (1.8×) | ≈405 B (2.2×) | ≈698 B (3.8×) | ≈430 B (2.3×) | ≈519 B (2.8×) | ≈480 B (2.6×) |
| medium (~80 chars) | 232 B | ≈379 B (1.6×) | ≈462 B (2.0×) | ≈754 B (3.2×) | ≈487 B (2.1×) | ≈593 B (2.6×) | ≈582 B (2.5×) |
| long (~296 chars) | 423 B | ≈572 B (1.4×) | ≈858 B (2.0×) | ≈1.1 KB (2.7×) | ≈887 B (2.1×) | ≈987 B (2.3×) | ≈943 B (2.2×) |
| very long (~2 718 chars) | 1.8 KB | ≈1.9 KB (1.1×) | ≈4.4 KB (2.5×) | ≈4.7 KB (2.7×) | ≈4.5 KB (2.5×) | ≈4.5 KB (2.5×) | ≈4.7 KB (2.6×) |
| new connection + first message | 773 B | ≈923 B (1.2×) | ≈4.2 KB (5.6×) | ≈4.1 KB (5.5×) | ≈4.3 KB (5.7×) | ≈4.3 KB (5.7×) | ≈49 KB (64.4×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 2.4 ms | 2.5 ms | 2.5 ms | 2.5 ms | 2.5 ms | 2.5 ms | 2.5 ms |
| medium (~80 chars) | 2.4 ms | 2.5 ms | 2.4 ms | 2.5 ms | 2.5 ms | 2.5 ms | 2.4 ms |
| long (~296 chars) | 2.6 ms | 2.5 ms | 2.4 ms | 2.5 ms | 2.4 ms | 2.5 ms | 2.4 ms |
| very long (~2 718 chars) | 2.7 ms | 2.9 ms | 2.4 ms | 2.5 ms | 2.5 ms | 2.5 ms | 2.4 ms |
| new connection + first message | 4.7 ms | 3.8 ms | 7.8 ms | 5.7 ms | 5.8 ms | 7.8 ms | 6.6 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 21528 (0) | 20206 (0) | 21421 (0) | 18790 (0) | 20720 (0) | 21145 (0) | 20566 (0) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 66 ms | 64 ms | 68 ms | 75 ms | 67 ms | 68 ms | 73 ms |
| medium (~80 chars) | 67 ms | 66 ms | 68 ms | 76 ms | 69 ms | 70 ms | 74 ms |
| long (~296 chars) | 71 ms | 69 ms | 74 ms | 82 ms | 74 ms | 76 ms | 80 ms |
| very long (~2 718 chars) | 100 ms | 97 ms | 141 ms | 150 ms | 143 ms | 139 ms | 152 ms |
| new connection + first message | 75 ms | 72 ms | 282 ms | 217 ms | 217 ms | 278 ms | 273 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 247 (0) | 300 (0) | 201 (0) | 634 (0) | 217 (0) | 209 (0) | 117 (0) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 82 ms | 81 ms | 82 ms | 83 ms | 83 ms | 82 ms | 83 ms |
| medium (~80 chars) | 81 ms | 81 ms | 82 ms | 83 ms | 83 ms | 82 ms | 83 ms |
| long (~296 chars) | 82 ms | 82 ms | 82 ms | 83 ms | 83 ms | 82 ms | 83 ms |
| very long (~2 718 chars) | 82 ms | 81 ms | 83 ms | 83 ms | 83 ms | 82 ms | 83 ms |
| new connection + first message | 84 ms | 83 ms | 246 ms | 168 ms | 165 ms | 247 ms | 166 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 487 (0) | 610 (0) | 605 (0) | 591 (0) | 606 (0) | 608 (0) | 558 (0) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 202 ms | 203 ms | 203 ms | 203 ms | 203 ms | 203 ms | 203 ms |
| medium (~80 chars) | 203 ms | 203 ms | 203 ms | 204 ms | 204 ms | 203 ms | 202 ms |
| long (~296 chars) | 203 ms | 204 ms | 204 ms | 205 ms | 203 ms | 203 ms | 203 ms |
| very long (~2 718 chars) | 206 ms | 206 ms | 210 ms | 210 ms | 210 ms | 208 ms | 208 ms |
| new connection + first message | 205 ms | 206 ms | 616 ms | 412 ms | 413 ms | 617 ms | 417 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 147 (0) | 245 (0) | 245 (0) | 239 (0) | 245 (0) | 244 (0) | 203 (0) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


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

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 2.4 ms | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) |
| medium (~80 chars) | 2.4 ms | 2.5 ms (1.0×) | 2.4 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.4 ms (1.0×) |
| long (~296 chars) | 2.6 ms | 2.5 ms (1.0×) | 2.4 ms (0.9×) | 2.5 ms (0.9×) | 2.4 ms (0.9×) | 2.5 ms (1.0×) | 2.4 ms (0.9×) |
| very long (~2 718 chars) | 2.7 ms | 2.9 ms (1.1×) | 2.4 ms (0.9×) | 2.5 ms (0.9×) | 2.5 ms (0.9×) | 2.5 ms (0.9×) | 2.4 ms (0.9×) |
| new connection + first message | 4.7 ms | 3.8 ms (0.8×) | 7.8 ms (1.7×) | 5.7 ms (1.2×) | 5.8 ms (1.3×) | 7.8 ms (1.7×) | 6.6 ms (1.4×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 21528 (0 errors) | 20206 (0 errors) (1.1×) | 21421 (0 errors) (1.0×) | 18790 (0 errors) (1.1×) | 20720 (0 errors) (1.0×) | 21145 (0 errors) (1.0×) | 20566 (0 errors) (1.0×) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 66 ms | 64 ms (1.0×) | 68 ms (1.0×) | 75 ms (1.1×) | 67 ms (1.0×) | 68 ms (1.0×) | 73 ms (1.1×) |
| medium (~80 chars) | 67 ms | 66 ms (1.0×) | 68 ms (1.0×) | 76 ms (1.1×) | 69 ms (1.0×) | 70 ms (1.0×) | 74 ms (1.1×) |
| long (~296 chars) | 71 ms | 69 ms (1.0×) | 74 ms (1.1×) | 82 ms (1.2×) | 74 ms (1.0×) | 76 ms (1.1×) | 80 ms (1.1×) |
| very long (~2 718 chars) | 100 ms | 97 ms (1.0×) | 141 ms (1.4×) | 150 ms (1.5×) | 143 ms (1.4×) | 139 ms (1.4×) | 152 ms (1.5×) |
| new connection + first message | 75 ms | 72 ms (1.0×) | 282 ms (3.8×) | 217 ms (2.9×) | 217 ms (2.9×) | 278 ms (3.7×) | 273 ms (3.6×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 247 (0 errors) | 300 (0 errors) (0.8×) | 201 (0 errors) (1.2×) | 634 (0 errors) (0.4×) | 217 (0 errors) (1.1×) | 209 (0 errors) (1.2×) | 117 (0 errors) (2.1×) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 82 ms | 81 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| medium (~80 chars) | 81 ms | 81 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| long (~296 chars) | 82 ms | 82 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| very long (~2 718 chars) | 82 ms | 81 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| new connection + first message | 84 ms | 83 ms (1.0×) | 246 ms (2.9×) | 168 ms (2.0×) | 165 ms (2.0×) | 247 ms (3.0×) | 166 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 487 (0 errors) | 610 (0 errors) (0.8×) | 605 (0 errors) (0.8×) | 591 (0 errors) (0.8×) | 606 (0 errors) (0.8×) | 608 (0 errors) (0.8×) | 558 (0 errors) (0.9×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 202 ms | 203 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) |
| medium (~80 chars) | 203 ms | 203 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) | 204 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) |
| long (~296 chars) | 203 ms | 204 ms (1.0×) | 204 ms (1.0×) | 205 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) |
| very long (~2 718 chars) | 206 ms | 206 ms (1.0×) | 210 ms (1.0×) | 210 ms (1.0×) | 210 ms (1.0×) | 208 ms (1.0×) | 208 ms (1.0×) |
| new connection + first message | 205 ms | 206 ms (1.0×) | 616 ms (3.0×) | 412 ms (2.0×) | 413 ms (2.0×) | 617 ms (3.0×) | 417 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 147 (0 errors) | 245 (0 errors) (0.6×) | 245 (0 errors) (0.6×) | 239 (0 errors) (0.6×) | 245 (0 errors) (0.6×) | 244 (0 errors) (0.6×) | 203 (0 errors) (0.7×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


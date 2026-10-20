## Human messages

What people type to an assistant: OpenAssistant oasst2 prompts, all languages.

Data: OpenAssistant oasst2 (Apache-2.0), human. Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 178 B | ≈332 B (1.9×) | ≈405 B (2.3×) | ≈698 B (3.9×) | ≈431 B (2.4×) | ≈539 B (3.0×) | ≈919 B (5.2×) |
| medium (~80 chars) | 226 B | ≈379 B (1.7×) | ≈462 B (2.0×) | ≈754 B (3.3×) | ≈486 B (2.2×) | ≈600 B (2.7×) | ≈591 B (2.6×) |
| long (~296 chars) | 423 B | ≈572 B (1.4×) | ≈858 B (2.0×) | ≈1.1 KB (2.7×) | ≈887 B (2.1×) | ≈962 B (2.3×) | ≈989 B (2.3×) |
| very long (~2 718 chars) | 1.8 KB | ≈1.9 KB (1.1×) | ≈4.4 KB (2.5×) | ≈4.7 KB (2.7×) | ≈4.5 KB (2.5×) | ≈4.5 KB (2.5×) | ≈5.0 KB (2.8×) |
| new connection + first message | 741 B | ≈923 B (1.2×) | ≈4.2 KB (5.9×) | ≈4.1 KB (5.7×) | ≈4.3 KB (5.9×) | ≈4.3 KB (6.0×) | ≈55 KB (75.4×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 2.7 ms | 2.8 ms (1.0×) | 2.7 ms (1.0×) | 2.7 ms (1.0×) | 2.7 ms (1.0×) | 2.7 ms (1.0×) | 2.7 ms (1.0×) |
| medium (~80 chars) | 2.8 ms | 2.7 ms (1.0×) | 2.7 ms (1.0×) | 2.8 ms (1.0×) | 2.8 ms (1.0×) | 2.7 ms (1.0×) | 2.7 ms (1.0×) |
| long (~296 chars) | 2.8 ms | 2.9 ms (1.0×) | 2.7 ms (0.9×) | 2.8 ms (1.0×) | 2.8 ms (1.0×) | 2.8 ms (1.0×) | 2.7 ms (0.9×) |
| very long (~2 718 chars) | 3.2 ms | 3.2 ms (1.0×) | 2.7 ms (0.8×) | 2.9 ms (0.9×) | 2.8 ms (0.9×) | 2.7 ms (0.9×) | 2.8 ms (0.9×) |
| new connection + first message | 5.4 ms | 4.2 ms (0.8×) | 9.5 ms (1.7×) | 7.1 ms (1.3×) | 6.3 ms (1.2×) | 9.4 ms (1.7×) | 7.4 ms (1.4×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 19035 (0 errors) | 19731 (0 errors) (1.0×) | 19714 (0 errors) (1.0×) | 16820 (0 errors) (1.1×) | 20576 (0 errors) (0.9×) | 18243 (0 errors) (1.0×) | 18718 (0 errors) (1.0×) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 67 ms | 66 ms (1.0×) | 68 ms (1.0×) | 75 ms (1.1×) | 67 ms (1.0×) | 69 ms (1.0×) | 73 ms (1.1×) |
| medium (~80 chars) | 69 ms | 67 ms (1.0×) | 69 ms (1.0×) | 76 ms (1.1×) | 69 ms (1.0×) | 70 ms (1.0×) | 74 ms (1.1×) |
| long (~296 chars) | 72 ms | 71 ms (1.0×) | 75 ms (1.0×) | 82 ms (1.1×) | 75 ms (1.0×) | 76 ms (1.1×) | 81 ms (1.1×) |
| very long (~2 718 chars) | 101 ms | 99 ms (1.0×) | 140 ms (1.4×) | 148 ms (1.5×) | 142 ms (1.4×) | 139 ms (1.4×) | 152 ms (1.5×) |
| new connection + first message | 75 ms | 74 ms (1.0×) | 283 ms (3.8×) | 218 ms (2.9×) | 217 ms (2.9×) | 279 ms (3.7×) | 273 ms (3.6×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 187 (0 errors) | 204 (0 errors) (0.9×) | 201 (0 errors) (0.9×) | 631 (0 errors) (0.3×) | 217 (0 errors) (0.9×) | 209 (0 errors) (0.9×) | 118 (0 errors) (1.6×) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 81 ms | 82 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| medium (~80 chars) | 82 ms | 82 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| long (~296 chars) | 82 ms | 82 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) |
| very long (~2 718 chars) | 83 ms | 83 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| new connection + first message | 86 ms | 84 ms (1.0×) | 247 ms (2.9×) | 166 ms (1.9×) | 167 ms (1.9×) | 248 ms (2.9×) | 167 ms (1.9×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 594 (0 errors) | 607 (0 errors) (1.0×) | 606 (0 errors) (1.0×) | 591 (0 errors) (1.0×) | 606 (0 errors) (1.0×) | 607 (0 errors) (1.0×) | 547 (0 errors) (1.1×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 203 ms | 202 ms (1.0×) | 203 ms (1.0×) | 205 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) |
| medium (~80 chars) | 204 ms | 203 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) |
| long (~296 chars) | 204 ms | 204 ms (1.0×) | 204 ms (1.0×) | 205 ms (1.0×) | 204 ms (1.0×) | 203 ms (1.0×) | 205 ms (1.0×) |
| very long (~2 718 chars) | 207 ms | 206 ms (1.0×) | 210 ms (1.0×) | 210 ms (1.0×) | 210 ms (1.0×) | 209 ms (1.0×) | 210 ms (1.0×) |
| new connection + first message | 205 ms | 206 ms (1.0×) | 616 ms (3.0×) | 412 ms (2.0×) | 413 ms (2.0×) | 618 ms (3.0×) | 418 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 233 (0 errors) | 245 (0 errors) (1.0×) | 245 (0 errors) (1.0×) | 239 (0 errors) (1.0×) | 245 (0 errors) (1.0×) | 241 (0 errors) (1.0×) | 190 (0 errors) (1.2×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


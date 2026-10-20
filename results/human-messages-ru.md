## Human messages, Russian

Russian-language prompts from oasst2: Cyrillic takes two bytes per character.

Data: OpenAssistant oasst2 (Apache-2.0), human, language ru. Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 185 B | ≈340 B (1.8×) | ≈412 B (2.2×) | ≈706 B (3.8×) | ≈438 B (2.4×) | ≈537 B (2.9×) | ≈707 B (3.8×) |
| medium (~75 chars) | 267 B | ≈421 B (1.6×) | ≈529 B (2.0×) | ≈822 B (3.1×) | ≈554 B (2.1×) | ≈663 B (2.5×) | ≈679 B (2.5×) |
| long (~329 chars) | 542 B | ≈690 B (1.3×) | ≈1.1 KB (2.1×) | ≈1.4 KB (2.6×) | ≈1.1 KB (2.1×) | ≈1.2 KB (2.3×) | ≈1.2 KB (2.3×) |
| very long (~3 305 chars) | 2.5 KB | ≈2.6 KB (1.0×) | ≈6.5 KB (2.6×) | ≈6.8 KB (2.8×) | ≈6.6 KB (2.7×) | ≈6.6 KB (2.7×) | ≈6.7 KB (2.7×) |
| new connection + first message | 779 B | ≈959 B (1.2×) | ≈4.2 KB (5.6×) | ≈4.2 KB (5.5×) | ≈4.3 KB (5.7×) | ≈4.3 KB (5.7×) | ≈47 KB (62.0×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 2.8 ms | 2.7 ms (1.0×) | 2.7 ms (1.0×) | 2.7 ms (1.0×) | 2.7 ms (1.0×) | 2.8 ms (1.0×) | 2.7 ms (1.0×) |
| medium (~75 chars) | 2.8 ms | 2.8 ms (1.0×) | 2.7 ms (1.0×) | 2.7 ms (1.0×) | 2.7 ms (1.0×) | 2.7 ms (1.0×) | 2.7 ms (1.0×) |
| long (~329 chars) | 2.9 ms | 2.8 ms (1.0×) | 2.7 ms (0.9×) | 2.8 ms (1.0×) | 2.8 ms (1.0×) | 2.7 ms (0.9×) | 2.7 ms (0.9×) |
| very long (~3 305 chars) | 3.4 ms | 3.5 ms (1.0×) | 2.9 ms (0.8×) | 3.0 ms (0.9×) | 2.8 ms (0.8×) | 2.7 ms (0.8×) | 2.7 ms (0.8×) |
| new connection + first message | 4.8 ms | 4.7 ms (1.0×) | 9.5 ms (2.0×) | 6.3 ms (1.3×) | 6.6 ms (1.4×) | 9.3 ms (1.9×) | 6.1 ms (1.3×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 18025 (0 errors) | 19496 (0 errors) (0.9×) | 19101 (0 errors) (0.9×) | 18025 (0 errors) (1.0×) | 21314 (0 errors) (0.8×) | 18949 (0 errors) (1.0×) | 19313 (0 errors) (0.9×) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 67 ms | 66 ms (1.0×) | 68 ms (1.0×) | 75 ms (1.1×) | 68 ms (1.0×) | 69 ms (1.0×) | 73 ms (1.1×) |
| medium (~75 chars) | 69 ms | 67 ms (1.0×) | 71 ms (1.0×) | 79 ms (1.1×) | 70 ms (1.0×) | 72 ms (1.0×) | 76 ms (1.1×) |
| long (~329 chars) | 75 ms | 73 ms (1.0×) | 80 ms (1.1×) | 87 ms (1.2×) | 80 ms (1.1×) | 81 ms (1.1×) | 85 ms (1.1×) |
| very long (~3 305 chars) | 128 ms | 126 ms (1.0×) | 182 ms (1.4×) | 191 ms (1.5×) | 182 ms (1.4×) | 183 ms (1.4×) | 194 ms (1.5×) |
| new connection + first message | 76 ms | 74 ms (1.0×) | 285 ms (3.7×) | 218 ms (2.9×) | 220 ms (2.9×) | 280 ms (3.7×) | 274 ms (3.6×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 162 (0 errors) | 242 (0 errors) (0.7×) | 158 (0 errors) (1.0×) | 624 (0 errors) (0.3×) | 167 (0 errors) (1.0×) | 162 (0 errors) (1.0×) | 101 (0 errors) (1.6×) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 82 ms | 82 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) |
| medium (~75 chars) | 82 ms | 82 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| long (~329 chars) | 82 ms | 82 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| very long (~3 305 chars) | 82 ms | 84 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| new connection + first message | 86 ms | 86 ms (1.0×) | 247 ms (2.9×) | 168 ms (2.0×) | 168 ms (2.0×) | 247 ms (2.9×) | 166 ms (1.9×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 526 (0 errors) | 606 (0 errors) (0.9×) | 607 (0 errors) (0.9×) | 589 (0 errors) (0.9×) | 605 (0 errors) (0.9×) | 607 (0 errors) (0.9×) | 572 (0 errors) (0.9×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 202 ms | 203 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) |
| medium (~75 chars) | 203 ms | 203 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) |
| long (~329 chars) | 204 ms | 203 ms (1.0×) | 204 ms (1.0×) | 205 ms (1.0×) | 205 ms (1.0×) | 204 ms (1.0×) | 204 ms (1.0×) |
| very long (~3 305 chars) | 208 ms | 208 ms (1.0×) | 214 ms (1.0×) | 214 ms (1.0×) | 213 ms (1.0×) | 213 ms (1.0×) | 212 ms (1.0×) |
| new connection + first message | 207 ms | 206 ms (1.0×) | 616 ms (3.0×) | 413 ms (2.0×) | 412 ms (2.0×) | 616 ms (3.0×) | 418 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 164 (0 errors) | 244 (0 errors) (0.7×) | 245 (0 errors) (0.7×) | 239 (0 errors) (0.7×) | 245 (0 errors) (0.7×) | 241 (0 errors) (0.7×) | 176 (0 errors) (0.9×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


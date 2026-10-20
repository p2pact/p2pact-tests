## Human messages, Russian

Russian-language prompts from oasst2: Cyrillic takes two bytes per character.

Data: OpenAssistant oasst2 (Apache-2.0), human, language ru. Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 185 B | ≈340 B (1.8×) | ≈412 B (2.2×) | ≈706 B (3.8×) | ≈437 B (2.4×) | ≈540 B (2.9×) | ≈932 B (5.0×) |
| medium (~75 chars) | 267 B | ≈421 B (1.6×) | ≈529 B (2.0×) | ≈822 B (3.1×) | ≈554 B (2.1×) | ≈658 B (2.5×) | ≈600 B (2.2×) |
| long (~329 chars) | 536 B | ≈690 B (1.3×) | ≈1.1 KB (2.1×) | ≈1.4 KB (2.6×) | ≈1.1 KB (2.2×) | ≈1.2 KB (2.3×) | ≈1.2 KB (2.3×) |
| very long (~3 305 chars) | 2.5 KB | ≈2.6 KB (1.1×) | ≈6.5 KB (2.6×) | ≈6.8 KB (2.8×) | ≈6.6 KB (2.7×) | ≈6.6 KB (2.7×) | ≈6.7 KB (2.7×) |
| new connection + first message | 638 B | ≈901 B (1.4×) | ≈4.2 KB (6.8×) | ≈4.2 KB (6.7×) | ≈4.3 KB (6.9×) | ≈4.3 KB (6.9×) | ≈47 KB (75.8×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 2.6 ms | 2.6 ms (1.0×) | 2.5 ms (0.9×) | 2.7 ms (1.0×) | 2.7 ms (1.0×) | 2.6 ms (1.0×) | 2.6 ms (1.0×) |
| medium (~75 chars) | 2.9 ms | 2.5 ms (0.9×) | 2.5 ms (0.8×) | 2.7 ms (0.9×) | 2.6 ms (0.9×) | 2.7 ms (0.9×) | 2.6 ms (0.9×) |
| long (~329 chars) | 2.7 ms | 2.7 ms (1.0×) | 2.7 ms (1.0×) | 2.6 ms (1.0×) | 2.7 ms (1.0×) | 2.6 ms (1.0×) | 2.6 ms (1.0×) |
| very long (~3 305 chars) | 3.0 ms | 2.9 ms (1.0×) | 2.7 ms (0.9×) | 2.7 ms (0.9×) | 2.7 ms (0.9×) | 2.7 ms (0.9×) | 2.6 ms (0.9×) |
| new connection + first message | 4.4 ms | 4.4 ms (1.0×) | 8.0 ms (1.8×) | 5.6 ms (1.3×) | 5.9 ms (1.3×) | 8.4 ms (1.9×) | 6.1 ms (1.4×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 19386 (0 errors) | 19332 (0 errors) (1.0×) | 17737 (0 errors) (1.1×) | 18788 (0 errors) (1.0×) | 20668 (0 errors) (0.9×) | 18382 (0 errors) (1.1×) | 20220 (0 errors) (1.0×) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 66 ms | 64 ms (1.0×) | 67 ms (1.0×) | 74 ms (1.1×) | 67 ms (1.0×) | 68 ms (1.0×) | 72 ms (1.1×) |
| medium (~75 chars) | 68 ms | 67 ms (1.0×) | 69 ms (1.0×) | 77 ms (1.1×) | 70 ms (1.0×) | 71 ms (1.0×) | 74 ms (1.1×) |
| long (~329 chars) | 73 ms | 71 ms (1.0×) | 78 ms (1.1×) | 85 ms (1.2×) | 78 ms (1.1×) | 79 ms (1.1×) | 83 ms (1.1×) |
| very long (~3 305 chars) | 127 ms | 115 ms (0.9×) | 182 ms (1.4×) | 189 ms (1.5×) | 183 ms (1.4×) | 183 ms (1.4×) | 192 ms (1.5×) |
| new connection + first message | 74 ms | 72 ms (1.0×) | 280 ms (3.8×) | 216 ms (2.9×) | 216 ms (2.9×) | 276 ms (3.7×) | 271 ms (3.7×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 203 (0 errors) | 241 (0 errors) (0.8×) | 158 (0 errors) (1.3×) | 630 (0 errors) (0.3×) | 167 (0 errors) (1.2×) | 162 (0 errors) (1.3×) | 102 (0 errors) (2.0×) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 81 ms | 81 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) |
| medium (~75 chars) | 81 ms | 81 ms (1.0×) | 82 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) |
| long (~329 chars) | 81 ms | 81 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) |
| very long (~3 305 chars) | 82 ms | 81 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) |
| new connection + first message | 82 ms | 82 ms (1.0×) | 245 ms (3.0×) | 164 ms (2.0×) | 165 ms (2.0×) | 246 ms (3.0×) | 165 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 491 (0 errors) | 612 (0 errors) (0.8×) | 614 (0 errors) (0.8×) | 597 (0 errors) (0.8×) | 612 (0 errors) (0.8×) | 613 (0 errors) (0.8×) | 585 (0 errors) (0.8×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 201 ms | 203 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) |
| medium (~75 chars) | 201 ms | 203 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) |
| long (~329 chars) | 202 ms | 203 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) |
| very long (~3 305 chars) | 206 ms | 207 ms (1.0×) | 212 ms (1.0×) | 212 ms (1.0×) | 211 ms (1.0×) | 211 ms (1.0×) | 214 ms (1.0×) |
| new connection + first message | 203 ms | 204 ms (1.0×) | 616 ms (3.0×) | 410 ms (2.0×) | 411 ms (2.0×) | 613 ms (3.0×) | 415 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 127 (0 errors) | 246 (0 errors) (0.5×) | 246 (0 errors) (0.5×) | 240 (0 errors) (0.5×) | 247 (0 errors) (0.5×) | 242 (0 errors) (0.5×) | 171 (0 errors) (0.7×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


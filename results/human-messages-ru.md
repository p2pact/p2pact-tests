## Human messages, Russian

Russian-language prompts from oasst2: Cyrillic takes two bytes per character.

Data: OpenAssistant oasst2 (Apache-2.0), human, language ru. Messages per size: 20 (2 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 171 B | ≈327 B (1.9×) | ≈389 B (2.3×) | ≈683 B (4.0×) | ≈414 B (2.4×) | ≈542 B (3.2×) | ≈934 B (5.5×) |
| medium (~75 chars) | 253 B | ≈410 B (1.6×) | ≈506 B (2.0×) | ≈799 B (3.2×) | ≈531 B (2.1×) | ≈650 B (2.6×) | ≈607 B (2.4×) |
| long (~329 chars) | 521 B | ≈677 B (1.3×) | ≈1.1 KB (2.1×) | ≈1.4 KB (2.7×) | ≈1.1 KB (2.2×) | ≈1.2 KB (2.4×) | ≈1.2 KB (2.3×) |
| very long (~3 305 chars) | 2.4 KB | ≈2.6 KB (1.1×) | ≈6.5 KB (2.7×) | ≈6.8 KB (2.8×) | ≈6.5 KB (2.7×) | ≈6.6 KB (2.7×) | ≈6.6 KB (2.7×) |
| new connection + first message | 772 B | ≈920 B (1.2×) | ≈4.4 KB (5.8×) | ≈4.3 KB (5.7×) | ≈4.4 KB (5.9×) | ≈4.4 KB (5.8×) | ≈12 KB (16.5×) |

**Time from sending a message to its ack** (median):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 2.5 ms | 2.4 ms | 2.4 ms | 2.4 ms | 2.5 ms | 2.5 ms | 2.5 ms |
| medium (~75 chars) | 2.5 ms | 2.5 ms | 2.4 ms | 2.5 ms | 2.5 ms | 2.5 ms | 2.5 ms |
| long (~329 chars) | 2.6 ms | 2.5 ms | 2.4 ms | 2.5 ms | 2.5 ms | 2.5 ms | 2.5 ms |
| very long (~3 305 chars) | 2.7 ms | 2.7 ms | 2.5 ms | 2.5 ms | 2.6 ms | 2.5 ms | 2.5 ms |
| new connection + first message | 3.2 ms | 3.2 ms | 8.2 ms | 5.5 ms | 5.8 ms | 8.0 ms | 6.1 ms |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 65 ms | 64 ms | 65 ms | 74 ms | 66 ms | 68 ms | 71 ms |
| medium (~75 chars) | 67 ms | 66 ms | 68 ms | 76 ms | 69 ms | 71 ms | 75 ms |
| long (~329 chars) | 72 ms | 71 ms | 77 ms | 84 ms | 77 ms | 80 ms | 85 ms |
| very long (~3 305 chars) | 125 ms | 122 ms | 180 ms | 189 ms | 183 ms | 183 ms | 223 ms |
| new connection + first message | 72 ms | 69 ms | 277 ms | 214 ms | 215 ms | 276 ms | 268 ms |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 81 ms | 81 ms | 81 ms | 82 ms | 82 ms | 82 ms | 82 ms |
| medium (~75 chars) | 81 ms | 81 ms | 81 ms | 82 ms | 82 ms | 82 ms | 82 ms |
| long (~329 chars) | 81 ms | 81 ms | 82 ms | 82 ms | 82 ms | 82 ms | 82 ms |
| very long (~3 305 chars) | 81 ms | 81 ms | 82 ms | 82 ms | 82 ms | 82 ms | 83 ms |
| new connection + first message | 81 ms | 81 ms | 242 ms | 163 ms | 164 ms | 246 ms | 164 ms |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 202 ms | 202 ms | 203 ms | 203 ms | 203 ms | 202 ms | 203 ms |
| medium (~75 chars) | 202 ms | 202 ms | 203 ms | 204 ms | 203 ms | 202 ms | 202 ms |
| long (~329 chars) | 203 ms | 203 ms | 203 ms | 204 ms | 203 ms | 202 ms | 202 ms |
| very long (~3 305 chars) | 207 ms | 206 ms | 212 ms | 212 ms | 212 ms | 210 ms | 214 ms |
| new connection + first message | 203 ms | 202 ms | 614 ms | 411 ms | 411 ms | 613 ms | 414 ms |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


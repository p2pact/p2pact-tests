## Human messages, Russian

Russian-language prompts from oasst2: Cyrillic takes two bytes per character.

Data: OpenAssistant oasst2 (Apache-2.0), human, language ru. Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 175 B | ≈330 B (1.9×) | ≈417 B (2.4×) | ≈710 B (4.1×) | ≈442 B (2.5×) | ≈555 B (3.2×) | ≈681 B (3.9×) |
| medium (~75 chars) | 196 B | ≈349 B (1.8×) | ≈508 B (2.6×) | ≈801 B (4.1×) | ≈533 B (2.7×) | ≈630 B (3.2×) | ≈627 B (3.2×) |
| long (~333 chars) | 397 B | ≈552 B (1.4×) | ≈1.1 KB (2.9×) | ≈1.4 KB (3.7×) | ≈1.2 KB (3.0×) | ≈1.3 KB (3.3×) | ≈1.3 KB (3.3×) |
| very long (~2 567 chars) | 2.0 KB | ≈2.1 KB (1.1×) | ≈6.4 KB (3.2×) | ≈6.7 KB (3.4×) | ≈6.5 KB (3.3×) | ≈6.5 KB (3.3×) | ≈6.6 KB (3.3×) |
| new connection + first message | 732 B | ≈961 B (1.3×) | ≈4.2 KB (5.9×) | ≈4.2 KB (5.8×) | ≈4.3 KB (6.0×) | ≈4.3 KB (6.1×) | ≈41 KB (57.9×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 2.7 ms | 2.8 ms (1.0×) | 2.7 ms (1.0×) | 2.8 ms (1.0×) | 2.8 ms (1.0×) | 2.7 ms (1.0×) | 2.7 ms (1.0×) |
| medium (~75 chars) | 2.8 ms | 2.9 ms (1.0×) | 2.7 ms (1.0×) | 2.8 ms (1.0×) | 2.8 ms (1.0×) | 2.7 ms (1.0×) | 2.6 ms (0.9×) |
| long (~333 chars) | 2.8 ms | 3.0 ms (1.1×) | 2.7 ms (0.9×) | 2.8 ms (1.0×) | 2.9 ms (1.0×) | 2.7 ms (1.0×) | 2.7 ms (0.9×) |
| very long (~2 567 chars) | 3.4 ms | 2.9 ms (0.9×) | 2.8 ms (0.8×) | 3.0 ms (0.9×) | 2.8 ms (0.8×) | 2.7 ms (0.8×) | 2.7 ms (0.8×) |
| new connection + first message | 5.0 ms | 5.2 ms (1.0×) | 8.5 ms (1.7×) | 7.1 ms (1.4×) | 6.0 ms (1.2×) | 10 ms (2.1×) | 7.5 ms (1.5×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 16507 (0 errors) | 15526 (0 errors) (1.1×) | 19859 (0 errors) (0.8×) | 16789 (0 errors) (1.0×) | 19767 (0 errors) (0.8×) | 17903 (0 errors) (0.9×) | 19082 (0 errors) (0.9×) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 67 ms | 66 ms (1.0×) | 67 ms (1.0×) | 74 ms (1.1×) | 67 ms (1.0×) | 69 ms (1.0×) | 73 ms (1.1×) |
| medium (~75 chars) | 68 ms | 65 ms (1.0×) | 69 ms (1.0×) | 76 ms (1.1×) | 69 ms (1.0×) | 71 ms (1.0×) | 75 ms (1.1×) |
| long (~333 chars) | 72 ms | 70 ms (1.0×) | 80 ms (1.1×) | 87 ms (1.2×) | 80 ms (1.1×) | 81 ms (1.1×) | 86 ms (1.2×) |
| very long (~2 567 chars) | 102 ms | 86 ms (0.8×) | 170 ms (1.7×) | 178 ms (1.8×) | 170 ms (1.7×) | 171 ms (1.7×) | 182 ms (1.8×) |
| new connection + first message | 79 ms | 79 ms (1.0×) | 281 ms (3.6×) | 217 ms (2.8×) | 217 ms (2.7×) | 278 ms (3.5×) | 272 ms (3.5×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 122 (0 errors) | 243 (0 errors) (0.5×) | 158 (0 errors) (0.8×) | 629 (0 errors) (0.2×) | 167 (0 errors) (0.7×) | 162 (0 errors) (0.8×) | 98 (0 errors) (1.2×) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 81 ms | 81 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) |
| medium (~75 chars) | 81 ms | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) |
| long (~333 chars) | 82 ms | 81 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) |
| very long (~2 567 chars) | 82 ms | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 81 ms (1.0×) |
| new connection + first message | 89 ms | 86 ms (1.0×) | 247 ms (2.8×) | 164 ms (1.8×) | 165 ms (1.8×) | 248 ms (2.8×) | 165 ms (1.9×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 517 (0 errors) | 612 (0 errors) (0.8×) | 613 (0 errors) (0.8×) | 595 (0 errors) (0.9×) | 610 (0 errors) (0.8×) | 612 (0 errors) (0.8×) | 563 (0 errors) (0.9×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 202 ms | 202 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) | 204 ms (1.0×) |
| medium (~75 chars) | 202 ms | 201 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) |
| long (~333 chars) | 203 ms | 203 ms (1.0×) | 204 ms (1.0×) | 204 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) |
| very long (~2 567 chars) | 204 ms | 205 ms (1.0×) | 211 ms (1.0×) | 211 ms (1.0×) | 210 ms (1.0×) | 211 ms (1.0×) | 211 ms (1.0×) |
| new connection + first message | 209 ms | 207 ms (1.0×) | 615 ms (2.9×) | 411 ms (2.0×) | 412 ms (2.0×) | 614 ms (2.9×) | 616 ms (2.9×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 151 (0 errors) | 246 (0 errors) (0.6×) | 246 (0 errors) (0.6×) | 240 (0 errors) (0.6×) | 246 (0 errors) (0.6×) | 242 (0 errors) (0.6×) | 187 (0 errors) (0.8×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


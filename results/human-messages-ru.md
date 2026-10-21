## Human messages, Russian

Russian-language prompts from oasst2: Cyrillic takes two bytes per character.

Data: OpenAssistant oasst2 (Apache-2.0), human, language ru. Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 176 B | ≈331 B (1.9×) | ≈417 B (2.4×) | ≈710 B (4.0×) | ≈442 B (2.5×) | ≈528 B (3.0×) | ≈925 B (5.3×) |
| medium (~75 chars) | 204 B | ≈358 B (1.8×) | ≈508 B (2.5×) | ≈801 B (3.9×) | ≈533 B (2.6×) | ≈613 B (3.0×) | ≈573 B (2.8×) |
| long (~333 chars) | 411 B | ≈566 B (1.4×) | ≈1.1 KB (2.8×) | ≈1.4 KB (3.5×) | ≈1.2 KB (2.9×) | ≈1.2 KB (3.1×) | ≈1.2 KB (3.0×) |
| very long (~2 567 chars) | 2.0 KB | ≈2.1 KB (1.1×) | ≈6.4 KB (3.2×) | ≈6.7 KB (3.4×) | ≈6.5 KB (3.3×) | ≈6.5 KB (3.3×) | ≈6.6 KB (3.4×) |
| new connection + first message | 867 B | ≈949 B (1.1×) | ≈4.2 KB (5.0×) | ≈4.2 KB (4.9×) | ≈4.3 KB (5.1×) | ≈4.3 KB (5.1×) | ≈47 KB (55.8×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 2.5 ms | 2.4 ms (1.0×) | 2.4 ms (1.0×) | 2.4 ms (1.0×) | 2.6 ms (1.0×) | 2.4 ms (1.0×) | 2.4 ms (1.0×) |
| medium (~75 chars) | 2.5 ms | 2.4 ms (1.0×) | 2.4 ms (1.0×) | 2.5 ms (1.0×) | 2.7 ms (1.1×) | 2.5 ms (1.0×) | 2.4 ms (1.0×) |
| long (~333 chars) | 2.5 ms | 2.5 ms (1.0×) | 2.4 ms (1.0×) | 2.6 ms (1.0×) | 2.5 ms (1.0×) | 2.4 ms (1.0×) | 2.4 ms (1.0×) |
| very long (~2 567 chars) | 2.6 ms | 2.6 ms (1.0×) | 2.4 ms (0.9×) | 2.7 ms (1.0×) | 2.5 ms (0.9×) | 2.4 ms (0.9×) | 2.4 ms (0.9×) |
| new connection + first message | 4.3 ms | 5.0 ms (1.2×) | 8.3 ms (1.9×) | 5.3 ms (1.2×) | 6.1 ms (1.4×) | 7.7 ms (1.8×) | 5.4 ms (1.3×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 20174 (0 errors) | 20455 (0 errors) (1.0×) | 21310 (0 errors) (0.9×) | 17548 (0 errors) (1.1×) | 21882 (0 errors) (0.9×) | 20655 (0 errors) (1.0×) | 21323 (0 errors) (0.9×) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 66 ms | 64 ms (1.0×) | 67 ms (1.0×) | 74 ms (1.1×) | 67 ms (1.0×) | 68 ms (1.0×) | 72 ms (1.1×) |
| medium (~75 chars) | 67 ms | 65 ms (1.0×) | 69 ms (1.0×) | 76 ms (1.1×) | 69 ms (1.0×) | 70 ms (1.1×) | 75 ms (1.1×) |
| long (~333 chars) | 71 ms | 69 ms (1.0×) | 79 ms (1.1×) | 87 ms (1.2×) | 81 ms (1.1×) | 81 ms (1.1×) | 85 ms (1.2×) |
| very long (~2 567 chars) | 101 ms | 98 ms (1.0×) | 170 ms (1.7×) | 178 ms (1.8×) | 171 ms (1.7×) | 171 ms (1.7×) | 180 ms (1.8×) |
| new connection + first message | 75 ms | 78 ms (1.0×) | 280 ms (3.8×) | 217 ms (2.9×) | 217 ms (2.9×) | 277 ms (3.7×) | 270 ms (3.6×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 169 (0 errors) | 348 (0 errors) (0.5×) | 158 (0 errors) (1.1×) | 628 (0 errors) (0.3×) | 167 (0 errors) (1.0×) | 163 (0 errors) (1.0×) | 99 (0 errors) (1.7×) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 81 ms | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) |
| medium (~75 chars) | 81 ms | 81 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) |
| long (~333 chars) | 81 ms | 82 ms (1.0×) | 82 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) |
| very long (~2 567 chars) | 82 ms | 81 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) |
| new connection + first message | 83 ms | 83 ms (1.0×) | 247 ms (3.0×) | 163 ms (2.0×) | 164 ms (2.0×) | 245 ms (3.0×) | 163 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 510 (0 errors) | 613 (0 errors) (0.8×) | 613 (0 errors) (0.8×) | 597 (0 errors) (0.9×) | 614 (0 errors) (0.8×) | 614 (0 errors) (0.8×) | 563 (0 errors) (0.9×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 201 ms | 202 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) |
| medium (~75 chars) | 202 ms | 202 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) |
| long (~333 chars) | 202 ms | 202 ms (1.0×) | 204 ms (1.0×) | 204 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) |
| very long (~2 567 chars) | 205 ms | 206 ms (1.0×) | 211 ms (1.0×) | 212 ms (1.0×) | 211 ms (1.0×) | 210 ms (1.0×) | 210 ms (1.0×) |
| new connection + first message | 208 ms | 209 ms (1.0×) | 614 ms (3.0×) | 412 ms (2.0×) | 411 ms (2.0×) | 613 ms (2.9×) | 416 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 208 (0 errors) | 246 (0 errors) (0.8×) | 246 (0 errors) (0.8×) | 240 (0 errors) (0.9×) | 247 (0 errors) (0.8×) | 245 (0 errors) (0.9×) | 178 (0 errors) (1.2×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


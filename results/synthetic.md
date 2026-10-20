## Generated text

Reproducible offline: generated text with natural-language word frequencies, four network profiles.

Data: generated text (Zipf-distributed words). Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 166 B | ≈318 B (1.9×) | ≈387 B (2.3×) | ≈681 B (4.1×) | ≈411 B (2.5×) | ≈498 B (3.0×) | ≈603 B (3.6×) |
| medium (~100 chars) | 233 B | ≈385 B (1.7×) | ≈469 B (2.0×) | ≈762 B (3.3×) | ≈490 B (2.1×) | ≈609 B (2.6×) | ≈566 B (2.4×) |
| long (~600 chars) | 496 B | ≈648 B (1.3×) | ≈969 B (2.0×) | ≈1.2 KB (2.5×) | ≈995 B (2.0×) | ≈1.1 KB (2.2×) | ≈1.0 KB (2.2×) |
| very long (~5 000 chars) | 2.4 KB | ≈2.6 KB (1.1×) | ≈5.4 KB (2.2×) | ≈5.8 KB (2.4×) | ≈5.5 KB (2.3×) | ≈5.6 KB (2.3×) | ≈5.8 KB (2.4×) |
| huge (~200 000 chars) | 77 KB | ≈78 KB (1.0×) | ≈206 KB (2.7×) | ≈207 KB (2.7×) | ≈207 KB (2.7×) | ≈207 KB (2.7×) | ≈222 KB (2.9×) |
| new connection + first message | 753 B | ≈933 B (1.2×) | ≈4.2 KB (5.7×) | ≈4.1 KB (5.6×) | ≈4.3 KB (5.8×) | ≈4.3 KB (5.9×) | ≈48 KB (65.7×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 2.5 ms | 2.6 ms | 2.4 ms | 2.5 ms | 2.6 ms | 2.5 ms | 2.5 ms |
| medium (~100 chars) | 2.5 ms | 2.6 ms | 2.5 ms | 2.6 ms | 2.7 ms | 2.5 ms | 2.5 ms |
| long (~600 chars) | 2.6 ms | 2.7 ms | 2.5 ms | 2.7 ms | 2.7 ms | 2.5 ms | 2.5 ms |
| very long (~5 000 chars) | 2.8 ms | 2.9 ms | 2.5 ms | 2.5 ms | 2.5 ms | 2.5 ms | 2.4 ms |
| huge (~200 000 chars) | 6.2 ms | 8.7 ms | 6.4 ms | 5.3 ms | 6.8 ms | 6.7 ms | 6.8 ms |
| new connection + first message | 4.7 ms | 5.8 ms | 8.6 ms | 6.1 ms | 6.3 ms | 7.8 ms | 6.8 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 21059 (0) | 19663 (0) | 20694 (0) | 19343 (0) | 21795 (0) | 18960 (0) | 20988 (0) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 65 ms | 64 ms | 66 ms | 73 ms | 67 ms | 68 ms | 73 ms |
| medium (~100 chars) | 67 ms | 66 ms | 68 ms | 75 ms | 69 ms | 70 ms | 74 ms |
| long (~600 chars) | 73 ms | 72 ms | 80 ms | 88 ms | 81 ms | 83 ms | 86 ms |
| very long (~5 000 chars) | 122 ms | 120 ms | 188 ms | 196 ms | 189 ms | 190 ms | 200 ms |
| huge (~200 000 chars) | 1.96 s | 1.92 s | 4.96 s | 4.97 s | 4.96 s | 4.96 s | 5.17 s |
| new connection + first message | 74 ms | 72 ms | 280 ms | 216 ms | 215 ms | 276 ms | 271 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 191 (0) | 309 (0) | 202 (0) | 632 (0) | 223 (0) | 211 (0) | 118 (0) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 81 ms | 81 ms | 82 ms | 82 ms | 82 ms | 81 ms | 82 ms |
| medium (~100 chars) | 81 ms | 81 ms | 81 ms | 82 ms | 81 ms | 81 ms | 82 ms |
| long (~600 chars) | 81 ms | 82 ms | 82 ms | 81 ms | 81 ms | 81 ms | 82 ms |
| very long (~5 000 chars) | 82 ms | 82 ms | 82 ms | 82 ms | 82 ms | 82 ms | 82 ms |
| huge (~200 000 chars) | 180 ms | 91 ms | 84 ms | 85 ms | 84 ms | 83 ms | 601 ms |
| new connection + first message | 83 ms | 83 ms | 247 ms | 164 ms | 164 ms | 247 ms | 164 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 560 (0) | 613 (0) | 613 (0) | 596 (0) | 611 (0) | 612 (0) | 553 (0) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 202 ms | 202 ms | 202 ms | 203 ms | 202 ms | 202 ms | 203 ms |
| medium (~100 chars) | 202 ms | 202 ms | 202 ms | 203 ms | 202 ms | 202 ms | 203 ms |
| long (~600 chars) | 203 ms | 203 ms | 203 ms | 204 ms | 203 ms | 202 ms | 203 ms |
| very long (~5 000 chars) | 207 ms | 207 ms | 212 ms | 213 ms | 212 ms | 212 ms | 212 ms |
| huge (~200 000 chars) | 722 ms | 356 ms | 597 ms | 597 ms | 597 ms | 597 ms | 3.07 s |
| new connection + first message | 206 ms | 205 ms | 614 ms | 412 ms | 410 ms | 612 ms | 416 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 137 (0) | 247 (0) | 245 (0) | 240 (0) | 246 (0) | 244 (0) | 181 (0) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


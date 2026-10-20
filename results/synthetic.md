## Generated text

Reproducible offline: generated text with natural-language word frequencies, four network profiles.

Data: generated text (Zipf-distributed words). Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 159 B | ≈314 B (2.0×) | ≈387 B (2.4×) | ≈681 B (4.3×) | ≈410 B (2.6×) | ≈503 B (3.2×) | ≈468 B (2.9×) |
| medium (~100 chars) | 227 B | ≈380 B (1.7×) | ≈469 B (2.1×) | ≈762 B (3.4×) | ≈490 B (2.2×) | ≈593 B (2.6×) | ≈581 B (2.6×) |
| long (~600 chars) | 496 B | ≈648 B (1.3×) | ≈969 B (2.0×) | ≈1.2 KB (2.5×) | ≈996 B (2.0×) | ≈1.1 KB (2.2×) | ≈1.0 KB (2.1×) |
| very long (~5 000 chars) | 2.5 KB | ≈2.6 KB (1.0×) | ≈5.4 KB (2.2×) | ≈5.8 KB (2.3×) | ≈5.5 KB (2.2×) | ≈5.6 KB (2.3×) | ≈5.8 KB (2.3×) |
| huge (~200 000 chars) | 77 KB | ≈78 KB (1.0×) | ≈206 KB (2.7×) | ≈207 KB (2.7×) | ≈207 KB (2.7×) | ≈207 KB (2.7×) | ≈211 KB (2.7×) |
| new connection + first message | 753 B | ≈933 B (1.2×) | ≈4.2 KB (5.7×) | ≈4.1 KB (5.6×) | ≈4.3 KB (5.8×) | ≈4.3 KB (5.9×) | ≈53 KB (71.8×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 2.8 ms | 2.7 ms (1.0×) | 2.7 ms (1.0×) | 2.8 ms (1.0×) | 2.9 ms (1.0×) | 2.8 ms (1.0×) | 2.8 ms (1.0×) |
| medium (~100 chars) | 2.8 ms | 2.8 ms (1.0×) | 2.8 ms (1.0×) | 2.8 ms (1.0×) | 2.8 ms (1.0×) | 2.7 ms (1.0×) | 2.7 ms (1.0×) |
| long (~600 chars) | 2.9 ms | 2.8 ms (1.0×) | 2.7 ms (0.9×) | 2.8 ms (1.0×) | 2.9 ms (1.0×) | 2.7 ms (0.9×) | 2.7 ms (0.9×) |
| very long (~5 000 chars) | 3.5 ms | 3.5 ms (1.0×) | 2.8 ms (0.8×) | 2.8 ms (0.8×) | 2.9 ms (0.8×) | 2.8 ms (0.8×) | 2.7 ms (0.8×) |
| huge (~200 000 chars) | 9.5 ms | 8.3 ms (0.9×) | 6.5 ms (0.7×) | 7.6 ms (0.8×) | 7.4 ms (0.8×) | 5.6 ms (0.6×) | 5.9 ms (0.6×) |
| new connection + first message | 4.6 ms | 4.7 ms (1.0×) | 8.9 ms (2.0×) | 6.6 ms (1.4×) | 5.7 ms (1.3×) | 9.1 ms (2.0×) | 7.3 ms (1.6×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 17823 (0 errors) | 18260 (0 errors) (1.0×) | 19135 (0 errors) (0.9×) | 17338 (0 errors) (1.0×) | 20524 (0 errors) (0.9×) | 17376 (0 errors) (1.0×) | 20303 (0 errors) (0.9×) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 66 ms | 66 ms (1.0×) | 68 ms (1.0×) | 75 ms (1.1×) | 66 ms (1.0×) | 68 ms (1.0×) | 73 ms (1.1×) |
| medium (~100 chars) | 68 ms | 67 ms (1.0×) | 69 ms (1.0×) | 77 ms (1.1×) | 69 ms (1.0×) | 71 ms (1.0×) | 76 ms (1.1×) |
| long (~600 chars) | 75 ms | 73 ms (1.0×) | 82 ms (1.1×) | 89 ms (1.2×) | 82 ms (1.1×) | 84 ms (1.1×) | 88 ms (1.2×) |
| very long (~5 000 chars) | 122 ms | 122 ms (1.0×) | 189 ms (1.5×) | 197 ms (1.6×) | 189 ms (1.6×) | 191 ms (1.6×) | 201 ms (1.6×) |
| huge (~200 000 chars) | 1.96 s | 1.92 s (1.0×) | 4.96 s (2.5×) | 4.97 s (2.5×) | 4.97 s (2.5×) | 4.97 s (2.5×) | 5.17 s (2.6×) |
| new connection + first message | 76 ms | 74 ms (1.0×) | 284 ms (3.7×) | 220 ms (2.9×) | 217 ms (2.8×) | 279 ms (3.7×) | 273 ms (3.6×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 76 (0 errors) | 309 (0 errors) (0.2×) | 202 (0 errors) (0.4×) | 621 (0 errors) (0.1×) | 223 (0 errors) (0.3×) | 208 (0 errors) (0.4×) | 119 (0 errors) (0.6×) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 82 ms | 82 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) |
| medium (~100 chars) | 82 ms | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| long (~600 chars) | 82 ms | 83 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| very long (~5 000 chars) | 83 ms | 84 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| huge (~200 000 chars) | 95 ms | 91 ms (1.0×) | 89 ms (0.9×) | 90 ms (0.9×) | 88 ms (0.9×) | 88 ms (0.9×) | 756 ms (7.9×) |
| new connection + first message | 84 ms | 84 ms (1.0×) | 250 ms (3.0×) | 168 ms (2.0×) | 167 ms (2.0×) | 247 ms (3.0×) | 168 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 534 (0 errors) | 605 (0 errors) (0.9×) | 605 (0 errors) (0.9×) | 590 (0 errors) (0.9×) | 604 (0 errors) (0.9×) | 606 (0 errors) (0.9×) | 548 (0 errors) (1.0×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 203 ms | 203 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 205 ms (1.0×) |
| medium (~100 chars) | 203 ms | 203 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) | 204 ms (1.0×) |
| long (~600 chars) | 204 ms | 203 ms (1.0×) | 205 ms (1.0×) | 205 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) |
| very long (~5 000 chars) | 208 ms | 208 ms (1.0×) | 213 ms (1.0×) | 215 ms (1.0×) | 214 ms (1.0×) | 214 ms (1.0×) | 212 ms (1.0×) |
| huge (~200 000 chars) | 653 ms | 357 ms (0.5×) | 600 ms (0.9×) | 601 ms (0.9×) | 600 ms (0.9×) | 599 ms (0.9×) | 2.75 s (4.2×) |
| new connection + first message | 205 ms | 205 ms (1.0×) | 616 ms (3.0×) | 412 ms (2.0×) | 413 ms (2.0×) | 615 ms (3.0×) | 418 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 167 (0 errors) | 245 (0 errors) (0.7×) | 245 (0 errors) (0.7×) | 239 (0 errors) (0.7×) | 246 (0 errors) (0.7×) | 241 (0 errors) (0.7×) | 180 (0 errors) (0.9×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


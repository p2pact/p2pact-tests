## Generated text

Reproducible offline: generated text with natural-language word frequencies, four network profiles.

Data: generated text (Zipf-distributed words). Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 159 B | ≈314 B (2.0×) | ≈387 B (2.4×) | ≈681 B (4.3×) | ≈410 B (2.6×) | ≈527 B (3.3×) | ≈490 B (3.1×) |
| medium (~100 chars) | 208 B | ≈355 B (1.7×) | ≈469 B (2.3×) | ≈762 B (3.7×) | ≈490 B (2.4×) | ≈602 B (2.9×) | ≈567 B (2.7×) |
| long (~600 chars) | 408 B | ≈556 B (1.4×) | ≈969 B (2.4×) | ≈1.2 KB (3.1×) | ≈996 B (2.4×) | ≈1.1 KB (2.7×) | ≈1.1 KB (2.7×) |
| very long (~5 000 chars) | 2.0 KB | ≈2.2 KB (1.1×) | ≈5.4 KB (2.7×) | ≈5.8 KB (2.9×) | ≈5.5 KB (2.7×) | ≈5.6 KB (2.8×) | ≈5.9 KB (2.9×) |
| huge (~200 000 chars) | 74 KB | ≈76 KB (1.0×) | ≈206 KB (2.8×) | ≈207 KB (2.8×) | ≈207 KB (2.8×) | ≈207 KB (2.8×) | ≈214 KB (2.9×) |
| new connection + first message | 722 B | ≈900 B (1.2×) | ≈4.2 KB (6.0×) | ≈4.1 KB (5.9×) | ≈4.3 KB (6.0×) | ≈4.3 KB (6.1×) | ≈54 KB (77.2×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 2.8 ms | 2.8 ms (1.0×) | 2.7 ms (1.0×) | 2.8 ms (1.0×) | 2.8 ms (1.0×) | 2.7 ms (1.0×) | 2.7 ms (1.0×) |
| medium (~100 chars) | 2.8 ms | 2.9 ms (1.0×) | 2.7 ms (1.0×) | 2.9 ms (1.0×) | 2.9 ms (1.0×) | 2.7 ms (1.0×) | 2.7 ms (1.0×) |
| long (~600 chars) | 2.9 ms | 2.9 ms (1.0×) | 2.8 ms (1.0×) | 2.8 ms (1.0×) | 2.9 ms (1.0×) | 2.7 ms (0.9×) | 2.7 ms (0.9×) |
| very long (~5 000 chars) | 3.6 ms | 3.4 ms (0.9×) | 3.0 ms (0.8×) | 3.0 ms (0.8×) | 3.0 ms (0.8×) | 2.8 ms (0.8×) | 2.7 ms (0.8×) |
| huge (~200 000 chars) | 9.6 ms | 12 ms (1.2×) | 9.0 ms (0.9×) | 9.4 ms (1.0×) | 9.4 ms (1.0×) | 8.0 ms (0.8×) | 9.1 ms (1.0×) |
| new connection + first message | 9.6 ms | 7.2 ms (0.8×) | 9.7 ms (1.0×) | 6.3 ms (0.7×) | 6.6 ms (0.7×) | 10 ms (1.1×) | 7.6 ms (0.8×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 18764 (0 errors) | 19204 (0 errors) (1.0×) | 20262 (0 errors) (0.9×) | 17928 (0 errors) (1.0×) | 20085 (0 errors) (0.9×) | 17858 (0 errors) (1.1×) | 18807 (0 errors) (1.0×) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 66 ms | 66 ms (1.0×) | 67 ms (1.0×) | 75 ms (1.1×) | 66 ms (1.0×) | 68 ms (1.0×) | 73 ms (1.1×) |
| medium (~100 chars) | 67 ms | 67 ms (1.0×) | 69 ms (1.0×) | 76 ms (1.1×) | 68 ms (1.0×) | 70 ms (1.1×) | 75 ms (1.1×) |
| long (~600 chars) | 72 ms | 72 ms (1.0×) | 81 ms (1.1×) | 88 ms (1.2×) | 82 ms (1.1×) | 83 ms (1.2×) | 88 ms (1.2×) |
| very long (~5 000 chars) | 113 ms | 111 ms (1.0×) | 189 ms (1.7×) | 197 ms (1.8×) | 190 ms (1.7×) | 191 ms (1.7×) | 201 ms (1.8×) |
| huge (~200 000 chars) | 1.91 s | 1.87 s (1.0×) | 4.96 s (2.6×) | 4.97 s (2.6×) | 4.97 s (2.6×) | 4.97 s (2.6×) | 5.18 s (2.7×) |
| new connection + first message | 79 ms | 77 ms (1.0×) | 284 ms (3.6×) | 219 ms (2.8×) | 217 ms (2.8×) | 279 ms (3.6×) | 273 ms (3.5×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 157 (0 errors) | 388 (0 errors) (0.4×) | 202 (0 errors) (0.8×) | 622 (0 errors) (0.3×) | 223 (0 errors) (0.7×) | 209 (0 errors) (0.8×) | 117 (0 errors) (1.3×) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 81 ms | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) |
| medium (~100 chars) | 81 ms | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) |
| long (~600 chars) | 82 ms | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) |
| very long (~5 000 chars) | 83 ms | 83 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) |
| huge (~200 000 chars) | 180 ms | 96 ms (0.5×) | 92 ms (0.5×) | 92 ms (0.5×) | 89 ms (0.5×) | 89 ms (0.5×) | 829 ms (4.6×) |
| new connection + first message | 89 ms | 89 ms (1.0×) | 249 ms (2.8×) | 167 ms (1.9×) | 168 ms (1.9×) | 246 ms (2.8×) | 166 ms (1.9×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 519 (0 errors) | 606 (0 errors) (0.9×) | 605 (0 errors) (0.9×) | 589 (0 errors) (0.9×) | 606 (0 errors) (0.9×) | 610 (0 errors) (0.9×) | 552 (0 errors) (0.9×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 201 ms | 202 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) |
| medium (~100 chars) | 202 ms | 202 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) |
| long (~600 chars) | 202 ms | 203 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) |
| very long (~5 000 chars) | 206 ms | 207 ms (1.0×) | 212 ms (1.0×) | 213 ms (1.0×) | 213 ms (1.0×) | 212 ms (1.0×) | 212 ms (1.0×) |
| huge (~200 000 chars) | 678 ms | 356 ms (0.5×) | 602 ms (0.9×) | 601 ms (0.9×) | 601 ms (0.9×) | 600 ms (0.9×) | 2.75 s (4.1×) |
| new connection + first message | 208 ms | 209 ms (1.0×) | 615 ms (3.0×) | 411 ms (2.0×) | 411 ms (2.0×) | 614 ms (3.0×) | 418 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 177 (0 errors) | 246 (0 errors) (0.7×) | 246 (0 errors) (0.7×) | 240 (0 errors) (0.7×) | 247 (0 errors) (0.7×) | 245 (0 errors) (0.7×) | 200 (0 errors) (0.9×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


## Generated text

Reproducible offline: generated text with natural-language word frequencies, four network profiles.

Data: generated text (Zipf-distributed words). Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 159 B | ≈314 B (2.0×) | ≈387 B (2.4×) | ≈681 B (4.3×) | ≈411 B (2.6×) | ≈522 B (3.3×) | ≈805 B (5.1×) |
| medium (~100 chars) | 200 B | ≈354 B (1.8×) | ≈469 B (2.3×) | ≈762 B (3.8×) | ≈490 B (2.5×) | ≈597 B (3.0×) | ≈587 B (2.9×) |
| long (~600 chars) | 395 B | ≈550 B (1.4×) | ≈969 B (2.5×) | ≈1.2 KB (3.2×) | ≈995 B (2.5×) | ≈1.1 KB (2.8×) | ≈1.1 KB (2.8×) |
| very long (~5 000 chars) | 2.0 KB | ≈2.2 KB (1.1×) | ≈5.4 KB (2.7×) | ≈5.8 KB (2.8×) | ≈5.5 KB (2.7×) | ≈5.6 KB (2.8×) | ≈5.9 KB (2.9×) |
| huge (~200 000 chars) | 74 KB | ≈76 KB (1.0×) | ≈206 KB (2.8×) | ≈207 KB (2.8×) | ≈207 KB (2.8×) | ≈207 KB (2.8×) | ≈209 KB (2.8×) |
| new connection + first message | 756 B | ≈934 B (1.2×) | ≈4.2 KB (5.7×) | ≈4.1 KB (5.6×) | ≈4.3 KB (5.8×) | ≈4.3 KB (5.8×) | ≈44 KB (60.1×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 2.5 ms | 2.5 ms (1.0×) | 2.4 ms (1.0×) | 2.4 ms (1.0×) | 2.4 ms (0.9×) | 2.4 ms (0.9×) | 2.3 ms (0.9×) |
| medium (~100 chars) | 2.5 ms | 2.6 ms (1.0×) | 2.4 ms (1.0×) | 2.4 ms (0.9×) | 2.4 ms (1.0×) | 2.4 ms (0.9×) | 2.3 ms (0.9×) |
| long (~600 chars) | 2.9 ms | 2.8 ms (1.0×) | 2.4 ms (0.8×) | 2.4 ms (0.8×) | 2.4 ms (0.8×) | 2.4 ms (0.8×) | 2.3 ms (0.8×) |
| very long (~5 000 chars) | 3.5 ms | 3.5 ms (1.0×) | 2.7 ms (0.8×) | 2.5 ms (0.7×) | 2.4 ms (0.7×) | 2.4 ms (0.7×) | 2.4 ms (0.7×) |
| huge (~200 000 chars) | 11 ms | 8.0 ms (0.7×) | 9.3 ms (0.9×) | 4.9 ms (0.5×) | 4.9 ms (0.5×) | 4.5 ms (0.4×) | 5.5 ms (0.5×) |
| new connection + first message | 9.4 ms | 4.4 ms (0.5×) | 8.4 ms (0.9×) | 6.1 ms (0.7×) | 5.1 ms (0.5×) | 7.6 ms (0.8×) | 5.2 ms (0.6×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 14169 (0 errors) | 16451 (0 errors) (0.9×) | 21163 (0 errors) (0.7×) | 20799 (0 errors) (0.7×) | 22038 (0 errors) (0.6×) | 21163 (0 errors) (0.7×) | 21662 (0 errors) (0.7×) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 65 ms | 64 ms (1.0×) | 66 ms (1.0×) | 73 ms (1.1×) | 66 ms (1.0×) | 68 ms (1.0×) | 73 ms (1.1×) |
| medium (~100 chars) | 65 ms | 64 ms (1.0×) | 68 ms (1.1×) | 75 ms (1.2×) | 68 ms (1.1×) | 70 ms (1.1×) | 75 ms (1.2×) |
| long (~600 chars) | 71 ms | 65 ms (0.9×) | 80 ms (1.1×) | 88 ms (1.2×) | 81 ms (1.1×) | 82 ms (1.2×) | 88 ms (1.2×) |
| very long (~5 000 chars) | 111 ms | 104 ms (0.9×) | 188 ms (1.7×) | 196 ms (1.8×) | 189 ms (1.7×) | 190 ms (1.7×) | 201 ms (1.8×) |
| huge (~200 000 chars) | 1.90 s | 1.86 s (1.0×) | 4.96 s (2.6×) | 4.97 s (2.6×) | 4.97 s (2.6×) | 4.96 s (2.6×) | 5.17 s (2.7×) |
| new connection + first message | 81 ms | 79 ms (1.0×) | 280 ms (3.4×) | 217 ms (2.7×) | 216 ms (2.7×) | 275 ms (3.4×) | 271 ms (3.3×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 173 (0 errors) | 388 (0 errors) (0.4×) | 202 (0 errors) (0.9×) | 631 (0 errors) (0.3×) | 223 (0 errors) (0.8×) | 209 (0 errors) (0.8×) | 115 (0 errors) (1.5×) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 81 ms | 82 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) |
| medium (~100 chars) | 81 ms | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) |
| long (~600 chars) | 82 ms | 82 ms (1.0×) | 82 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) |
| very long (~5 000 chars) | 82 ms | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) |
| huge (~200 000 chars) | 180 ms | 86 ms (0.5×) | 84 ms (0.5×) | 84 ms (0.5×) | 90 ms (0.5×) | 87 ms (0.5×) | 686 ms (3.8×) |
| new connection + first message | 83 ms | 90 ms (1.1×) | 247 ms (3.0×) | 164 ms (2.0×) | 165 ms (2.0×) | 247 ms (3.0×) | 165 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 545 (0 errors) | 614 (0 errors) (0.9×) | 611 (0 errors) (0.9×) | 597 (0 errors) (0.9×) | 611 (0 errors) (0.9×) | 611 (0 errors) (0.9×) | 551 (0 errors) (1.0×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 201 ms | 202 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) |
| medium (~100 chars) | 202 ms | 202 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) |
| long (~600 chars) | 203 ms | 202 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) |
| very long (~5 000 chars) | 205 ms | 205 ms (1.0×) | 212 ms (1.0×) | 213 ms (1.0×) | 212 ms (1.0×) | 212 ms (1.0×) | 212 ms (1.0×) |
| huge (~200 000 chars) | 705 ms | 351 ms (0.5×) | 601 ms (0.9×) | 599 ms (0.8×) | 599 ms (0.8×) | 599 ms (0.8×) | 3.12 s (4.4×) |
| new connection + first message | 207 ms | 210 ms (1.0×) | 614 ms (3.0×) | 412 ms (2.0×) | 411 ms (2.0×) | 612 ms (3.0×) | 615 ms (3.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 231 (0 errors) | 247 (0 errors) (0.9×) | 245 (0 errors) (0.9×) | 240 (0 errors) (1.0×) | 247 (0 errors) (0.9×) | 243 (0 errors) (1.0×) | 183 (0 errors) (1.3×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


## Generated text

Reproducible offline: generated text with natural-language word frequencies, four network profiles.

Data: generated text (Zipf-distributed words). Messages per size: 20 (2 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 145 B | ≈301 B (2.1×) | ≈364 B (2.5×) | ≈658 B (4.5×) | ≈387 B (2.7×) | ≈495 B (3.4×) | ≈1.1 KB (8.0×) |
| medium (~100 chars) | 213 B | ≈370 B (1.7×) | ≈446 B (2.1×) | ≈739 B (3.5×) | ≈467 B (2.2×) | ≈590 B (2.8×) | ≈575 B (2.7×) |
| long (~600 chars) | 476 B | ≈632 B (1.3×) | ≈946 B (2.0×) | ≈1.2 KB (2.6×) | ≈972 B (2.0×) | ≈1.1 KB (2.3×) | ≈1.1 KB (2.3×) |
| very long (~5 000 chars) | 2.4 KB | ≈2.5 KB (1.1×) | ≈5.4 KB (2.3×) | ≈5.7 KB (2.4×) | ≈5.5 KB (2.3×) | ≈5.6 KB (2.4×) | ≈5.7 KB (2.4×) |
| huge (~200 000 chars) | 77 KB | ≈78 KB (1.0×) | ≈206 KB (2.7×) | ≈207 KB (2.7×) | ≈207 KB (2.7×) | ≈207 KB (2.7×) | ≈230 KB (3.0×) |
| new connection + first message | 756 B | ≈904 B (1.2×) | ≈4.4 KB (5.9×) | ≈4.3 KB (5.8×) | ≈4.4 KB (6.0×) | ≈4.3 KB (5.9×) | ≈15 KB (21.0×) |

**Time from sending a message to its ack** (median):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 2.4 ms | 2.4 ms | 2.4 ms | 2.4 ms | 2.5 ms | 2.5 ms | 2.5 ms |
| medium (~100 chars) | 2.5 ms | 2.5 ms | 2.4 ms | 2.5 ms | 2.4 ms | 2.5 ms | 2.5 ms |
| long (~600 chars) | 2.5 ms | 2.5 ms | 2.4 ms | 2.5 ms | 2.6 ms | 2.4 ms | 2.5 ms |
| very long (~5 000 chars) | 3.0 ms | 2.9 ms | 2.5 ms | 2.6 ms | 2.7 ms | 2.4 ms | 2.4 ms |
| huge (~200 000 chars) | 21 ms | 19 ms | 5.1 ms | 5.6 ms | 6.4 ms | 11 ms | 11 ms |
| new connection + first message | 3.8 ms | 3.6 ms | 7.8 ms | 5.4 ms | 5.7 ms | 8.3 ms | 5.6 ms |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 65 ms | 64 ms | 66 ms | 73 ms | 66 ms | 67 ms | 72 ms |
| medium (~100 chars) | 67 ms | 65 ms | 68 ms | 75 ms | 68 ms | 70 ms | 74 ms |
| long (~600 chars) | 74 ms | 72 ms | 80 ms | 88 ms | 81 ms | 83 ms | 87 ms |
| very long (~5 000 chars) | 121 ms | 119 ms | 188 ms | 196 ms | 189 ms | 190 ms | 200 ms |
| huge (~200 000 chars) | 2.02 s | 1.97 s | 4.96 s | 4.97 s | 4.97 s | 5.03 s | 5.21 s |
| new connection + first message | 73 ms | 70 ms | 281 ms | 215 ms | 215 ms | 276 ms | 270 ms |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 82 ms | 81 ms | 82 ms | 82 ms | 82 ms | 82 ms | 82 ms |
| medium (~100 chars) | 82 ms | 82 ms | 82 ms | 82 ms | 82 ms | 82 ms | 82 ms |
| long (~600 chars) | 82 ms | 82 ms | 82 ms | 82 ms | 83 ms | 82 ms | 82 ms |
| very long (~5 000 chars) | 82 ms | 82 ms | 82 ms | 82 ms | 83 ms | 82 ms | 82 ms |
| huge (~200 000 chars) | 182 ms | 181 ms | 85 ms | 85 ms | 87 ms | 249 ms | 1.07 s |
| new connection + first message | 83 ms | 82 ms | 247 ms | 164 ms | 165 ms | 246 ms | 164 ms |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 202 ms | 202 ms | 203 ms | 203 ms | 203 ms | 203 ms | 203 ms |
| medium (~100 chars) | 202 ms | 202 ms | 203 ms | 204 ms | 203 ms | 203 ms | 203 ms |
| long (~600 chars) | 203 ms | 203 ms | 204 ms | 205 ms | 203 ms | 203 ms | 203 ms |
| very long (~5 000 chars) | 207 ms | 207 ms | 212 ms | 213 ms | 212 ms | 213 ms | 212 ms |
| huge (~200 000 chars) | 674 ms | 354 ms | 598 ms | 598 ms | 596 ms | 905 ms | 4.10 s |
| new connection + first message | 202 ms | 203 ms | 615 ms | 412 ms | 411 ms | 614 ms | 415 ms |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


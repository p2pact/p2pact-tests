## Coding agent output

A coding agent at work: tool outputs (files, logs, diffs) from SWE-Gym OpenHands trajectories.

Data: SWE-Gym OpenHands trajectories (MIT), tool messages. Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 288 B | ≈443 B (1.5×) | ≈547 B (1.9×) | ≈840 B (2.9×) | ≈570 B (2.0×) | ≈650 B (2.3×) | ≈639 B (2.2×) |
| long (~660 chars) | 598 B | ≈752 B (1.3×) | ≈1.3 KB (2.2×) | ≈1.6 KB (2.7×) | ≈1.3 KB (2.3×) | ≈1.4 KB (2.4×) | ≈1.5 KB (2.5×) |
| very long (~5 302 chars) | 2.5 KB | ≈2.6 KB (1.1×) | ≈8.3 KB (3.4×) | ≈8.6 KB (3.5×) | ≈8.4 KB (3.4×) | ≈8.1 KB (3.2×) | ≈8.6 KB (3.5×) |
| huge (~23 355 chars) | 6.3 KB | ≈6.5 KB (1.0×) | ≈29 KB (4.7×) | ≈29 KB (4.7×) | ≈29 KB (4.7×) | ≈27 KB (4.4×) | ≈30 KB (4.7×) |
| new connection + first message | 830 B | ≈1.0 KB (1.3×) | ≈4.4 KB (5.4×) | ≈4.3 KB (5.3×) | ≈4.4 KB (5.5×) | ≈4.5 KB (5.5×) | ≈39 KB (47.7×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 2.7 ms | 2.8 ms (1.0×) | 2.7 ms (1.0×) | 2.8 ms (1.0×) | 2.8 ms (1.0×) | 2.7 ms (1.0×) | 2.7 ms (1.0×) |
| long (~660 chars) | 2.7 ms | 2.8 ms (1.0×) | 2.8 ms (1.0×) | 2.7 ms (1.0×) | 2.8 ms (1.0×) | 2.7 ms (1.0×) | 2.7 ms (1.0×) |
| very long (~5 302 chars) | 3.0 ms | 3.0 ms (1.0×) | 2.8 ms (0.9×) | 2.7 ms (0.9×) | 2.8 ms (0.9×) | 2.7 ms (0.9×) | 2.7 ms (0.9×) |
| huge (~23 355 chars) | 4.1 ms | 3.5 ms (0.9×) | 3.0 ms (0.7×) | 3.1 ms (0.8×) | 3.0 ms (0.7×) | 2.7 ms (0.7×) | 2.7 ms (0.7×) |
| new connection + first message | 4.3 ms | 5.3 ms (1.2×) | 8.5 ms (2.0×) | 6.1 ms (1.4×) | 6.5 ms (1.5×) | 9.1 ms (2.1×) | 7.0 ms (1.6×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 20207 (0 errors) | 19673 (0 errors) (1.0×) | 19283 (0 errors) (1.0×) | 17697 (0 errors) (1.1×) | 20342 (0 errors) (1.0×) | 18623 (0 errors) (1.1×) | 18544 (0 errors) (1.1×) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 69 ms | 69 ms (1.0×) | 72 ms (1.0×) | 79 ms (1.1×) | 71 ms (1.0×) | 73 ms (1.1×) | 77 ms (1.1×) |
| long (~660 chars) | 79 ms | 76 ms (1.0×) | 90 ms (1.1×) | 98 ms (1.2×) | 92 ms (1.2×) | 92 ms (1.2×) | 94 ms (1.2×) |
| very long (~5 302 chars) | 107 ms | 103 ms (1.0×) | 190 ms (1.8×) | 198 ms (1.9×) | 191 ms (1.8×) | 185 ms (1.7×) | 200 ms (1.9×) |
| huge (~23 355 chars) | 212 ms | 207 ms (1.0×) | 855 ms (4.0×) | 862 ms (4.1×) | 857 ms (4.0×) | 805 ms (3.8×) | 894 ms (4.2×) |
| new connection + first message | 78 ms | 77 ms (1.0×) | 287 ms (3.7×) | 222 ms (2.8×) | 222 ms (2.8×) | 284 ms (3.6×) | 276 ms (3.5×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 91 (0 errors) | 109 (0 errors) (0.8×) | 159 (0 errors) (0.6×) | 622 (0 errors) (0.1×) | 170 (0 errors) (0.5×) | 165 (0 errors) (0.6×) | 102 (0 errors) (0.9×) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 83 ms | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) |
| long (~660 chars) | 83 ms | 82 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| very long (~5 302 chars) | 83 ms | 83 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| huge (~23 355 chars) | 85 ms | 85 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 176 ms (2.1×) |
| new connection + first message | 86 ms | 85 ms (1.0×) | 248 ms (2.9×) | 167 ms (2.0×) | 168 ms (2.0×) | 248 ms (2.9×) | 168 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 562 (0 errors) | 609 (0 errors) (0.9×) | 607 (0 errors) (0.9×) | 590 (0 errors) (1.0×) | 604 (0 errors) (0.9×) | 604 (0 errors) (0.9×) | 573 (0 errors) (1.0×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 204 ms | 202 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) |
| long (~660 chars) | 205 ms | 204 ms (1.0×) | 205 ms (1.0×) | 205 ms (1.0×) | 205 ms (1.0×) | 205 ms (1.0×) | 205 ms (1.0×) |
| very long (~5 302 chars) | 207 ms | 206 ms (1.0×) | 213 ms (1.0×) | 214 ms (1.0×) | 213 ms (1.0×) | 213 ms (1.0×) | 214 ms (1.0×) |
| huge (~23 355 chars) | 216 ms | 216 ms (1.0×) | 266 ms (1.2×) | 266 ms (1.2×) | 266 ms (1.2×) | 261 ms (1.2×) | 672 ms (3.1×) |
| new connection + first message | 206 ms | 206 ms (1.0×) | 615 ms (3.0×) | 414 ms (2.0×) | 413 ms (2.0×) | 615 ms (3.0×) | 418 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 170 (0 errors) | 246 (0 errors) (0.7×) | 244 (0 errors) (0.7×) | 239 (0 errors) (0.7×) | 245 (0 errors) (0.7×) | 240 (0 errors) (0.7×) | 167 (0 errors) (1.0×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


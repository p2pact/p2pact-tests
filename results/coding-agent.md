## Coding agent output

A coding agent at work: tool outputs (files, logs, diffs) from SWE-Gym OpenHands trajectories.

Data: SWE-Gym OpenHands trajectories (MIT), tool messages. Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 288 B | ≈447 B (1.6×) | ≈547 B (1.9×) | ≈840 B (2.9×) | ≈570 B (2.0×) | ≈687 B (2.4×) | ≈644 B (2.2×) |
| long (~660 chars) | 598 B | ≈756 B (1.3×) | ≈1.3 KB (2.2×) | ≈1.6 KB (2.7×) | ≈1.3 KB (2.3×) | ≈1.4 KB (2.3×) | ≈1.4 KB (2.4×) |
| very long (~5 302 chars) | 2.5 KB | ≈2.6 KB (1.1×) | ≈8.3 KB (3.3×) | ≈8.6 KB (3.5×) | ≈8.4 KB (3.4×) | ≈8.1 KB (3.2×) | ≈8.6 KB (3.5×) |
| huge (~23 355 chars) | 6.3 KB | ≈6.5 KB (1.0×) | ≈29 KB (4.7×) | ≈29 KB (4.7×) | ≈29 KB (4.7×) | ≈27 KB (4.4×) | ≈31 KB (4.9×) |
| new connection + first message | 830 B | ≈965 B (1.2×) | ≈4.4 KB (5.4×) | ≈4.3 KB (5.3×) | ≈4.4 KB (5.5×) | ≈4.4 KB (5.5×) | ≈51 KB (62.7×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 2.5 ms | 2.5 ms (1.0×) | 2.4 ms (1.0×) | 2.5 ms (1.0×) | 2.4 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) |
| long (~660 chars) | 2.7 ms | 2.6 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.4 ms (0.9×) | 2.5 ms (0.9×) |
| very long (~5 302 chars) | 2.5 ms | 2.6 ms (1.0×) | 2.5 ms (1.0×) | 2.7 ms (1.1×) | 2.5 ms (1.0×) | 2.4 ms (1.0×) | 2.5 ms (1.0×) |
| huge (~23 355 chars) | 3.0 ms | 3.0 ms (1.0×) | 2.7 ms (0.9×) | 2.7 ms (0.9×) | 3.0 ms (1.0×) | 2.6 ms (0.8×) | 2.9 ms (0.9×) |
| new connection + first message | 3.8 ms | 4.0 ms (1.1×) | 7.7 ms (2.0×) | 6.0 ms (1.6×) | 5.6 ms (1.5×) | 8.9 ms (2.4×) | 5.8 ms (1.5×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 20902 (0 errors) | 16985 (0 errors) (1.2×) | 20721 (0 errors) (1.0×) | 19600 (0 errors) (1.1×) | 21263 (0 errors) (1.0×) | 16951 (0 errors) (1.2×) | 20505 (0 errors) (1.0×) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 69 ms | 67 ms (1.0×) | 70 ms (1.0×) | 77 ms (1.1×) | 70 ms (1.0×) | 73 ms (1.1×) | 76 ms (1.1×) |
| long (~660 chars) | 77 ms | 75 ms (1.0×) | 91 ms (1.2×) | 97 ms (1.3×) | 91 ms (1.2×) | 91 ms (1.2×) | 95 ms (1.2×) |
| very long (~5 302 chars) | 106 ms | 104 ms (1.0×) | 189 ms (1.8×) | 196 ms (1.8×) | 191 ms (1.8×) | 184 ms (1.7×) | 202 ms (1.9×) |
| huge (~23 355 chars) | 212 ms | 206 ms (1.0×) | 854 ms (4.0×) | 861 ms (4.1×) | 858 ms (4.0×) | 804 ms (3.8×) | 894 ms (4.2×) |
| new connection + first message | 77 ms | 74 ms (1.0×) | 284 ms (3.7×) | 218 ms (2.8×) | 218 ms (2.8×) | 279 ms (3.6×) | 272 ms (3.5×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 195 (0 errors) | 221 (0 errors) (0.9×) | 159 (0 errors) (1.2×) | 630 (0 errors) (0.3×) | 170 (0 errors) (1.1×) | 165 (0 errors) (1.2×) | 100 (0 errors) (1.9×) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 82 ms | 81 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 81 ms (1.0×) | 83 ms (1.0×) |
| long (~660 chars) | 81 ms | 82 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| very long (~5 302 chars) | 82 ms | 82 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 84 ms (1.0×) |
| huge (~23 355 chars) | 82 ms | 83 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 243 ms (2.9×) |
| new connection + first message | 83 ms | 82 ms (1.0×) | 244 ms (3.0×) | 163 ms (2.0×) | 163 ms (2.0×) | 245 ms (3.0×) | 167 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 498 (0 errors) | 611 (0 errors) (0.8×) | 614 (0 errors) (0.8×) | 599 (0 errors) (0.8×) | 613 (0 errors) (0.8×) | 610 (0 errors) (0.8×) | 544 (0 errors) (0.9×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 202 ms | 202 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) |
| long (~660 chars) | 203 ms | 203 ms (1.0×) | 205 ms (1.0×) | 205 ms (1.0×) | 204 ms (1.0×) | 204 ms (1.0×) | 204 ms (1.0×) |
| very long (~5 302 chars) | 206 ms | 205 ms (1.0×) | 212 ms (1.0×) | 214 ms (1.0×) | 212 ms (1.0×) | 211 ms (1.0×) | 213 ms (1.0×) |
| huge (~23 355 chars) | 216 ms | 215 ms (1.0×) | 264 ms (1.2×) | 266 ms (1.2×) | 265 ms (1.2×) | 260 ms (1.2×) | 637 ms (2.9×) |
| new connection + first message | 203 ms | 203 ms (1.0×) | 616 ms (3.0×) | 412 ms (2.0×) | 411 ms (2.0×) | 615 ms (3.0×) | 416 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 138 (0 errors) | 245 (0 errors) (0.6×) | 244 (0 errors) (0.6×) | 239 (0 errors) (0.6×) | 246 (0 errors) (0.6×) | 240 (0 errors) (0.6×) | 168 (0 errors) (0.8×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


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

**Time from sending a message to its ack** (median):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 2.5 ms | 2.5 ms | 2.4 ms | 2.5 ms | 2.4 ms | 2.5 ms | 2.5 ms |
| long (~660 chars) | 2.7 ms | 2.6 ms | 2.5 ms | 2.5 ms | 2.5 ms | 2.4 ms | 2.5 ms |
| very long (~5 302 chars) | 2.5 ms | 2.6 ms | 2.5 ms | 2.7 ms | 2.5 ms | 2.4 ms | 2.5 ms |
| huge (~23 355 chars) | 3.0 ms | 3.0 ms | 2.7 ms | 2.7 ms | 3.0 ms | 2.6 ms | 2.9 ms |
| new connection + first message | 3.8 ms | 4.0 ms | 7.7 ms | 6.0 ms | 5.6 ms | 8.9 ms | 5.8 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 20902 (0) | 16985 (0) | 20721 (0) | 19600 (0) | 21263 (0) | 16951 (0) | 20505 (0) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 69 ms | 67 ms | 70 ms | 77 ms | 70 ms | 73 ms | 76 ms |
| long (~660 chars) | 77 ms | 75 ms | 91 ms | 97 ms | 91 ms | 91 ms | 95 ms |
| very long (~5 302 chars) | 106 ms | 104 ms | 189 ms | 196 ms | 191 ms | 184 ms | 202 ms |
| huge (~23 355 chars) | 212 ms | 206 ms | 854 ms | 861 ms | 858 ms | 804 ms | 894 ms |
| new connection + first message | 77 ms | 74 ms | 284 ms | 218 ms | 218 ms | 279 ms | 272 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 195 (0) | 221 (0) | 159 (0) | 630 (0) | 170 (0) | 165 (0) | 100 (0) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 82 ms | 81 ms | 81 ms | 81 ms | 82 ms | 81 ms | 83 ms |
| long (~660 chars) | 81 ms | 82 ms | 81 ms | 81 ms | 81 ms | 82 ms | 83 ms |
| very long (~5 302 chars) | 82 ms | 82 ms | 81 ms | 82 ms | 82 ms | 82 ms | 84 ms |
| huge (~23 355 chars) | 82 ms | 83 ms | 82 ms | 82 ms | 82 ms | 82 ms | 243 ms |
| new connection + first message | 83 ms | 82 ms | 244 ms | 163 ms | 163 ms | 245 ms | 167 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 498 (0) | 611 (0) | 614 (0) | 599 (0) | 613 (0) | 610 (0) | 544 (0) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 202 ms | 202 ms | 203 ms | 204 ms | 203 ms | 203 ms | 203 ms |
| long (~660 chars) | 203 ms | 203 ms | 205 ms | 205 ms | 204 ms | 204 ms | 204 ms |
| very long (~5 302 chars) | 206 ms | 205 ms | 212 ms | 214 ms | 212 ms | 211 ms | 213 ms |
| huge (~23 355 chars) | 216 ms | 215 ms | 264 ms | 266 ms | 265 ms | 260 ms | 637 ms |
| new connection + first message | 203 ms | 203 ms | 616 ms | 412 ms | 411 ms | 615 ms | 416 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 138 (0) | 245 (0) | 244 (0) | 239 (0) | 246 (0) | 240 (0) | 168 (0) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


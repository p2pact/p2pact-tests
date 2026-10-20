## Coding agent output

A coding agent at work: tool outputs (files, logs, diffs) from SWE-Gym OpenHands trajectories.

Data: SWE-Gym OpenHands trajectories (MIT), tool messages. Messages per size: 20 (2 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 273 B | ≈432 B (1.6×) | ≈524 B (1.9×) | ≈817 B (3.0×) | ≈547 B (2.0×) | ≈646 B (2.4×) | ≈1.0 KB (3.8×) |
| long (~660 chars) | 585 B | ≈738 B (1.3×) | ≈1.3 KB (2.2×) | ≈1.6 KB (2.7×) | ≈1.3 KB (2.3×) | ≈1.4 KB (2.4×) | ≈1.4 KB (2.5×) |
| very long (~5 302 chars) | 2.4 KB | ≈2.6 KB (1.1×) | ≈8.3 KB (3.4×) | ≈8.6 KB (3.6×) | ≈8.3 KB (3.5×) | ≈8.0 KB (3.3×) | ≈8.7 KB (3.6×) |
| huge (~23 355 chars) | 5.8 KB | ≈6.0 KB (1.0×) | ≈28 KB (4.8×) | ≈28 KB (4.8×) | ≈28 KB (4.8×) | ≈26 KB (4.5×) | ≈28 KB (4.8×) |
| new connection + first message | 885 B | ≈1.0 KB (1.2×) | ≈4.5 KB (5.2×) | ≈4.4 KB (5.1×) | ≈4.6 KB (5.3×) | ≈4.5 KB (5.2×) | ≈16 KB (18.1×) |

**Time from sending a message to its ack** (median):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 2.5 ms | 2.5 ms | 2.4 ms | 2.4 ms | 2.4 ms | 2.5 ms | 2.4 ms |
| long (~660 chars) | 2.9 ms | 2.5 ms | 2.4 ms | 2.4 ms | 2.4 ms | 2.4 ms | 2.4 ms |
| very long (~5 302 chars) | 3.8 ms | 2.7 ms | 2.4 ms | 2.4 ms | 2.4 ms | 2.4 ms | 2.4 ms |
| huge (~23 355 chars) | 3.4 ms | 3.4 ms | 2.6 ms | 2.5 ms | 2.5 ms | 2.4 ms | 2.8 ms |
| new connection + first message | 3.4 ms | 5.6 ms | 7.9 ms | 5.3 ms | 5.2 ms | 7.9 ms | 5.7 ms |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 68 ms | 67 ms | 69 ms | 76 ms | 69 ms | 72 ms | 74 ms |
| long (~660 chars) | 78 ms | 75 ms | 90 ms | 96 ms | 91 ms | 90 ms | 94 ms |
| very long (~5 302 chars) | 104 ms | 103 ms | 188 ms | 195 ms | 190 ms | 185 ms | 201 ms |
| huge (~23 355 chars) | 208 ms | 205 ms | 854 ms | 863 ms | 857 ms | 805 ms | 894 ms |
| new connection + first message | 75 ms | 73 ms | 284 ms | 218 ms | 217 ms | 279 ms | 267 ms |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 81 ms | 81 ms | 82 ms | 82 ms | 82 ms | 81 ms | 81 ms |
| long (~660 chars) | 82 ms | 82 ms | 82 ms | 82 ms | 82 ms | 81 ms | 81 ms |
| very long (~5 302 chars) | 82 ms | 82 ms | 82 ms | 82 ms | 82 ms | 81 ms | 82 ms |
| huge (~23 355 chars) | 200 ms | 83 ms | 82 ms | 82 ms | 82 ms | 82 ms | 163 ms |
| new connection + first message | 82 ms | 82 ms | 247 ms | 165 ms | 163 ms | 244 ms | 163 ms |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 202 ms | 201 ms | 203 ms | 204 ms | 203 ms | 203 ms | 203 ms |
| long (~660 chars) | 202 ms | 202 ms | 204 ms | 205 ms | 203 ms | 203 ms | 204 ms |
| very long (~5 302 chars) | 204 ms | 205 ms | 212 ms | 212 ms | 212 ms | 211 ms | 213 ms |
| huge (~23 355 chars) | 214 ms | 212 ms | 266 ms | 266 ms | 266 ms | 262 ms | 430 ms |
| new connection + first message | 203 ms | 202 ms | 612 ms | 412 ms | 411 ms | 614 ms | 415 ms |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


## Coding agent output

A coding agent at work: tool outputs (files, logs, diffs) from SWE-Gym OpenHands trajectories.

Data: SWE-Gym OpenHands trajectories (MIT), tool messages. Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~125 chars) | 183 B | ≈342 B (1.9×) | ≈512 B (2.8×) | ≈805 B (4.4×) | ≈536 B (2.9×) | ≈610 B (3.3×) | ≈626 B (3.4×) |
| long (~749 chars) | 369 B | ≈516 B (1.4×) | ≈1.2 KB (3.4×) | ≈1.5 KB (4.2×) | ≈1.3 KB (3.5×) | ≈1.3 KB (3.6×) | ≈1.4 KB (3.8×) |
| very long (~5 101 chars) | 1.7 KB | ≈1.9 KB (1.1×) | ≈7.0 KB (4.0×) | ≈7.3 KB (4.2×) | ≈7.0 KB (4.1×) | ≈6.8 KB (4.0×) | ≈7.2 KB (4.2×) |
| huge (~24 207 chars) | 6.2 KB | ≈6.1 KB (1.0×) | ≈26 KB (4.2×) | ≈26 KB (4.2×) | ≈26 KB (4.2×) | ≈24 KB (3.9×) | ≈26 KB (4.3×) |
| new connection + first message | 806 B | ≈1.1 KB (1.4×) | ≈4.3 KB (5.5×) | ≈4.2 KB (5.4×) | ≈4.4 KB (5.6×) | ≈4.4 KB (5.6×) | ≈57 KB (71.8×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~125 chars) | 2.8 ms | 2.7 ms (1.0×) | 2.7 ms (1.0×) | 2.8 ms (1.0×) | 2.9 ms (1.0×) | 2.8 ms (1.0×) | 2.7 ms (0.9×) |
| long (~749 chars) | 3.0 ms | 2.9 ms (1.0×) | 2.7 ms (0.9×) | 2.8 ms (0.9×) | 2.7 ms (0.9×) | 2.7 ms (0.9×) | 2.7 ms (0.9×) |
| very long (~5 101 chars) | 3.4 ms | 3.5 ms (1.0×) | 2.8 ms (0.8×) | 2.9 ms (0.8×) | 2.9 ms (0.8×) | 2.7 ms (0.8×) | 2.7 ms (0.8×) |
| huge (~24 207 chars) | 3.1 ms | 3.1 ms (1.0×) | 3.1 ms (1.0×) | 3.2 ms (1.1×) | 2.8 ms (0.9×) | 2.7 ms (0.9×) | 2.7 ms (0.9×) |
| new connection + first message | 8.0 ms | 7.3 ms (0.9×) | 9.5 ms (1.2×) | 6.5 ms (0.8×) | 7.3 ms (0.9×) | 9.5 ms (1.2×) | 5.9 ms (0.7×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 21412 (0 errors) | 21073 (0 errors) (1.0×) | 19110 (0 errors) (1.1×) | 18252 (0 errors) (1.2×) | 19522 (0 errors) (1.1×) | 18471 (0 errors) (1.2×) | 19084 (0 errors) (1.1×) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~125 chars) | 67 ms | 66 ms (1.0×) | 70 ms (1.1×) | 78 ms (1.2×) | 70 ms (1.1×) | 71 ms (1.1×) | 76 ms (1.1×) |
| long (~749 chars) | 70 ms | 70 ms (1.0×) | 90 ms (1.3×) | 98 ms (1.4×) | 90 ms (1.3×) | 91 ms (1.3×) | 95 ms (1.4×) |
| very long (~5 101 chars) | 96 ms | 96 ms (1.0×) | 175 ms (1.8×) | 181 ms (1.9×) | 175 ms (1.8×) | 169 ms (1.8×) | 184 ms (1.9×) |
| huge (~24 207 chars) | 209 ms | 209 ms (1.0×) | 600 ms (2.9×) | 606 ms (2.9×) | 602 ms (2.9×) | 563 ms (2.7×) | 625 ms (3.0×) |
| new connection + first message | 85 ms | 83 ms (1.0×) | 287 ms (3.4×) | 221 ms (2.6×) | 220 ms (2.6×) | 283 ms (3.3×) | 274 ms (3.2×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 412 (0 errors) | 626 (0 errors) (0.7×) | 170 (0 errors) (2.4×) | 618 (0 errors) (0.7×) | 182 (0 errors) (2.3×) | 175 (0 errors) (2.4×) | 106 (0 errors) (3.9×) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~125 chars) | 82 ms | 82 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) |
| long (~749 chars) | 82 ms | 81 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) |
| very long (~5 101 chars) | 82 ms | 82 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| huge (~24 207 chars) | 84 ms | 83 ms (1.0×) | 84 ms (1.0×) | 84 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 166 ms (2.0×) |
| new connection + first message | 89 ms | 86 ms (1.0×) | 247 ms (2.8×) | 168 ms (1.9×) | 168 ms (1.9×) | 249 ms (2.8×) | 168 ms (1.9×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 543 (0 errors) | 609 (0 errors) (0.9×) | 606 (0 errors) (0.9×) | 589 (0 errors) (0.9×) | 606 (0 errors) (0.9×) | 606 (0 errors) (0.9×) | 573 (0 errors) (0.9×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~125 chars) | 203 ms | 202 ms (1.0×) | 204 ms (1.0×) | 204 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 205 ms (1.0×) |
| long (~749 chars) | 204 ms | 203 ms (1.0×) | 205 ms (1.0×) | 206 ms (1.0×) | 204 ms (1.0×) | 204 ms (1.0×) | 205 ms (1.0×) |
| very long (~5 101 chars) | 206 ms | 206 ms (1.0×) | 211 ms (1.0×) | 213 ms (1.0×) | 212 ms (1.0×) | 211 ms (1.0×) | 210 ms (1.0×) |
| huge (~24 207 chars) | 214 ms | 214 ms (1.0×) | 247 ms (1.2×) | 246 ms (1.2×) | 247 ms (1.2×) | 244 ms (1.1×) | 621 ms (2.9×) |
| new connection + first message | 209 ms | 208 ms (1.0×) | 617 ms (3.0×) | 413 ms (2.0×) | 414 ms (2.0×) | 617 ms (3.0×) | 417 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 161 (0 errors) | 247 (0 errors) (0.7×) | 245 (0 errors) (0.7×) | 239 (0 errors) (0.7×) | 245 (0 errors) (0.7×) | 240 (0 errors) (0.7×) | 207 (0 errors) (0.8×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


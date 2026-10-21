## Coding agent output

A coding agent at work: tool outputs (files, logs, diffs) from SWE-Gym OpenHands trajectories.

Data: SWE-Gym OpenHands trajectories (MIT), tool messages. Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~125 chars) | 186 B | ≈344 B (1.8×) | ≈512 B (2.8×) | ≈805 B (4.3×) | ≈535 B (2.9×) | ≈624 B (3.4×) | ≈1.0 KB (5.6×) |
| long (~749 chars) | 374 B | ≈528 B (1.4×) | ≈1.2 KB (3.4×) | ≈1.5 KB (4.1×) | ≈1.3 KB (3.4×) | ≈1.3 KB (3.6×) | ≈1.3 KB (3.7×) |
| very long (~5 101 chars) | 1.7 KB | ≈1.9 KB (1.1×) | ≈7.0 KB (4.0×) | ≈7.3 KB (4.2×) | ≈7.0 KB (4.1×) | ≈6.8 KB (3.9×) | ≈7.3 KB (4.2×) |
| huge (~24 207 chars) | 6.3 KB | ≈6.5 KB (1.0×) | ≈26 KB (4.1×) | ≈26 KB (4.2×) | ≈26 KB (4.2×) | ≈24 KB (3.9×) | ≈26 KB (4.2×) |
| new connection + first message | 771 B | ≈1003 B (1.3×) | ≈4.3 KB (5.7×) | ≈4.2 KB (5.6×) | ≈4.4 KB (5.8×) | ≈4.4 KB (5.9×) | ≈43 KB (57.0×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~125 chars) | 2.5 ms | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.4 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) |
| long (~749 chars) | 2.6 ms | 2.6 ms (1.0×) | 2.4 ms (0.9×) | 2.6 ms (1.0×) | 2.6 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) |
| very long (~5 101 chars) | 2.9 ms | 2.9 ms (1.0×) | 2.5 ms (0.9×) | 2.6 ms (0.9×) | 2.6 ms (0.9×) | 2.5 ms (0.9×) | 2.5 ms (0.9×) |
| huge (~24 207 chars) | 3.4 ms | 2.9 ms (0.8×) | 2.9 ms (0.8×) | 2.6 ms (0.7×) | 2.6 ms (0.8×) | 2.5 ms (0.7×) | 2.6 ms (0.8×) |
| new connection + first message | 8.9 ms | 4.7 ms (0.5×) | 8.8 ms (1.0×) | 5.8 ms (0.7×) | 5.5 ms (0.6×) | 8.0 ms (0.9×) | 5.7 ms (0.6×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 19831 (0 errors) | 21157 (0 errors) (0.9×) | 21056 (0 errors) (0.9×) | 19295 (0 errors) (1.0×) | 21036 (0 errors) (0.9×) | 19847 (0 errors) (1.0×) | 20639 (0 errors) (1.0×) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~125 chars) | 66 ms | 65 ms (1.0×) | 69 ms (1.0×) | 77 ms (1.2×) | 69 ms (1.0×) | 71 ms (1.1×) | 75 ms (1.1×) |
| long (~749 chars) | 70 ms | 68 ms (1.0×) | 89 ms (1.3×) | 97 ms (1.4×) | 90 ms (1.3×) | 91 ms (1.3×) | 95 ms (1.4×) |
| very long (~5 101 chars) | 94 ms | 92 ms (1.0×) | 172 ms (1.8×) | 181 ms (1.9×) | 173 ms (1.9×) | 169 ms (1.8×) | 185 ms (2.0×) |
| huge (~24 207 chars) | 221 ms | 217 ms (1.0×) | 597 ms (2.7×) | 604 ms (2.7×) | 601 ms (2.7×) | 562 ms (2.5×) | 624 ms (2.8×) |
| new connection + first message | 77 ms | 75 ms (1.0×) | 282 ms (3.7×) | 218 ms (2.8×) | 218 ms (2.8×) | 279 ms (3.6×) | 270 ms (3.5×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 271 (0 errors) | 445 (0 errors) (0.6×) | 170 (0 errors) (1.6×) | 630 (0 errors) (0.4×) | 182 (0 errors) (1.5×) | 177 (0 errors) (1.5×) | 104 (0 errors) (2.6×) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~125 chars) | 81 ms | 82 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) |
| long (~749 chars) | 82 ms | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) |
| very long (~5 101 chars) | 82 ms | 82 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) |
| huge (~24 207 chars) | 83 ms | 83 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 165 ms (2.0×) |
| new connection + first message | 86 ms | 83 ms (1.0×) | 246 ms (2.9×) | 164 ms (1.9×) | 163 ms (1.9×) | 246 ms (2.9×) | 165 ms (1.9×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 557 (0 errors) | 610 (0 errors) (0.9×) | 612 (0 errors) (0.9×) | 597 (0 errors) (0.9×) | 609 (0 errors) (0.9×) | 613 (0 errors) (0.9×) | 563 (0 errors) (1.0×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~125 chars) | 202 ms | 202 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) |
| long (~749 chars) | 203 ms | 203 ms (1.0×) | 204 ms (1.0×) | 204 ms (1.0×) | 204 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) |
| very long (~5 101 chars) | 205 ms | 205 ms (1.0×) | 210 ms (1.0×) | 211 ms (1.0×) | 211 ms (1.0×) | 210 ms (1.0×) | 210 ms (1.0×) |
| huge (~24 207 chars) | 215 ms | 214 ms (1.0×) | 245 ms (1.1×) | 246 ms (1.1×) | 245 ms (1.1×) | 241 ms (1.1×) | 617 ms (2.9×) |
| new connection + first message | 210 ms | 208 ms (1.0×) | 614 ms (2.9×) | 410 ms (2.0×) | 411 ms (2.0×) | 613 ms (2.9×) | 414 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 173 (0 errors) | 247 (0 errors) (0.7×) | 246 (0 errors) (0.7×) | 240 (0 errors) (0.7×) | 247 (0 errors) (0.7×) | 242 (0 errors) (0.7×) | 191 (0 errors) (0.9×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


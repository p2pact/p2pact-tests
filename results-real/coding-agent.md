## Coding agent output

A coding agent at work: tool outputs (files, logs, diffs) from SWE-Gym OpenHands trajectories.

Data: SWE-Gym OpenHands trajectories (MIT), tool messages. Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 291 B | ≈519 B (1.8×) | ≈547 B (1.9×) | ≈840 B (2.9×) | ≈573 B (2.0×) | ≈670 B (2.3×) | ≈1.2 KB (4.4×) |
| long (~660 chars) | 598 B | ≈828 B (1.4×) | ≈1.3 KB (2.2×) | ≈1.6 KB (2.7×) | ≈1.3 KB (2.3×) | ≈1.4 KB (2.3×) | ≈1.4 KB (2.4×) |
| very long (~5 302 chars) | 3.8 KB | ≈2.7 KB (0.7×) | ≈8.3 KB (2.2×) | ≈8.6 KB (2.3×) | ≈8.4 KB (2.2×) | ≈8.1 KB (2.1×) | ≈8.9 KB (2.3×) |
| huge (~23 355 chars) | 6.2 KB | ≈6.7 KB (1.1×) | ≈29 KB (4.7×) | ≈29 KB (4.8×) | ≈29 KB (4.7×) | ≈27 KB (4.4×) | ≈29 KB (4.8×) |
| new connection + first message | 836 B | ≈1.2 KB (1.4×) | ≈4.4 KB (5.4×) | ≈4.3 KB (5.3×) | ≈4.5 KB (5.5×) | ≈4.4 KB (5.4×) | ≈21 KB (26.3×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 125 ms | 122 ms (1.0×) | 138 ms (1.1×) | 140 ms (1.1×) | 138 ms (1.1×) | 126 ms (1.0×) | 128 ms (1.0×) |
| long (~660 chars) | 133 ms | 119 ms (0.9×) | 139 ms (1.0×) | 141 ms (1.1×) | 133 ms (1.0×) | 129 ms (1.0×) | 123 ms (0.9×) |
| very long (~5 302 chars) | 129 ms | 123 ms (0.9×) | 138 ms (1.1×) | 136 ms (1.1×) | 154 ms (1.2×) | 132 ms (1.0×) | 134 ms (1.0×) |
| huge (~23 355 chars) | 129 ms | 123 ms (1.0×) | 151 ms (1.2×) | 156 ms (1.2×) | 141 ms (1.1×) | 152 ms (1.2×) | 145 ms (1.1×) |
| new connection + first message | 140 ms | 259 ms (1.8×) | 498 ms (3.6×) | 357 ms (2.6×) | 405 ms (2.9×) | 504 ms (3.6×) | 253 ms (1.8×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 362 (0 errors) | 363 (0 errors) (1.0×) | 341 (0 errors) (1.1×) | 233 (0 errors) (1.6×) | 307 (0 errors) (1.2×) | 361 (0 errors) (1.0×) | 380 (0 errors) (1.0×) |


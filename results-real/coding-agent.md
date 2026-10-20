## Coding agent output

A coding agent at work: tool outputs (files, logs, diffs) from SWE-Gym OpenHands trajectories.

Data: SWE-Gym OpenHands trajectories (MIT), tool messages. Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 298 B | ≈444 B (1.5×) | ≈547 B (1.8×) | ≈840 B (2.8×) | ≈572 B (1.9×) | ≈662 B (2.2×) | ≈1.2 KB (4.1×) |
| long (~660 chars) | 716 B | ≈775 B (1.1×) | ≈1.3 KB (1.8×) | ≈1.6 KB (2.3×) | ≈1.3 KB (1.9×) | ≈1.4 KB (2.0×) | ≈1.4 KB (2.0×) |
| very long (~5 302 chars) | 3.9 KB | ≈2.7 KB (0.7×) | ≈8.3 KB (2.1×) | ≈8.6 KB (2.2×) | ≈8.4 KB (2.1×) | ≈8.1 KB (2.1×) | ≈8.7 KB (2.2×) |
| huge (~23 355 chars) | 6.3 KB | ≈6.5 KB (1.0×) | ≈29 KB (4.6×) | ≈29 KB (4.7×) | ≈29 KB (4.6×) | ≈27 KB (4.4×) | ≈30 KB (4.7×) |
| new connection + first message | 782 B | ≈951 B (1.2×) | ≈4.4 KB (5.7×) | ≈4.3 KB (5.6×) | ≈4.4 KB (5.8×) | ≈4.4 KB (5.8×) | ≈21 KB (28.1×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 116 ms | 109 ms | 109 ms | 117 ms | 97 ms | 108 ms | 98 ms |
| long (~660 chars) | 116 ms | 109 ms | 109 ms | 117 ms | 97 ms | 108 ms | 98 ms |
| very long (~5 302 chars) | 117 ms | 110 ms | 113 ms | 120 ms | 98 ms | 111 ms | 100 ms |
| huge (~23 355 chars) | 119 ms | 112 ms | 116 ms | 123 ms | 102 ms | 115 ms | 112 ms |
| new connection + first message | 119 ms | 223 ms | 425 ms | 309 ms | 334 ms | 438 ms | 198 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 415 (0) | 439 (0) | 443 (0) | 277 (0) | 491 (0) | 452 (0) | 495 (0) |


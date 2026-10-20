## Coding agent output

A coding agent at work: tool outputs (files, logs, diffs) from SWE-Gym OpenHands trajectories.

Data: SWE-Gym OpenHands trajectories (MIT), tool messages. Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 291 B | ≈444 B (1.5×) | ≈547 B (1.9×) | ≈840 B (2.9×) | ≈573 B (2.0×) | ≈666 B (2.3×) | ≈1.2 KB (4.1×) |
| long (~660 chars) | 598 B | ≈776 B (1.3×) | ≈1.3 KB (2.2×) | ≈1.6 KB (2.7×) | ≈1.3 KB (2.3×) | ≈1.4 KB (2.3×) | ≈1.4 KB (2.4×) |
| very long (~5 302 chars) | 3.1 KB | ≈2.7 KB (0.9×) | ≈8.3 KB (2.7×) | ≈8.6 KB (2.8×) | ≈8.4 KB (2.7×) | ≈8.0 KB (2.6×) | ≈8.7 KB (2.8×) |
| huge (~23 355 chars) | 6.4 KB | ≈6.5 KB (1.0×) | ≈29 KB (4.5×) | ≈29 KB (4.6×) | ≈29 KB (4.6×) | ≈27 KB (4.3×) | ≈29 KB (4.6×) |
| new connection + first message | 836 B | ≈951 B (1.1×) | ≈4.4 KB (5.4×) | ≈4.3 KB (5.3×) | ≈4.4 KB (5.4×) | ≈4.4 KB (5.4×) | ≈21 KB (26.2×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~157 chars) | 140 ms | 136 ms | 161 ms | 133 ms | 146 ms | 146 ms | 137 ms |
| long (~660 chars) | 144 ms | 153 ms | 154 ms | 144 ms | 140 ms | 137 ms | 140 ms |
| very long (~5 302 chars) | 147 ms | 151 ms | 160 ms | 133 ms | 147 ms | 151 ms | 152 ms |
| huge (~23 355 chars) | 156 ms | 160 ms | 169 ms | 164 ms | 160 ms | 156 ms | 159 ms |
| new connection + first message | 182 ms | 288 ms | 593 ms | 420 ms | 429 ms | 595 ms | 303 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 339 (0) | 331 (0) | 312 (0) | 224 (0) | 306 (0) | 331 (0) | 310 (0) |


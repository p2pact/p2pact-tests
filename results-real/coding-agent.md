## Coding agent output

A coding agent at work: tool outputs (files, logs, diffs) from SWE-Gym OpenHands trajectories.

Data: SWE-Gym OpenHands trajectories (MIT), tool messages. Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~125 chars) | 277 B | ≈430 B (1.6×) | ≈512 B (1.8×) | ≈805 B (2.9×) | ≈538 B (1.9×) | ≈633 B (2.3×) | ≈1.2 KB (4.4×) |
| long (~749 chars) | 555 B | ≈709 B (1.3×) | ≈1.2 KB (2.3×) | ≈1.5 KB (2.8×) | ≈1.3 KB (2.3×) | ≈1.3 KB (2.4×) | ≈1.5 KB (2.8×) |
| very long (~5 101 chars) | 2.8 KB | ≈2.2 KB (0.8×) | ≈7.0 KB (2.5×) | ≈7.3 KB (2.6×) | ≈7.0 KB (2.5×) | ≈6.8 KB (2.5×) | ≈7.7 KB (2.8×) |
| huge (~24 207 chars) | 6.7 KB | ≈6.8 KB (1.0×) | ≈26 KB (3.9×) | ≈26 KB (3.9×) | ≈26 KB (3.9×) | ≈24 KB (3.6×) | ≈27 KB (4.0×) |
| new connection + first message | 871 B | ≈1018 B (1.2×) | ≈4.3 KB (5.1×) | ≈4.2 KB (5.0×) | ≈4.4 KB (5.2×) | ≈4.4 KB (5.2×) | ≈21 KB (24.8×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~125 chars) | 123 ms | 112 ms (0.9×) | 126 ms (1.0×) | 130 ms (1.1×) | 133 ms (1.1×) | 136 ms (1.1×) | 129 ms (1.0×) |
| long (~749 chars) | 122 ms | 113 ms (0.9×) | 136 ms (1.1×) | 133 ms (1.1×) | 138 ms (1.1×) | 130 ms (1.1×) | 133 ms (1.1×) |
| very long (~5 101 chars) | 137 ms | 115 ms (0.8×) | 135 ms (1.0×) | 131 ms (1.0×) | 145 ms (1.1×) | 144 ms (1.1×) | 149 ms (1.1×) |
| huge (~24 207 chars) | 131 ms | 133 ms (1.0×) | 152 ms (1.2×) | 146 ms (1.1×) | 145 ms (1.1×) | 142 ms (1.1×) | 149 ms (1.1×) |
| new connection + first message | 132 ms | 230 ms (1.7×) | 492 ms (3.7×) | 383 ms (2.9×) | 407 ms (3.1×) | 507 ms (3.8×) | 268 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 362 (0 errors) | 395 (0 errors) (0.9×) | 350 (0 errors) (1.0×) | 231 (0 errors) (1.6×) | 323 (0 errors) (1.1×) | 344 (0 errors) (1.1×) | 344 (0 errors) (1.1×) |


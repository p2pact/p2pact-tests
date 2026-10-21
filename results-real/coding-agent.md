## Coding agent output

A coding agent at work: tool outputs (files, logs, diffs) from SWE-Gym OpenHands trajectories.

Data: SWE-Gym OpenHands trajectories (MIT), tool messages. Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~125 chars) | 290 B | ≈430 B (1.5×) | ≈512 B (1.8×) | ≈805 B (2.8×) | ≈538 B (1.9×) | ≈629 B (2.2×) | ≈1.2 KB (4.1×) |
| long (~749 chars) | 568 B | ≈709 B (1.2×) | ≈1.2 KB (2.2×) | ≈1.5 KB (2.7×) | ≈1.3 KB (2.3×) | ≈1.3 KB (2.4×) | ≈1.5 KB (2.7×) |
| very long (~5 101 chars) | 2.6 KB | ≈2.2 KB (0.8×) | ≈7.0 KB (2.6×) | ≈7.3 KB (2.8×) | ≈7.0 KB (2.7×) | ≈6.8 KB (2.6×) | ≈7.3 KB (2.8×) |
| huge (~24 207 chars) | 7.4 KB | ≈6.9 KB (0.9×) | ≈26 KB (3.5×) | ≈26 KB (3.5×) | ≈26 KB (3.5×) | ≈24 KB (3.3×) | ≈26 KB (3.6×) |
| new connection + first message | 925 B | ≈1.3 KB (1.4×) | ≈4.3 KB (4.8×) | ≈4.2 KB (4.7×) | ≈4.4 KB (4.8×) | ≈4.4 KB (4.9×) | ≈21 KB (23.1×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| medium (~125 chars) | 137 ms | 135 ms (1.0×) | 131 ms (1.0×) | 114 ms (0.8×) | 130 ms (0.9×) | 132 ms (1.0×) | 105 ms (0.8×) |
| long (~749 chars) | 122 ms | 131 ms (1.1×) | 124 ms (1.0×) | 113 ms (0.9×) | 101 ms (0.8×) | 108 ms (0.9×) | 106 ms (0.9×) |
| very long (~5 101 chars) | 120 ms | 113 ms (0.9×) | 128 ms (1.1×) | 119 ms (1.0×) | 105 ms (0.9×) | 111 ms (0.9×) | 107 ms (0.9×) |
| huge (~24 207 chars) | 125 ms | 116 ms (0.9×) | 133 ms (1.1×) | 119 ms (1.0×) | 109 ms (0.9×) | 116 ms (0.9×) | 116 ms (0.9×) |
| new connection + first message | 134 ms | 253 ms (1.9×) | 501 ms (3.7×) | 402 ms (3.0×) | 374 ms (2.8×) | 557 ms (4.1×) | 221 ms (1.6×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 407 (0 errors) | 415 (0 errors) (1.0×) | 427 (0 errors) (1.0×) | 261 (0 errors) (1.6×) | 433 (0 errors) (0.9×) | 401 (0 errors) (1.0×) | 403 (0 errors) (1.0×) |


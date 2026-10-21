## Assistant replies

What an assistant answers: oasst2 assistant messages.

Data: OpenAssistant oasst2 (Apache-2.0), assistant. Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 174 B | ≈329 B (1.9×) | ≈401 B (2.3×) | ≈695 B (4.0×) | ≈428 B (2.5×) | ≈518 B (3.0×) | ≈1.0 KB (6.1×) |
| medium (~104 chars) | 299 B | ≈422 B (1.4×) | ≈513 B (1.7×) | ≈806 B (2.7×) | ≈540 B (1.8×) | ≈630 B (2.1×) | ≈713 B (2.4×) |
| long (~694 chars) | 826 B | ≈881 B (1.1×) | ≈1.4 KB (1.7×) | ≈1.7 KB (2.1×) | ≈1.4 KB (1.8×) | ≈1.5 KB (1.9×) | ≈1.6 KB (1.9×) |
| very long (~2 476 chars) | 1.3 KB | ≈1.5 KB (1.1×) | ≈3.2 KB (2.4×) | ≈3.5 KB (2.6×) | ≈3.2 KB (2.4×) | ≈3.3 KB (2.5×) | ≈3.6 KB (2.7×) |
| new connection + first message | 874 B | ≈1018 B (1.2×) | ≈4.2 KB (4.9×) | ≈4.1 KB (4.8×) | ≈4.3 KB (5.0×) | ≈4.3 KB (5.0×) | ≈21 KB (24.2×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 106 ms | 111 ms (1.1×) | 104 ms (1.0×) | 108 ms (1.0×) | 102 ms (1.0×) | 97 ms (0.9×) | 100 ms (1.0×) |
| medium (~104 chars) | 114 ms | 112 ms (1.0×) | 104 ms (0.9×) | 108 ms (0.9×) | 102 ms (0.9×) | 95 ms (0.8×) | 101 ms (0.9×) |
| long (~694 chars) | 124 ms | 112 ms (0.9×) | 104 ms (0.8×) | 109 ms (0.9×) | 103 ms (0.8×) | 96 ms (0.8×) | 100 ms (0.8×) |
| very long (~2 476 chars) | 111 ms | 115 ms (1.0×) | 105 ms (1.0×) | 114 ms (1.0×) | 103 ms (0.9×) | 95 ms (0.9×) | 101 ms (0.9×) |
| new connection + first message | 114 ms | 235 ms (2.1×) | 417 ms (3.6×) | 430 ms (3.8×) | 338 ms (3.0×) | 433 ms (3.8×) | 203 ms (1.8×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 282 (0 errors) | 422 (0 errors) (0.7×) | 458 (0 errors) (0.6×) | 263 (0 errors) (1.1×) | 448 (0 errors) (0.6×) | 504 (0 errors) (0.6×) | 479 (0 errors) (0.6×) |


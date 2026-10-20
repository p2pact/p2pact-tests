## Generated text

Reproducible offline: generated text with natural-language word frequencies, four network profiles.

Data: generated text (Zipf-distributed words). Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 162 B | ≈315 B (1.9×) | ≈387 B (2.4×) | ≈681 B (4.2×) | ≈413 B (2.5×) | ≈512 B (3.2×) | ≈1.0 KB (6.5×) |
| medium (~100 chars) | 227 B | ≈389 B (1.7×) | ≈469 B (2.1×) | ≈762 B (3.4×) | ≈493 B (2.2×) | ≈595 B (2.6×) | ≈547 B (2.4×) |
| long (~600 chars) | 511 B | ≈644 B (1.3×) | ≈969 B (1.9×) | ≈1.2 KB (2.5×) | ≈998 B (2.0×) | ≈1.1 KB (2.1×) | ≈1.1 KB (2.3×) |
| very long (~5 000 chars) | 3.0 KB | ≈2.6 KB (0.9×) | ≈5.4 KB (1.8×) | ≈5.8 KB (1.9×) | ≈5.5 KB (1.9×) | ≈5.6 KB (1.9×) | ≈5.8 KB (2.0×) |
| huge (~200 000 chars) | 80 KB | ≈101 KB (1.3×) | ≈206 KB (2.6×) | ≈207 KB (2.6×) | ≈207 KB (2.6×) | ≈207 KB (2.6×) | ≈207 KB (2.6×) |
| new connection + first message | 703 B | ≈953 B (1.4×) | ≈4.2 KB (6.2×) | ≈4.1 KB (6.0×) | ≈4.3 KB (6.2×) | ≈4.3 KB (6.2×) | ≈23 KB (34.1×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 114 ms | 120 ms (1.1×) | 136 ms (1.2×) | 128 ms (1.1×) | 121 ms (1.1×) | 134 ms (1.2×) | 135 ms (1.2×) |
| medium (~100 chars) | 115 ms | 126 ms (1.1×) | 126 ms (1.1×) | 133 ms (1.2×) | 121 ms (1.1×) | 134 ms (1.2×) | 137 ms (1.2×) |
| long (~600 chars) | 128 ms | 129 ms (1.0×) | 129 ms (1.0×) | 129 ms (1.0×) | 133 ms (1.0×) | 140 ms (1.1×) | 134 ms (1.1×) |
| very long (~5 000 chars) | 121 ms | 135 ms (1.1×) | 137 ms (1.1×) | 124 ms (1.0×) | 129 ms (1.1×) | 139 ms (1.1×) | 143 ms (1.2×) |
| huge (~200 000 chars) | 160 ms | 158 ms (1.0×) | 171 ms (1.1×) | 184 ms (1.1×) | 186 ms (1.2×) | 201 ms (1.3×) | 220 ms (1.4×) |
| new connection + first message | 121 ms | 241 ms (2.0×) | 534 ms (4.4×) | 405 ms (3.4×) | 383 ms (3.2×) | 557 ms (4.6×) | 320 ms (2.7×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 388 (0 errors) | 371 (0 errors) (1.0×) | 358 (0 errors) (1.1×) | 233 (0 errors) (1.7×) | 362 (0 errors) (1.1×) | 350 (0 errors) (1.1×) | 337 (0 errors) (1.1×) |


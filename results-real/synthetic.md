## Generated text

Reproducible offline: generated text with natural-language word frequencies, four network profiles.

Data: generated text (Zipf-distributed words). Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 162 B | ≈315 B (1.9×) | ≈387 B (2.4×) | ≈681 B (4.2×) | ≈413 B (2.5×) | ≈507 B (3.1×) | ≈957 B (5.9×) |
| medium (~100 chars) | 260 B | ≈380 B (1.5×) | ≈469 B (1.8×) | ≈762 B (2.9×) | ≈493 B (1.9×) | ≈591 B (2.3×) | ≈533 B (2.0×) |
| long (~600 chars) | 520 B | ≈665 B (1.3×) | ≈969 B (1.9×) | ≈1.2 KB (2.4×) | ≈998 B (1.9×) | ≈1.1 KB (2.1×) | ≈1.0 KB (2.0×) |
| very long (~5 000 chars) | 3.0 KB | ≈2.6 KB (0.9×) | ≈5.4 KB (1.8×) | ≈5.8 KB (1.9×) | ≈5.5 KB (1.8×) | ≈5.6 KB (1.8×) | ≈5.8 KB (1.9×) |
| huge (~200 000 chars) | 92 KB | ≈84 KB (0.9×) | ≈206 KB (2.3×) | ≈207 KB (2.3×) | ≈207 KB (2.3×) | ≈207 KB (2.3×) | ≈210 KB (2.3×) |
| new connection + first message | 703 B | ≈846 B (1.2×) | ≈4.2 KB (6.1×) | ≈4.1 KB (6.0×) | ≈4.3 KB (6.2×) | ≈4.3 KB (6.2×) | ≈21 KB (30.2×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 139 ms | 130 ms (0.9×) | 126 ms (0.9×) | 159 ms (1.1×) | 136 ms (1.0×) | 133 ms (1.0×) | 135 ms (1.0×) |
| medium (~100 chars) | 145 ms | 141 ms (1.0×) | 135 ms (0.9×) | 157 ms (1.1×) | 123 ms (0.9×) | 129 ms (0.9×) | 136 ms (0.9×) |
| long (~600 chars) | 141 ms | 140 ms (1.0×) | 129 ms (0.9×) | 154 ms (1.1×) | 134 ms (0.9×) | 130 ms (0.9×) | 144 ms (1.0×) |
| very long (~5 000 chars) | 144 ms | 144 ms (1.0×) | 144 ms (1.0×) | 161 ms (1.1×) | 140 ms (1.0×) | 143 ms (1.0×) | 143 ms (1.0×) |
| huge (~200 000 chars) | 181 ms | 197 ms (1.1×) | 199 ms (1.1×) | 231 ms (1.3×) | 206 ms (1.1×) | 188 ms (1.0×) | 440 ms (2.4×) |
| new connection + first message | 150 ms | 287 ms (1.9×) | 528 ms (3.5×) | 420 ms (2.8×) | 430 ms (2.9×) | 567 ms (3.8×) | 291 ms (1.9×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 330 (0 errors) | 339 (0 errors) (1.0×) | 347 (0 errors) (1.0×) | 217 (0 errors) (1.5×) | 332 (0 errors) (1.0×) | 357 (0 errors) (0.9×) | 348 (0 errors) (0.9×) |


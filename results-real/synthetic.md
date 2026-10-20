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

**Time from sending a message to its ack** (median):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 139 ms | 130 ms | 126 ms | 159 ms | 136 ms | 133 ms | 135 ms |
| medium (~100 chars) | 145 ms | 141 ms | 135 ms | 157 ms | 123 ms | 129 ms | 136 ms |
| long (~600 chars) | 141 ms | 140 ms | 129 ms | 154 ms | 134 ms | 130 ms | 144 ms |
| very long (~5 000 chars) | 144 ms | 144 ms | 144 ms | 161 ms | 140 ms | 143 ms | 143 ms |
| huge (~200 000 chars) | 181 ms | 197 ms | 199 ms | 231 ms | 206 ms | 188 ms | 440 ms |
| new connection + first message | 150 ms | 287 ms | 528 ms | 420 ms | 430 ms | 567 ms | 291 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 330 (0) | 339 (0) | 347 (0) | 217 (0) | 332 (0) | 357 (0) | 348 (0) |


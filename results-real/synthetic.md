## Generated text

Reproducible offline: generated text with natural-language word frequencies, four network profiles.

Data: generated text (Zipf-distributed words). Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 168 B | ≈315 B (1.9×) | ≈387 B (2.3×) | ≈681 B (4.1×) | ≈413 B (2.5×) | ≈516 B (3.1×) | ≈1.0 KB (6.2×) |
| medium (~100 chars) | 282 B | ≈380 B (1.3×) | ≈469 B (1.7×) | ≈762 B (2.7×) | ≈493 B (1.7×) | ≈595 B (2.1×) | ≈561 B (2.0×) |
| long (~600 chars) | 496 B | ≈665 B (1.3×) | ≈969 B (2.0×) | ≈1.2 KB (2.5×) | ≈997 B (2.0×) | ≈1.1 KB (2.2×) | ≈1.1 KB (2.3×) |
| very long (~5 000 chars) | 2.7 KB | ≈2.6 KB (0.9×) | ≈5.4 KB (2.0×) | ≈5.8 KB (2.1×) | ≈5.5 KB (2.0×) | ≈5.6 KB (2.1×) | ≈5.8 KB (2.1×) |
| huge (~200 000 chars) | 92 KB | ≈79 KB (0.9×) | ≈206 KB (2.2×) | ≈207 KB (2.2×) | ≈207 KB (2.2×) | ≈207 KB (2.2×) | ≈208 KB (2.3×) |
| new connection + first message | 771 B | ≈914 B (1.2×) | ≈4.2 KB (5.6×) | ≈4.1 KB (5.5×) | ≈4.3 KB (5.7×) | ≈4.3 KB (5.7×) | ≈21 KB (27.2×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 108 ms | 111 ms (1.0×) | 116 ms (1.1×) | 120 ms (1.1×) | 113 ms (1.0×) | 114 ms (1.1×) | 129 ms (1.2×) |
| medium (~100 chars) | 109 ms | 119 ms (1.1×) | 120 ms (1.1×) | 113 ms (1.0×) | 117 ms (1.1×) | 115 ms (1.1×) | 129 ms (1.2×) |
| long (~600 chars) | 113 ms | 116 ms (1.0×) | 119 ms (1.1×) | 118 ms (1.0×) | 121 ms (1.1×) | 118 ms (1.0×) | 128 ms (1.1×) |
| very long (~5 000 chars) | 121 ms | 124 ms (1.0×) | 127 ms (1.0×) | 118 ms (1.0×) | 125 ms (1.0×) | 123 ms (1.0×) | 129 ms (1.1×) |
| huge (~200 000 chars) | 137 ms | 148 ms (1.1×) | 192 ms (1.4×) | 169 ms (1.2×) | 173 ms (1.3×) | 272 ms (2.0×) | 174 ms (1.3×) |
| new connection + first message | 130 ms | 226 ms (1.7×) | 513 ms (3.9×) | 375 ms (2.9×) | 381 ms (2.9×) | 516 ms (4.0×) | 257 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 405 (0 errors) | 392 (0 errors) (1.0×) | 365 (0 errors) (1.1×) | 231 (0 errors) (1.8×) | 363 (0 errors) (1.1×) | 393 (0 errors) (1.0×) | 373 (0 errors) (1.1×) |


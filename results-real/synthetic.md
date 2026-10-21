## Generated text

Reproducible offline: generated text with natural-language word frequencies, four network profiles.

Data: generated text (Zipf-distributed words). Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 162 B | ≈315 B (1.9×) | ≈387 B (2.4×) | ≈681 B (4.2×) | ≈412 B (2.5×) | ≈504 B (3.1×) | ≈965 B (6.0×) |
| medium (~100 chars) | 243 B | ≈389 B (1.6×) | ≈469 B (1.9×) | ≈762 B (3.1×) | ≈493 B (2.0×) | ≈585 B (2.4×) | ≈615 B (2.5×) |
| long (~600 chars) | 586 B | ≈644 B (1.1×) | ≈969 B (1.7×) | ≈1.2 KB (2.2×) | ≈998 B (1.7×) | ≈1.1 KB (1.9×) | ≈1.0 KB (1.8×) |
| very long (~5 000 chars) | 2.5 KB | ≈2.6 KB (1.0×) | ≈5.4 KB (2.2×) | ≈5.8 KB (2.3×) | ≈5.5 KB (2.2×) | ≈5.6 KB (2.3×) | ≈5.8 KB (2.4×) |
| huge (~200 000 chars) | 82 KB | ≈80 KB (1.0×) | ≈206 KB (2.5×) | ≈207 KB (2.5×) | ≈207 KB (2.5×) | ≈207 KB (2.5×) | ≈211 KB (2.6×) |
| new connection + first message | 835 B | ≈989 B (1.2×) | ≈4.2 KB (5.2×) | ≈4.1 KB (5.1×) | ≈4.3 KB (5.2×) | ≈4.3 KB (5.3×) | ≈26 KB (31.7×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 139 ms | 141 ms (1.0×) | 137 ms (1.0×) | 121 ms (0.9×) | 136 ms (1.0×) | 131 ms (0.9×) | 112 ms (0.8×) |
| medium (~100 chars) | 113 ms | 125 ms (1.1×) | 139 ms (1.2×) | 129 ms (1.1×) | 124 ms (1.1×) | 119 ms (1.1×) | 122 ms (1.1×) |
| long (~600 chars) | 112 ms | 130 ms (1.2×) | 126 ms (1.1×) | 113 ms (1.0×) | 117 ms (1.0×) | 112 ms (1.0×) | 115 ms (1.0×) |
| very long (~5 000 chars) | 113 ms | 120 ms (1.1×) | 118 ms (1.0×) | 103 ms (0.9×) | 121 ms (1.1×) | 111 ms (1.0×) | 117 ms (1.0×) |
| huge (~200 000 chars) | 133 ms | 161 ms (1.2×) | 136 ms (1.0×) | 132 ms (1.0×) | 133 ms (1.0×) | 144 ms (1.1×) | 567 ms (4.3×) |
| new connection + first message | 147 ms | 253 ms (1.7×) | 513 ms (3.5×) | 376 ms (2.6×) | 402 ms (2.7×) | 504 ms (3.4×) | 274 ms (1.9×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 402 (0 errors) | 324 (0 errors) (1.2×) | 390 (0 errors) (1.0×) | 252 (0 errors) (1.6×) | 391 (0 errors) (1.0×) | 380 (0 errors) (1.1×) | 387 (0 errors) (1.0×) |


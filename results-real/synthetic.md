## Generated text

Reproducible offline: generated text with natural-language word frequencies, four network profiles.

Data: generated text (Zipf-distributed words). Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 185 B | ≈319 B (1.7×) | ≈387 B (2.1×) | ≈681 B (3.7×) | ≈414 B (2.2×) | ≈503 B (2.7×) | ≈1019 B (5.5×) |
| medium (~100 chars) | 247 B | ≈385 B (1.6×) | ≈469 B (1.9×) | ≈762 B (3.1×) | ≈493 B (2.0×) | ≈584 B (2.4×) | ≈527 B (2.1×) |
| long (~600 chars) | 553 B | ≈671 B (1.2×) | ≈969 B (1.8×) | ≈1.2 KB (2.3×) | ≈998 B (1.8×) | ≈1.1 KB (2.0×) | ≈1.0 KB (1.9×) |
| very long (~5 000 chars) | 2.6 KB | ≈2.6 KB (1.0×) | ≈5.4 KB (2.1×) | ≈5.8 KB (2.2×) | ≈5.5 KB (2.1×) | ≈5.6 KB (2.1×) | ≈5.8 KB (2.2×) |
| huge (~200 000 chars) | 95 KB | ≈79 KB (0.8×) | ≈206 KB (2.2×) | ≈207 KB (2.2×) | ≈207 KB (2.2×) | ≈207 KB (2.2×) | ≈209 KB (2.2×) |
| new connection + first message | 747 B | ≈818 B (1.1×) | ≈4.2 KB (5.8×) | ≈4.1 KB (5.7×) | ≈4.3 KB (5.8×) | ≈4.3 KB (5.9×) | ≈21 KB (28.8×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~20 chars) | 99 ms | 100 ms | 104 ms | 118 ms | 109 ms | 96 ms | 105 ms |
| medium (~100 chars) | 99 ms | 100 ms | 104 ms | 118 ms | 107 ms | 96 ms | 105 ms |
| long (~600 chars) | 99 ms | 100 ms | 104 ms | 118 ms | 108 ms | 96 ms | 105 ms |
| very long (~5 000 chars) | 100 ms | 102 ms | 107 ms | 120 ms | 111 ms | 99 ms | 107 ms |
| huge (~200 000 chars) | 125 ms | 117 ms | 117 ms | 133 ms | 120 ms | 109 ms | 145 ms |
| new connection + first message | 102 ms | 226 ms | 438 ms | 340 ms | 325 ms | 434 ms | 212 ms |

Load: 2000 messages, 50 at a time — messages per second (errors):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 348 (0) | 477 (0) | 461 (0) | 273 (0) | 243 (0) | 507 (0) | 459 (0) |


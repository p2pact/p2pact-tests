## Human messages

What people type to an assistant: OpenAssistant oasst2 prompts, all languages.

Data: OpenAssistant oasst2 (Apache-2.0), human. Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 178 B | ≈331 B (1.9×) | ≈403 B (2.3×) | ≈697 B (3.9×) | ≈431 B (2.4×) | ≈526 B (3.0×) | ≈1.0 KB (6.0×) |
| medium (~80 chars) | 232 B | ≈386 B (1.7×) | ≈469 B (2.0×) | ≈761 B (3.3×) | ≈495 B (2.1×) | ≈592 B (2.6×) | ≈558 B (2.4×) |
| long (~298 chars) | 533 B | ≈688 B (1.3×) | ≈1.1 KB (2.1×) | ≈1.4 KB (2.6×) | ≈1.1 KB (2.1×) | ≈1.2 KB (2.3×) | ≈1.2 KB (2.3×) |
| very long (~2 750 chars) | 2.0 KB | ≈1.8 KB (0.9×) | ≈4.2 KB (2.1×) | ≈4.5 KB (2.2×) | ≈4.3 KB (2.1×) | ≈4.3 KB (2.1×) | ≈4.4 KB (2.2×) |
| new connection + first message | 788 B | ≈1.1 KB (1.4×) | ≈4.2 KB (5.5×) | ≈4.1 KB (5.4×) | ≈4.3 KB (5.6×) | ≈4.3 KB (5.6×) | ≈21 KB (27.6×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 116 ms | 114 ms (1.0×) | 115 ms (1.0×) | 132 ms (1.1×) | 133 ms (1.1×) | 130 ms (1.1×) | 132 ms (1.1×) |
| medium (~80 chars) | 116 ms | 113 ms (1.0×) | 114 ms (1.0×) | 128 ms (1.1×) | 132 ms (1.1×) | 133 ms (1.1×) | 134 ms (1.2×) |
| long (~298 chars) | 128 ms | 117 ms (0.9×) | 116 ms (0.9×) | 129 ms (1.0×) | 143 ms (1.1×) | 135 ms (1.1×) | 130 ms (1.0×) |
| very long (~2 750 chars) | 133 ms | 118 ms (0.9×) | 122 ms (0.9×) | 133 ms (1.0×) | 143 ms (1.1×) | 140 ms (1.1×) | 137 ms (1.0×) |
| new connection + first message | 134 ms | 247 ms (1.8×) | 631 ms (4.7×) | 396 ms (3.0×) | 395 ms (2.9×) | 488 ms (3.6×) | 276 ms (2.1×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 367 (0 errors) | 349 (0 errors) (1.1×) | 370 (0 errors) (1.0×) | 234 (0 errors) (1.6×) | 333 (0 errors) (1.1×) | 351 (0 errors) (1.0×) | 325 (0 errors) (1.1×) |


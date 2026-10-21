## Human messages, Russian

Russian-language prompts from oasst2: Cyrillic takes two bytes per character.

Data: OpenAssistant oasst2 (Apache-2.0), human, language ru. Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 316 B | ≈419 B (1.3×) | ≈417 B (1.3×) | ≈710 B (2.2×) | ≈445 B (1.4×) | ≈549 B (1.7×) | ≈1.1 KB (3.7×) |
| medium (~75 chars) | 393 B | ≈483 B (1.2×) | ≈508 B (1.3×) | ≈801 B (2.0×) | ≈537 B (1.4×) | ≈639 B (1.6×) | ≈679 B (1.7×) |
| long (~333 chars) | 684 B | ≈836 B (1.2×) | ≈1.1 KB (1.7×) | ≈1.4 KB (2.1×) | ≈1.2 KB (1.7×) | ≈1.2 KB (1.9×) | ≈1.3 KB (1.9×) |
| very long (~2 567 chars) | 2.6 KB | ≈2.6 KB (1.0×) | ≈6.4 KB (2.5×) | ≈6.7 KB (2.6×) | ≈6.5 KB (2.5×) | ≈6.5 KB (2.5×) | ≈6.6 KB (2.6×) |
| new connection + first message | 1.2 KB | ≈1.2 KB (1.0×) | ≈4.2 KB (3.6×) | ≈4.2 KB (3.5×) | ≈4.3 KB (3.7×) | ≈4.3 KB (3.7×) | ≈23 KB (19.7×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 110 ms | 117 ms (1.1×) | 127 ms (1.2×) | 127 ms (1.2×) | 129 ms (1.2×) | 127 ms (1.2×) | 130 ms (1.2×) |
| medium (~75 chars) | 114 ms | 118 ms (1.0×) | 129 ms (1.1×) | 126 ms (1.1×) | 125 ms (1.1×) | 123 ms (1.1×) | 129 ms (1.1×) |
| long (~333 chars) | 118 ms | 124 ms (1.1×) | 131 ms (1.1×) | 127 ms (1.1×) | 131 ms (1.1×) | 126 ms (1.1×) | 133 ms (1.1×) |
| very long (~2 567 chars) | 123 ms | 120 ms (1.0×) | 133 ms (1.1×) | 138 ms (1.1×) | 131 ms (1.1×) | 127 ms (1.0×) | 140 ms (1.1×) |
| new connection + first message | 123 ms | 242 ms (2.0×) | 509 ms (4.1×) | 385 ms (3.1×) | 374 ms (3.0×) | 480 ms (3.9×) | 280 ms (2.3×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 384 (0 errors) | 359 (0 errors) (1.1×) | 359 (0 errors) (1.1×) | 234 (0 errors) (1.6×) | 338 (0 errors) (1.1×) | 378 (0 errors) (1.0×) | 354 (0 errors) (1.1×) |


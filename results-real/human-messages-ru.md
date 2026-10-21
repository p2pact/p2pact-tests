## Human messages, Russian

Russian-language prompts from oasst2: Cyrillic takes two bytes per character.

Data: OpenAssistant oasst2 (Apache-2.0), human, language ru. Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 204 B | ≈344 B (1.7×) | ≈417 B (2.0×) | ≈710 B (3.5×) | ≈445 B (2.2×) | ≈533 B (2.6×) | ≈1.1 KB (5.3×) |
| medium (~75 chars) | 282 B | ≈416 B (1.5×) | ≈508 B (1.8×) | ≈801 B (2.8×) | ≈537 B (1.9×) | ≈623 B (2.2×) | ≈599 B (2.1×) |
| long (~333 chars) | 569 B | ≈749 B (1.3×) | ≈1.1 KB (2.0×) | ≈1.4 KB (2.6×) | ≈1.2 KB (2.1×) | ≈1.2 KB (2.2×) | ≈1.3 KB (2.4×) |
| very long (~2 567 chars) | 4.2 KB | ≈2.6 KB (0.6×) | ≈6.4 KB (1.5×) | ≈6.7 KB (1.6×) | ≈6.5 KB (1.6×) | ≈6.5 KB (1.6×) | ≈6.7 KB (1.6×) |
| new connection + first message | 847 B | ≈1.1 KB (1.3×) | ≈4.2 KB (5.1×) | ≈4.2 KB (5.0×) | ≈4.3 KB (5.2×) | ≈4.3 KB (5.2×) | ≈21 KB (25.0×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~27 chars) | 114 ms | 106 ms (0.9×) | 110 ms (1.0×) | 114 ms (1.0×) | 137 ms (1.2×) | 106 ms (0.9×) | 100 ms (0.9×) |
| medium (~75 chars) | 114 ms | 105 ms (0.9×) | 110 ms (1.0×) | 115 ms (1.0×) | 122 ms (1.1×) | 105 ms (0.9×) | 100 ms (0.9×) |
| long (~333 chars) | 114 ms | 106 ms (0.9×) | 111 ms (1.0×) | 115 ms (1.0×) | 114 ms (1.0×) | 106 ms (0.9×) | 101 ms (0.9×) |
| very long (~2 567 chars) | 116 ms | 108 ms (0.9×) | 113 ms (1.0×) | 117 ms (1.0×) | 112 ms (1.0×) | 107 ms (0.9×) | 102 ms (0.9×) |
| new connection + first message | 121 ms | 231 ms (1.9×) | 457 ms (3.8×) | 320 ms (2.6×) | 335 ms (2.8×) | 440 ms (3.6×) | 206 ms (1.7×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 415 (0 errors) | 452 (0 errors) (0.9×) | 433 (0 errors) (1.0×) | 278 (0 errors) (1.5×) | 432 (0 errors) (1.0×) | 439 (0 errors) (0.9×) | 480 (0 errors) (0.9×) |


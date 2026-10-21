## Human messages

What people type to an assistant: OpenAssistant oasst2 prompts, all languages.

Data: OpenAssistant oasst2 (Apache-2.0), human. Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (real network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 194 B | ≈331 B (1.7×) | ≈403 B (2.1×) | ≈697 B (3.6×) | ≈430 B (2.2×) | ≈520 B (2.7×) | ≈1.0 KB (5.4×) |
| medium (~80 chars) | 277 B | ≈392 B (1.4×) | ≈469 B (1.7×) | ≈761 B (2.7×) | ≈495 B (1.8×) | ≈584 B (2.1×) | ≈598 B (2.2×) |
| long (~298 chars) | 594 B | ≈770 B (1.3×) | ≈1.1 KB (1.8×) | ≈1.4 KB (2.3×) | ≈1.1 KB (1.9×) | ≈1.2 KB (2.0×) | ≈1.2 KB (2.1×) |
| very long (~2 750 chars) | 1.9 KB | ≈1.8 KB (1.0×) | ≈4.2 KB (2.3×) | ≈4.5 KB (2.4×) | ≈4.3 KB (2.3×) | ≈4.3 KB (2.3×) | ≈4.5 KB (2.4×) |
| new connection + first message | 874 B | ≈1022 B (1.2×) | ≈4.2 KB (5.0×) | ≈4.2 KB (4.9×) | ≈4.3 KB (5.0×) | ≈4.3 KB (5.0×) | ≈21 KB (24.6×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*real network — a real network, nothing emulated*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 105 ms | 115 ms (1.1×) | 102 ms (1.0×) | 106 ms (1.0×) | 100 ms (1.0×) | 93 ms (0.9×) | 107 ms (1.0×) |
| medium (~80 chars) | 107 ms | 114 ms (1.1×) | 103 ms (1.0×) | 106 ms (1.0×) | 102 ms (1.0×) | 93 ms (0.9×) | 108 ms (1.0×) |
| long (~298 chars) | 105 ms | 117 ms (1.1×) | 103 ms (1.0×) | 107 ms (1.0×) | 102 ms (1.0×) | 93 ms (0.9×) | 108 ms (1.0×) |
| very long (~2 750 chars) | 105 ms | 115 ms (1.1×) | 105 ms (1.0×) | 107 ms (1.0×) | 104 ms (1.0×) | 95 ms (0.9×) | 108 ms (1.0×) |
| new connection + first message | 112 ms | 222 ms (2.0×) | 451 ms (4.0×) | 323 ms (2.9×) | 339 ms (3.0×) | 434 ms (3.9×) | 218 ms (1.9×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 466 (0 errors) | 404 (0 errors) (1.2×) | 397 (0 errors) (1.2×) | 282 (0 errors) (1.7×) | 425 (0 errors) (1.1×) | 523 (0 errors) (0.9×) | 446 (0 errors) (1.0×) |


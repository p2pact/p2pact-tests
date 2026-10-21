## Assistant replies

What an assistant answers: oasst2 assistant messages.

Data: OpenAssistant oasst2 (Apache-2.0), assistant. Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 179 B | ≈327 B (1.8×) | ≈401 B (2.2×) | ≈695 B (3.9×) | ≈426 B (2.4×) | ≈545 B (3.0×) | ≈602 B (3.4×) |
| medium (~104 chars) | 258 B | ≈405 B (1.6×) | ≈513 B (2.0×) | ≈806 B (3.1×) | ≈539 B (2.1×) | ≈653 B (2.5×) | ≈615 B (2.4×) |
| long (~694 chars) | 575 B | ≈730 B (1.3×) | ≈1.4 KB (2.5×) | ≈1.7 KB (3.0×) | ≈1.4 KB (2.5×) | ≈1.5 KB (2.7×) | ≈1.6 KB (2.9×) |
| very long (~2 476 chars) | 1.1 KB | ≈1.3 KB (1.1×) | ≈3.2 KB (2.9×) | ≈3.5 KB (3.1×) | ≈3.2 KB (2.8×) | ≈3.2 KB (2.8×) | ≈3.4 KB (3.0×) |
| new connection + first message | 853 B | ≈1003 B (1.2×) | ≈4.2 KB (5.1×) | ≈4.1 KB (5.0×) | ≈4.3 KB (5.1×) | ≈4.3 KB (5.2×) | ≈56 KB (67.4×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 2.5 ms | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.6 ms (1.0×) | 2.6 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) |
| medium (~104 chars) | 2.5 ms | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.6 ms (1.0×) | 2.6 ms (1.0×) | 2.4 ms (1.0×) | 2.5 ms (1.0×) |
| long (~694 chars) | 2.9 ms | 2.8 ms (1.0×) | 2.5 ms (0.9×) | 2.6 ms (0.9×) | 2.6 ms (0.9×) | 2.5 ms (0.9×) | 2.5 ms (0.9×) |
| very long (~2 476 chars) | 3.2 ms | 2.7 ms (0.9×) | 2.5 ms (0.8×) | 2.7 ms (0.8×) | 2.6 ms (0.8×) | 2.6 ms (0.8×) | 2.6 ms (0.8×) |
| new connection + first message | 6.0 ms | 8.3 ms (1.4×) | 9.3 ms (1.5×) | 6.1 ms (1.0×) | 7.0 ms (1.2×) | 9.4 ms (1.6×) | 7.9 ms (1.3×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 19635 (0 errors) | 19302 (0 errors) (1.0×) | 20839 (0 errors) (0.9×) | 18917 (0 errors) (1.0×) | 20508 (0 errors) (1.0×) | 19062 (0 errors) (1.0×) | 17950 (0 errors) (1.1×) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 66 ms | 64 ms (1.0×) | 66 ms (1.0×) | 74 ms (1.1×) | 67 ms (1.0×) | 68 ms (1.0×) | 72 ms (1.1×) |
| medium (~104 chars) | 67 ms | 66 ms (1.0×) | 68 ms (1.0×) | 75 ms (1.1×) | 68 ms (1.0×) | 70 ms (1.0×) | 74 ms (1.1×) |
| long (~694 chars) | 77 ms | 76 ms (1.0×) | 93 ms (1.2×) | 103 ms (1.3×) | 95 ms (1.2×) | 95 ms (1.2×) | 99 ms (1.3×) |
| very long (~2 476 chars) | 88 ms | 87 ms (1.0×) | 124 ms (1.4×) | 132 ms (1.5×) | 125 ms (1.4×) | 125 ms (1.4×) | 132 ms (1.5×) |
| new connection + first message | 79 ms | 76 ms (1.0×) | 280 ms (3.6×) | 216 ms (2.8×) | 215 ms (2.7×) | 277 ms (3.5×) | 272 ms (3.5×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 95 (0 errors) | 273 (0 errors) (0.3×) | 166 (0 errors) (0.6×) | 629 (0 errors) (0.2×) | 176 (0 errors) (0.5×) | 172 (0 errors) (0.6×) | 102 (0 errors) (0.9×) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 82 ms | 82 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) |
| medium (~104 chars) | 81 ms | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) |
| long (~694 chars) | 82 ms | 82 ms (1.0×) | 82 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) |
| very long (~2 476 chars) | 82 ms | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) |
| new connection + first message | 86 ms | 86 ms (1.0×) | 247 ms (2.9×) | 165 ms (1.9×) | 165 ms (1.9×) | 246 ms (2.8×) | 165 ms (1.9×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 491 (0 errors) | 614 (0 errors) (0.8×) | 612 (0 errors) (0.8×) | 597 (0 errors) (0.8×) | 610 (0 errors) (0.8×) | 613 (0 errors) (0.8×) | 566 (0 errors) (0.9×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 202 ms | 201 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) |
| medium (~104 chars) | 202 ms | 202 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) |
| long (~694 chars) | 202 ms | 202 ms (1.0×) | 204 ms (1.0×) | 205 ms (1.0×) | 204 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) |
| very long (~2 476 chars) | 204 ms | 204 ms (1.0×) | 207 ms (1.0×) | 207 ms (1.0×) | 207 ms (1.0×) | 207 ms (1.0×) | 207 ms (1.0×) |
| new connection + first message | 208 ms | 203 ms (1.0×) | 613 ms (2.9×) | 412 ms (2.0×) | 410 ms (2.0×) | 612 ms (2.9×) | 417 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 148 (0 errors) | 247 (0 errors) (0.6×) | 246 (0 errors) (0.6×) | 240 (0 errors) (0.6×) | 246 (0 errors) (0.6×) | 242 (0 errors) (0.6×) | 178 (0 errors) (0.8×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


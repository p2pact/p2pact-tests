## Assistant replies

What an assistant answers: oasst2 assistant messages.

Data: OpenAssistant oasst2 (Apache-2.0), assistant. Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 172 B | ≈327 B (1.9×) | ≈401 B (2.3×) | ≈695 B (4.0×) | ≈426 B (2.5×) | ≈541 B (3.1×) | ≈491 B (2.9×) |
| medium (~104 chars) | 248 B | ≈401 B (1.6×) | ≈513 B (2.1×) | ≈806 B (3.2×) | ≈538 B (2.2×) | ≈649 B (2.6×) | ≈636 B (2.6×) |
| long (~694 chars) | 581 B | ≈736 B (1.3×) | ≈1.4 KB (2.5×) | ≈1.7 KB (3.0×) | ≈1.4 KB (2.5×) | ≈1.5 KB (2.7×) | ≈1.6 KB (2.7×) |
| very long (~2 476 chars) | 1.2 KB | ≈1.3 KB (1.1×) | ≈3.2 KB (2.8×) | ≈3.5 KB (3.1×) | ≈3.2 KB (2.8×) | ≈3.3 KB (2.9×) | ≈3.4 KB (3.0×) |
| new connection + first message | 753 B | ≈933 B (1.2×) | ≈4.2 KB (5.7×) | ≈4.1 KB (5.6×) | ≈4.3 KB (5.8×) | ≈4.3 KB (5.8×) | ≈47 KB (64.2×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 2.5 ms | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) |
| medium (~104 chars) | 2.7 ms | 2.5 ms (0.9×) | 2.5 ms (0.9×) | 2.5 ms (0.9×) | 2.6 ms (1.0×) | 2.5 ms (0.9×) | 2.5 ms (0.9×) |
| long (~694 chars) | 2.9 ms | 2.9 ms (1.0×) | 2.4 ms (0.8×) | 2.6 ms (0.9×) | 2.5 ms (0.8×) | 2.5 ms (0.8×) | 2.5 ms (0.9×) |
| very long (~2 476 chars) | 3.1 ms | 3.1 ms (1.0×) | 2.4 ms (0.8×) | 2.6 ms (0.8×) | 2.6 ms (0.8×) | 2.5 ms (0.8×) | 2.5 ms (0.8×) |
| new connection + first message | 4.9 ms | 5.2 ms (1.1×) | 8.7 ms (1.8×) | 5.5 ms (1.1×) | 6.9 ms (1.4×) | 9.4 ms (1.9×) | 5.6 ms (1.1×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 15742 (0 errors) | 14871 (0 errors) (1.1×) | 21051 (0 errors) (0.7×) | 18961 (0 errors) (0.8×) | 21031 (0 errors) (0.7×) | 19358 (0 errors) (0.8×) | 19616 (0 errors) (0.8×) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 66 ms | 64 ms (1.0×) | 67 ms (1.0×) | 74 ms (1.1×) | 67 ms (1.0×) | 68 ms (1.0×) | 73 ms (1.1×) |
| medium (~104 chars) | 67 ms | 65 ms (1.0×) | 69 ms (1.0×) | 75 ms (1.1×) | 68 ms (1.0×) | 71 ms (1.1×) | 74 ms (1.1×) |
| long (~694 chars) | 77 ms | 73 ms (1.0×) | 94 ms (1.2×) | 100 ms (1.3×) | 95 ms (1.2×) | 95 ms (1.2×) | 99 ms (1.3×) |
| very long (~2 476 chars) | 88 ms | 83 ms (0.9×) | 124 ms (1.4×) | 131 ms (1.5×) | 125 ms (1.4×) | 125 ms (1.4×) | 132 ms (1.5×) |
| new connection + first message | 79 ms | 79 ms (1.0×) | 280 ms (3.5×) | 217 ms (2.7×) | 216 ms (2.7×) | 276 ms (3.5×) | 272 ms (3.4×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 123 (0 errors) | 281 (0 errors) (0.4×) | 166 (0 errors) (0.7×) | 630 (0 errors) (0.2×) | 176 (0 errors) (0.7×) | 172 (0 errors) (0.7×) | 103 (0 errors) (1.2×) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 81 ms | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| medium (~104 chars) | 81 ms | 82 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| long (~694 chars) | 82 ms | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 81 ms (1.0×) | 84 ms (1.0×) |
| very long (~2 476 chars) | 82 ms | 82 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) |
| new connection + first message | 89 ms | 84 ms (0.9×) | 247 ms (2.8×) | 166 ms (1.9×) | 166 ms (1.9×) | 247 ms (2.8×) | 166 ms (1.9×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 531 (0 errors) | 610 (0 errors) (0.9×) | 612 (0 errors) (0.9×) | 596 (0 errors) (0.9×) | 610 (0 errors) (0.9×) | 611 (0 errors) (0.9×) | 573 (0 errors) (0.9×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 202 ms | 203 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) |
| medium (~104 chars) | 203 ms | 203 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) |
| long (~694 chars) | 204 ms | 204 ms (1.0×) | 206 ms (1.0×) | 206 ms (1.0×) | 204 ms (1.0×) | 204 ms (1.0×) | 205 ms (1.0×) |
| very long (~2 476 chars) | 204 ms | 204 ms (1.0×) | 209 ms (1.0×) | 209 ms (1.0×) | 208 ms (1.0×) | 207 ms (1.0×) | 208 ms (1.0×) |
| new connection + first message | 210 ms | 210 ms (1.0×) | 615 ms (2.9×) | 414 ms (2.0×) | 413 ms (2.0×) | 617 ms (2.9×) | 419 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 190 (0 errors) | 245 (0 errors) (0.8×) | 244 (0 errors) (0.8×) | 239 (0 errors) (0.8×) | 245 (0 errors) (0.8×) | 240 (0 errors) (0.8×) | 189 (0 errors) (1.0×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


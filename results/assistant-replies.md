## Assistant replies

What an assistant answers: oasst2 assistant messages.

Data: OpenAssistant oasst2 (Apache-2.0), assistant. Messages per size: 20 (10 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 163 B | ≈318 B (2.0×) | ≈391 B (2.4×) | ≈685 B (4.2×) | ≈416 B (2.6×) | ≈494 B (3.0×) | ≈568 B (3.5×) |
| medium (~103 chars) | 305 B | ≈462 B (1.5×) | ≈583 B (1.9×) | ≈876 B (2.9×) | ≈610 B (2.0×) | ≈688 B (2.3×) | ≈677 B (2.2×) |
| long (~697 chars) | 744 B | ≈893 B (1.2×) | ≈1.6 KB (2.2×) | ≈1.9 KB (2.7×) | ≈1.6 KB (2.2×) | ≈1.7 KB (2.3×) | ≈1.8 KB (2.4×) |
| very long (~2 491 chars) | 1.4 KB | ≈1.5 KB (1.1×) | ≈3.7 KB (2.7×) | ≈3.9 KB (2.9×) | ≈3.7 KB (2.7×) | ≈3.7 KB (2.7×) | ≈3.8 KB (2.8×) |
| new connection + first message | 639 B | ≈871 B (1.4×) | ≈4.2 KB (6.8×) | ≈4.1 KB (6.6×) | ≈4.3 KB (6.8×) | ≈4.3 KB (6.9×) | ≈55 KB (87.4×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 2.8 ms | 2.7 ms (1.0×) | 2.7 ms (0.9×) | 2.7 ms (0.9×) | 2.7 ms (1.0×) | 2.7 ms (0.9×) | 2.6 ms (0.9×) |
| medium (~103 chars) | 2.9 ms | 2.9 ms (1.0×) | 2.7 ms (0.9×) | 2.7 ms (0.9×) | 2.6 ms (0.9×) | 2.7 ms (0.9×) | 2.7 ms (0.9×) |
| long (~697 chars) | 3.0 ms | 2.8 ms (0.9×) | 2.7 ms (0.9×) | 2.7 ms (0.9×) | 2.7 ms (0.9×) | 2.7 ms (0.9×) | 2.6 ms (0.9×) |
| very long (~2 491 chars) | 2.9 ms | 3.1 ms (1.0×) | 2.7 ms (0.9×) | 2.7 ms (0.9×) | 2.7 ms (0.9×) | 2.6 ms (0.9×) | 2.7 ms (0.9×) |
| new connection + first message | 4.6 ms | 7.4 ms (1.6×) | 9.7 ms (2.1×) | 5.6 ms (1.2×) | 5.8 ms (1.3×) | 8.2 ms (1.8×) | 5.7 ms (1.2×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 19696 (0 errors) | 19622 (0 errors) (1.0×) | 19308 (0 errors) (1.0×) | 18106 (0 errors) (1.1×) | 21822 (0 errors) (0.9×) | 18954 (0 errors) (1.0×) | 19748 (0 errors) (1.0×) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 66 ms | 65 ms (1.0×) | 67 ms (1.0×) | 75 ms (1.1×) | 67 ms (1.0×) | 69 ms (1.0×) | 72 ms (1.1×) |
| medium (~103 chars) | 70 ms | 68 ms (1.0×) | 71 ms (1.0×) | 79 ms (1.1×) | 71 ms (1.0×) | 73 ms (1.0×) | 77 ms (1.1×) |
| long (~697 chars) | 81 ms | 81 ms (1.0×) | 102 ms (1.3×) | 110 ms (1.4×) | 103 ms (1.3×) | 102 ms (1.3×) | 108 ms (1.3×) |
| very long (~2 491 chars) | 97 ms | 97 ms (1.0×) | 138 ms (1.4×) | 146 ms (1.5×) | 139 ms (1.4×) | 139 ms (1.4×) | 146 ms (1.5×) |
| new connection + first message | 74 ms | 72 ms (1.0×) | 283 ms (3.8×) | 217 ms (2.9×) | 217 ms (2.9×) | 278 ms (3.7×) | 271 ms (3.6×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 108 (0 errors) | 91 (0 errors) (1.2×) | 167 (0 errors) (0.6×) | 626 (0 errors) (0.2×) | 177 (0 errors) (0.6×) | 172 (0 errors) (0.6×) | 107 (0 errors) (1.0×) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 82 ms | 81 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) |
| medium (~103 chars) | 81 ms | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) |
| long (~697 chars) | 82 ms | 82 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| very long (~2 491 chars) | 82 ms | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) |
| new connection + first message | 85 ms | 84 ms (1.0×) | 248 ms (2.9×) | 168 ms (2.0×) | 166 ms (2.0×) | 248 ms (2.9×) | 167 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 507 (0 errors) | 609 (0 errors) (0.8×) | 607 (0 errors) (0.8×) | 590 (0 errors) (0.9×) | 606 (0 errors) (0.8×) | 606 (0 errors) (0.8×) | 533 (0 errors) (1.0×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 202 ms | 202 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) | 204 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) |
| medium (~103 chars) | 203 ms | 203 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) |
| long (~697 chars) | 204 ms | 204 ms (1.0×) | 205 ms (1.0×) | 206 ms (1.0×) | 205 ms (1.0×) | 205 ms (1.0×) | 205 ms (1.0×) |
| very long (~2 491 chars) | 205 ms | 205 ms (1.0×) | 209 ms (1.0×) | 210 ms (1.0×) | 209 ms (1.0×) | 208 ms (1.0×) | 208 ms (1.0×) |
| new connection + first message | 207 ms | 204 ms (1.0×) | 615 ms (3.0×) | 414 ms (2.0×) | 412 ms (2.0×) | 616 ms (3.0×) | 416 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 148 (0 errors) | 244 (0 errors) (0.6×) | 245 (0 errors) (0.6×) | 239 (0 errors) (0.6×) | 245 (0 errors) (0.6×) | 240 (0 errors) (0.6×) | 200 (0 errors) (0.7×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


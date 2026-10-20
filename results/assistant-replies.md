## Assistant replies

What an assistant answers: oasst2 assistant messages.

Data: OpenAssistant oasst2 (Apache-2.0), assistant. Messages per size: 20 (2 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 149 B | ≈305 B (2.0×) | ≈368 B (2.5×) | ≈662 B (4.4×) | ≈394 B (2.6×) | ≈513 B (3.4×) | ≈945 B (6.3×) |
| medium (~103 chars) | 292 B | ≈448 B (1.5×) | ≈560 B (1.9×) | ≈853 B (2.9×) | ≈586 B (2.0×) | ≈709 B (2.4×) | ≈649 B (2.2×) |
| long (~697 chars) | 724 B | ≈880 B (1.2×) | ≈1.6 KB (2.2×) | ≈1.9 KB (2.7×) | ≈1.6 KB (2.3×) | ≈1.7 KB (2.4×) | ≈1.8 KB (2.5×) |
| very long (~2 491 chars) | 1.3 KB | ≈1.5 KB (1.1×) | ≈3.6 KB (2.7×) | ≈3.9 KB (3.0×) | ≈3.7 KB (2.8×) | ≈3.7 KB (2.8×) | ≈3.8 KB (2.9×) |
| new connection + first message | 782 B | ≈930 B (1.2×) | ≈4.4 KB (5.7×) | ≈4.3 KB (5.6×) | ≈4.4 KB (5.8×) | ≈4.4 KB (5.7×) | ≈15 KB (20.2×) |

**Time from sending a message to its ack** (median):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 2.4 ms | 2.4 ms | 2.4 ms | 2.5 ms | 2.5 ms | 2.5 ms | 2.5 ms |
| medium (~103 chars) | 2.5 ms | 2.5 ms | 2.4 ms | 2.5 ms | 2.5 ms | 2.4 ms | 2.4 ms |
| long (~697 chars) | 2.5 ms | 2.5 ms | 2.4 ms | 2.5 ms | 2.5 ms | 2.4 ms | 2.5 ms |
| very long (~2 491 chars) | 2.6 ms | 2.6 ms | 2.4 ms | 2.5 ms | 2.5 ms | 2.4 ms | 2.5 ms |
| new connection + first message | 3.3 ms | 3.8 ms | 8.2 ms | 5.8 ms | 5.8 ms | 9.2 ms | 7.4 ms |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 65 ms | 64 ms | 66 ms | 73 ms | 65 ms | 68 ms | 72 ms |
| medium (~103 chars) | 69 ms | 67 ms | 70 ms | 76 ms | 69 ms | 72 ms | 76 ms |
| long (~697 chars) | 81 ms | 81 ms | 100 ms | 107 ms | 100 ms | 102 ms | 107 ms |
| very long (~2 491 chars) | 98 ms | 95 ms | 137 ms | 144 ms | 138 ms | 137 ms | 146 ms |
| new connection + first message | 73 ms | 71 ms | 277 ms | 216 ms | 214 ms | 276 ms | 267 ms |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 81 ms | 81 ms | 81 ms | 82 ms | 83 ms | 82 ms | 82 ms |
| medium (~103 chars) | 81 ms | 81 ms | 82 ms | 82 ms | 82 ms | 82 ms | 82 ms |
| long (~697 chars) | 81 ms | 81 ms | 82 ms | 82 ms | 82 ms | 83 ms | 82 ms |
| very long (~2 491 chars) | 81 ms | 81 ms | 82 ms | 82 ms | 82 ms | 82 ms | 81 ms |
| new connection + first message | 82 ms | 81 ms | 244 ms | 162 ms | 165 ms | 246 ms | 164 ms |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~18 chars) | 201 ms | 201 ms | 201 ms | 203 ms | 202 ms | 202 ms | 202 ms |
| medium (~103 chars) | 202 ms | 202 ms | 202 ms | 204 ms | 203 ms | 201 ms | 202 ms |
| long (~697 chars) | 203 ms | 203 ms | 204 ms | 206 ms | 204 ms | 204 ms | 205 ms |
| very long (~2 491 chars) | 204 ms | 204 ms | 208 ms | 209 ms | 208 ms | 207 ms | 208 ms |
| new connection + first message | 202 ms | 202 ms | 611 ms | 411 ms | 410 ms | 611 ms | 416 ms |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


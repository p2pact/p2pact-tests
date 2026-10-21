## Human messages

What people type to an assistant: OpenAssistant oasst2 prompts, all languages.

Data: OpenAssistant oasst2 (Apache-2.0), human. Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 173 B | ≈327 B (1.9×) | ≈403 B (2.3×) | ≈697 B (4.0×) | ≈429 B (2.5×) | ≈514 B (3.0×) | ≈956 B (5.5×) |
| medium (~80 chars) | 208 B | ≈362 B (1.7×) | ≈469 B (2.3×) | ≈761 B (3.7×) | ≈492 B (2.4×) | ≈605 B (2.9×) | ≈539 B (2.6×) |
| long (~298 chars) | 462 B | ≈616 B (1.3×) | ≈1.1 KB (2.4×) | ≈1.4 KB (3.0×) | ≈1.1 KB (2.4×) | ≈1.2 KB (2.7×) | ≈1.2 KB (2.6×) |
| very long (~2 750 chars) | 1.5 KB | ≈1.6 KB (1.1×) | ≈4.2 KB (2.9×) | ≈4.5 KB (3.1×) | ≈4.3 KB (2.9×) | ≈4.3 KB (2.9×) | ≈4.4 KB (3.0×) |
| new connection + first message | 715 B | ≈936 B (1.3×) | ≈4.2 KB (6.1×) | ≈4.2 KB (5.9×) | ≈4.3 KB (6.1×) | ≈4.3 KB (6.2×) | ≈44 KB (63.2×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 2.5 ms | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.6 ms (1.0×) | 2.4 ms (1.0×) | 2.4 ms (1.0×) |
| medium (~80 chars) | 2.5 ms | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.6 ms (1.0×) | 2.6 ms (1.0×) | 2.5 ms (1.0×) | 2.4 ms (1.0×) |
| long (~298 chars) | 2.6 ms | 2.6 ms (1.0×) | 2.5 ms (1.0×) | 2.6 ms (1.0×) | 2.4 ms (0.9×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) |
| very long (~2 750 chars) | 2.6 ms | 2.9 ms (1.1×) | 2.5 ms (1.0×) | 2.6 ms (1.0×) | 2.6 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) |
| new connection + first message | 6.7 ms | 9.2 ms (1.4×) | 9.1 ms (1.4×) | 6.8 ms (1.0×) | 6.8 ms (1.0×) | 9.4 ms (1.4×) | 7.3 ms (1.1×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 20664 (0 errors) | 20884 (0 errors) (1.0×) | 20971 (0 errors) (1.0×) | 19000 (0 errors) (1.1×) | 21093 (0 errors) (1.0×) | 19457 (0 errors) (1.1×) | 21382 (0 errors) (1.0×) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 66 ms | 64 ms (1.0×) | 66 ms (1.0×) | 74 ms (1.1×) | 66 ms (1.0×) | 68 ms (1.0×) | 73 ms (1.1×) |
| medium (~80 chars) | 67 ms | 65 ms (1.0×) | 68 ms (1.0×) | 75 ms (1.1×) | 68 ms (1.0×) | 69 ms (1.0×) | 74 ms (1.1×) |
| long (~298 chars) | 72 ms | 70 ms (1.0×) | 78 ms (1.1×) | 85 ms (1.2×) | 78 ms (1.1×) | 80 ms (1.1×) | 83 ms (1.2×) |
| very long (~2 750 chars) | 100 ms | 97 ms (1.0×) | 160 ms (1.6×) | 168 ms (1.7×) | 161 ms (1.6×) | 162 ms (1.6×) | 171 ms (1.7×) |
| new connection + first message | 78 ms | 77 ms (1.0×) | 280 ms (3.6×) | 216 ms (2.8×) | 216 ms (2.8×) | 277 ms (3.6×) | 272 ms (3.5×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 197 (0 errors) | 338 (0 errors) (0.6×) | 200 (0 errors) (1.0×) | 638 (0 errors) (0.3×) | 216 (0 errors) (0.9×) | 207 (0 errors) (1.0×) | 115 (0 errors) (1.7×) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 82 ms | 82 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) |
| medium (~80 chars) | 82 ms | 82 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| long (~298 chars) | 82 ms | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) |
| very long (~2 750 chars) | 83 ms | 82 ms (1.0×) | 82 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 83 ms (1.0×) | 82 ms (1.0×) |
| new connection + first message | 88 ms | 88 ms (1.0×) | 247 ms (2.8×) | 169 ms (1.9×) | 166 ms (1.9×) | 248 ms (2.8×) | 167 ms (1.9×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 565 (0 errors) | 609 (0 errors) (0.9×) | 607 (0 errors) (0.9×) | 590 (0 errors) (1.0×) | 605 (0 errors) (0.9×) | 606 (0 errors) (0.9×) | 574 (0 errors) (1.0×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 202 ms | 202 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 203 ms (1.0×) | 205 ms (1.0×) |
| medium (~80 chars) | 202 ms | 202 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) | 204 ms (1.0×) | 202 ms (1.0×) | 204 ms (1.0×) |
| long (~298 chars) | 203 ms | 203 ms (1.0×) | 204 ms (1.0×) | 205 ms (1.0×) | 204 ms (1.0×) | 204 ms (1.0×) | 205 ms (1.0×) |
| very long (~2 750 chars) | 205 ms | 205 ms (1.0×) | 209 ms (1.0×) | 211 ms (1.0×) | 210 ms (1.0×) | 210 ms (1.0×) | 210 ms (1.0×) |
| new connection + first message | 207 ms | 208 ms (1.0×) | 616 ms (3.0×) | 413 ms (2.0×) | 415 ms (2.0×) | 616 ms (3.0×) | 419 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 224 (0 errors) | 246 (0 errors) (0.9×) | 245 (0 errors) (0.9×) | 239 (0 errors) (0.9×) | 246 (0 errors) (0.9×) | 241 (0 errors) (0.9×) | 183 (0 errors) (1.2×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


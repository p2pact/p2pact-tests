## Human messages

What people type to an assistant: OpenAssistant oasst2 prompts, all languages.

Data: OpenAssistant oasst2 (Apache-2.0), human. Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (local network; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 172 B | ≈327 B (1.9×) | ≈403 B (2.3×) | ≈697 B (4.1×) | ≈428 B (2.5×) | ≈548 B (3.2×) | ≈529 B (3.1×) |
| medium (~80 chars) | 208 B | ≈362 B (1.7×) | ≈469 B (2.3×) | ≈761 B (3.7×) | ≈493 B (2.4×) | ≈583 B (2.8×) | ≈580 B (2.8×) |
| long (~298 chars) | 461 B | ≈616 B (1.3×) | ≈1.1 KB (2.4×) | ≈1.4 KB (3.0×) | ≈1.1 KB (2.4×) | ≈1.2 KB (2.7×) | ≈1.2 KB (2.7×) |
| very long (~2 750 chars) | 1.5 KB | ≈1.6 KB (1.1×) | ≈4.2 KB (2.9×) | ≈4.5 KB (3.1×) | ≈4.3 KB (2.9×) | ≈4.3 KB (2.9×) | ≈4.7 KB (3.2×) |
| new connection + first message | 792 B | ≈955 B (1.2×) | ≈4.2 KB (5.5×) | ≈4.2 KB (5.4×) | ≈4.3 KB (5.6×) | ≈4.3 KB (5.6×) | ≈42 KB (54.7×) |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*local network — 1 ms one way*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 2.6 ms | 2.4 ms (0.9×) | 2.4 ms (0.9×) | 2.5 ms (1.0×) | 2.6 ms (1.0×) | 2.4 ms (0.9×) | 2.5 ms (1.0×) |
| medium (~80 chars) | 2.4 ms | 2.4 ms (1.0×) | 2.4 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) | 2.5 ms (1.0×) |
| long (~298 chars) | 2.8 ms | 2.7 ms (1.0×) | 2.5 ms (0.9×) | 2.5 ms (0.9×) | 2.6 ms (0.9×) | 2.5 ms (0.9×) | 2.5 ms (0.9×) |
| very long (~2 750 chars) | 3.4 ms | 2.9 ms (0.9×) | 2.5 ms (0.7×) | 2.6 ms (0.8×) | 2.7 ms (0.8×) | 2.4 ms (0.7×) | 2.5 ms (0.8×) |
| new connection + first message | 8.0 ms | 5.0 ms (0.6×) | 8.0 ms (1.0×) | 5.7 ms (0.7×) | 7.0 ms (0.9×) | 8.6 ms (1.1×) | 5.7 ms (0.7×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 19139 (0 errors) | 19806 (0 errors) (1.0×) | 20890 (0 errors) (0.9×) | 17454 (0 errors) (1.1×) | 21777 (0 errors) (0.9×) | 20088 (0 errors) (1.0×) | 21451 (0 errors) (0.9×) |

*bluetooth — 30 ms one way, 40 KB/s*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 66 ms | 65 ms (1.0×) | 67 ms (1.0×) | 74 ms (1.1×) | 67 ms (1.0×) | 68 ms (1.0×) | 73 ms (1.1×) |
| medium (~80 chars) | 67 ms | 64 ms (1.0×) | 68 ms (1.0×) | 75 ms (1.1×) | 68 ms (1.0×) | 70 ms (1.1×) | 74 ms (1.1×) |
| long (~298 chars) | 71 ms | 69 ms (1.0×) | 77 ms (1.1×) | 85 ms (1.2×) | 79 ms (1.1×) | 81 ms (1.1×) | 84 ms (1.2×) |
| very long (~2 750 chars) | 100 ms | 90 ms (0.9×) | 160 ms (1.6×) | 169 ms (1.7×) | 161 ms (1.6×) | 161 ms (1.6×) | 171 ms (1.7×) |
| new connection + first message | 77 ms | 79 ms (1.0×) | 281 ms (3.7×) | 217 ms (2.8×) | 216 ms (2.8×) | 278 ms (3.6×) | 272 ms (3.5×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 175 (0 errors) | 349 (0 errors) (0.5×) | 200 (0 errors) (0.9×) | 637 (0 errors) (0.3×) | 216 (0 errors) (0.8×) | 209 (0 errors) (0.8×) | 113 (0 errors) (1.5×) |

*internet — 40 ms one way, 1% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 81 ms | 81 ms (1.0×) | 81 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 81 ms (1.0×) |
| medium (~80 chars) | 81 ms | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 81 ms (1.0×) |
| long (~298 chars) | 81 ms | 82 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) |
| very long (~2 750 chars) | 82 ms | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 82 ms (1.0×) | 81 ms (1.0×) | 81 ms (1.0×) |
| new connection + first message | 89 ms | 88 ms (1.0×) | 247 ms (2.8×) | 165 ms (1.8×) | 165 ms (1.8×) | 247 ms (2.8×) | 163 ms (1.8×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 511 (0 errors) | 610 (0 errors) (0.8×) | 613 (0 errors) (0.8×) | 595 (0 errors) (0.9×) | 612 (0 errors) (0.8×) | 614 (0 errors) (0.8×) | 569 (0 errors) (0.9×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*mobile — 100 ms one way, 500 KB/s, 3% loss*

| Message | P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---|---:|---:|---:|---:|---:|---:|---:|
| very short (~24 chars) | 201 ms | 202 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) |
| medium (~80 chars) | 201 ms | 201 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) | 202 ms (1.0×) |
| long (~298 chars) | 202 ms | 202 ms (1.0×) | 203 ms (1.0×) | 204 ms (1.0×) | 203 ms (1.0×) | 202 ms (1.0×) | 203 ms (1.0×) |
| very long (~2 750 chars) | 204 ms | 204 ms (1.0×) | 210 ms (1.0×) | 210 ms (1.0×) | 209 ms (1.0×) | 209 ms (1.0×) | 210 ms (1.0×) |
| new connection + first message | 205 ms | 211 ms (1.0×) | 614 ms (3.0×) | 411 ms (2.0×) | 411 ms (2.0×) | 613 ms (3.0×) | 417 ms (2.0×) |

Load: 2000 messages, 50 at a time — messages per second (errors; in brackets — how many times fewer than P2PACT UDP):

| P2PACT UDP | P2PACT TCP | WebSocket | HTTPS+JSON | HTTP/2 | gRPC | HTTP/3 |
|---:|---:|---:|---:|---:|---:|---:|
| 181 (0 errors) | 247 (0 errors) (0.7×) | 246 (0 errors) (0.7×) | 240 (0 errors) (0.8×) | 246 (0 errors) (0.7×) | 243 (0 errors) (0.7×) | 198 (0 errors) (0.9×) |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


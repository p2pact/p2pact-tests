# P2PACT — summary

Each cell: the value, and in brackets how many times worse it is than P2PACT over UDP (longer time, more bytes, fewer messages a second). Above 1× P2PACT is better, below 1× it is worse. Every message was checked by a fingerprint of what the receiver got; errors are listed in the detailed results.

## Real network

A laptop on home Wi-Fi to a cloud server, nothing emulated. Details: [results-real/RESULTS.md](results-real/RESULTS.md).

| Protocol | Medium message | Long message | Huge message (200 000 chars) | New connection + first message | Load: messages a second (50 at a time) | Bytes per medium message |
|---|---:|---:|---:|---:|---:|---:|
| P2PACT UDP | **107 ms** | **105 ms** | **133 ms** | **112 ms** | **466/s** | **277 B** |
| P2PACT TCP | 114 ms (1.1×) | 117 ms (1.1×) | 161 ms (1.2×) | 222 ms (2.0×) | 404/s (1.2×) | 392 B (1.4×) |
| WebSocket | 103 ms (1.0×) | 103 ms (1.0×) | 136 ms (1.0×) | 451 ms (4.0×) | 397/s (1.2×) | 469 B (1.7×) |
| HTTPS+JSON | 106 ms (1.0×) | 107 ms (1.0×) | 132 ms (1.0×) | 323 ms (2.9×) | 282/s (1.7×) | 761 B (2.7×) |
| HTTP/2 | 102 ms (1.0×) | 102 ms (1.0×) | 133 ms (1.0×) | 339 ms (3.0×) | 425/s (1.1×) | 495 B (1.8×) |
| gRPC | 93 ms (0.9×) | 93 ms (0.9×) | 144 ms (1.1×) | 434 ms (3.9×) | 523/s (0.9×) | 584 B (2.1×) |
| HTTP/3 | 108 ms (1.0×) | 108 ms (1.0×) | 567 ms (4.3×) | 218 ms (1.9×) | 446/s (1.0×) | 598 B (2.2×) |

## Emulated links

One machine, delay, bandwidth and loss emulated. Details: [results/RESULTS.md](results/RESULTS.md).

| Protocol | Long, Bluetooth | Huge, Bluetooth | Huge, local network | Huge, internet 1% loss * | Huge, mobile 3% loss * | New connection, mobile | Load, internet 1% loss * | Load, local network |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| P2PACT UDP | **96 ms** | **1.90 s** | **11 ms** | **180 ms** | **705 ms** | **205 ms** | **511/s** | **19139/s** |
| P2PACT TCP | 96 ms (1.0×) | 1.86 s (1.0×) | 8.0 ms (0.7×) | 86 ms (0.5×) | 351 ms (0.5×) | 211 ms (1.0×) | 610/s (0.8×) | 19806/s (1.0×) |
| WebSocket | 175 ms (1.8×) | 4.96 s (2.6×) | 9.3 ms (0.9×) | 84 ms (0.5×) | 601 ms (0.9×) | 614 ms (3.0×) | 613/s (0.8×) | 20890/s (0.9×) |
| HTTPS+JSON | 181 ms (1.9×) | 4.97 s (2.6×) | 4.9 ms (0.5×) | 84 ms (0.5×) | 599 ms (0.8×) | 411 ms (2.0×) | 595/s (0.9×) | 17454/s (1.1×) |
| HTTP/2 | 175 ms (1.8×) | 4.97 s (2.6×) | 4.9 ms (0.5×) | 90 ms (0.5×) | 599 ms (0.8×) | 411 ms (2.0×) | 612/s (0.8×) | 21777/s (0.9×) |
| gRPC | 169 ms (1.8×) | 4.96 s (2.6×) | 4.5 ms (0.4×) | 87 ms (0.5×) | 599 ms (0.8×) | 613 ms (3.0×) | 614/s (0.8×) | 20088/s (1.0×) |
| HTTP/3 | 184 ms (1.9×) | 5.17 s (2.7×) | 5.5 ms (0.5×) | 686 ms (3.8×) | 3.12 s (4.4×) | 417 ms (2.0×) | 569/s (0.9×) | 21451/s (0.9×) |

\* TCP-based protocols (P2PACT TCP, WebSocket, HTTPS, HTTP/2, gRPC) run without loss in the emulation, and the emulating proxy terminates TCP, so these columns are optimistic for them. On lossy links the fair comparison is with HTTP/3, which goes through the same emulated UDP link as P2PACT.

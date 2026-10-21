# P2PACT — summary

Each cell: the value, and in brackets how many times worse it is than P2PACT over UDP (longer time, more bytes, fewer messages a second). Above 1× P2PACT is better, below 1× it is worse. Every message was checked by a fingerprint of what the receiver got; errors are listed in the detailed results.

## Real network

A laptop on home Wi-Fi to a cloud server, nothing emulated. Details: [results-real/RESULTS.md](results-real/RESULTS.md).

| Protocol | Medium message | Long message | Huge message (200 000 chars) | New connection + first message | Load: messages a second (50 at a time) | Bytes per medium message |
|---|---:|---:|---:|---:|---:|---:|
| P2PACT UDP | **116 ms** | **128 ms** | **137 ms** | **134 ms** | **367/s** | **232 B** |
| P2PACT TCP | 113 ms (1.0×) | 117 ms (0.9×) | 148 ms (1.1×) | 247 ms (1.8×) | 349/s (1.1×) | 386 B (1.7×) |
| WebSocket | 114 ms (1.0×) | 116 ms (0.9×) | 192 ms (1.4×) | 631 ms (4.7×) | 370/s (1.0×) | 469 B (2.0×) |
| HTTPS+JSON | 128 ms (1.1×) | 129 ms (1.0×) | 169 ms (1.2×) | 396 ms (3.0×) | 234/s (1.6×) | 761 B (3.3×) |
| HTTP/2 | 132 ms (1.1×) | 143 ms (1.1×) | 173 ms (1.3×) | 395 ms (2.9×) | 333/s (1.1×) | 495 B (2.1×) |
| gRPC | 133 ms (1.1×) | 135 ms (1.1×) | 272 ms (2.0×) | 488 ms (3.6×) | 351/s (1.0×) | 592 B (2.6×) |
| HTTP/3 | 134 ms (1.2×) | 130 ms (1.0×) | 174 ms (1.3×) | 276 ms (2.1×) | 325/s (1.1×) | 558 B (2.4×) |

## Emulated links

One machine, delay, bandwidth and loss emulated. Details: [results/RESULTS.md](results/RESULTS.md).

| Protocol | Long, Bluetooth | Huge, Bluetooth | Huge, local network | Huge, internet 1% loss * | Huge, mobile 3% loss * | New connection, mobile | Load, internet 1% loss * | Load, local network |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| P2PACT UDP | **94 ms** | **1.91 s** | **9.6 ms** | **180 ms** | **678 ms** | **207 ms** | **565/s** | **20664/s** |
| P2PACT TCP | 92 ms (1.0×) | 1.87 s (1.0×) | 12 ms (1.2×) | 96 ms (0.5×) | 356 ms (0.5×) | 208 ms (1.0×) | 609/s (0.9×) | 20884/s (1.0×) |
| WebSocket | 172 ms (1.8×) | 4.96 s (2.6×) | 9.0 ms (0.9×) | 92 ms (0.5×) | 602 ms (0.9×) | 616 ms (3.0×) | 607/s (0.9×) | 20971/s (1.0×) |
| HTTPS+JSON | 181 ms (1.9×) | 4.97 s (2.6×) | 9.4 ms (1.0×) | 92 ms (0.5×) | 601 ms (0.9×) | 413 ms (2.0×) | 590/s (1.0×) | 19000/s (1.1×) |
| HTTP/2 | 173 ms (1.9×) | 4.97 s (2.6×) | 9.4 ms (1.0×) | 89 ms (0.5×) | 601 ms (0.9×) | 415 ms (2.0×) | 605/s (0.9×) | 21093/s (1.0×) |
| gRPC | 169 ms (1.8×) | 4.97 s (2.6×) | 8.0 ms (0.8×) | 89 ms (0.5×) | 600 ms (0.9×) | 616 ms (3.0×) | 606/s (0.9×) | 19457/s (1.1×) |
| HTTP/3 | 185 ms (2.0×) | 5.18 s (2.7×) | 9.1 ms (1.0×) | 829 ms (4.6×) | 2.75 s (4.1×) | 419 ms (2.0×) | 574/s (1.0×) | 21382/s (1.0×) |

\* TCP-based protocols (P2PACT TCP, WebSocket, HTTPS, HTTP/2, gRPC) run without loss in the emulation, and the emulating proxy terminates TCP, so these columns are optimistic for them. On lossy links the fair comparison is with HTTP/3, which goes through the same emulated UDP link as P2PACT.

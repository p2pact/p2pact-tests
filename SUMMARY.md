# P2PACT — summary

Each cell: the value, and in brackets how many times worse it is than P2PACT over UDP (longer time, more bytes, fewer messages a second). Above 1× P2PACT is better, below 1× it is worse. Every message was checked by a fingerprint of what the receiver got; errors are listed in the detailed results.

## Real network

A laptop on home Wi-Fi to a cloud server, nothing emulated. Details: [results-real/RESULTS.md](results-real/RESULTS.md).

| Protocol | Medium message | Long message | Huge message (200 000 chars) | New connection + first message | Load: messages a second (50 at a time) | Bytes per medium message |
|---|---:|---:|---:|---:|---:|---:|
| P2PACT UDP | **120 ms** | **134 ms** | **160 ms** | **121 ms** | **355/s** | **232 B** |
| P2PACT TCP | 134 ms (1.1×) | 142 ms (1.1×) | 158 ms (1.0×) | 242 ms (2.0×) | 343/s (1.0×) | 386 B (1.7×) |
| WebSocket | 135 ms (1.1×) | 131 ms (1.0×) | 171 ms (1.1×) | 478 ms (3.9×) | 357/s (1.0×) | 462 B (2.0×) |
| HTTPS+JSON | 137 ms (1.1×) | 141 ms (1.1×) | 184 ms (1.1×) | 401 ms (3.3×) | 230/s (1.5×) | 754 B (3.2×) |
| HTTP/2 | 125 ms (1.0×) | 120 ms (0.9×) | 186 ms (1.2×) | 397 ms (3.3×) | 368/s (1.0×) | 489 B (2.1×) |
| gRPC | 121 ms (1.0×) | 129 ms (1.0×) | 201 ms (1.3×) | 517 ms (4.3×) | 398/s (0.9×) | 588 B (2.5×) |
| HTTP/3 | 119 ms (1.0×) | 117 ms (0.9×) | 220 ms (1.4×) | 244 ms (2.0×) | 392/s (0.9×) | 610 B (2.6×) |

## Emulated links

One machine, delay, bandwidth and loss emulated. Details: [results/RESULTS.md](results/RESULTS.md).

| Protocol | Long, Bluetooth | Huge, Bluetooth | Huge, local network | Huge, internet 1% loss * | Huge, mobile 3% loss * | New connection, mobile | Load, internet 1% loss * | Load, local network |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| P2PACT UDP | **107 ms** | **1.96 s** | **9.5 ms** | **95 ms** | **653 ms** | **205 ms** | **594/s** | **19035/s** |
| P2PACT TCP | 103 ms (1.0×) | 1.92 s (1.0×) | 8.3 ms (0.9×) | 91 ms (1.0×) | 357 ms (0.5×) | 206 ms (1.0×) | 607/s (1.0×) | 19731/s (1.0×) |
| WebSocket | 190 ms (1.8×) | 4.96 s (2.5×) | 6.5 ms (0.7×) | 89 ms (0.9×) | 600 ms (0.9×) | 616 ms (3.0×) | 606/s (1.0×) | 19714/s (1.0×) |
| HTTPS+JSON | 198 ms (1.9×) | 4.97 s (2.5×) | 7.6 ms (0.8×) | 90 ms (0.9×) | 601 ms (0.9×) | 412 ms (2.0×) | 591/s (1.0×) | 16820/s (1.1×) |
| HTTP/2 | 191 ms (1.8×) | 4.97 s (2.5×) | 7.4 ms (0.8×) | 88 ms (0.9×) | 600 ms (0.9×) | 413 ms (2.0×) | 606/s (1.0×) | 20576/s (0.9×) |
| gRPC | 185 ms (1.7×) | 4.97 s (2.5×) | 5.6 ms (0.6×) | 88 ms (0.9×) | 599 ms (0.9×) | 618 ms (3.0×) | 607/s (1.0×) | 18243/s (1.0×) |
| HTTP/3 | 200 ms (1.9×) | 5.17 s (2.6×) | 5.9 ms (0.6×) | 756 ms (7.9×) | 2.75 s (4.2×) | 418 ms (2.0×) | 547/s (1.1×) | 18718/s (1.0×) |

\* TCP-based protocols (P2PACT TCP, WebSocket, HTTPS, HTTP/2, gRPC) run without loss in the emulation, and the emulating proxy terminates TCP, so these columns are optimistic for them. On lossy links the fair comparison is with HTTP/3, which goes through the same emulated UDP link as P2PACT.

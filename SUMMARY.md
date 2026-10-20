# P2PACT — summary

Each cell: the value, and in brackets how many times worse it is than P2PACT over UDP (longer time, more bytes, fewer messages a second). Above 1× P2PACT is better, below 1× it is worse. Every message was checked by a fingerprint of what the receiver got; errors are listed in the detailed results.

## Real network

A laptop on home Wi-Fi to a cloud server, nothing emulated. Details: [results-real/RESULTS.md](results-real/RESULTS.md).

| Protocol | Medium message | Long message | Huge message (200 000 chars) | New connection + first message | Load: messages a second (50 at a time) | Bytes per medium message |
|---|---:|---:|---:|---:|---:|---:|
| P2PACT UDP | **136 ms** | **138 ms** | **181 ms** | **138 ms** | **368/s** | **226 B** |
| P2PACT TCP | 142 ms (1.0×) | 151 ms (1.1×) | 197 ms (1.1×) | 286 ms (2.1×) | 311/s (1.2×) | 388 B (1.7×) |
| WebSocket | 145 ms (1.1×) | 139 ms (1.0×) | 199 ms (1.1×) | 572 ms (4.1×) | 339/s (1.1×) | 462 B (2.0×) |
| HTTPS+JSON | 141 ms (1.0×) | 154 ms (1.1×) | 231 ms (1.3×) | 450 ms (3.3×) | 209/s (1.8×) | 754 B (3.3×) |
| HTTP/2 | 156 ms (1.1×) | 153 ms (1.1×) | 206 ms (1.1×) | 464 ms (3.4×) | 308/s (1.2×) | 490 B (2.2×) |
| gRPC | 148 ms (1.1×) | 154 ms (1.1×) | 188 ms (1.0×) | 553 ms (4.0×) | 297/s (1.2×) | 578 B (2.6×) |
| HTTP/3 | 145 ms (1.1×) | 155 ms (1.1×) | 440 ms (2.4×) | 304 ms (2.2×) | 333/s (1.1×) | 634 B (2.8×) |

## Emulated links

One machine, delay, bandwidth and loss emulated. Details: [results/RESULTS.md](results/RESULTS.md).

| Protocol | Long, Bluetooth | Huge, Bluetooth | Huge, local network | Huge, internet 1% loss * | Huge, mobile 3% loss * | New connection, mobile | Load, internet 1% loss * | Load, local network |
|---|---:|---:|---:|---:|---:|---:|---:|---:|
| P2PACT UDP | **106 ms** | **1.96 s** | **6.2 ms** | **180 ms** | **722 ms** | **205 ms** | **487/s** | **21528/s** |
| P2PACT TCP | 104 ms (1.0×) | 1.92 s (1.0×) | 8.7 ms (1.4×) | 91 ms (0.5×) | 356 ms (0.5×) | 206 ms (1.0×) | 610/s (0.8×) | 20206/s (1.1×) |
| WebSocket | 189 ms (1.8×) | 4.96 s (2.5×) | 6.4 ms (1.0×) | 84 ms (0.5×) | 597 ms (0.8×) | 616 ms (3.0×) | 605/s (0.8×) | 21421/s (1.0×) |
| HTTPS+JSON | 196 ms (1.8×) | 4.97 s (2.5×) | 5.3 ms (0.9×) | 85 ms (0.5×) | 597 ms (0.8×) | 412 ms (2.0×) | 591/s (0.8×) | 18790/s (1.1×) |
| HTTP/2 | 191 ms (1.8×) | 4.96 s (2.5×) | 6.8 ms (1.1×) | 84 ms (0.5×) | 597 ms (0.8×) | 413 ms (2.0×) | 606/s (0.8×) | 20720/s (1.0×) |
| gRPC | 184 ms (1.7×) | 4.96 s (2.5×) | 6.7 ms (1.1×) | 83 ms (0.5×) | 597 ms (0.8×) | 617 ms (3.0×) | 608/s (0.8×) | 21145/s (1.0×) |
| HTTP/3 | 202 ms (1.9×) | 5.17 s (2.6×) | 6.8 ms (1.1×) | 601 ms (3.3×) | 3.07 s (4.3×) | 417 ms (2.0×) | 558/s (0.9×) | 20566/s (1.0×) |

\* TCP-based protocols (P2PACT TCP, WebSocket, HTTPS, HTTP/2, gRPC) run without loss in the emulation, and the emulating proxy terminates TCP, so these columns are optimistic for them. On lossy links the fair comparison is with HTTP/3, which goes through the same emulated UDP link as P2PACT.

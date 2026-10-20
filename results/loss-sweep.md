## Packet loss sweep

P2PACT over UDP as loss grows (TCP loss cannot be emulated without root).

Data: generated text (Zipf-distributed words). Messages per size: 20 (2 for 10 000+ characters). Compression: true. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-20.

**Bytes on the wire per message, both directions** (0% loss; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP |
|---|---:|
| very short (~20 chars) | 147 B |
| medium (~100 chars) | 213 B |
| long (~600 chars) | 476 B |
| very long (~5 000 chars) | 2.4 KB |
| huge (~200 000 chars) | 94 KB |
| new connection + first message | 592 B |

**Time from sending a message to its ack** (median):

*0% loss — 40 ms one way*

| Message | P2PACT UDP |
|---|---:|
| very short (~20 chars) | 81 ms |
| medium (~100 chars) | 83 ms |
| long (~600 chars) | 82 ms |
| very long (~5 000 chars) | 82 ms |
| huge (~200 000 chars) | 240 ms |
| new connection + first message | 81 ms |

*1% loss — 40 ms one way, 1% loss*

| Message | P2PACT UDP |
|---|---:|
| very short (~20 chars) | 81 ms |
| medium (~100 chars) | 82 ms |
| long (~600 chars) | 82 ms |
| very long (~5 000 chars) | 83 ms |
| huge (~200 000 chars) | 280 ms |
| new connection + first message | 82 ms |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*5% loss — 40 ms one way, 5% loss*

| Message | P2PACT UDP |
|---|---:|
| very short (~20 chars) | 82 ms |
| medium (~100 chars) | 82 ms |
| long (~600 chars) | 81 ms |
| very long (~5 000 chars) | 82 ms |
| huge (~200 000 chars) | 399 ms |
| new connection + first message | 83 ms |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*10% loss — 40 ms one way, 10% loss*

| Message | P2PACT UDP |
|---|---:|
| very short (~20 chars) | 81 ms |
| medium (~100 chars) | 82 ms |
| long (~600 chars) | 81 ms |
| very long (~5 000 chars) | 83 ms |
| huge (~200 000 chars) | 700 ms |
| new connection + first message | 82 ms |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*20% loss — 40 ms one way, 20% loss*

| Message | P2PACT UDP |
|---|---:|
| very short (~20 chars) | 81 ms |
| medium (~100 chars) | 178 ms |
| long (~600 chars) | 81 ms |
| very long (~5 000 chars) | 176 ms |
| huge (~200 000 chars) | 2.90 s |
| new connection + first message | 602 ms |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


## Packet loss sweep

P2PACT over UDP as loss grows (TCP loss cannot be emulated without root).

Data: generated text (Zipf-distributed words). Messages per size: 20 (10 for 10 000+ characters). P2PACT compression: zstd with dictionaries — the common one (fba2b14e) and the client's own, trained on a warm-up conversation from the other half of the data. Seed: 1.  
Machine: darwin/arm64, 15 CPUs, go1.27.0, 2026-10-21.

**Bytes on the wire per message, both directions** (0% loss; the number in brackets is how many times more than P2PACT UDP):

| Message | P2PACT UDP |
|---|---:|
| very short (~20 chars) | 293 B |
| medium (~100 chars) | 333 B |
| long (~600 chars) | 530 B |
| very long (~5 000 chars) | 2.2 KB |
| huge (~200 000 chars) | 86 KB |
| new connection + first message | 1.0 KB |

**Delivery**: every message arrived intact (each one checked by a fingerprint of what the receiver got).

**Time from sending a message to its ack** (median; in brackets — how many times longer than P2PACT UDP):

*0% loss — 40 ms one way*

| Message | P2PACT UDP |
|---|---:|
| very short (~20 chars) | 81 ms |
| medium (~100 chars) | 81 ms |
| long (~600 chars) | 82 ms |
| very long (~5 000 chars) | 83 ms |
| huge (~200 000 chars) | 87 ms |
| new connection + first message | 85 ms |

*1% loss — 40 ms one way, 1% loss*

| Message | P2PACT UDP |
|---|---:|
| very short (~20 chars) | 81 ms |
| medium (~100 chars) | 82 ms |
| long (~600 chars) | 82 ms |
| very long (~5 000 chars) | 83 ms |
| huge (~200 000 chars) | 96 ms |
| new connection + first message | 88 ms |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*5% loss — 40 ms one way, 5% loss*

| Message | P2PACT UDP |
|---|---:|
| very short (~20 chars) | 82 ms |
| medium (~100 chars) | 82 ms |
| long (~600 chars) | 82 ms |
| very long (~5 000 chars) | 83 ms |
| huge (~200 000 chars) | 340 ms |
| new connection + first message | 87 ms |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*10% loss — 40 ms one way, 10% loss*

| Message | P2PACT UDP |
|---|---:|
| very short (~20 chars) | 82 ms |
| medium (~100 chars) | 82 ms |
| long (~600 chars) | 81 ms |
| very long (~5 000 chars) | 82 ms |
| huge (~200 000 chars) | 300 ms |
| new connection + first message | 83 ms |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.

*20% loss — 40 ms one way, 20% loss*

| Message | P2PACT UDP |
|---|---:|
| very short (~20 chars) | 81 ms |
| medium (~100 chars) | 81 ms |
| long (~600 chars) | 81 ms |
| very long (~5 000 chars) | 204 ms |
| huge (~200 000 chars) | 1.79 s |
| new connection + first message | 86 ms |

TCP-based protocols run without loss here (emulating it needs root), so their times are optimistic.


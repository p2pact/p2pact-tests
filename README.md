# p2pact-tests

Tests for [P2PACT](https://github.com/p2pact/p2pact) implementations:

- **interop** — every pair of implementations (Go, TypeScript, Python, and yours) pairs by invite
  and exchanges messages both ways, over UDP and TCP: short, compressible, Russian text and
  300 KB of random bytes;
- **p2pact-bench** — bytes and time against WebSocket, HTTPS, HTTP/2, gRPC and HTTP/3.

**Results at a glance: [SUMMARY.md](SUMMARY.md)** — one table for the real network, one for
the emulated links; every cell is the value and how many times worse the protocol is than
P2PACT over UDP. Full tables: [results-real/RESULTS.md](results-real/RESULTS.md) and
[results/RESULTS.md](results/RESULTS.md). `go run ./cmd/p2pact-bench report` rebuilds all of
them from the saved JSON.

## Interop

```sh
go test ./interop                                        # Go with itself
P2PACT_TS=../p2pact-ts/dist/src/interop.js \
P2PACT_PY=../p2pact-py/.venv/bin/python go test -v ./interop   # all pairs of three languages
```

Each implementation ships an interop peer: a node driven by JSON lines on stdin
(`listen`, `invite`, `join`, `known`, `call`, `dict`, `close` — see
[p2pact-go/cmd/p2pact-interop](https://github.com/p2pact/p2pact-go/tree/main/cmd/p2pact-interop)).
Write one for your implementation and add it to `interop_test.go`.

## Findings

### Real network

[`results-real/RESULTS.md`](results-real/RESULTS.md): a laptop on home Wi-Fi to a 1-CPU cloud
server, about 100–140 ms round trip, nothing emulated. Every message was checked and every one
arrived intact. P2PACT over UDP against the six other protocols:

- **First message on a new connection**: 112–147 ms against 203–274 ms for HTTP/3, 320–430 ms
  for HTTPS and HTTP/2, 417–557 ms for WebSocket and gRPC — **2–4× faster** (one round trip:
  the request rides in the first handshake packet).
- **Bytes**: a medium message (about 100 characters) with its ack takes 243–299 bytes against
  469–540 for WebSocket and HTTP/2, 584–630 for gRPC, 598–1181 for HTTP/3 and 761–806 for
  HTTPS — **1.6–4× fewer**: short headers, and each device's own compression dictionary.
- **Huge messages** (200 000 characters): 133 ms against 132–161 ms for the TCP-based protocols
  and 567 ms for HTTP/3 — **level, 4× faster than HTTP/3**.
- **Everyday messages**: the same for everyone, 93–139 ms; the network's round trip decides.
- **Load** (2000 messages, 50 at a time, like 50 chats): 282–466 messages a second against
  380–523 for gRPC, WebSocket, HTTP/2 and HTTP/3 and 252–282 for HTTPS — **level**; a real
  network varies from run to run (a repeat of the 282 scenario gave 410).

### Emulated links

[`results/RESULTS.md`](results/RESULTS.md): local network, Bluetooth (30 ms, 40 KB/s),
internet (40 ms, 1 % loss), mobile (100 ms, 500 KB/s, 3 % loss). All messages arrived intact.

- **Faster**: a new connection 2–3× everywhere; Bluetooth 1.1–1.9× on very long messages and
  2.6× on huge ones (fewer bytes on a slow link); huge messages against HTTP/3 on the same lossy
  link 4–7× (96–180 ms against 686 ms at 1 % loss, 0.7 s against 3.1 s at 3 %).
- **Level**: everyday messages; the local network under load (14 000–21 000 messages a second
  for every protocol); load against HTTP/3 on the lossy profiles (internet 511–545 a second
  against 551–573, mobile 151–231 against 183–207).
- **Slower**: load over Bluetooth (122–175 a second against 158–223 for WebSocket, HTTP/2 and
  gRPC; 412 against 170–182 on the coding agent's output, which compresses well); huge
  messages at 10–20 % loss take 0.3–1.8 s and vary a lot from run to run (4.1–4.5 s in the
  previous run): a lost last packet waits for the retransmission timer. TCP-based protocols
  look faster on lossy profiles only because they run without loss there (see the method).
- A huge message at 1 % loss is about 60 packets: in about half the runs one of them is lost
  and costs a round trip, so its median is 96 ms in one run and 180 ms in another.

### What the tests found in the protocol

Each of these was a real bug or weakness, fixed in all implementations and in the spec:

1. Queued messages went out on a 20 ms tick instead of at once.
2. A handshake counted its first message as sent when no handshake packet went out, so the
   message waited for its retransmission timer.
3. Data that overtook the handshake reply reset the session.
4. The reference implementation used keyed BLAKE2s where the spec says HMAC-BLAKE2s, so
   invite ids differed from other implementations (found by the interop test).
5. Acks for a lone part waited 40 ms; now every second part is acknowledged at once, as in TCP.
6. A fixed window and timer: no growth on fast links, spurious retransmissions on slow ones.
   Now a congestion window that grows while no queue builds (as BBR probes), timers from
   the measured round trip, a delivery-rate limit in bytes while a queue builds.
7. A random loss halved the window as if the network were congested.
8. Replies went to a road the device had used before it reconnected (a new TCP connection,
   a restart, UDP given up for TCP) and arrived only by retransmission: now the road of the
   current session heard from last.
9. On TCP the protocol kept a window of its own on top of TCP's.
10. All messages between two devices shared one queue, so a loss on one chat held up all the
    others. Now **streams**, as in QUIC: one per chat, independent, with ack ranges.
11. The known-devices file was written in place and could be read half-written.
12. Losses were noticed only by the timer or three later parts. Now time-based loss detection
    (RACK, RFC 8985): huge messages at 1 % loss 180 → 95 ms, load on the mobile profile
    147 → 233 a second.
13. That loss detection also ran on TCP, where nothing is lost: parts waiting in a slow TCP
    connection's queue were resent, a huge message over Bluetooth sent 352 KB instead of 77 KB
    and took 9.2 s instead of 1.9 s. Now it runs only on UDP.
14. Compression used no dictionary at all in the benchmark: the spec described one, but none
    existed. Now each device trains its own for each device it talks to: a medium message
    8–35 % smaller, a long one 14–39 % (the coding agent's output gains most).
15. With a dictionary, zstd's best level cost 0.6 ms a message under the node's lock: load on the
    local network fell from 19 000 to 1 500 a second. The next level is 30× faster for 3 % more
    bytes.
16. A dictionary added at start-up queued a list announcement, which opened the session early,
    without the message: the first message waited one more round trip (mobile 205 → 404 ms).
    Now lists are announced only over a live session; a new one starts with the hello.
17. A new dictionary moved the others in the list, and the list travels on stream 0: a message
    on another stream could arrive after the new list and decompress with the wrong dictionary.
    Now a listed dictionary keeps its place (found by review).

## Real network

```sh
# on a server (ports 48731–48736, UDP and TCP)
p2pact-bench serve -bind 0.0.0.0 -advertise <public address>
# on your machine
go run ./cmd/p2pact-bench run -remote <server address> scenarios/*.json -out results-real
```

The scenario's profiles are replaced by one, "real network". Keep the connection up for the
whole run: lost messages are counted as errors and a run with many of them stops.

## p2pact-bench

Compare P2PACT with the usual ways of carrying chat messages — **WebSocket** (JSON over TLS)
and **HTTPS+JSON** (as MCP/A2A-style APIs do) — on bytes and time. Anyone can rerun every
number, and change any part of the experiment.

```sh
git clone https://github.com/p2pact/p2pact-tests && cd p2pact-tests
go run ./cmd/p2pact-bench run scenarios/*.json     # all scenarios → results/
go run ./cmd/p2pact-bench run scenarios/synthetic.json   # offline, no downloads
```

Results go to `results/<scenario>.md` and `.json` (with the machine and commit), and
`results/RESULTS.md` collects them. The committed run is [`results/RESULTS.md`](results/RESULTS.md).

## What is measured

The same action for every protocol: send a chat message, get "accepted" back.

- **P2PACT UDP / TCP** — two real P2PACT nodes, paired with an invite, calling
  `chat.send` (tool 1, method 2) over the UDP road or the TCP road.
- **WebSocket** — TLS 1.3, a real upgrade request, RFC 6455 frames:
  `{"type":"send","chat":…,"text":…}` → `{"type":"ok"}`.
- **HTTPS+JSON** — Go's HTTP server, TLS 1.3, HTTP/1.1 keep-alive:
  `POST /chats/<id>/messages {"text":…}` with a bearer token → `{"ok":true}`.
- **HTTP/2** — the same request, multiplexed on one HTTP/2 connection.
- **gRPC** — a unary call `Chat/Send` with a protobuf message `{chat, text}` (grpc-go).
- **HTTP/3** — the same request over QUIC (quic-go).

For every size bin: the median time from sending to the ack over `repeat` messages, and bytes
per message in both directions. Plus a **cold start**: a new connection (for P2PACT — a new
session between paired devices) and the first message.

### Compression

P2PACT compresses with zstd and dictionaries, as a device does (spec §5): every node has the
common dictionary (`CommonDict`, trained on oasst2) and trains its own for each device it talks
to. Before measuring, the client talks to the server a while — about 1.1 MB of messages, each
cut to 4 KB — so it has its own dictionary in use, as a device that has chatted before. The open
datasets are split in two by a hash of each text: that conversation and the common dictionary use
one half, the measurements send the other, so no dictionary has seen a message it is measured
on. `-dict=false` measures plain zstd instead. The other protocols send their usual bytes
(JSON, protobuf) without compression.

### How bytes are counted

- UDP: a counter on the road counts every datagram; + 28 bytes of IP and UDP headers each.
  Exact.
- TCP (P2PACT TCP, WebSocket, HTTPS): a counter on the socket counts everything above TCP
  (TLS, HTTP, frames). TCP/IP headers are **estimated**: 52 bytes per segment of up to 1448
  bytes, one ACK per two segments; a cold start adds the 3-packet TCP handshake. Hence "≈".

### How the network is emulated

Everything runs on one machine over 127.0.0.1, so a run is reproducible.

- P2PACT UDP: the UDP road delays each packet, paces it to the link rate and drops packets at
  random (seeded).
- TCP-based protocols (WebSocket, HTTPS, HTTP/2, gRPC, P2PACT TCP): a proxy delays and paces
  every chunk per direction. Two caveats, both in their favour: **it cannot drop TCP packets**
  (that needs root and a kernel emulator), and **the proxy terminates TCP**, so their congestion
  control sees a loopback link and sends a whole message at once instead of ramping up over
  several round trips as on a real link. For big messages and lossy links their numbers are
  optimistic; a real network (two machines) is the honest comparison.
- HTTP/3 (QUIC) runs through a UDP proxy with the same delay, rate and loss as P2PACT — the
  fair comparison on lossy links.
- Bluetooth is emulated by speed and latency; a real BLE road is planned.

## Scenarios — change anything

A scenario is a JSON file. Copy one, change a field, run it:

```json
{
  "name": "Human messages, Russian",
  "data": "oasst2:human:ru",
  "seed": 1,
  "repeat": 20,
  "sizes": [{"name": "very short", "min": 1, "max": 40, "chars": 20}, …],
  "profiles": [{"name": "internet", "delay_ms": 40, "rate_kbps": 0, "loss_percent": 1}, …],
  "protocols": ["p2pact-udp", "p2pact-tcp", "websocket", "https"],
  "compression": true
}
```

- `data` — where texts come from:
  - `synthetic` — generated text with natural word frequencies (offline);
  - `oasst2:human`, `oasst2:assistant`, `oasst2:human:<lang>` — [OpenAssistant oasst2](https://huggingface.co/datasets/OpenAssistant/oasst2) (Apache-2.0);
  - `swe-gym:agent`, `swe-gym:tool` — a coding agent's actions and tool outputs from
    [SWE-Gym OpenHands trajectories](https://huggingface.co/datasets/SWE-Gym/OpenHands-SFT-Trajectories) (MIT);
  - `file:<path>` — your own messages as a JSON array of strings. They never leave your machine.
- `sizes` — bins in **characters** `[min, max)`; with real data each bin gets the messages of
  that length, `chars` is for generated text.
- `profiles` — latency one way, bandwidth per direction (KB/s, 0 — unlimited), loss in percent.
- `protocols` — any of `p2pact-udp`, `p2pact-tcp`, `websocket`, `https`, `http2`, `grpc`, `http3`;
- `repeat`, `repeat_large` (for 10 000+ characters), `seed`, `compression`.

Flags override files: `run -repeat 50 -seed 7 -data oasst2:assistant scenarios/synthetic.json`.

Datasets are downloaded once into the cache (`p2pact-bench data` shows where;
`P2PACT_BENCH_CACHE` changes it).

## Adding a protocol

Write `RunX(sc *Scenario, p Profile) (*Result, error)` next to `RunHTTPS` in `tcp.go`: count the
bytes, time the exchanges with the scenario's messages, and add its name to `Run`. Pull requests
with gRPC, HTTP/3, MQTT or anything else are welcome.

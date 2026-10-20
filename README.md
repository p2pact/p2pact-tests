# p2pact-tests

Tests for [P2PACT](https://github.com/p2pact/p2pact) implementations:

- **interop** — every pair of implementations (Go, TypeScript, Python, and yours) pairs by invite
  and exchanges messages both ways, over UDP and TCP: short, compressible, Russian text and
  300 KB of random bytes;
- **p2pact-bench** — bytes and time against WebSocket, HTTPS, HTTP/2, gRPC and HTTP/3.

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

## Findings so far

From [`results/RESULTS.md`](results/RESULTS.md), P2PACT over UDP against the other protocols:

- **Bytes**: 2–4.5× fewer than HTTPS, 2–3× fewer than WebSocket, HTTP/2 and gRPC, on every
  message size; a new connection costs about 0.7 KB against 4–16 KB.
- **First message on a new connection**: 2× faster on the internet and mobile (one round
  trip instead of two), 3× faster over Bluetooth.
- **Everyday messages** (up to a few thousand characters) on a live connection: the same time
  as everyone else — the network's round trip dominates.
- **Slow links** (Bluetooth): long messages 1.3–4× faster, because there are fewer bytes.
- **Huge messages** (200 000 characters) on fast links: slower than TCP-based protocols on a
  local network (compression costs CPU), and in the emulated internet. But TCP-based
  protocols are measured in better-than-real conditions here (next section); against HTTP/3,
  the only protocol measured on the same lossy UDP link, P2PACT is 1–6× faster.

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

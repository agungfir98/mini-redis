
# mini-redis

A minimal Redis clone written in Go — supports a subset of the Redis protocol (RESP).  

> [!WARNING]
> Don't use this on productions!, WTF is wrong with you

---

## Features

- [x] RESP protocol parsing
- [x] Basic commands:
- [x] `PING` – Ping the server and will reply with PONG
- [x] `SET key value` – Store a value by key
- [x] `GET key` – Retrieve a value by key
- [x] `DEL key` – Delete a key
- [ ] `KEYS pattern` – Pattern match key lookup
- [ ] `HSET key field value` – Set hash field
- [ ] `HGET key field` – Get hash field

---

## Getting Started

### 1. Clone and Build

```bash
git clone https://github.com/agungfir98/mini-redis.git
cd mini-redis
go build -o mini-redis
```

## Connecting with mini-redis

### using redis-cli or redli

```bash
redis-cli -p 6380
# or redli
redli -p 6380
```

### example usage


```bash
127.0.0.1:6379> SET foo bar
OK

127.0.0.1:6379> GET foo
"bar"

127.0.0.1:6379> DEL foo
(integer) 1

127.0.0.1:6379> GET foo
(nil)
```


## Implementation Notes
- written in pure Go
- parses raw Resp manually
- handle each connection with go routines
- uses map[string]string internally for storage



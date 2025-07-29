
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
- [ ] make SET accept expiration
- [x] `GET key` – Retrieve a value by key
- [x] `DEL key` – Delete a key
- [x] `TTL` - Return remaining time to live of a key that has a timeout.
- [x] `KEYS` - Return all keys matching pattern
- [x] `HSET key field value` – Set hash field
- [x] `HGET key field` – Get hash field
- [x] `HDEL key field` – Delete specified field from the hash
- [x] `HGETALL key field` – Delete specified field from the hash

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



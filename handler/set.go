package handler

import (
	"time"

	"github.com/agungfir98/mini-redis/proto"
	"github.com/agungfir98/mini-redis/store"
)

func Set(args []proto.RespMessage) proto.RespMessage {
	if len(args) < 2 {
		return proto.RespMessage{Typ: "error", Error: "wrong number of arguments for 'set' command"}
	}

	key := args[0].String
	value := args[1].String
	opts, err := parseOptions(args[2:])
	if err != nil {
		return proto.RespMessage{Typ: "error", Error: err.Error()}
	}

	ttl := time.Time{}
	_, ok := store.SETs[key]
	if ok && opts.NX {
		return proto.RespMessage{Typ: "null"}
	}
	if !ok && opts.XX {
		return proto.RespMessage{Typ: "null"}
	}
	if opts.EX || opts.PX {
		expireAt := time.Now().Add(opts.ttl)
		ttl = expireAt
	}

	store.SetRaw(key, value, ttl)
	return proto.RespMessage{Typ: "status", Status: "OK"}
}

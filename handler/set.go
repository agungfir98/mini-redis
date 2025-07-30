package handler

import (
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

	_, ok := store.SETs[key]
	if ok && opts.NX {
		return proto.RespMessage{Typ: "null"}
	}
	if !ok && opts.XX {
		return proto.RespMessage{Typ: "null"}
	}

	store.SetRaw(key, value, opts.ttl)
	return proto.RespMessage{Typ: "status", Status: "OK"}
}

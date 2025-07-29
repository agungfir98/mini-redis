package handler

import (
	"time"

	"github.com/agungfir98/mini-redis/proto"
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

	SetMu.Lock()
	defer SetMu.Unlock()
	_, ok := SETs[key]
	if ok && opts.NX {
		return proto.RespMessage{Typ: "null"}
	}
	if !ok && opts.XX {
		return proto.RespMessage{Typ: "null"}
	}
	data := Sets{value: value}
	if opts.EX || opts.PX {
		data.expireAt = time.Now().Add(opts.ttl)
	}
	SETs[key] = data

	return proto.RespMessage{Typ: "status", Status: "OK"}
}

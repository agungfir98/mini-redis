package handler

import (
	"github.com/agungfir98/mini-redis/proto"
)

func Get(args []proto.RespMessage) proto.RespMessage {
	if len(args) != 1 {
		return proto.RespMessage{Typ: "error", Error: "wrong number of arguments for 'get' command"}
	}

	key := args[0].String
	SetMu.RLock()
	val, ok := SETs[key]
	SetMu.RUnlock()

	if !ok {
		return proto.RespMessage{Typ: "null"}
	}

	return proto.RespMessage{Typ: "string", String: val.value}
}

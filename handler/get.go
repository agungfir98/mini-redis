package handler

import (
	"github.com/agungfir98/mini-redis/proto"
	"github.com/agungfir98/mini-redis/store"
)

func Get(args []proto.RespMessage) proto.RespMessage {
	if len(args) != 1 {
		return proto.RespMessage{Typ: "error", Error: "wrong number of arguments for 'get' command"}
	}

	key := args[0].String
	val, ok := store.GetRaw(key)
	if !ok {
		return proto.RespMessage{Typ: "null"}
	}
	return proto.RespMessage{Typ: "string", String: val.Value}
}

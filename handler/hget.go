package handler

import (
	"github.com/agungfir98/mini-redis/proto"
	"github.com/agungfir98/mini-redis/store"
)

func Hget(args []proto.RespMessage) proto.RespMessage {
	if len(args) != 2 {
		return WrongArgNumber("hget")
	}

	hash := args[0].String
	key := args[1].String

	val, ok := store.HygetRaw(hash, key)
	if !ok {
		return proto.RespMessage{Typ: "null"}
	}

	return proto.RespMessage{Typ: "string", String: val}
}

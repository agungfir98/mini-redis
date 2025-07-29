package handler

import (
	"github.com/agungfir98/mini-redis/proto"
	"github.com/agungfir98/mini-redis/store"
)

func Keys(args []proto.RespMessage) proto.RespMessage {
	if len(args) < 1 {
		return WrongArgNumber("keys")
	}

	keys := store.GetKeys(args[0].String)
	return proto.RespMessage{Typ: "array", Array: keys}
}

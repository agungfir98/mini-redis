package handler

import (
	"github.com/agungfir98/mini-redis/proto"
	"github.com/agungfir98/mini-redis/store"
)

func HgetAll(args []proto.RespMessage) proto.RespMessage {
	if len(args) > 1 {
		return WrongArgNumber("hgetall")
	}

	hash := args[0].String
	items := store.HgetAllRaw(hash)

	return proto.RespMessage{Typ: "array", Array: items}
}

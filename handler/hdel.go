package handler

import (
	"github.com/agungfir98/mini-redis/proto"
	"github.com/agungfir98/mini-redis/store"
)

func Hdel(args []proto.RespMessage) proto.RespMessage {
	if len(args) < 2 {
		return WrongArgNumber("hdel")
	}
	hash := args[0].String

	n := store.HdellRaw(hash, args[1:])

	return proto.RespMessage{Typ: "integer", Num: n}
}

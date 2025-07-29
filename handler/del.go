package handler

import (
	"github.com/agungfir98/mini-redis/proto"
	"github.com/agungfir98/mini-redis/store"
)

func Del(args []proto.RespMessage) proto.RespMessage {
	if len(args) == 0 {
		return WrongArgNumber("del")
	}

	n := store.DelRaw(args)

	return proto.RespMessage{Typ: "integer", Num: n}
}

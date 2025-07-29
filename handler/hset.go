package handler

import (
	"github.com/agungfir98/mini-redis/proto"
	"github.com/agungfir98/mini-redis/store"
)

// hash-key-value store that return the number of item afffected
func Hset(args []proto.RespMessage) proto.RespMessage {
	if len(args)%2 != 1 {
		return WrongArgNumber("hset")
	}

	hash := args[0].String
	items := args[1:]
	n := store.HsetRaw(hash, items)
	return proto.RespMessage{Typ: "integer", Num: n}
}

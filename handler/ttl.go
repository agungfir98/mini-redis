package handler

import (
	"time"

	"github.com/agungfir98/mini-redis/proto"
	"github.com/agungfir98/mini-redis/store"
)

func TTL(args []proto.RespMessage) proto.RespMessage {
	if len(args) != 1 {
		return WrongArgNumber("ttl")
	}
	key := args[0].String

	value, ok := store.GetRaw(key)
	if !ok {
		return proto.RespMessage{Typ: "integer", Num: -2}
	}
	if value.ExpireAt.IsZero() {
		return proto.RespMessage{Typ: "integer", Num: -1}
	}

	now := time.Now()
	TTL := max(value.ExpireAt.Sub(now), 0)
	return proto.RespMessage{Typ: "integer", Num: int(TTL.Seconds())}

}

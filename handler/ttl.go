package handler

import (
	"time"

	"github.com/agungfir98/mini-redis/proto"
)

func TTL(args []proto.RespMessage) proto.RespMessage {
	if len(args) != 1 {
		return WrongArgNumber("ttl")
	}
	key := args[0].String

	SetMu.RLock()
	defer SetMu.RUnlock()

	value, ok := SETs[key]
	if !ok {
		return proto.RespMessage{Typ: "integer", Num: -2}
	}

	if value.expireAt.IsZero() {
		return proto.RespMessage{Typ: "integer", Num: -1}
	}

	now := time.Now()
	TTL := max(value.expireAt.Sub(now), 0)
	return proto.RespMessage{Typ: "integer", Num: int(TTL.Seconds())}

}

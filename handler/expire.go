package handler

import (
	"strconv"

	"github.com/agungfir98/mini-redis/proto"
	"github.com/agungfir98/mini-redis/store"
)

func Expire(args []proto.RespMessage) proto.RespMessage {
	if len(args) != 3 {
		return WrongArgNumber("expire")
	}
	key := args[0].String
	seconds := args[1].String
	opt := args[2].String // options: NX XX GT LT

	second, err := strconv.Atoi(seconds)
	if err != nil {
		return proto.RespMessage{Typ: "error", Error: "ERR value is not an integer or out of range"}
	}

	n, err := store.ExpireRaw(key, second, opt)
	if err != nil {
		return proto.RespMessage{Typ: "error", Error: err.Error()}
	}

	return proto.RespMessage{Typ: "integer", Num: n}
}

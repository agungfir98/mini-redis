package handler

import (
	"github.com/agungfir98/mini-redis/proto"
)

func Set(args []proto.RespMessage) proto.RespMessage {
	if len(args) != 2 {
		return proto.RespMessage{Typ: "error", Error: "wrong number of arguments for 'set' command"}
	}
	key := args[0].String
	value := args[1].String

	SetMu.Lock()
	SETs[key] = value
	SetMu.Unlock()

	return proto.RespMessage{Typ: "status", Status: "OK"}
}

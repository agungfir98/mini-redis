package handler

import "github.com/agungfir98/mini-redis/proto"

func Ping(args []proto.RespMessage) proto.RespMessage {
	return proto.RespMessage{Typ: "status", Status: "PONG"}
}

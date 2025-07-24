package handler

import "github.com/agungfir98/mini-redis/proto"

func Hget(args []proto.RespMessage) proto.RespMessage {
	if len(args) != 2 {
		return WrongArgNumber("hget")
	}

	hash := args[0].String
	key := args[1].String

	HsetMu.RLock()
	if _, ok := HSETs[hash]; !ok {
		HsetMu.RUnlock()
		return proto.RespMessage{Typ: "null"}
	}
	val, ok := HSETs[hash][key]
	HsetMu.RUnlock()

	if !ok {
		return proto.RespMessage{Typ: "null"}
	}

	return proto.RespMessage{Typ: "string", String: val}
}

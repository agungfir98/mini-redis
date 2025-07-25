package handler

import "github.com/agungfir98/mini-redis/proto"

func HgetAll(args []proto.RespMessage) proto.RespMessage {
	if len(args) > 1 {
		return WrongArgNumber("hgetall")
	}

	var items []proto.RespMessage
	hash := args[0].String

	data, ok := HSETs[hash]
	if !ok {
		return proto.RespMessage{Typ: "array", Array: items}
	}
	HsetMu.RLock()
	for key, val := range data {
		items = append(items, proto.RespMessage{Typ: "string", String: key})
		items = append(items, proto.RespMessage{Typ: "string", String: val})
	}
	HsetMu.RUnlock()

	return proto.RespMessage{Typ: "array", Array: items}
}

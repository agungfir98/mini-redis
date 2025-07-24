package handler

import "github.com/agungfir98/mini-redis/proto"

func Hdel(args []proto.RespMessage) proto.RespMessage {
	if len(args) < 2 {
		return WrongArgNumber("hdel")
	}

	var n int

	hash := args[0].String
	_, ok := HSETs[hash]
	if !ok {
		return proto.RespMessage{Typ: "integer", Num: n}
	}

	HsetMu.Lock()
	for _, key := range args[1:] {
		_, ok := HSETs[hash][key.String]
		if !ok {
			continue
		}
		delete(HSETs[hash], key.String)
		n += 1
	}

	HsetMu.Unlock()

	return proto.RespMessage{Typ: "integer", Num: n}
}

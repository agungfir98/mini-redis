package handler

import "github.com/agungfir98/mini-redis/proto"

func Del(args []proto.RespMessage) proto.RespMessage {
	if len(args) == 0 {
		return WrongArgNumber("del")
	}
	var n int

	SetMu.Lock()
	for _, arg := range args {
		key := arg.String
		_, ok := SETs[key]
		if !ok {
			continue
		}
		delete(SETs, key)
		n += 1
	}
	SetMu.Unlock()

	return proto.RespMessage{Typ: "integer", Num: n}
}

package handler

import (
	"github.com/agungfir98/mini-redis/proto"
)

// hash-key-value store that return the number of item afffected
func Hset(args []proto.RespMessage) proto.RespMessage {
	var n int

	if len(args)%2 != 1 {
		return WrongArgNumber("hset")
	}

	hash := args[0].String
	items := args[1:]

	HsetMu.Lock()
	for i := 0; i < len(items); i += 2 {
		if _, ok := HSETs[hash]; !ok {
			HSETs[hash] = map[string]string{}
		}

		HSETs[hash][items[i].String] = items[i+1].String
		n += 1
	}
	HsetMu.Unlock()

	return proto.RespMessage{Typ: "integer", Num: n}
}

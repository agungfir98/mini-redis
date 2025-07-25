package handler

import (
	"fmt"
	"regexp"

	"github.com/agungfir98/mini-redis/proto"
)

func Keys(args []proto.RespMessage) proto.RespMessage {
	if len(args) < 1 {
		return WrongArgNumber("keys")
	}

	pattern := regexp.MustCompile(args[0].String)

	keys := []proto.RespMessage{}

	SetMu.RLock()
	for k := range SETs {
		fmt.Println(k)
		if pattern.MatchString(k) {
			keys = append(keys, proto.RespMessage{Typ: "string", String: k})
		}
	}
	SetMu.RUnlock()

	return proto.RespMessage{Typ: "array", Array: keys}
}

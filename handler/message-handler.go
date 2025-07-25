package handler

import (
	"fmt"
	"sync"

	"github.com/agungfir98/mini-redis/proto"
)

var Message = map[string]func([]proto.RespMessage) proto.RespMessage{
	"PING":    Ping,
	"SET":     Set,
	"GET":     Get,
	"DEL":     Del,
	"HSET":    Hset,
	"HGET":    Hget,
	"HDEL":    Hdel,
	"HGETALL": HgetAll,
	"KEYS":    Keys,
}

var SETs = map[string]string{}
var SetMu = sync.RWMutex{}

var HSETs = map[string]map[string]string{}
var HsetMu = sync.RWMutex{}

func WrongArgNumber(cmd string) proto.RespMessage {
	return proto.RespMessage{Typ: "error", Error: fmt.Sprintf("wrong number of arguments for '%v' command", cmd)}
}

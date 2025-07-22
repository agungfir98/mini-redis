package handler

import (
	"sync"

	"github.com/agungfir98/mini-redis/proto"
)

var Message = map[string]func([]proto.RespMessage) proto.RespMessage{
	"SET":  Set,
	"GET":  Get,
	"PING": Ping,
}

var SETs = map[string]string{}
var SetMu = sync.RWMutex{}

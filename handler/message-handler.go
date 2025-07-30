package handler

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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
	"TTL":     TTL,
}

func WrongArgNumber(cmd string) proto.RespMessage {
	return proto.RespMessage{Typ: "error", Error: fmt.Sprintf("wrong number of arguments for '%v' command", cmd)}
}

type SetOptions struct {
	ttl  time.Time
	EX   bool // second
	PX   bool // millisecond
	EXAT bool // timestamp second
	PXAT bool // timestamp millisecond
	NX   bool // only set if not exists
	XX   bool // only set if key exists
}

func parseOptions(args []proto.RespMessage) (opts SetOptions, err error) {
	var i int = 0

	for i < len(args) {
		opt := strings.ToUpper(args[i].String)
		switch opt {
		case "EX", "PX", "EXAT", "PXAT":
			if i+1 >= len(args) {
				return opts, fmt.Errorf("ERR missing expire time")
			}
			numStr := args[i+1].String
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return opts, fmt.Errorf("invalid expire time")
			}

			var ttl time.Time
			now := time.Now()
			switch opt {
			case "EX":
				opts.EX = true
				ttl = now.Add(time.Duration(num) * time.Second)
			case "PX":
				opts.PX = true
				ttl = now.Add(time.Duration(num) * time.Millisecond)
			case "EXAT":
				opts.EXAT = true
				ttl = time.Unix(int64(num), 0)
			case "PXAT":
				opts.PXAT = true
				ttl = time.UnixMilli(int64(num))
			}
			opts.ttl = ttl
			i += 2
			return opts, nil
		case "NX":
			opts.NX = true
			i++
			return opts, nil
		case "XX":
			i++
			opts.XX = true
			return opts, nil
		default:
			opts = SetOptions{}
			return opts, fmt.Errorf("ERR invalid options")
		}
	}
	return opts, nil
}

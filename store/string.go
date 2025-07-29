package store

import (
	"container/heap"
	"regexp"
	"time"

	"github.com/agungfir98/mini-redis/proto"
)

func SetRaw(key, value string, ttl time.Time) {
	SetMu.Lock()
	data := Sets{Value: value, ExpireAt: ttl}
	SETs[key] = data
	heap.Push(expQ, ExpiryItem{key: key, expireAt: ttl})
	SetMu.Unlock()
}

func GetRaw(key string) (Sets, bool) {
	SetMu.RLock()
	val, ok := SETs[key]
	SetMu.RUnlock()
	return val, ok
}

func DelRaw(keys []proto.RespMessage) int {
	var n int
	SetMu.Lock()
	for _, key := range keys {
		key := key.String
		_, ok := SETs[key]
		if !ok {
			continue
		}
		delete(SETs, key)
		n += 1
	}
	SetMu.Unlock()

	return n
}

func GetKeys(key string) []proto.RespMessage {
	keys := []proto.RespMessage{}

	pattern := regexp.MustCompile(key)
	SetMu.RLock()
	for k := range SETs {
		if pattern.MatchString(k) {
			keys = append(keys, proto.RespMessage{Typ: "string", String: k})
		}
	}
	SetMu.RUnlock()

	return keys
}

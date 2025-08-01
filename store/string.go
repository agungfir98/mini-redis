package store

import (
	"container/heap"
	"fmt"
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

func ExpireRaw(key string, seconds int, opt string) (int, error) {
	SetMu.Lock()
	defer SetMu.Unlock()

	val, ok := SETs[key]
	if !ok {
		return 0, nil
	}
	expireAt := time.Now().Add(time.Duration(seconds) * time.Second)

	switch opt {
	case "NX":
		if !val.ExpireAt.IsZero() {
			return 0, nil
		}
	case "XX":
		if val.ExpireAt.IsZero() {
			return 0, nil
		}
	case "GT":
		if val.ExpireAt.After(expireAt) {
			return 0, nil
		}
	case "LT":
		if val.ExpireAt.Before(expireAt) {
			return 0, nil
		}
	default:
		return 0, fmt.Errorf("ERR unsupported option %v\n", seconds)
	}

	val.ExpireAt = expireAt
	SETs[key] = val

	return 1, nil
}

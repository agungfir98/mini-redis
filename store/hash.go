package store

import "github.com/agungfir98/mini-redis/proto"

func HsetRaw(hash string, items []proto.RespMessage) int {
	var n int
	HsetMu.Lock()
	for i := 0; i < len(items); i += 2 {
		if _, ok := HSETs[hash]; !ok {
			HSETs[hash] = map[string]string{}
		}

		HSETs[hash][items[i].String] = items[i+1].String
		n += 1
	}
	HsetMu.Unlock()

	return n
}

func HygetRaw(hash, key string) (string, bool) {
	HsetMu.RLock()
	defer HsetMu.RUnlock()

	if _, ok := HSETs[hash]; !ok {
		return "", ok
	}
	val, ok := HSETs[hash][key]
	return val, ok
}

func HgetAllRaw(hash string) []proto.RespMessage {
	var items []proto.RespMessage
	HsetMu.RLock()
	defer HsetMu.RUnlock()
	data, ok := HSETs[hash]
	if !ok {
		return items
	}
	for key, val := range data {
		items = append(items, proto.RespMessage{Typ: "string", String: key}, proto.RespMessage{Typ: "string", String: val})
	}
	return items
}

func HdellRaw(hash string, entries []proto.RespMessage) int {
	var n int

	HsetMu.Lock()
	defer HsetMu.Unlock()
	_, ok := HSETs[hash]
	if !ok {
		return n
	}

	for _, key := range entries {
		_, ok := HSETs[hash][key.String]
		if !ok {
			continue
		}
		delete(HSETs[hash], key.String)
		n += 1
	}
	return n
}

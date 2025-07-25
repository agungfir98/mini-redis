package handler

import (
	"testing"

	"github.com/agungfir98/mini-redis/proto"
)

func TestMessageHandler(t *testing.T) {
	testStringData(t) // get set del ....
	testKeys(t)
	testHashData(t) // hget hset hgetl ...
}

type StringTestCase struct {
	name     string
	setArgs  []proto.RespMessage
	getArgs  []proto.RespMessage
	delArgs  []proto.RespMessage
	keysArgs []proto.RespMessage
	setWant  proto.RespMessage
	getWant  proto.RespMessage
	delWant  proto.RespMessage
	keysWant proto.RespMessage
}

func testStringData(t *testing.T) {
	tc := []StringTestCase{
		{
			name: "case 1",
			setArgs: []proto.RespMessage{
				{Typ: "string", String: "SET"},
				{Typ: "string", String: "foo"},
				{Typ: "string", String: "bar"},
			},
			getArgs: []proto.RespMessage{
				{Typ: "string", String: "GET"},
				{Typ: "string", String: "foo"},
			},
			delArgs: []proto.RespMessage{
				{Typ: "string", String: "DEL"},
				{Typ: "string", String: "foo"},
			},
			keysArgs: []proto.RespMessage{
				{Typ: "string", String: "KEYS"},
				{Typ: "string", String: "f?"},
			},
			setWant:  proto.RespMessage{Typ: "status", Status: "OK"},
			getWant:  proto.RespMessage{Typ: "string", String: "bar"},
			delWant:  proto.RespMessage{Typ: "integer", Num: 1},
			keysWant: proto.RespMessage{Typ: "string", String: "foo"},
		},
	}

	for _, c := range tc {
		t.Run("string command", func(t *testing.T) {
			t.Run(c.name, func(t *testing.T) {
				var skip bool
				runStep(t, "set", &skip, func(t *testing.T) { testSet(t, c) })
				runStep(t, "get", &skip, func(t *testing.T) { testGet(t, c) })
				runStep(t, "del", &skip, func(t *testing.T) { testDel(t, c) })
			})
		})
	}
}

type HashTestCase struct {
	name        string
	setArgs     []proto.RespMessage
	getArgs     []proto.RespMessage
	delArgs     []proto.RespMessage
	hgetallArgs []proto.RespMessage
	setWant     proto.RespMessage
	getWant     proto.RespMessage
	delWant     proto.RespMessage
	hgetallWant proto.RespMessage
}

func testHashData(t *testing.T) {
	tc := []HashTestCase{
		{
			name: "Hash command",
			setArgs: []proto.RespMessage{
				{Typ: "string", String: "HSET"},
				{Typ: "string", String: "user"},
				{Typ: "string", String: "u1"},
				{Typ: "string", String: "foo"},
				{Typ: "string", String: "u2"},
				{Typ: "string", String: "bar"},
				{Typ: "string", String: "u3"},
				{Typ: "string", String: "fizz"},
				{Typ: "string", String: "u4"},
				{Typ: "string", String: "buzz"},
			},
			getArgs: []proto.RespMessage{
				{Typ: "string", String: "HGET"},
				{Typ: "string", String: "user"},
				{Typ: "string", String: "u1"},
			},
			delArgs: []proto.RespMessage{
				{Typ: "string", String: "HDEL"},
				{Typ: "string", String: "user"},
				{Typ: "string", String: "u1"},
				{Typ: "string", String: "u2"},
				{Typ: "string", String: "u8"},
			},
			hgetallArgs: []proto.RespMessage{
				{Typ: "string", String: "HGETALL"},
				{Typ: "string", String: "user"},
			},
			setWant: proto.RespMessage{Typ: "integer", Num: 4},
			getWant: proto.RespMessage{Typ: "string", String: "foo"},
			delWant: proto.RespMessage{Typ: "integer", Num: 2},
			hgetallWant: proto.RespMessage{
				Typ: "array",
				Array: []proto.RespMessage{
					{Typ: "string", String: "u1"},
					{Typ: "string", String: "foo"},
					{Typ: "string", String: "u2"},
					{Typ: "string", String: "bar"},
					{Typ: "string", String: "u3"},
					{Typ: "string", String: "fizz"},
					{Typ: "string", String: "u4"},
					{Typ: "string", String: "buzz"},
				},
			},
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			var skip bool
			runStep(t, "hset", &skip, func(t *testing.T) { testHset(t, c) })
			runStep(t, "hget", &skip, func(t *testing.T) { testHget(t, c) })
			runStep(t, "hgetall", &skip, func(t *testing.T) { testHgetAll(t, c) })
			runStep(t, "hdel", &skip, func(t *testing.T) { testHDel(t, c) })
		})
	}
}

func runStep(t *testing.T, name string, skip *bool, fn func(t *testing.T)) {
	t.Run(name, func(t *testing.T) {
		if *skip {
			t.Skipf("skipping %s because earlier step failed", name)
			return
		}

		fn(t)
		if t.Failed() {
			*skip = true
		}
	})
}

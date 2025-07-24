package handler

import (
	"testing"

	"github.com/agungfir98/mini-redis/proto"
)

func TestMessageHandler(t *testing.T) {
	testStringData(t) // get set del ....
}

type SetTestCase struct {
	name    string
	setArgs []proto.RespMessage
	getArgs []proto.RespMessage
	delArgs []proto.RespMessage
	setWant proto.RespMessage
	getWant proto.RespMessage
	delWant proto.RespMessage
}

func testStringData(t *testing.T) {
	tc := []SetTestCase{
		{
			name: "String command",
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
			setWant: proto.RespMessage{Typ: "status", Status: "OK"},
			getWant: proto.RespMessage{Typ: "string", String: "bar"},
			delWant: proto.RespMessage{Typ: "integer", Num: 1},
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			var skip bool
			runStep(t, "set", &skip, func(t *testing.T) { testSet(t, c) })
			runStep(t, "get", &skip, func(t *testing.T) { testGet(t, c) })
			runStep(t, "del", &skip, func(t *testing.T) { testDel(t, c) })
		})
	}
}

func runStep(t *testing.T, name string, skip *bool, fn func(t *testing.T)) {
	t.Run(name, func(t *testing.T) {
		if *skip {
			t.Skipf("skipping %s because earlier step failed", name)
		}

		fn(t)
		if t.Failed() {
			*skip = true
		}
	})
}

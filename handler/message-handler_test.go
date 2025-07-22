package handler

import (
	"testing"

	"github.com/agungfir98/mini-redis/proto"
)

func TestMessageHandler(t *testing.T) {
	testSetGetDel(t)
}

type SetTestCase struct {
	name string
	args []proto.RespMessage
	want proto.RespMessage
}

func testSetGetDel(t *testing.T) {
	tc := []SetTestCase{
		{name: "set foo bar", args: []proto.RespMessage{{Typ: "string", String: "SET"}, {Typ: "string", String: "foo"}, {Typ: "string", String: "bar"}}, want: proto.RespMessage{Typ: "status", Status: "OK"}},
	}

	// TODO:
	// This testing flow is awful, gotta refactor sometimes later
	t.Run("test set get del", func(t *testing.T) {
		testSet(t, tc)
		testGet(t, tc)
	})
}

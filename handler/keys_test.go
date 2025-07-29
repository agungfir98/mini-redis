package handler

import (
	"reflect"
	"strings"
	"testing"

	"github.com/agungfir98/mini-redis/proto"
)

func testKeys(t *testing.T) {
	listSet := []proto.RespMessage{
		{
			Typ: "array", Array: []proto.RespMessage{
				{Typ: "string", String: "SET"},
				{Typ: "string", String: "hallo"},
				{Typ: "string", String: "sekai"},
			},
		},
		{
			Typ: "array", Array: []proto.RespMessage{
				{Typ: "string", String: "SET"},
				{Typ: "string", String: "hello"},
				{Typ: "string", String: "sekai"},
			},
		},
		{
			Typ: "array", Array: []proto.RespMessage{
				{Typ: "string", String: "SET"},
				{Typ: "string", String: "hillo"},
				{Typ: "string", String: "sekai"},
			},
		},
	}

	for _, msg := range listSet {
		setCmd := strings.ToUpper(msg.Array[0].String)
		args := msg.Array[1:]
		Set := Message[setCmd]
		Set(args)
	}

	tests := []struct {
		name string
		args []proto.RespMessage
		want proto.RespMessage
	}{
		{
			name: "keys command",
			args: []proto.RespMessage{
				{Typ: "string", String: "KEYS"},
				{Typ: "string", String: "h[ae]llo"},
			},
			want: proto.RespMessage{Typ: "array", Array: []proto.RespMessage{
				{Typ: "string", String: "hallo"},
				{Typ: "string", String: "hello"},
			}},
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			cmd := c.args[0].String
			handler := Message[cmd]

			got := handler(c.args[1:])

			if !reflect.DeepEqual(got, c.want) {
				t.Fatalf("expected: %v, got: %v\n", c.want, got)
			}
		})
	}
}

package proto

import (
	"strconv"
	"testing"
)

// test function to validate the marshal implementation
func TestMarshaller(t *testing.T) {
	type testCase struct {
		name    string
		payload RespMessage
		want    string
	}
	tc := []testCase{
		{name: "marshal status", payload: RespMessage{Typ: "status", Status: "OK"}, want: "+OK\r\n"},
		{name: "marshal null", payload: RespMessage{Typ: "null"}, want: "$-1\r\n"},
		{name: "marshal nil", payload: RespMessage{Typ: "nil"}, want: "_\r\n"},
		{name: "marshal error", payload: RespMessage{Typ: "error", Error: "ERR"}, want: "-ERR\r\n"},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {

			res := c.payload.Marshal()

			if string(res) != c.want {
				t.Fatalf("expected: %v, got: %v\n", strconv.Quote(c.want), strconv.Quote(string(res)))
			}
		})
	}
}

package handler

import (
	"reflect"
	"strings"
	"testing"
)

func testDel(t *testing.T, c StringTestCase) {
	cmd := strings.ToUpper(c.delArgs[0].String)
	args := c.delArgs[1:]
	handler, ok := Message[cmd]
	if !ok {
		t.Fatalf("no such command %v, test: %v\n", cmd, c.name)
	}

	got := handler(args)
	if !reflect.DeepEqual(c.delWant, got) {
		t.Fatalf("expected: %v, got: %v\n", c.delWant, got)
	}
}

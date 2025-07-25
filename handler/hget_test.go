package handler

import (
	"reflect"
	"strings"
	"testing"
)

func testHget(t *testing.T, c HashTestCase) {
	cmd := strings.ToUpper(c.getArgs[0].String)
	args := c.getArgs[1:]

	handler, ok := Message[cmd]
	if !ok {
		t.Fatalf("command not found: %v\n", cmd)
	}

	got := handler(args)
	if !reflect.DeepEqual(got, c.getWant) {
		t.Fatalf("expected: %v, got: %v\n", c.getWant, got)
	}
}

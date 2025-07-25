package handler

import (
	"reflect"
	"strings"
	"testing"
)

func testHgetAll(t *testing.T, c HashTestCase) {
	cmd := strings.ToUpper(c.hgetallArgs[0].String)
	args := c.hgetallArgs[1:]

	handler, ok := Message[cmd]
	if !ok {
		t.Fatalf("command not found: %v\n", cmd)
	}

	got := handler(args)
	if !reflect.DeepEqual(got, c.hgetallWant) {
		t.Fatalf("expected: %v, got: %v\n", c.hgetallWant, got)
	}
}

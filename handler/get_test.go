package handler

import (
	"reflect"
	"strings"
	"testing"
)

func testGet(t *testing.T, c StringTestCase) {
	cmd := strings.ToUpper(c.getArgs[0].String)
	handler, ok := Message[cmd]
	if !ok {
		t.Fatalf("no such command %v, test: %v\n", cmd, c.name)
	}

	got := handler(c.getArgs[1:])
	if !reflect.DeepEqual(got, c.getWant) {
		t.Fatalf("expected: %v, got: %v\n", c.getWant, got)
	}
}

package handler

import (
	"reflect"
	"strings"
	"testing"
)

func testSet(t *testing.T, c StringTestCase) {
	cmd := strings.ToUpper(c.setArgs[0].String)
	handler, ok := Message[cmd]
	if !ok {
		t.Fatalf("no such command %v, test: %v\n", cmd, c.name)
	}

	got := handler(c.setArgs[1:])
	if !reflect.DeepEqual(got, c.setWant) {
		t.Fatalf("expected: %v, got: %v\n", c.setWant, got)
	}
}

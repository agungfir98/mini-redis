package handler

import (
	"reflect"
	"strings"
	"testing"
)

func testHset(t *testing.T, c SetTestCase) {
	cmd := strings.ToUpper(c.setArgs[0].String)
	args := c.setArgs[1:]

	handler, ok := Message[cmd]
	if !ok {
		t.Fatalf("command not found: %v\n", cmd)
	}
	res := handler(args)

	if !reflect.DeepEqual(res, c.setWant) {
		t.Fatalf("expected:%v, got: %v\n", c.setWant, res)
	}

}

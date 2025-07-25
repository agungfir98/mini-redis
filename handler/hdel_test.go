package handler

import (
	"reflect"
	"strings"
	"testing"
)

func testHDel(t *testing.T, c HashTestCase) {
	cmd := strings.ToUpper(c.delArgs[0].String)

	handler, ok := Message[cmd]
	if !ok {
		t.Fatalf("no such command '%v'\n", cmd)
	}

	got := handler(c.delArgs[1:])

	if !reflect.DeepEqual(got, c.delWant) {
		t.Errorf("expected:%v, got %v", c.delWant, got)
	}
}

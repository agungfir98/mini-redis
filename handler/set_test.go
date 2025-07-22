package handler

import (
	"reflect"
	"strings"
	"testing"
)

func testSet(t *testing.T, tc []SetTestCase) {
	for _, c := range tc {
		t.Run("set", func(t *testing.T) {
			cmd := c.args[0].String

			handler, ok := Message[strings.ToUpper(cmd)]
			if !ok {
				t.Fatalf("no such command %v, test: %v\n", cmd, c.name)
			}

			result := handler(c.args[1:])
			if !reflect.DeepEqual(result, c.want) {
				t.Fatalf("expected: %v, got: %v\n", c.want, result)
			}

		})
	}
}

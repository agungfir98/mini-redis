package handler

import "testing"

func testGet(t *testing.T, tc []SetTestCase) {
	for _, c := range tc {
		t.Run("get", func(t *testing.T) {

			key := c.args[1].String
			value := c.args[2].String

			SetMu.RLock()
			val, ok := SETs[key]
			SetMu.RUnlock()

			if !ok {
				t.Fatalf("key not found: %v\n", key)
			}

			if val != value {
				t.Fatalf("expected value: %v, got: %v\n", value, val)
			}
		})
	}
}

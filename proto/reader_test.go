package proto

import (
	"bytes"
	"reflect"
	"testing"
)

func TestAll(t *testing.T) {
	t.Run("testing readLine", testReadLine)
	t.Run("testing readArray", testReadArray)
	t.Run("testing readString", testString)

	t.Run("testing Read", func(t *testing.T) {
		if t.Failed() {
			t.Skip("prerequsite test failed, skipping testRead")
		}
		testRead(t)
	})

}

func testReadLine(t *testing.T) {
	type testCase struct {
		name    string
		payload string
		want    string
		wantErr bool
	}
	tc := []testCase{
		{name: "case 1", payload: "*2\r\n", want: "*2"},
		{name: "case 2", payload: "%2\r\n", want: "%2"},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			resp := NewResp(bytes.NewReader([]byte(c.payload)))
			line, _, err := resp.readLine()
			if (err != nil) != c.wantErr {
				t.Fatalf("expected to error")
			}

			if string(line) != c.want {
				t.Fatalf("expected %v, got: %q\n", c.want, string(line))
			}
		})
	}

}

func testReadArray(t *testing.T) {
	type testCase struct {
		name      string
		payload   string
		typWant   string
		arrayWant Value
		wantErr   bool
	}

	tc := []testCase{
		{name: "case 1", payload: "*2\r\n", typWant: "array", arrayWant: Value{Typ: "array", Array: make([]Value, 2)}},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			resp := NewResp(bytes.NewReader([]byte(c.payload)))
			resp.reader.ReadByte() // this just for mimic Read() to read the first byte so that next its reading the array length *<length>

			val, err := resp.readArray()
			if err != nil {
				t.Fatalf("supposed to not error, got: %v\n", err)
			}

			if val.Typ != c.typWant {
				t.Fatalf("expected %v, got: %v\n", c.typWant, val.Typ)
			}

			if !reflect.DeepEqual(val, c.arrayWant) {
				t.Fatalf("expected %v, got: %v\n", c.arrayWant, val)
			}

		})
	}
}

func testString(t *testing.T) {
	type testCase struct {
		name       string
		payload    string
		typWant    string
		stringWant string
		wantErr    bool
	}

	tc := []testCase{
		{name: "case 1", payload: "$3\r\nfoo\r\n", typWant: "string", stringWant: "foo"},
		//edge case when string length doesn't match actual string length
		{name: "case 2", payload: "$3\r\nbarrrrr\r\n", typWant: "string", stringWant: "bar"},
		{name: "case 3", payload: "$3\r\nba\r\n", typWant: "string", stringWant: "ba"},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			resp := NewResp(bytes.NewReader([]byte(c.payload)))
			resp.reader.ReadByte() // this just for mimic Read() to read the first byte so that next its reading the array length *<length>

			val, err := resp.readString()
			if err != nil {
				t.Fatalf("expected to not error, got: %v\n", err)
			}

			if val.Typ != c.typWant {
				t.Fatalf("expected: %v, got: %v\n", c.typWant, val.Typ)
			}

			if val.String != c.stringWant {
				t.Fatalf("expected: %v, got: %v\n", c.stringWant, val.String)
			}

		})
	}
}

func testRead(t *testing.T) {
	type testCase struct {
		name    string
		payload string
		want    Value
		wantErr bool
	}
	tc := []testCase{
		{
			name:    "set foo bar",
			payload: "*3\r\n$3\r\nset\r\n$3\r\nfoo\r\n$3\r\nbar\r\n",
			want: Value{Typ: "array", Array: []Value{
				{Typ: "string", String: "set"},
				{Typ: "string", String: "foo"},
				{Typ: "string", String: "bar"},
			}},
		},
	}

	for _, c := range tc {
		t.Run(c.name, func(t *testing.T) {
			resp := NewResp(bytes.NewReader([]byte(c.payload)))
			val, err := resp.Read()
			if err != nil {
				t.Fatalf("Expected to not error, got: %v, supposedly: %v", err, c.wantErr)
			}

			if !reflect.DeepEqual(val, c.want) {
				t.Fatalf("Expected: %v, got: %v\n", c.want, val)
			}

		})
	}
}

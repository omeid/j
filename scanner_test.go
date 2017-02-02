package json

import (
	"reflect"
	"testing"
)

type simpleExamples struct {
	in  []byte
	err error
}

func makeSyntaxError(offset, line, column int, msg string) *SyntaxError {
	return &SyntaxError{
		Position: Position{
			Offset: offset,
			Line:   line,
			Column: column,
		},
		Message: msg,
	}
}

var examples = []simpleExamples{
	{
		in:  []byte(`}`),
		err: makeSyntaxError(1, 1, 1, "invalid character '}' expected beginning of json ('{' or '[')"),
	},
	{
		in:  []byte("       \n\r\t {}"),
		err: nil,
	},
	{
		in:  []byte(`{`),
		err: makeSyntaxError(2, 1, 2, "Unexpected end of json"),
	},
	{
		in:  []byte(`]`),
		err: makeSyntaxError(1, 1, 1, "invalid character ']' expected beginning of json ('{' or '[')"),
	},
	{
		in:  []byte(`[]`),
		err: nil,
	},
	{
		in:  []byte(`{,`),
		err: makeSyntaxError(2, 1, 2, "invalid character ',' expected a pair or end of object"),
	},
	{
		in:  []byte(`{"key",`),
		err: makeSyntaxError(7, 1, 7, "invalid character ',' expected a colon"),
	},
	{
		in:  []byte(`{"key":}`),
		err: makeSyntaxError(8, 1, 8, "invalid character '}' expected a value type"),
	},
	{
		in:  []byte(`{"key":{}}`),
		err: nil,
	},
	{
		in:  []byte(`{"multi": {}, "key": {}}`),
		err: nil,
	},
	{
		in:  []byte(`{"key":{"nested": {}}}`),
		err: nil,
	},
	{
		in:  []byte(`{"nexted":{"multi": {}, "key": []}}`),
		err: nil,
	},
}

func TestValid(t *testing.T) {
	for _, tt := range examples {
		err := Valid(tt.in)
		if !reflect.DeepEqual(tt.err, err) {
			t.Errorf("For %s\nExpected: %#v\nGot:      %#v", tt.in, tt.err, err)
		}
	}
}

func BenchmarkValid(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for _, tt := range examples {
		err := Valid(tt.in)
		if !reflect.DeepEqual(tt.err, err) {
			b.Fatalf("For %s\nExpected: %#v\nGot:      %#v", tt.in, tt.err, err)
		}
	}
}

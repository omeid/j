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
	{
		in:  []byte(`["simple"]`),
		err: nil,
	},
	{
		in:  []byte(`["multi"  , "value"]`),
		err: nil,
	},
	{
		in:  []byte(`{"simple": true, "wrong": false, "value" : null }`),
		err: nil,
	},
	{
		in:  []byte(`{"a": true, "b": [true] }`),
		err: nil,
	},
	{
		in:  []byte(`{"escaped": "\"\\\/\b\f\n\r\t", "bad": "\x" }`),
		err: makeSyntaxError(42, 1, 42, "invalid character 'x' Expected a qoutation mark, reverse solidus, or a control character"),
	},
	{
		in:  []byte(`{"escaped": "\\\n\r\u000000NotHex\uFFFFFF\uffffff", "bad": "\uL" }`),
		err: makeSyntaxError(63, 1, 63, "invalid character 'L' Expected a Hexadecimal digit."),
	},
	{
		in:  []byte(`{"\\\/\b\f\uF00F00key":{"nested\n\t": {}}}`),
		err: nil,
	},
	{
		in:  []byte(`{"\\\/\b\f\uF00F00key":{"nested\n\t": ["BadEscape: \uX"]}}`),
		err: makeSyntaxError(54, 1, 54, "invalid character 'X' Expected a Hexadecimal digit."),
	},

	{
		in: []byte(`[
0,
-3,
-234324,
5,
324132432,
-7,
-0.8,
-1.9,
0.10,
1.11,
-0.9e-12,
-0.10E+13
]`),
		err: nil,
	},
}

func TestValid(t *testing.T) {
	for _, tt := range examples {
		// fmt.Printf("\n\n%s\n\n", tt.in)
		err := Valid(tt.in)
		if !reflect.DeepEqual(tt.err, err) {
			t.Errorf("For:\n%s\nExpected: %#v\nGot:      %#v", tt.in, tt.err, err)
			break
		}
	}
}

func BenchmarkValid(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, tt := range examples {
			err := Valid(tt.in)
			//if !reflect.DeepEqual(tt.err, err) {
			//	b.Fatalf("For %s\nExpected: %#v\nGot:      %#v", tt.in, tt.err, err)
			//}
			_ = err
		}
	}
}

package j

import (
	"reflect"
	"testing"
)

func makeParseError(offset int, err error) error {
	return ParseError{
		offset: offset,
		err:    err,
	}
}

type b []byte
type bb []b

type parseTest struct {
	srcs   []b
	expect Value
	err    error
}

var parseTests = []parseTest{
	//fuzz stuff
	parseTest{
		srcs: bb{b(`[""]`)},
		expect: NewArray([]Value{
			NewString(""),
		}),
		err: nil,
	},
	parseTest{
		srcs: bb{b(`["\ufffd0"]`)},
		expect: NewArray([]Value{
			NewString("\ufffd0"),
		}),
		err: nil,
	},

	parseTest{
		srcs:   bb{b(`[[]]`)},
		expect: NewArray([]Value{NewArray(nil)}),
		err:    nil,
	},

	parseTest{
		srcs:   bb{b(`[{}]`)},
		expect: NewArray([]Value{NewObject(nil)}),
		err:    nil,
	},

	parseTest{
		srcs:   bb{b(`[[0]]`)},
		expect: NewArray([]Value{NewArray([]Value{NewNumber(b(`0`))})}),
		err:    nil,
	},

	parseTest{
		srcs:   bb{b(`[]`), b(`        []`), b(`   [   ]    `), b("[] \t \r \n"), b("\r \t []")},
		expect: NewArray(nil),
		err:    nil,
	},
	parseTest{
		srcs:   bb{b(`["hello"]`)},
		expect: NewArray([]Value{NewString("hello")}),
		err:    nil,
	},
	parseTest{
		srcs: bb{
			b(`["hello", "world", ["and", "the under", ["world"]], "boo yeah"]`),
			b(`[      "hello",             "world",				 ["and", "the under", ["world"]  ], "boo yeah"    ]`),
		},
		expect: NewArray([]Value{
			NewString("hello"),
			NewString("world"),

			NewArray([]Value{
				NewString("and"),
				NewString("the under"),
				NewArray([]Value{
					NewString("world"),
				}),
			}),
			NewString("boo yeah"),
		}),
		err: nil,
	},
	parseTest{
		srcs: bb{
			b(`[0, 1, 10, 11, -2, -22,  [1,2,3, "test", 3]]`),
		},
		expect: NewArray([]Value{
			NewNumber(b(`0`)),
			NewNumber(b(`1`)),
			NewNumber(b(`10`)),
			NewNumber(b(`11`)),
			NewNumber(b(`-2`)),
			NewNumber(b(`-22`)),

			NewArray([]Value{
				NewNumber(b(`1`)),
				NewNumber(b(`2`)),
				NewNumber(b(`3`)),
				NewString("test"),
				NewNumber(b(`3`)),
			}),
		}),
		err: nil,
	},
	parseTest{
		srcs:   bb{b(`{    "hello"    :    "world"}`), b(`{"hello":"world"}`)},
		expect: NewObject([]Member{NewMember("", "hello", NewString("world"))}),
		err:    nil,
	},

	parseTest{
		srcs: bb{b(`{"alpha": [{"beta": ["delta", 25.5 ]}]}`)},
		expect: NewObject([]Member{
			NewMember("", "alpha", NewArray([]Value{
				NewObject([]Member{
					NewMember(
						"",
						"beta",
						NewArray([]Value{
							NewString("delta"),
							NewNumberFloat64(25.5),
						}),
					),
				}),
			})),
		}),
		err: nil,
	},

	parseTest{
		srcs: bb{
			b(`[ 0, 0 ,  0 ,  0, -3, -234324, 5, 324132432, -7, -0.8, -1.9, 0.10, 1.11, -0.9e-12, -0.10E+13 ]`),
		},
		expect: NewArray([]Value{
			NewNumber(b(`0`)),
			NewNumber(b(`0`)),
			NewNumber(b(`0`)),
			NewNumber(b(`0`)),

			NewNumber(b(`-3`)),
			NewNumber(b(`-234324`)),
			NewNumber(b(`5`)),
			NewNumber(b(`324132432`)),
			NewNumber(b(`-7`)),
			NewNumber(b(`-0.8`)),
			NewNumber(b(`-1.9`)),
			NewNumber(b(`0.10`)),
			NewNumber(b(`1.11`)),
			NewNumber(b(`-0.9e-12`)),
			NewNumber(b(`-0.10E+13`)),
		}),
		err: nil,
	},
}

func TestDecode(t *testing.T) {
	for _, pt := range parseTests {

		for _, src := range pt.srcs {
			value, err := Decode(src)

			if !reflect.DeepEqual(pt.err, err) {
				t.Errorf("For Decode: `%s`\nExpected: %#v\nGot:      %#v", src, pt.err, err)
				return
			}

			exp, err := Encode(pt.expect)
			if err != nil {
				t.Errorf("Failed to Encode Expection: `%s`\n%s", src, err)
				return
			}

			enc, err := Encode(value)
			if err != nil {
				t.Errorf("Failed to Encode Just-Decoded Value: `%s`\n%s", value, err)
				return
			}

			if !Match(value, pt.expect) {
				t.Errorf("Decoded Value and Expection mismatch: `%s`\nExpected  %#v\nGot       %#v\nExp:\n%s\nEnc:\n%s", src, pt.expect, value, exp, enc)
				return
			}

			value, err = Decode(enc)
			if err != nil {
				t.Errorf("Failed to Decode just Encoded-Decoded of src: `%s`\n%s", value, err)
				return
			}

			if !Match(value, pt.expect) {
				t.Errorf("Mismatch for Decode of just Encoded-Decoded of src and expection\nFor:  `%s`\nExpected  %#v\nGot       %#v\nExp:\n%s\nEnc:\n%s", src, pt.expect, value, exp, enc)
				// return
			}

		}
	}
}

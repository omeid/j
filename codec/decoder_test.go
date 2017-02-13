package codec

import (
	"reflect"
	"testing"

	"github.com/omeid/j"
	"github.com/omeid/j/internal/valuetest"
	"github.com/omeid/j/mutable"
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
	expect j.Value
	err    error
}

var parseTests = []parseTest{
	//fuzz stuff
	parseTest{
		srcs: bb{b(`[""]`)},
		expect: func() j.Value {
			ar := mutable.NewArray()
			ar.Add(mutable.NewString("").Value())
			return ar.Value()
		}(),
		err: nil,
	},
	parseTest{
		srcs: bb{b(`["\ufffd0"]`)},
		expect: func() j.Value {
			ar := mutable.NewArray()
			ar.Add(mutable.NewString("\ufffd0").Value())
			return ar.Value()
		}(),
		err: nil,
	},

	parseTest{
		srcs: bb{b(`[[]]`)},
		expect: func() j.Value {
			ar := mutable.NewArray()
			ar.Add(mutable.NewArray().Value())
			return ar.Value()
		}(),
		err: nil,
	},

	parseTest{
		srcs: bb{b(`[{}]`)},
		expect: func() j.Value {
			ar := mutable.NewArray()
			ar.Add(mutable.NewObject().Value())
			return ar.Value()
		}(),
		err: nil,
	},

	parseTest{
		srcs: bb{b(`[[0]]`)},
		expect: func() j.Value {

			in := mutable.NewArray()
			in.Add(mutable.NewNumber(b(`0`)).Value())

			out := mutable.NewArray()
			out.Add(in.Value())

			return out.Value()
		}(),
		err: nil,
	},

	parseTest{
		srcs:   bb{b(`[]`), b(`        []`), b(`   [   ]    `), b("[] \t \r \n"), b("\r \t []")},
		expect: mutable.NewArray().Value(),
		err:    nil,
	},
	parseTest{
		srcs: bb{b(`["hello"]`)},
		expect: func() j.Value {
			ar := mutable.NewArray()
			ar.Add(mutable.NewString("hello").Value())
			return ar.Value()
		}(),
		err: nil,
	},
	parseTest{
		srcs: bb{
			b(`["hello", "world", ["and", "the under", ["world"]], "boo yeah"]`),
			b(`[      "hello",             "world",				 ["and", "the under", ["world"]  ], "boo yeah"    ]`),
		},
		expect: func() j.Value {

			world := mutable.NewArray()
			world.Add(mutable.NewString("world").Value())

			andunderworld := mutable.NewArray()
			andunderworld.Add(mutable.NewString("and").Value())
			andunderworld.Add(mutable.NewString("the under").Value())
			andunderworld.Add(world.Value())

			ar := mutable.NewArray()
			ar.Add(mutable.NewString("hello").Value())
			ar.Add(mutable.NewString("world").Value())

			ar.Add(andunderworld.Value())
			ar.Add(mutable.NewString("boo yeah").Value())

			return ar.Value()
		}(),
		err: nil,
	},
	parseTest{
		srcs: bb{
			b(`[0, 1, 10, 11, -2, -22,  [1,2,3, "test", 3]]`),
		},
		expect: func() j.Value {
			ar := mutable.NewArray()

			ar.Add(mutable.NewNumber(b(`0`)).Value())
			ar.Add(mutable.NewNumber(b(`1`)).Value())
			ar.Add(mutable.NewNumber(b(`10`)).Value())
			ar.Add(mutable.NewNumber(b(`11`)).Value())
			ar.Add(mutable.NewNumber(b(`-2`)).Value())
			ar.Add(mutable.NewNumber(b(`-22`)).Value())

			arr := mutable.NewArray()
			arr.Add(mutable.NewNumber(b(`1`)).Value())
			arr.Add(mutable.NewNumber(b(`2`)).Value())
			arr.Add(mutable.NewNumber(b(`3`)).Value())
			arr.Add(mutable.NewString("test").Value())
			arr.Add(mutable.NewNumber(b(`3`)).Value())

			ar.Add(arr.Value())
			return ar.Value()
		}(),
		err: nil,
	},
	parseTest{
		srcs: bb{b(`{"hello":"world"}`)},
		expect: func() j.Value {

			obj := mutable.NewObject()
			hw := mutable.NewMember("", "hello", mutable.NewString("world").Value())
			err := obj.Add(hw)
			if err != nil {
				panic(err)
			}

			return obj.Value()
		}(),
		err: nil,
	},

	parseTest{
		srcs: bb{
			b(`[ 0, 0 ,  0 ,  0, -3, -234324, 5, 324132432, -7, -0.8, -1.9, 0.10, 1.11, -0.9e-12, -0.10E+13 ]`),
		},
		expect: func() j.Value {
			ar := mutable.NewArray()

			ar.Add(mutable.NewNumber(b(`0`)).Value())
			ar.Add(mutable.NewNumber(b(`0`)).Value())
			ar.Add(mutable.NewNumber(b(`0`)).Value())
			ar.Add(mutable.NewNumber(b(`0`)).Value())

			ar.Add(mutable.NewNumber(b(`-3`)).Value())
			ar.Add(mutable.NewNumber(b(`-234324`)).Value())
			ar.Add(mutable.NewNumber(b(`5`)).Value())
			ar.Add(mutable.NewNumber(b(`324132432`)).Value())
			ar.Add(mutable.NewNumber(b(`-7`)).Value())
			ar.Add(mutable.NewNumber(b(`-0.8`)).Value())
			ar.Add(mutable.NewNumber(b(`-1.9`)).Value())
			ar.Add(mutable.NewNumber(b(`0.10`)).Value())
			ar.Add(mutable.NewNumber(b(`1.11`)).Value())
			ar.Add(mutable.NewNumber(b(`-0.9e-12`)).Value())
			ar.Add(mutable.NewNumber(b(`-0.10E+13`)).Value())

			return ar.Value()
		}(),
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

			if !valuetest.Match(value, pt.expect) {
				t.Errorf("Decoded Value and Expection mismatch: `%s`\nExpected  %#v\nGot       %#v\nExp:\n%s\nEnc:\n%s", src, pt.expect, value, exp, enc)
				return
			}

			value, err = Decode(enc)
			if err != nil {
				t.Errorf("Failed to Decode just Encoded-Decoded of src: `%s`\n%s", value, err)
				return
			}

			if !valuetest.Match(value, pt.expect) {
				t.Errorf("Mismatch for Decode of just Encoded-Decoded of src and expection\nFor:  `%s`\nExpected  %#v\nGot       %#v\nExp:\n%s\nEnc:\n%s", src, pt.expect, value, exp, enc)
				// return
			}

		}
	}
}

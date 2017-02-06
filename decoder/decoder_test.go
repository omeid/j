package decoder

import (
	"reflect"
	"testing"

	"github.com/omeid/j"
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
	parseTest{
		srcs:   bb{b(`[]`), b(`        []`), b(`   [   ]    `), b("[] \t \r \n"), b("\r \t []")},
		expect: mutable.NewArray().Value(),
		err:    nil,
	},
	parseTest{
		srcs: bb{b(`["hello"]`)},
		expect: func() j.Value {
			ar := mutable.NewArray()
			err := ar.Add(mutable.NewString(b(`hello`)).Value())
			if err != nil {
				panic(err)
			}
			return ar.Value()
		}(),
		err: nil,
	},
	parseTest{
		srcs: bb{
			b(`["hello", "world", ["and", "the under", ["world"]]]`),
			b(`           ["hello",             "world",				 ["and", "the under", ["world"]  ]    ]     `),
		},
		expect: func() j.Value {

			world := mutable.NewArray()
			err := world.Add(mutable.NewString(b(`world`)).Value())
			if err != nil {
				panic(err)
			}

			andunderworld := mutable.NewArray()
			err = andunderworld.Add(mutable.NewString(b(`and`)).Value())
			if err != nil {
				panic(err)
			}
			err = andunderworld.Add(mutable.NewString(b(`the under`)).Value())
			if err != nil {
				panic(err)
			}
			err = andunderworld.Add(world.Value())
			if err != nil {
				panic(err)
			}

			ar := mutable.NewArray()
			err = ar.Add(mutable.NewString(b(`hello`)).Value())
			if err != nil {
				panic(err)
			}
			err = ar.Add(mutable.NewString(b(`world`)).Value())
			if err != nil {
				panic(err)
			}

			err = ar.Add(andunderworld.Value())
			if err != nil {
				panic(err)
			}

			return ar.Value()
		}(),
		err: nil,
	},
	// parseTest{
	// 	srcs:   bb{b(`{}`)},
	// 	expect: mutable.NewObject().Value(),
	// 	err:    nil,
	// },
}

func TestDecode(t *testing.T) {
	for _, pt := range parseTests {

		for _, src := range pt.srcs {
			value, err := Decode(src)

			if !reflect.DeepEqual(pt.err, err) {
				t.Errorf("For: `%s`\nExpected: %#v\nGot:      %#v", src, pt.err, err)
				break
			}

			if !Match(value, pt.expect) {
				t.Errorf("For:  `%s`\nExpected  %#v\nGot       %#v", src, pt.expect, value)
				break
			}
		}
	}
}

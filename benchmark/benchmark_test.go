package benchmark

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/omeid/j"
	"github.com/omeid/j/benchmark"
)

var largeStructText []byte

func init() {
	var err error
	largeStructText, err = ioutil.ReadFile("example.json")
	if err != nil {
		log.Fatal(err)
	}
}

func TestLargeStruct(t *testing.T) {
	jv, err := j.Decode(largeStructText)
	if err != nil {
		t.Error(err)
		return
	}

	if jv == nil {
		panic("wot")
	}

	s := LargeStruct{}
	err = s.FromJSON(jv)
	if err != nil {
		t.Error(err)
		return
	}

	jv2, err := s.ToJSON()
	if err != nil {
		t.Error(err)
		return
	}
}

func BenchmarLargeStruct(b *testing.B) {
	b.SetBytes(int64(len(LargeStructText)))

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		jv, err := j.Decode(largeStructText)
		if err != nil {
			b.Error(err)
			return
		}

		if jv == nil {
			panic("wot")
		}

		s := benchmark.LargeStruct{}
		err = s.FromJSON(jv)
		if err != nil {
			b.Error(err)
			return
		}
	}
	b.SetBytes(int64(len(largeStructText)))

}

func BenchmarValid(b *testing.B) {
	b.SetBytes(int64(len(LargeStructText)))

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := j.Vecode(largeStructText)
		if err != nil {
			b.Error(err)
		}
	}
}

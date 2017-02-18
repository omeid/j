package benchmark

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/omeid/j"
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

	dec, err := j.Encode(jv2)
	if err != nil {
		t.Error(err)
		return
	}

	if !j.Match(jv, jv2) {
		t.Errorf("Value missmatch.\nsrc: %s\n\ndec: %s\n\nphase 1: %v\n phase 2: %v\n", largeStructText, dec, jv, jv2)
	}
}

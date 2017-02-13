// +build gofuzz

package codec

import (
	"fmt"

	"github.com/omeid/j/internal/valuetest"
)

// Fuzz is used for fuzzer testing, ignore it.
func Fuzz(input []byte) int {
	err := Valid(input)
	if err != nil {
		return -1
	}

	value, err := Decode(input)
	if err != nil {
		panic("valid input failed to decode.")
	}

	enc, err := Encode(value)

	if err != nil {
		fmt.Println(err)
		panic("valid value failed to encode")
	}

	value2, err := Decode(enc)

	if err != nil {
		fmt.Println(err)
		panic("valid input failed to decode")
	}

	if !valuetest.Match(value, value2) {
		fmt.Println(err)
		panic("Second level mismatch!")
	}

	return 1
}

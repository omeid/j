// Package numbers provides JSON number encoding and decode for use in codec and mutable package. Other implementation may use the functions here for a consistent encoding of numbers.
package numbers

import "strconv"

type floatEncoder int // number of bits

// EncodeUint64 encodes the provided value to a byte slice.
func EncodeUint64(n uint64) []byte {
	return []byte(strconv.FormatUint(n, 10))
}

// EncodeInt64 encodes the provided value to a byte slice.
func EncodeInt64(n int64) []byte {
	return []byte(strconv.FormatInt(n, 10))
}

// EncodeFloat64 encodes the provided value to a byte slice.
func EncodeFloat64(n float64) []byte {
	return []byte(strconv.FormatFloat(n, 'g', -1, 64))
}

// DecodeInt64 attempts parse the provided value to int64.
func DecodeInt64(b []byte) (int64, error) {
	return strconv.ParseInt(string(b), 10, 64)
}

// DecodeFloat64 attempts parse the provided value to float64.
func DecodeFloat64(b []byte) (float64, error) {
	return strconv.ParseFloat(string(b), 64)
}

// DecodeUint64 attempts parse the provided value to int64.
func DecodeUint64(b []byte) (uint64, error) {
	return strconv.ParseUint(string(b), 10, 64)
}

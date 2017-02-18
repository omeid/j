package j

import "github.com/omeid/j/numbers"

// NewNumber creates a number type from provided raw value.
// It is the callers responsiblity to make sure it is correctly
// encoded JSON numbers.
func NewNumber(raw []byte) Value {
	return &number{value: baseValue, raw: raw}
}

// NewNumberFloat64 creates a number from the provided value.
func NewNumberFloat64(v float64) Value {
	n := &number{value: baseValue}
	n.raw = numbers.EncodeFloat64(v)
	return n
}

// NewNumberInt64 creates a number from the provided value.
func NewNumberInt64(v int64) Value {
	n := &number{value: baseValue}
	n.raw = numbers.EncodeInt64(v)
	return n
}

// NewNumberUint64 creates a number from the provided value.
func NewNumberUint64(v uint64) Value {
	n := &number{value: baseValue}
	n.raw = numbers.EncodeUint64(v)
	return n
}

type number struct {
	*value
	raw []byte
}

func (n number) Type() Type {
	return NumberType
}

func (n *number) Raw() []byte {
	return n.raw
}

func (n *number) Float64() (float64, error) {
	return numbers.DecodeFloat64(n.raw)
}

func (n *number) Int64() (int64, error) {
	return numbers.DecodeInt64(n.raw)
}

func (n *number) Uint64() (uint64, error) {
	return numbers.DecodeUint64(n.raw)
}

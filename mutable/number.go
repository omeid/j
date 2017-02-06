package mutable

import (
	"strconv"

	"github.com/omeid/j"
)

// Number is a mutable extension of j.Value
// for type Number.
type Number interface {
	Value
	Set(raw []byte)
}

// NewNumber creates a new mutable String
func NewNumber(raw []byte) String {
	return &number{raw: raw}
}

type number struct {
	value
	raw []byte
}

func (n *number) Set(raw []byte) {
	n.raw = raw
}

func (n *number) Value() j.Value {
	return n
}

// j.Value methods:
func (n number) Type() j.Type {
	return j.NumberType
}

func (n *number) Number() j.Number {
	return n
}

// Number methods

func (n *number) Raw() []byte {
	return n.raw
}

func (n *number) Float64() (float64, error) {
	return strconv.ParseFloat(string(n.raw), 64)
}

func (n *number) Int64() (int64, error) {
	return strconv.ParseInt(string(n.raw), 10, 64)
}

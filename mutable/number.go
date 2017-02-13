package mutable

import (
	"github.com/omeid/j"
	"github.com/omeid/j/codec/numbers"
)

// Number is a mutable extension of j.Value
// for type num.
type Number interface {
	Value
	Set(raw []byte)
	SetFloat64(float64)
	SetInt64(int64)
	SetUint64(uint64)
}

// NewNumber creates a number type from provided raw value.
// It is the callers responsiblity to make sure it is correctly
// encoded JSON numbers.
func NewNumber(raw []byte) Number {
	return &number{raw: raw}
}

// NewNumberFloat64 creates a number from the provided value.
func NewNumberFloat64(v float64) Number {
	n := &number{}
	n.SetFloat64(v)
	return n
}

// NewNumberInt64 creates a number from the provided value.
func NewNumberInt64(v int64) Number {
	n := &number{}
	n.SetInt64(v)
	return n
}

// NewNumberUint64 creates a number from the provided value.
func NewNumberUint64(v uint64) Number {
	n := &number{}
	n.SetUint64(v)
	return n
}

type number struct {
	value
	raw []byte
}

func (n *number) Set(raw []byte) {
	n.raw = raw
}

func (n *number) SetFloat64(f float64) {
	n.raw = numbers.EncodeFloat64(f)
}

func (n *number) SetInt64(v int64) {
	n.raw = numbers.EncodeInt64(v)
}

func (n *number) SetUint64(u uint64) {
	n.raw = numbers.EncodeUint64(u)
}

func (n *number) Value() j.Value {
	return n
}

// j.Value methods:
func (n number) Type() j.Type {
	return j.NumberType
}

// num methods

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

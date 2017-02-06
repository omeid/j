package mutable

import "github.com/omeid/j"

// Null is a mutable extension of j.Value
// for type Bool.
type Null interface {
	Value
}

// NewNull creates a new mutable null
func NewNull() Null {
	return &null{}
}

type null struct {
	value
}

func (n *null) Value() j.Value {
	return n
}

// j.Value methods:
func (n null) Type() j.Type {
	return j.NullType
}

package mutable

import "github.com/omeid/j"

// Bool is a mutable extension of j.Value
// for type Bool.
type Bool interface {
	Value
	Set(bool)
}

// NewBool creates a new mutable object
func NewBool(value bool) Bool {
	return &boolean{boolean: value}
}

type boolean struct {
	value
	boolean bool
}

func (b *boolean) Set(v bool) {
	b.boolean = v
}

func (b *boolean) Value() j.Value {
	return b
}

// j.Value methods:
func (b boolean) Type() j.Type {
	return j.BoolType
}

func (b *boolean) Bool() bool {
	return b.boolean
}

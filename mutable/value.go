package mutable

import (
	"fmt"

	"github.com/omeid/j"
)

// value implements all the j.Value
// methods with a simple panic, this type
// is embedded as a shortcut to satisfying
// j.Value interface.
type value struct {
	typ j.Type
}

func (b value) Type() j.Type {
	panic(fmt.Sprintf("Called on type %v", b.typ))
}
func (b value) Members() []j.Member {
	panic(fmt.Sprintf("Called Object on type %v", b.typ))
}
func (b value) Values() []j.Value {
	panic(fmt.Sprintf("Called Array on type %v", b.typ))
}
func (b value) Bool() bool {
	panic(fmt.Sprintf("Called Bool on type %v", b.typ))
}
func (b value) Number() j.Number {
	panic(fmt.Sprintf("Called Number on type %v", b.typ))
}
func (b value) String() j.String {
	panic(fmt.Sprintf("Called String on type %v", b.typ))
}

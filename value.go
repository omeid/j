package j

import "fmt"

// value implements all the Value
// methods with a simple panic, this type
// is embedded as a shortcut to satisfying
// Value interface.

var baseValue = &value{}

type value struct {
	typ Type
}

func (b value) Type() Type {
	panic(fmt.Sprintf("Called on type %v", b.typ))
}
func (b value) Members() []Member {
	panic(fmt.Sprintf("Called Members on type %v", b.typ))
}
func (b value) Member(name string) Value {
	panic(fmt.Sprintf("Called Member on type %v", b.typ))
}

func (b value) Values() []Value {
	panic(fmt.Sprintf("Called Array on type %v", b.typ))
}
func (b value) Bool() bool {
	panic(fmt.Sprintf("Called Bool on type %v", b.typ))
}

func (b value) Float64() (float64, error) {
	panic(fmt.Sprintf("Called Number on type %v", b.typ))
}

func (b value) Int64() (int64, error) {
	panic(fmt.Sprintf("Called Number on type %v", b.typ))
}

func (b value) Uint64() (uint64, error) {
	panic(fmt.Sprintf("Called Number on type %v", b.typ))
}

func (b value) Raw() []byte {
	panic(fmt.Sprintf("Called Number on type %v", b.typ))
}

func (b value) String() String {
	panic(fmt.Sprintf("Called String on type %v", b.typ))
}

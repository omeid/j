package j

// value implements all the Value
// methods with a simple panic, this type
// is embedded as a shortcut to satisfying
// Value interface.

var baseValue = &value{}

type value struct{}

func (b value) Type() Type {
	panic("Illegal Call")
}
func (b value) Members() []Member {
	panic("Illegal Call")
}
func (b value) Member(name string) Value {
	panic("Illegal Call")
}

func (b value) Values() []Value {
	panic("Illegal Call")
}
func (b value) Bool() bool {
	panic("Illegal Call")
}

func (b value) Float64() (float64, error) {
	panic("Illegal Call")
}

func (b value) Int64() (int64, error) {
	panic("Illegal Call")
}

func (b value) Uint64() (uint64, error) {
	panic("Illegal Call")
}

func (b value) Raw() []byte {
	panic("Illegal Call")
}

func (b value) String() String {
	panic("Illegal Call")
}

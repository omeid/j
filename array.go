package j

// NewArray creates a new mutable array value
func NewArray(values []Value) Value {
	return &array{value: baseValue, values: values}
}

type array struct {
	*value
	values []Value
}

// Value methods:
func (a array) Type() Type {
	return ArrayType
}

func (a *array) Values() []Value {
	return a.values
}

package j

// NewBool creates a new mutable object
func NewBool(value bool) Value {
	return &boolean{value: baseValue, boolean: value}
}

type boolean struct {
	*value
	boolean bool
}

// Value methods:
func (b boolean) Type() Type {
	return BoolType
}

func (b *boolean) Bool() bool {
	return b.boolean
}

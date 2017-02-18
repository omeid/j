package j

var n = &null{value: baseValue}

// NewNull creates a new null
func NewNull() Value {
	return n
}

type null struct{ *value }

func (n null) Type() Type {
	return NullType
}

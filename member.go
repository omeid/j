package j

// NewMember creates a new mutable object
func NewMember(tag string, name string, value Value) Member {
	return &member{tag: tag, name: name, value: value}
}

type member struct {
	tag   string
	name  string
	value Value
}

func (m *member) Tag() string {
	return m.tag
}

func (m *member) Name() *string {
	return &m.name
}

func (m *member) Value() Value {
	return m.value
}

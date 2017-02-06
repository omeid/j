package mutable

import "github.com/omeid/j"

// Member is a mutable extension of j.Member
// As Member is not a j.Value type, it does
// _not_ implement the mutable.Value interface.
type Member interface {
	j.Member
}

// NewMember creates a new mutable object
func NewMember(tag string, name string, value j.Value) Member {
	return &member{tag: tag, name: name, value: value}
}

type member struct {
	tag   string
	name  string
	value j.Value
}

func (m *member) Tag() string {
	return m.tag
}

func (m *member) Name() string {
	return m.name
}

func (m *member) Value() j.Value {
	return m.value
}

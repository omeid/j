package json

import "encoding/json"

type value struct {
}

type object struct {
	members []Member
}

func (o *object) Members() []Member {
	return o.members
}

type member struct {
	key   string
	tag   string
	value Value
}

func (m *member) Tag() string {
	return m.tag
}

func (m *member) Key() string {
	return m.key
}

func (m *member) Value() Value {
	return m.value
}

type array struct {
	values []Value
}

func (a *array) Values() []Value {
	return a.values
}

type number json.Number

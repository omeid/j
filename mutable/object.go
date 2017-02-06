// Package mutable provides a mutable implementation of the j.Value
package mutable

import (
	"errors"

	"github.com/omeid/j"
)

// Object represents an interface to a mutable j.Value of
// Type j.ObjecType.
type Object interface {
	Value
	j.Value
	Add(name string, value j.Value) error
	Remove(name string) error
}

// NewObject creates a new mutable object
func NewObject() Object {
	return &object{members: []j.Member{}}
}

type object struct {
	value
	members []j.Member
}

func (o *object) Value() j.Value {
	return o
}

// Add the provided value at said key point, it is your responsiblity to make
// sure there is no duplicate keys.
func (o *object) Add(name string, value j.Value) error {
	// m := NewMember()
	// m.SetKey(key)
	// m.SetValue(value)

	// o.members = append(o.members, m)
	return nil
}

// Removes the Member at the provided key, it is your responsiblity to make sure
// the value exists.
func (o *object) Remove(name string) error {
	for i, m := range o.members {
		if m.Name() == name {
			o.members[i] = o.members[len(o.members)-1]
			o.members = o.members[:len(o.members)-1]
			return nil
		}
	}
	return errors.New("No Such Member")
}

// j.Value methods

func (o object) Type() j.Type {
	return j.ObjectType
}

func (o *object) Members() []j.Member {
	return o.members
}

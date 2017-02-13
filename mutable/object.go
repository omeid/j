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
	Add(Member) error
	Remove(name string) error
}

// NewObject creates a new mutable object
func NewObject() Object {
	return &object{members: map[string]j.Member{}}
}

type object struct {
	value
	members map[string]j.Member
}

func (o *object) Value() j.Value {
	return o
}

// Add the provided value at said key point, it is your responsiblity to make
// sure there is no duplicate keys.
func (o *object) Add(m Member) error {
	key := m.Name()

	if _, ok := o.members[key]; ok {
		return errors.New("Member already exists.")
	}
	o.members[key] = m
	return nil
}

// Removes the Member at the provided key, it is your responsiblity to make sure
// the value exists.
func (o *object) Remove(name string) error {
	if _, ok := o.members[name]; !ok {
		return errors.New("No Such Member")
	}
	delete(o.members, name)
	return nil
}

// j.Value methods

func (o object) Type() j.Type {
	return j.ObjectType
}

func (o *object) Members() []j.Member {
	ms := make([]j.Member, len(o.members))
	i := 0
	for _, m := range o.members {
		ms[i] = m
		i++
	}

	return ms
}

func (o *object) Member(name string) j.Value {
	m, ok := o.members[name]
	if !ok {
		return nil
	}

	return m.Value()
}

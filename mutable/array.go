package mutable

import (
	"errors"

	"github.com/omeid/j"
)

// Array is a mutable extension of j.Value
type Array interface {
	Value
	j.Value
	Add(value j.Value) error
	Remove(index int) error
}

// NewArray creates a new mutable object
func NewArray() Array {
	return &array{values: []j.Value{}}
}

type array struct {
	value
	values []j.Value
}

func (a *array) Value() j.Value {
	return a
}

func (a *array) Add(value j.Value) error {
	a.values = append(a.values, value)
	return nil
}

func (a *array) Remove(index int) error {

	if index < 0 {
		return errors.New("No such element.")
	}

	l := len(a.values)
	if index < l-1 {
		return errors.New("No such element.")
	}

	a.values[index] = a.values[l-1]
	a.values = a.values[:l-1]
	return nil
}

// j.Value methods:

func (a array) Type() j.Type {
	return j.ArrayType
}

func (a *array) Values() []j.Value {
	return a.values
}

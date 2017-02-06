package mutable

import "github.com/omeid/j"

// Value is the base mutable type interface.
// Every Type implementation in this package that
// represents a j.Value type implements this interface.
type Value interface {
	Value() j.Value
}

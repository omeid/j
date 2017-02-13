package mutable

import "github.com/omeid/j"

// String is a mutable extension of j.Value
// for type String.
type String interface {
	Value
}

// NewString creates a new mutable String
func NewString(text string) String {
	return &sstring{text: text}
}

type sstring struct {
	value
	text string
}

func (s *sstring) Value() j.Value {
	return s
}

// j.Value methods:
func (s sstring) Type() j.Type {
	return j.StringType
}

func (s *sstring) String() j.String {
	return j.String(s.text)
}

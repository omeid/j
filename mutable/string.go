package mutable

import "github.com/omeid/j"

// String is a mutable extension of j.Value
// for type String.
type String interface {
	Value
}

// NewString creates a new mutable String
func NewString(raw []byte) String {
	return &sstring{raw: raw}
}

type sstring struct {
	value
	raw []byte
}

func (s *sstring) Value() j.Value {
	return s
}

// j.Value methods:
func (s sstring) Type() j.Type {
	return j.StringType
}

func (s *sstring) String() j.String {
	//TODO: unescape
	return j.String(s.raw)
}

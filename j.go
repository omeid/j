package j

// A Type describes the type of a value
// as per JSON spec.
type Type uint

// Type of Objects as per the JSON spec.
const (
	InvalidType Type = iota
	ObjectType
	MemberType
	ArrayType
	BoolType
	NumberType
	StringType
	NullType
)

// Value is the json value.
// The Methods should be only called on specified type documented.
// Null type accepts no methods.
type Value interface {
	Type() Type

	// Object applies to Object type.
	Object() Object

	// Array applies to Array type.
	Array() Array

	// Bool applies to Bool type.
	Bool() bool

	// Number applies to Number type.
	Number() Number

	// String applies to string type.
	String() string
}

// Object represents a JSON object.
type Object interface {
	Members() []Member
}

// Array represents a JSON array.
type Array interface {
	Values() []Value
}

// Member is a JSON object Member.
type Member interface {
	// Tag is the struct member tag.
	Tag() string

	Name() string
	Value() Value
}

// Number is a json number type.
type Number interface {
	Float64() (float64, error)
	Int64() (int64, error)
	Raw() []byte
}

// Iterator is the common interface for JSON iterators.
type Iterator interface {
	// Next returns the next value if any or nil.
	// More indicates whatever there is any more
	// _possible_ values, so you must always check
	// to make sure the item is not nil.
	// Calling Next after the last item MUST return
	// `nil, false`.
	Next() (item Value, more bool)
	// Len returns -1 for indefinite streams or an advisory
	// number indicating the the number of values it holds.
	// The iterator _MAY_ return more or less items than what
	// is advised but _MUST_ always return a finite
	// number of items when Len returns 1 or more.
	// 0 indicates exactly 0 items.
	Len() int
}

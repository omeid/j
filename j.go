package j

// A Type describes the type of a value
// as per JSON spec.
type Type uint

//go:generate stringer -type=Type

// Type of Objects as per the JSON spec.
const (
	InvalidType Type = iota
	ObjectType
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
	Member(string) Value
	Members() []Member

	// Array applies to Array type.
	Values() []Value

	// Bool applies to Bool type.
	Bool() bool

	// Number applies to Number type.
	Float64() (float64, error)
	Int64() (int64, error)
	Uint64() (uint64, error)
	Raw() []byte

	// String applies to string type.
	String() String
}

// The String type is simply a string.
// It merely exists to keep the fmt package
// from calling it and resulting in panics.
type String *string

// Member is a JSON object Member.
type Member interface {
	// Tag is the struct member tag.
	Tag() string

	Name() *string
	Value() Value
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

package json

// A Kind describes the type of a value
// as per JSON spec.
type Kind uint

// sKind here.
const (
	InvalidKind Kind = iota
	ObjectKind
	MemberKind
	ArrayKind
	BoolKind
	NumberKind
	StringKind
	NullKind
)

// Value is the json value.
type Value interface {
	Kind() Kind

	// Object
	Object() Object

	// Array
	Array() Array

	// Bool
	Bool() bool

	// Number
	Number() Number

	// String
	String() string

	// Nil
	Null()
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

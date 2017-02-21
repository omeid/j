package j

import (
	"testing"
)

var Types = []Type{
	ObjectType,
	ArrayType,
	BoolType,
	NumberType,
	StringType,
	NullType,
}

func checkType(t *testing.T, v Value, typ Type) {
	ttyp := v.Type()
	if ttyp != typ {
		t.Errorf("Expected Type %v got %v", typ, ttyp)
	}
}

func TestNewNull(t *testing.T) {
	n := NewNull()

	checkType(t, n, NullType)
	checkPanics(t, n)
}

func TestNewString(t *testing.T) {
	n := NewString("")

	checkType(t, n, StringType)
	checkPanics(t, n)
}

func TestNewBool(t *testing.T) {
	n := NewBool(false)

	checkType(t, n, BoolType)
	checkPanics(t, n)
}

func TestNewArray(t *testing.T) {
	n := NewArray(nil)

	checkType(t, n, ArrayType)
	checkPanics(t, n)
}

func TestNewObject(t *testing.T) {
	n := NewObject(nil)

	checkType(t, n, ObjectType)
	checkPanics(t, n)
}

func TestNewMember(t *testing.T) {
	_ = NewMember("", "", nil)
}

func TestBool(t *testing.T) {

	expect := func(b Value, expect bool) {
		value := b.Bool()

		if value != expect {
			t.Errorf("Expected %v but go %v", expect, value)
		}
	}

	for _, v := range []bool{true, false} {
		b := NewBool(v)
		checkType(t, b, BoolType)
		expect(b, v)
	}

}

func TestNumber(t *testing.T) {

	// values := []b{
	// 	b(`0`),
	// 	b(`-1`),
	// 	b(`-2`),
	// 	b(`1.5`),
	// 	b(`-1.5`),
	// 	b(`1e1`),
	// 	b(`1e+1`),
	// 	b(`1e-1`),
	// }
}

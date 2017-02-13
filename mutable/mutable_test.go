package mutable

import (
	"testing"

	"github.com/omeid/j"
	"github.com/omeid/j/internal/valuetest"
)

var Types = []j.Type{
	j.ObjectType,
	j.ArrayType,
	j.BoolType,
	j.NumberType,
	j.StringType,
	j.NullType,
}

// Just used as reference to keep all "New" functions.
var _ = map[j.Type]interface{}{
	j.ObjectType: NewObject,
	j.ArrayType:  NewArray,
	j.BoolType:   NewBool,
	j.NumberType: NewNumber,
	j.StringType: NewString,
	j.NullType:   NewNull,
}

func checkValueType(t *testing.T, v Value, typ j.Type) {
	ttyp := v.Value().Type()
	if ttyp != typ {
		t.Errorf("Expected Type %v got %v", typ, ttyp)
	}
}

type valueMethod func(j.Value)

func TestNewNull(t *testing.T) {
	n := NewNull()

	checkValueType(t, n, j.NullType)
	valuetest.CheckPanics(t, n.Value())
}

func TestNewString(t *testing.T) {
	n := NewString("")

	checkValueType(t, n, j.StringType)
	valuetest.CheckPanics(t, n.Value())
}

func TestNewBool(t *testing.T) {
	n := NewBool(false)

	checkValueType(t, n, j.BoolType)
	valuetest.CheckPanics(t, n.Value())
}

func TestNewArray(t *testing.T) {
	n := NewArray()

	checkValueType(t, n, j.ArrayType)
	valuetest.CheckPanics(t, n.Value())
}

func TestNewObject(t *testing.T) {
	n := NewObject()

	checkValueType(t, n, j.ObjectType)
	valuetest.CheckPanics(t, n.Value())
}

func TestNewMember(t *testing.T) {
	_ = NewMember("", "", nil)
}

func TestBool(t *testing.T) {

	expect := func(b Bool, expect bool) {
		value := b.Value().Bool()

		if value != expect {
			t.Errorf("Expected %v but go %v", expect, value)
		}
	}

	for _, v := range []bool{true, false} {
		b := NewBool(v)
		checkValueType(t, b, j.BoolType)
		expect(b, v)

		b.Set(!v)
		expect(b, !v)
	}

}

type b []byte

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

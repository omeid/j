package valuetest

import (
	"testing"

	"github.com/omeid/j"
)

// CheckPanics makes sure the Value actually
// does panics when incorrect methods are called on.
func CheckPanics(t *testing.T, v j.Value) {

	typ := v.Type()

	checkMethod := func(mustPanic bool, method func(j.Value), name string) {

		defer func() {
			r := recover()

			if mustPanic && r == nil {
				t.Errorf("Expected panic on calling %s on %v type", name, typ)
			} else if !mustPanic && r != nil {
				t.Errorf("Unexpected panic on calling %s on %v", name, typ)
			}
		}()

		method(v)
	}

	type methodCheck struct {
		mustPanic bool
		method    func(j.Value)
		name      string
	}

	for _, mc := range []methodCheck{
		methodCheck{
			mustPanic: typ != j.ObjectType,
			method:    func(v j.Value) { v.Members() },
			name:      "Members",
		},
		methodCheck{
			mustPanic: typ != j.ArrayType,
			method:    func(v j.Value) { v.Values() },
			name:      "Values",
		},
		methodCheck{
			mustPanic: typ != j.BoolType,
			method:    func(v j.Value) { v.Bool() },
			name:      "Bool",
		},
		methodCheck{
			mustPanic: typ != j.NumberType,
			method:    func(v j.Value) { v.Int64() },
			name:      "Int64",
		},
		methodCheck{
			mustPanic: typ != j.NumberType,
			method:    func(v j.Value) { v.Float64() },
			name:      "Float64",
		},
		methodCheck{
			mustPanic: typ != j.NumberType,
			method:    func(v j.Value) { v.Raw() },
			name:      "Raw",
		},
		methodCheck{
			mustPanic: typ != j.StringType,
			method:    func(v j.Value) { _ = v.String() },
			name:      "String",
		},
	} {
		checkMethod(mc.mustPanic, mc.method, mc.name)
	}

}

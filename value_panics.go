package j

import (
	"testing"
)

// checkPanics makes sure the Value actually
// does panics when incorrect methods are called on.
func checkPanics(t *testing.T, v Value) {

	typ := v.Type()

	checkMethod := func(mustPanic bool, method func(Value), name string) {

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
		method    func(Value)
		name      string
	}

	for _, mc := range []methodCheck{
		methodCheck{
			mustPanic: typ != ObjectType,
			method:    func(v Value) { v.Members() },
			name:      "Members",
		},
		methodCheck{
			mustPanic: typ != ArrayType,
			method:    func(v Value) { v.Values() },
			name:      "Values",
		},
		methodCheck{
			mustPanic: typ != BoolType,
			method:    func(v Value) { v.Bool() },
			name:      "Bool",
		},
		methodCheck{
			mustPanic: typ != NumberType,
			method:    func(v Value) { v.Int64() },
			name:      "Int64",
		},
		methodCheck{
			mustPanic: typ != NumberType,
			method:    func(v Value) { v.Float64() },
			name:      "Float64",
		},
		methodCheck{
			mustPanic: typ != NumberType,
			method:    func(v Value) { v.Raw() },
			name:      "Raw",
		},
		methodCheck{
			mustPanic: typ != StringType,
			method:    func(v Value) { _ = v.String() },
			name:      "String",
		},
	} {
		checkMethod(mc.mustPanic, mc.method, mc.name)
	}

}

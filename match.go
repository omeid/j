package j

import (
	"bytes"
	"fmt"
	"sort"
)

// Match deep compares the provided values.
func Match(a Value, b Value) bool {
	if a.Type() != b.Type() {
		return false
	}

	// We know the types are equal, so a or b
	// doesn't matter.
	typ := a.Type()

	switch typ {
	case InvalidType:
		return false // Invalid Types are never the same.
	case ObjectType:
		return matchObject(a, b)
	case ArrayType:
		return matchArray(a, b)
	case BoolType:
		return a.Bool() == b.Bool()
	case NumberType:
		return matchNumber(a, b)
	case StringType:
		return a.String() == b.String()
	case NullType:
		return true // Null types are always the same.
	default:
		panic("Unexpected type.")
	}
}

// SortableMembers implements a
type sortableByNameMembers []Member

func (sm sortableByNameMembers) Len() int           { return len(sm) }
func (sm sortableByNameMembers) Swap(i, j int)      { sm[i], sm[j] = sm[j], sm[i] }
func (sm sortableByNameMembers) Less(i, j int) bool { return sm[i].Name() < sm[j].Name() }

func matchObject(a Value, b Value) bool {
	ams := a.Members()
	bms := b.Members()
	if len(ams) != len(bms) {
		return false
	}
	// TOOD: Compare the values.
	sort.Sort(sortableByNameMembers(ams))
	sort.Sort(sortableByNameMembers(bms))

	for i, v := range ams {
		if !matchMember(v, bms[i]) {
			fmt.Printf("::: a: %v %v, b:%v %v\n", v.Name(), v.Value().Type(), bms[i].Name(), bms[i].Value().Type())
			fmt.Printf("::: Miss match %v %v\n", v.Name(), bms[i].Name())
			return false
		}
	}

	return true
}

func matchArray(a Value, b Value) bool {
	avs := a.Values()
	bvs := b.Values()

	if len(avs) != len(bvs) {
		return false
	}

	for i, v := range avs {
		if !Match(v, bvs[i]) {
			return false
		}
	}

	return true
}

func matchNumber(a Value, b Value) bool {

	ar := a.Raw()
	br := b.Raw()

	return bytes.Compare(ar, br) == 0
}

func matchMember(a Member, b Member) bool {
	if a.Name() != b.Name() {
		return false
	}

	return Match(a.Value(), b.Value())
}

package decoder

import (
	"bytes"
	"sort"

	"github.com/omeid/j"
)

// Match deep compares the provided values.
func Match(a j.Value, b j.Value) bool {
	if a.Type() != b.Type() {
		return false
	}

	// We know the types are equal, so a or b
	// doesn't matter.
	typ := a.Type()

	switch typ {
	case j.InvalidType:
		return false // Invalid Types are never the same.
	case j.ObjectType:
		return matchObject(a, b)
	case j.ArrayType:
		return matchArray(a, b)
	case j.BoolType:
		return a.Bool() == b.Bool()
	case j.NumberType:
		return matchNumber(a, b)
	case j.StringType:
		return a.String() == b.String()
	case j.NullType:
		return true // Null types are always the same.
	default:
		panic("Unexpected type.")
	}
}

// SortableMembers implements a
type sortableByNameMembers []j.Member

func (sm sortableByNameMembers) Len() int           { return len(sm) }
func (sm sortableByNameMembers) Swap(i, j int)      { sm[i], sm[j] = sm[j], sm[i] }
func (sm sortableByNameMembers) Less(i, j int) bool { return sm[i].Name() < sm[j].Name() }

func matchObject(a j.Value, b j.Value) bool {
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
			return false
		}
	}

	return true
}

func matchArray(a j.Value, b j.Value) bool {
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

func matchNumber(a j.Value, b j.Value) bool {

	ar := a.Number().Raw()
	br := b.Number().Raw()

	return bytes.Compare(ar, br) == 0
}

func matchMember(a j.Member, b j.Member) bool {
	if a.Name() != b.Name() {
		return false
	}

	return Match(a.Value(), b.Value())
}

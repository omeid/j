package j

// NewObject creates a new mutable object
func NewObject(members []Member) Value {
	index := make(map[string]int, len(members))
	for i, m := range members {
		index[m.Name()] = i
	}
	return &object{value: baseValue, members: members, index: index}
}

type object struct {
	*value

	members []Member
	index   map[string]int
}

func (o object) Type() Type {
	return ObjectType
}

func (o *object) Members() []Member {
	return o.members
}

func (o *object) Member(name string) Value {
	i, ok := o.index[name]
	if !ok {
		return nil
	}

	return o.members[i].Value()
}

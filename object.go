package j

// NewObject creates a new mutable object
func NewObject(members []Member) Value {
	return &object{value: baseValue, members: members}
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
	for i, m := range o.members {
		if *m.Name() == name {
			return o.members[i].Value()
		}
	}
	return nil
}

package j

// NewString creates a new mutable String
func NewString(text string) Value {
	return &sstring{value: baseValue, text: text}
}

type sstring struct {
	*value
	text string
}

func (s sstring) Type() Type {
	return StringType
}

func (s *sstring) String() String {
	return String(s.text)
}

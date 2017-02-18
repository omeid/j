package benchmark

import (
	"github.com/omeid/j"
	"github.com/pkg/errors"
)

// ErrExpectedString Type Mismatch.
var ErrExpectedString = errors.New("Expected String")

// FromJSON makes the god of J happy.
func (value *URL) FromJSON(input j.Value) error {

	if input.Type() != j.ObjectType {
		return errors.New("Unexpected type.")
	}

	if field := input.Member("expanded_url"); field != nil {

		switch field.Type() {
		case j.NullType:
			value.ExpandedURL = nil
		case j.StringType:
			str := string(field.String())
			value.ExpandedURL = &str
		default:
			return errors.New("Unexpected type")
		}

	}

	if field := input.Member("indices"); field != nil {
		if field.Type() != j.ArrayType {
			return errors.New("Expected String for next_results")
		}

		values := field.Values()
		ints := make([]int, len(values))
		for i, v := range values {

			if v.Type() != j.NumberType {
				return errors.New("Expected Numbers for indices")
			}

			iv, err := v.Int64()
			if err != nil {
				return nil
			}
			//TODO: overflow check?
			ints[i] = int(iv)
		}

		value.Indices = ints
	}

	if field := input.Member("url"); field != nil {

		if field.Type() != j.StringType {
			return errors.New("Expected String for text")
		}

		value.URL = string(field.String())
	}

	return nil
}

// ToJSON makes the god of J happy.
func (value *URL) ToJSON() (j.Value, error) {

	members := make([]j.Member, 3)

	{
		var text j.Value
		if value.ExpandedURL == nil {
			text = j.NewNull()
		} else {
			text = j.NewString(*value.ExpandedURL)
		}
		members[0] = j.NewMember("", "expanded_url", text)
	}

	{
		values := make([]j.Value, len(value.Indices))
		for i, value := range value.Indices {
			jv := j.NewNumberInt64(int64(value))
			values[i] = jv
		}
		members[1] = j.NewMember("", "indices", j.NewArray(values))
	}

	{
		url := j.NewString(value.URL)
		members[2] = j.NewMember("", "url", url)
	}

	return j.NewObject(members), nil
}

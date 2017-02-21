package benchmark

import (
	"github.com/omeid/j"
	"github.com/pkg/errors"
)

func (value *UserEntityDescription) FromJSON(input j.Value) error {

	if input.Type() != j.ObjectType {
		return errors.New("Unexpected type.")
	}

	if field := input.Member("urls"); field != nil {
		if field.Type() != j.ArrayType {
			return errors.New("Expected Array for hasthags")
		}

		values := field.Values()
		strs := make([]*string, len(values))
		for i, v := range values {

			if v.Type() != j.NumberType {
				return errors.New("Expected Numbers for indices")
			}

			strs[i] = v.String()
		}

		value.Urls = strs
	}

	return nil
}

func (value *UserEntityDescription) ToJSON() (j.Value, error) {

	values := make([]j.Value, len(value.Urls))
	for i, value := range value.Urls {
		jv := j.NewString(*value)
		values[i] = jv
	}

	return j.NewObject([]j.Member{j.NewMember("", "urls", j.NewArray(values))}), nil
}

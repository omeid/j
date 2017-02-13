package codec

import (
	"github.com/omeid/j"
	"github.com/omeid/j/mutable"
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
		for _, v := range values {

			if v.Type() != j.NumberType {
				return errors.New("Expected Numbers for indices")
			}

			str := string(v.String())
			strs = append(strs, &str)
		}

		value.Urls = strs
	}

	return nil
}

func (value *UserEntityDescription) ToJSON() (j.Value, error) {

	obj := mutable.NewObject()

	{
		urls := mutable.NewArray()
		for _, value := range value.Urls {
			jv := mutable.NewString(*value)
			urls.Add(jv.Value())
		}
		obj.Add(mutable.NewMember("", "urls", urls.Value()))
	}

	return obj.Value(), nil
}

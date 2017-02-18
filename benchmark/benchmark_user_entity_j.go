package benchmark

import (
	"github.com/omeid/j"
	"github.com/pkg/errors"
)

func (value *UserEntityURL) FromJSON(input j.Value) error {

	if input.Type() != j.ObjectType {
		return errors.New("Unexpected type.")
	}

	if field := input.Member("urls"); field != nil {
		if field.Type() != j.ArrayType {
			return errors.New("Expected Array for hasthags")
		}

		values := field.Values()
		urls := make([]URL, len(values))
		for i, v := range values {

			url := URL{}

			err := url.FromJSON(v)
			if err != nil {
				return err
			}

			urls[i] = url
		}

		value.Urls = urls
	}

	return nil
}

func (value *UserEntityURL) ToJSON() (j.Value, error) {

	values := make([]j.Value, len(value.Urls))
	for i, value := range value.Urls {
		jv, err := value.ToJSON()
		if err != nil {
			return nil, err
		}
		values[i] = jv
	}
	return j.NewObject([]j.Member{j.NewMember("", "urls", j.NewArray(values))}), nil
}

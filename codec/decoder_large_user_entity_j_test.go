package codec

import (
	"github.com/omeid/j"
	"github.com/omeid/j/mutable"
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
		for _, v := range values {

			url := URL{}

			err := url.FromJSON(v)
			if err != nil {
				return err
			}

			urls = append(urls, url)
		}

		value.Urls = urls
	}

	return nil
}

func (value *UserEntityURL) ToJSON() (j.Value, error) {

	obj := mutable.NewObject()

	{
		urls := mutable.NewArray()
		for _, value := range value.Urls {
			jv, err := value.ToJSON()
			if err != nil {
				return nil, err
			}
			urls.Add(jv)
		}
		obj.Add(mutable.NewMember("", "urls", urls.Value()))
	}

	return obj.Value(), nil
}

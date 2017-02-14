package benchmark

import (
	"github.com/omeid/j"
	"github.com/omeid/j/mutable"
	"github.com/pkg/errors"
)

func (value *UserEntities) FromJSON(input j.Value) error {

	if input.Type() != j.ObjectType {
		return errors.New("Unexpected type.")
	}

	if field := input.Member("description"); field != nil {
		if field.Type() != j.ObjectType {
			return errors.New("Expected Object for description")
		}

		Description := UserEntityDescription{}
		err := Description.FromJSON(field)
		if err != nil {
			return err
		}
		value.Description = Description
	}

	if field := input.Member("url"); field != nil {
		if field.Type() != j.ObjectType {
			return errors.New("Expected Object for url")
		}

		URL := UserEntityURL{}
		err := URL.FromJSON(field)
		if err != nil {
			return err
		}
		value.URL = URL
	}

	return nil
}

func (value *UserEntities) ToJSON() (j.Value, error) {

	obj := mutable.NewObject()

	{
		description, err := value.Description.ToJSON()
		if err != nil {
			return nil, err
		}
		obj.Add(mutable.NewMember("", "hashtags", description))
	}

	{
		url, err := value.URL.ToJSON()
		if err != nil {
			return nil, err
		}
		obj.Add(mutable.NewMember("", "hashtags", url))
	}

	return obj.Value(), nil
}

package benchmark

import (
	"github.com/omeid/j"
	"github.com/pkg/errors"
)

// FromJSON is to make the lord of J happy.
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

// ToJSON is to make the lord of J happy.
func (value *UserEntities) ToJSON() (j.Value, error) {

	members := make([]j.Member, 2)

	{
		description, err := value.Description.ToJSON()
		if err != nil {
			return nil, err
		}
		members[0] = j.NewMember("", "description", description)
	}

	{
		url, err := value.URL.ToJSON()
		if err != nil {
			return nil, err
		}
		members[1] = j.NewMember("", "url", url)
	}

	return j.NewObject(members), nil
}

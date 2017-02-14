package benchmark

import (
	"github.com/omeid/j"
	"github.com/omeid/j/mutable"
	"github.com/pkg/errors"
)

func (value *LargeStruct) FromJSON(input j.Value) error {

	if input.Type() != j.ObjectType {
		return errors.New("Unexpected type.")
	}

	if field := input.Member("search_metadata"); field != nil {
		SearchMetadata := SearchMetadata{}

		err := SearchMetadata.FromJSON(field)
		if err != nil {
			return nil
		}

		value.SearchMetadata = SearchMetadata
	}

	if field := input.Member("statuses"); field != nil {
		if field.Type() != j.ArrayType {
			return errors.New("Expected Array for hasthags")
		}

		var Statuses []Status
		for _, v := range field.Values() {
			s := Status{}
			err := s.FromJSON(v)
			if err != nil {
				return err
			}
			Statuses = append(Statuses, s)
		}

		value.Statuses = Statuses
	}

	return nil
}

func (value *LargeStruct) ToJSON() (j.Value, error) {

	obj := mutable.NewObject()

	{
		SearchMetadata, err := value.SearchMetadata.ToJSON()
		if err != nil {
			return nil, err
		}

		obj.Add(mutable.NewMember("", "search_metadata", SearchMetadata))
	}

	{
		statuses := mutable.NewArray()
		for _, value := range value.Statuses {
			jv, err := value.ToJSON()
			if err != nil {
				return nil, err
			}
			statuses.Add(jv)
		}
		obj.Add(mutable.NewMember("", "statuses", statuses.Value()))
	}

	return obj.Value(), nil
}

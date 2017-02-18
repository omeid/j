package benchmark

import (
	"github.com/omeid/j"
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

		fields := field.Values()
		Statuses := make([]Status, len(fields))
		for i, v := range fields {
			s := Status{}
			err := s.FromJSON(v)
			if err != nil {
				return err
			}
			Statuses[i] = s
		}

		value.Statuses = Statuses
	}

	return nil
}

func (value *LargeStruct) ToJSON() (j.Value, error) {

	members := make([]j.Member, 2)

	{
		SearchMetadata, err := value.SearchMetadata.ToJSON()
		if err != nil {
			return nil, err
		}

		members[0] = j.NewMember("", "search_metadata", SearchMetadata)
	}

	{
		values := make([]j.Value, len(value.Statuses))
		for i, value := range value.Statuses {
			jv, err := value.ToJSON()
			if err != nil {
				return nil, err
			}
			values[i] = jv
		}
		members[1] = j.NewMember("", "statuses", j.NewArray(values))
	}

	return j.NewObject(members), nil
}

package benchmark

import (
	"github.com/omeid/j"
	"github.com/pkg/errors"
)

func (value *StatusMetadata) FromJSON(input j.Value) error {

	if input.Type() != j.ObjectType {
		return errors.New("Unexpected type.")
	}

	if field := input.Member("iso_language_code"); field != nil {

		if field.Type() != j.StringType {
			return errors.New("Expected String for max_id_str")
		}

		value.IsoLanguageCode = string(field.String())
	}

	if field := input.Member("result_type"); field != nil {

		if field.Type() != j.StringType {
			return errors.New("Expected String for next_results")
		}

		value.ResultType = string(field.String())
	}

	return nil
}

func (value *StatusMetadata) ToJSON() (j.Value, error) {

	members := make([]j.Member, 2)

	{
		IsoLanguageCode := j.NewString(value.IsoLanguageCode)
		members[0] = j.NewMember("", "iso_language_code", IsoLanguageCode)
	}

	{
		ResultType := j.NewString(value.ResultType)
		members[1] = j.NewMember("", "result_type", ResultType)
	}

	return j.NewObject(members), nil
}

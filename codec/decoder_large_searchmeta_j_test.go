package codec

import (
	"github.com/omeid/j"
	"github.com/omeid/j/mutable"
	"github.com/pkg/errors"
)

func (value *SearchMetadata) FromJSON(input j.Value) error {

	if input.Type() != j.ObjectType {
		return errors.New("Unexpected type.")
	}

	if field := input.Member("completed_in"); field != nil {
		var err error

		if field.Type() != j.NumberType {
			return errors.New("Expected Number for completed_in")
		}

		value.CompletedIn, err = field.Float64()
		if err != nil {
			return errors.WithMessage(err, "CompletedIn field")
		}
	}

	if field := input.Member("count"); field != nil {

		if field.Type() != j.NumberType {
			return errors.New("Expected Number for count")
		}

		i, err := field.Int64()
		if err != nil {
			return errors.WithMessage(err, "Count field")
		}
		value.Count = int(i)
	}

	if field := input.Member("max_id"); field != nil {

		if field.Type() != j.NumberType {
			return errors.New("Expected Number for max_id")
		}

		i, err := field.Int64()
		if err != nil {
			return errors.WithMessage(err, "Max field")
		}
		value.MaxID = int(i)
	}

	if field := input.Member("max_id_str"); field != nil {

		if field.Type() != j.StringType {
			return errors.New("Expected String for max_id_str")
		}

		value.MaxIDStr = string(field.String())
	}

	if field := input.Member("next_results"); field != nil {

		if field.Type() != j.StringType {
			return errors.New("Expected String for next_results")
		}

		value.NextResults = string(field.String())
	}

	if field := input.Member("query"); field != nil {

		if field.Type() != j.StringType {
			return errors.New("Expected String for next_results")
		}

		value.Query = string(field.String())
	}

	if field := input.Member("refresh_url"); field != nil {

		if field.Type() != j.StringType {
			return errors.New("Expected String for refresh_url")
		}

		value.RefreshURL = string(field.String())
	}

	if field := input.Member("since_id"); field != nil {

		if field.Type() != j.NumberType {
			return errors.New("Expected Number for since_id")
		}

		i, err := field.Int64()
		if err != nil {
			return errors.WithMessage(err, "Max ID field")
		}
		value.SinceID = int(i)
	}

	if field := input.Member("since_id_str"); field != nil {

		if field.Type() != j.StringType {
			return errors.New("Expected String for since_id_str")
		}

		value.SinceIDStr = string(field.String())
	}

	return nil
}

func (value *SearchMetadata) ToJSON() (j.Value, error) {

	obj := mutable.NewObject()

	obj.Add(mutable.NewMember("", "completed_in", mutable.NewNumberFloat64(value.CompletedIn).Value()))
	obj.Add(mutable.NewMember("", "count", mutable.NewNumberInt64(int64(value.Count)).Value()))
	obj.Add(mutable.NewMember("", "max_id", mutable.NewNumberInt64(int64(value.MaxID)).Value()))
	obj.Add(mutable.NewMember("", "max_id_str", mutable.NewString(value.MaxIDStr).Value()))
	obj.Add(mutable.NewMember("", "next_results", mutable.NewString(value.NextResults).Value()))
	obj.Add(mutable.NewMember("", "query", mutable.NewString(value.Query).Value()))
	obj.Add(mutable.NewMember("", "refresh_url", mutable.NewString(value.RefreshURL).Value()))
	obj.Add(mutable.NewMember("", "since_id", mutable.NewNumberInt64(int64(value.SinceID)).Value()))
	obj.Add(mutable.NewMember("", "since_id_str", mutable.NewString(value.SinceIDStr).Value()))

	return obj.Value(), nil
}

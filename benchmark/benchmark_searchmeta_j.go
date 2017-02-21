package benchmark

import (
	"github.com/omeid/j"
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

		value.MaxIDStr = string(*field.String())
	}

	if field := input.Member("next_results"); field != nil {

		if field.Type() != j.StringType {
			return errors.New("Expected String for next_results")
		}

		value.NextResults = string(*field.String())
	}

	if field := input.Member("query"); field != nil {

		if field.Type() != j.StringType {
			return errors.New("Expected String for next_results")
		}

		value.Query = string(*field.String())
	}

	if field := input.Member("refresh_url"); field != nil {

		if field.Type() != j.StringType {
			return errors.New("Expected String for refresh_url")
		}

		value.RefreshURL = string(*field.String())
	} else {
		panic("refresh url")
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

		value.SinceIDStr = string(*field.String())
	}

	return nil
}

func (value *SearchMetadata) ToJSON() (j.Value, error) {

	members := make([]j.Member, 9)

	members[0] = j.NewMember("", "completed_in", j.NewNumberFloat64(value.CompletedIn))
	members[1] = j.NewMember("", "count", j.NewNumberInt64(int64(value.Count)))
	members[2] = j.NewMember("", "max_id", j.NewNumberInt64(int64(value.MaxID)))
	members[3] = j.NewMember("", "max_id_str", j.NewString(value.MaxIDStr))
	members[4] = j.NewMember("", "next_results", j.NewString(value.NextResults))
	members[5] = j.NewMember("", "query", j.NewString(value.Query))
	members[6] = j.NewMember("", "refresh_url", j.NewString(value.RefreshURL))
	members[7] = j.NewMember("", "since_id", j.NewNumberInt64(int64(value.SinceID)))
	members[8] = j.NewMember("", "since_id_str", j.NewString(value.SinceIDStr))

	return j.NewObject(members), nil
}

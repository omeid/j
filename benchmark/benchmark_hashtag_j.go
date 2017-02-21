package benchmark

import (
	"github.com/omeid/j"
	"github.com/pkg/errors"
)

func (value *Hashtag) FromJSON(input j.Value) error {

	if input.Type() != j.ObjectType {
		return errors.New("Unexpected type.")
	}

	if field := input.Member("text"); field != nil {

		if field.Type() != j.StringType {
			return errors.New("Expected String for text")
		}

		value.Text = string(*field.String())
	}

	if field := input.Member("indices"); field != nil {

		if field.Type() != j.ArrayType {
			return errors.New("Expected String for next_results")
		}

		values := field.Values()

		ints := make([]int, len(values))
		for i, v := range values {

			if v.Type() != j.NumberType {
				return errors.New("Expected Numbers for indices")
			}

			iv, err := v.Int64()
			if err != nil {
				return err
			}

			ints[i] = int(iv)
		}

		value.Indices = ints
	}

	return nil
}

func (value *Hashtag) ToJSON() (j.Value, error) {

	members := make([]j.Member, 2)

	{
		text := j.NewString(value.Text)
		members[0] = j.NewMember("", "text", text)
	}

	{
		values := make([]j.Value, len(value.Indices))
		for i, index := range value.Indices {
			values[i] = j.NewNumberInt64(int64(index))
		}
		members[1] = j.NewMember("", "indices", j.NewArray(values))
	}

	return j.NewObject(members), nil
}

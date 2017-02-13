package codec

import (
	"github.com/omeid/j"
	"github.com/omeid/j/mutable"
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

		value.Text = string(field.String())
	}

	if field := input.Member("indices"); field != nil {

		if field.Type() != j.ArrayType {
			return errors.New("Expected String for next_results")
		}

		values := field.Values()

		ints := make([]int, len(values))
		for _, v := range values {

			if v.Type() != j.NumberType {
				return errors.New("Expected Numbers for indices")
			}

			i, err := v.Int64()
			if err != nil {
				return err
			}

			ints = append(ints, int(i))
		}

		value.Indices = ints
	}

	return nil
}

func (value *Hashtag) ToJSON() (j.Value, error) {

	obj := mutable.NewObject()

	{
		text := mutable.NewString(value.Text)
		obj.Add(mutable.NewMember("", "text", text.Value()))
	}

	{
		indices := mutable.NewArray()

		for _, index := range value.Indices {
			indices.Add(mutable.NewNumberInt64(int64(index)).Value())
		}
		obj.Add(mutable.NewMember("", "indices", indices.Value()))
	}

	return obj.Value(), nil
}

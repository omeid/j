package codec

import (
	"github.com/omeid/j"
	"github.com/omeid/j/mutable"
	"github.com/pkg/errors"
)

func (value *URL) FromJSON(input j.Value) error {

	if input.Type() != j.ObjectType {
		return errors.New("Unexpected type.")
	}

	if field := input.Member("expanded_url"); field != nil {

		if field.Type() != j.StringType {
			return errors.New("Expected String for text")
		}

		str := string(field.String())
		value.ExpandedURL = &str
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
				return nil
			}
			//TODO: overflow check?
			ints = append(ints, int(i))
		}

		value.Indices = ints
	}

	if field := input.Member("url"); field != nil {

		if field.Type() != j.StringType {
			return errors.New("Expected String for text")
		}

		value.URL = string(field.String())
	}

	return nil
}

func (value *URL) ToJSON() (j.Value, error) {

	obj := mutable.NewObject()

	{
		var text j.Value
		if value.ExpandedURL == nil {
			text = mutable.NewNull().Value()
		} else {
			text = mutable.NewString(*value.ExpandedURL).Value()
		}
		obj.Add(mutable.NewMember("", "text", text))
	}

	{
		urls := mutable.NewArray()
		for _, value := range value.Indices {
			jv := mutable.NewNumberInt64(int64(value))
			urls.Add(jv.Value())
		}
		obj.Add(mutable.NewMember("", "indices", urls.Value()))
	}

	{
		url := mutable.NewString(value.URL).Value()
		obj.Add(mutable.NewMember("", "text", url))
	}

	return obj.Value(), nil
}

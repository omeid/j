package codec

import (
	"github.com/omeid/j"
	"github.com/omeid/j/mutable"
	"github.com/pkg/errors"
)

func (value *Entities) FromJSON(input j.Value) error {

	if input.Type() != j.ObjectType {
		return errors.New("Unexpected type.")
	}

	if field := input.Member("hashtags"); field != nil {
		if field.Type() != j.ArrayType {
			return errors.New("Expected Array for hasthags")
		}

		var hashtags []Hashtag
		for _, v := range field.Values() {
			h := Hashtag{}
			err := h.FromJSON(v)
			if err != nil {
				return err
			}
			hashtags = append(hashtags, h)
		}

		value.Hashtags = hashtags
	}

	if field := input.Member("urls"); field != nil {
		if field.Type() != j.ArrayType {
			return errors.New("Expected String for next_results")
		}

		values := field.Values()
		strs := make([]*string, len(values))
		for _, v := range values {

			if v.Type() != j.NumberType {
				return errors.New("Expected Numbers for indices")
			}

			str := string(v.String())
			strs = append(strs, &str)
		}

		value.Urls = strs
	}

	if field := input.Member("user_mentions"); field != nil {
		if field.Type() != j.ArrayType {
			return errors.New("Expected String for next_results")
		}

		values := field.Values()
		strs := make([]*string, len(values))
		for _, v := range values {

			if v.Type() != j.NumberType {
				return errors.New("Expected Numbers for indices")
			}

			str := string(v.String())
			strs = append(strs, &str)
		}

		value.UserMentions = strs
	}

	return nil
}

func (value *Entities) ToJSON() (j.Value, error) {

	obj := mutable.NewObject()

	{
		hashtags := mutable.NewArray()
		for _, value := range value.Hashtags {
			jv, err := value.ToJSON()
			if err != nil {
				return nil, err
			}
			hashtags.Add(jv)
		}
		obj.Add(mutable.NewMember("", "hashtags", hashtags.Value()))
	}

	{
		urls := mutable.NewArray()
		for _, value := range value.Urls {
			jv := mutable.NewString(*value)
			urls.Add(jv.Value())
		}
		obj.Add(mutable.NewMember("", "urls", urls.Value()))
	}

	{
		usermentions := mutable.NewArray()
		for _, value := range value.UserMentions {
			jv := mutable.NewString(*value)
			usermentions.Add(jv.Value())
		}
		obj.Add(mutable.NewMember("", "urls", usermentions.Value()))
	}

	return obj.Value(), nil
}

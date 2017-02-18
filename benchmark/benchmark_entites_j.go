package benchmark

import (
	"github.com/omeid/j"
	"github.com/pkg/errors"
)

// FromJSON makes the lord of j happy.
func (value *Entities) FromJSON(input j.Value) error {

	if input.Type() != j.ObjectType {
		return errors.New("Unexpected type.")
	}

	if field := input.Member("hashtags"); field != nil {
		if field.Type() != j.ArrayType {
			return errors.New("Expected Array for hasthags")
		}

		fields := field.Values()
		hashtags := make([]Hashtag, len(fields))

		for i, v := range fields {
			h := Hashtag{}
			err := h.FromJSON(v)
			if err != nil {
				return err
			}
			hashtags[i] = h
		}

		value.Hashtags = hashtags
	}

	if field := input.Member("urls"); field != nil {
		if field.Type() != j.ArrayType {
			return errors.New("Expected String for next_results")
		}

		values := field.Values()
		strs := make([]*string, len(values))
		for i, v := range values {

			if v.Type() != j.NumberType {
				return errors.New("Expected Numbers for indices")
			}

			str := string(v.String())
			strs[i] = &str
		}

		value.Urls = strs
	}

	if field := input.Member("user_mentions"); field != nil {
		if field.Type() != j.ArrayType {
			return errors.New("Expected Array for user_mentions")
		}

		values := field.Values()
		strs := make([]*string, len(values))
		for i, v := range values {

			if v.Type() != j.NumberType {
				return errors.New("Expected Numbers for indices")
			}

			str := string(v.String())
			strs[i] = &str
		}

		value.UserMentions = strs
	}

	return nil
}

// ToJSON makes the lords of J happy.
func (value *Entities) ToJSON() (j.Value, error) {

	members := make([]j.Member, 3)
	{
		values := make([]j.Value, len(value.Hashtags))
		for i, value := range value.Hashtags {
			jv, err := value.ToJSON()
			if err != nil {
				return nil, err
			}
			values[i] = jv
		}
		members[0] = j.NewMember("", "hashtags", j.NewArray(values))
	}

	{
		values := make([]j.Value, len(value.Urls))
		for i, value := range value.Urls {
			jv := j.NewString(*value)
			values[i] = jv
		}
		members[1] = j.NewMember("", "urls", j.NewArray(values))
	}

	{
		values := make([]j.Value, len(value.UserMentions))
		for i, value := range value.UserMentions {
			jv := j.NewString(*value)
			values[i] = jv
		}

		members[2] = j.NewMember("", "user_mentions", j.NewArray(values))
	}

	return j.NewObject(members), nil
}

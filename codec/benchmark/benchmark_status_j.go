package benchmark

import (
	"github.com/omeid/j"
	"github.com/omeid/j/mutable"
	"github.com/pkg/errors"
)

func (value *Status) FromJSON(input j.Value) error {

	if field := input.Member("contributors"); field != nil {
		switch field.Type() {
		case j.NullType:
			value.Contributors = nil
		case j.StringType:
			str := string(field.String())
			value.Contributors = &str
		default:
			return errors.New("Unexpected type")
		}
	}

	if field := input.Member("coordinates"); field != nil {
		switch field.Type() {
		case j.NullType:
			value.Coordinates = nil
		case j.StringType:
			str := string(field.String())
			value.Coordinates = &str
		default:
			return errors.New("Unexpected type")
		}
	}

	if field := input.Member("created_at"); field != nil {
		if field.Type() != j.StringType {
			return errors.New("Unexpected type")
		}

		value.CreatedAt = string(field.String())
	}

	if field := input.Member("entities"); field != nil {
		if field.Type() != j.ObjectType {
			return errors.New("Unexpected type")
		}

		Entities := Entities{}

		if err := Entities.FromJSON(field); err != nil {
			return err
		}

		value.Entities = Entities

	}

	if field := input.Member("favorited"); field != nil {
		if field.Type() != j.BoolType {
			return errors.New("Unexpected type.")
		}

		value.Favorited = field.Bool()
	}

	if field := input.Member("geo"); field != nil {
		switch field.Type() {
		case j.NullType:
			value.Geo = nil
		case j.StringType:
			str := string(field.String())
			value.Geo = &str
		default:
			return errors.New("Unexpected type")
		}
	}

	if field := input.Member("id"); field != nil {
		if field.Type() != j.NumberType {
			return errors.New("Unexpected type")
		}

		n, err := field.Int64()
		if err != nil {
			return err
		}

		value.ID = n
	}

	if field := input.Member("id_str"); field != nil {

		if field.Type() != j.StringType {
			return errors.New("Unexpected type")
		}

		value.IDStr = string(field.String())
	}

	if field := input.Member("in_reply_to_screen_name"); field != nil {

		switch field.Type() {
		case j.NullType:
			value.InReplyToScreenName = nil
		case j.StringType:
			str := string(field.String())
			value.InReplyToScreenName = &str
		default:
			return errors.New("Unexpected type")
		}
	}

	if field := input.Member("in_reply_to_status_id"); field != nil {

		switch field.Type() {
		case j.NullType:
			value.InReplyToStatusID = nil
		case j.StringType:
			str := string(field.String())
			value.InReplyToStatusID = &str
		default:
			return errors.New("Unexpected type")
		}
	}

	if field := input.Member("in_reply_to_status_id_str"); field != nil {

		switch field.Type() {
		case j.NullType:
			value.InReplyToStatusIDStr = nil
		case j.StringType:
			str := string(field.String())
			value.InReplyToStatusIDStr = &str
		default:
			return errors.New("Unexpected type")
		}
	}

	if field := input.Member("in_reply_to_user_id"); field != nil {

		switch field.Type() {
		case j.NullType:
			value.InReplyToUserID = nil
		case j.StringType:
			str := string(field.String())
			value.InReplyToUserID = &str
		default:
			return errors.New("Unexpected type")
		}
	}

	if field := input.Member("in_reply_to_user_id_str"); field != nil {

		switch field.Type() {
		case j.NullType:
			value.InReplyToUserIDStr = nil
		case j.StringType:
			str := string(field.String())
			value.InReplyToUserIDStr = &str
		default:
			return errors.New("Unexpected type")
		}
	}

	if field := input.Member("metadata"); field != nil {
		if field.Type() != j.ObjectType {
			return errors.New("Unexpected type")
		}

		Metadata := StatusMetadata{}

		if err := Metadata.FromJSON(field); err != nil {
			return err
		}

		value.Metadata = Metadata
	}

	if field := input.Member("place"); field != nil {

		switch field.Type() {
		case j.NullType:
			value.Place = nil
		case j.StringType:
			str := string(field.String())
			value.Place = &str
		default:
			return errors.New("Unexpected type")
		}
	}

	if field := input.Member("retweet_count"); field != nil {

		if field.Type() != j.NumberType {
			return errors.New("Unexpected type")
		}

		n, err := field.Int64()
		if err != nil {
			return err
		}

		value.RetweetCount = int(n)
	}

	if field := input.Member("retweeted"); field != nil {

		if field.Type() != j.BoolType {
			return errors.New("Unexpected type.")
		}

		value.Retweeted = field.Bool()
	}

	if field := input.Member("source"); field != nil {

		if field.Type() != j.StringType {
			return errors.New("Unexpected type")
		}

		value.Source = string(field.String())
	}

	if field := input.Member("text"); field != nil {

		if field.Type() != j.StringType {
			return errors.New("Unexpected type")
		}

		value.Text = string(field.String())
	}

	if field := input.Member("truncated"); field != nil {

		if field.Type() != j.BoolType {
			return errors.New("Unexpected type.")
		}

		value.Truncated = field.Bool()
	}

	if field := input.Member("user"); field != nil {
		if field.Type() != j.ObjectType {
			return errors.New("Unexpected type")
		}

		User := User{}

		if err := User.FromJSON(field); err != nil {
			return err
		}

		value.User = User
	}

	return nil
}

func (value *Status) ToJSON() (j.Value, error) {

	obj := mutable.NewObject()

	//ResultType := mutable.NewString(value.ResultType)
	// obj.Add(mutable.NewMember("", "result_type", ResultType.Value(.Value())))

	{
		var Contributors j.Value
		if Contributors != nil {
			Contributors = mutable.NewString(*value.Contributors).Value()
		} else {
			Contributors = mutable.NewNull().Value()
		}
		obj.Add(mutable.NewMember("", "contributors", Contributors))

	}

	{
		var Coordinates j.Value
		if Coordinates != nil {
			Coordinates = mutable.NewString(*value.Coordinates).Value()
		} else {
			Coordinates = mutable.NewNull().Value()
		}
		obj.Add(mutable.NewMember("", "coordinates", Coordinates))

	}

	{
		CreatedAt := mutable.NewString(value.CreatedAt)
		obj.Add(mutable.NewMember("", "created_at", CreatedAt.Value()))

	}

	{
		Entities, err := value.Entities.ToJSON()
		if err != nil {
			return nil, err
		}
		obj.Add(mutable.NewMember("", "entities", Entities))
	}

	{
		Favorited := mutable.NewBool(value.Favorited)
		obj.Add(mutable.NewMember("", "favorited", Favorited.Value()))

	}

	{
		var Geo j.Value
		if Geo != nil {
			Geo = mutable.NewString(*value.Geo).Value()
		} else {
			Geo = mutable.NewNull().Value()
		}
		obj.Add(mutable.NewMember("", "geo", Geo))

	}

	{
		ID := mutable.NewNumberInt64(value.ID)
		obj.Add(mutable.NewMember("", "id", ID.Value()))

	}

	{
		IDStr := mutable.NewString(value.IDStr)
		obj.Add(mutable.NewMember("", "id_str", IDStr.Value()))

	}

	{
		var InReplyToScreenName j.Value
		if InReplyToScreenName != nil {
			InReplyToScreenName = mutable.NewString(*value.InReplyToScreenName).Value()
		} else {
			InReplyToScreenName = mutable.NewNull().Value()
		}
		obj.Add(mutable.NewMember("", "in_reply_to_screen_name", InReplyToScreenName))

	}

	{
		var InReplyToStatusID j.Value
		if InReplyToStatusID != nil {
			InReplyToStatusID = mutable.NewString(*value.InReplyToStatusID).Value()
		} else {
			InReplyToStatusID = mutable.NewNull().Value()
		}
		obj.Add(mutable.NewMember("", "in_reply_to_status_id", InReplyToStatusID))

	}

	{
		var InReplyToStatusIDStr j.Value
		if InReplyToStatusIDStr != nil {
			InReplyToStatusIDStr = mutable.NewString(*value.InReplyToStatusIDStr).Value()
		} else {
			InReplyToStatusIDStr = mutable.NewNull().Value()
		}
		obj.Add(mutable.NewMember("", "in_reply_to_status_id_str", InReplyToStatusIDStr))

	}

	{
		var InReplyToUserID j.Value
		if InReplyToUserID != nil {
			InReplyToUserID = mutable.NewString(*value.InReplyToUserID).Value()
		} else {
			InReplyToUserID = mutable.NewNull().Value()
		}
		obj.Add(mutable.NewMember("", "in_reply_to_user_id", InReplyToUserID))

	}

	{
		var InReplyToUserIDStr j.Value
		if InReplyToUserIDStr != nil {
			InReplyToUserIDStr = mutable.NewString(*value.InReplyToUserIDStr).Value()
		} else {
			InReplyToUserIDStr = mutable.NewNull().Value()
		}
		obj.Add(mutable.NewMember("", "in_reply_to_user_id_str", InReplyToUserIDStr))

	}

	{
		Metadata, err := value.Metadata.ToJSON()
		if err != nil {
			return nil, err
		}
		obj.Add(mutable.NewMember("", "metadata", Metadata))
	}

	{
		var Place j.Value
		if Place != nil {
			Place = mutable.NewString(*value.Place).Value()
		} else {
			Place = mutable.NewNull().Value()
		}
		obj.Add(mutable.NewMember("", "place", Place))

	}

	{
		RetweetCount := mutable.NewNumberInt64(int64(value.RetweetCount))
		obj.Add(mutable.NewMember("", "retweet_count", RetweetCount.Value()))

	}

	{
		Retweeted := mutable.NewBool(value.Retweeted)
		obj.Add(mutable.NewMember("", "retweeted", Retweeted.Value()))

	}

	{
		Source := mutable.NewString(value.Source)
		obj.Add(mutable.NewMember("", "source", Source.Value()))

	}

	{
		Text := mutable.NewString(value.Text)
		obj.Add(mutable.NewMember("", "text", Text.Value()))

	}

	{
		Truncated := mutable.NewBool(value.Truncated)
		obj.Add(mutable.NewMember("", "truncated", Truncated.Value()))

	}

	{
		User, err := value.User.ToJSON()
		if err != nil {
			return nil, err
		}
		obj.Add(mutable.NewMember("", "user", User))
	}

	return obj.Value(), nil
}

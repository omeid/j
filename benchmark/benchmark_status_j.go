package benchmark

import (
	"github.com/omeid/j"
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

	members := make([]j.Member, 21)

	//ResultType := j.NewString(value.ResultType)
	// members[] = j.NewMember("", "result_type", ResultType.Value(.Value()))

	{
		var Contributors j.Value
		if Contributors != nil {
			Contributors = j.NewString(*value.Contributors)
		} else {
			Contributors = j.NewNull()
		}
		members[0] = j.NewMember("", "contributors", Contributors)

	}

	{
		var Coordinates j.Value
		if Coordinates != nil {
			Coordinates = j.NewString(*value.Coordinates)
		} else {
			Coordinates = j.NewNull()
		}
		members[1] = j.NewMember("", "coordinates", Coordinates)

	}

	{
		CreatedAt := j.NewString(value.CreatedAt)
		members[2] = j.NewMember("", "created_at", CreatedAt)

	}

	{
		Entities, err := value.Entities.ToJSON()
		if err != nil {
			return nil, err
		}
		members[3] = j.NewMember("", "entities", Entities)
	}

	{
		Favorited := j.NewBool(value.Favorited)
		members[4] = j.NewMember("", "favorited", Favorited)

	}

	{
		var Geo j.Value
		if Geo != nil {
			Geo = j.NewString(*value.Geo)
		} else {
			Geo = j.NewNull()
		}
		members[5] = j.NewMember("", "geo", Geo)

	}

	{
		ID := j.NewNumberInt64(value.ID)
		members[6] = j.NewMember("", "id", ID)

	}

	{
		IDStr := j.NewString(value.IDStr)
		members[7] = j.NewMember("", "id_str", IDStr)

	}

	{
		var InReplyToScreenName j.Value
		if InReplyToScreenName != nil {
			InReplyToScreenName = j.NewString(*value.InReplyToScreenName)
		} else {
			InReplyToScreenName = j.NewNull()
		}
		members[8] = j.NewMember("", "in_reply_to_screen_name", InReplyToScreenName)

	}

	{
		var InReplyToStatusID j.Value
		if InReplyToStatusID != nil {
			InReplyToStatusID = j.NewString(*value.InReplyToStatusID)
		} else {
			InReplyToStatusID = j.NewNull()
		}
		members[9] = j.NewMember("", "in_reply_to_status_id", InReplyToStatusID)

	}

	{
		var InReplyToStatusIDStr j.Value
		if InReplyToStatusIDStr != nil {
			InReplyToStatusIDStr = j.NewString(*value.InReplyToStatusIDStr)
		} else {
			InReplyToStatusIDStr = j.NewNull()
		}
		members[10] = j.NewMember("", "in_reply_to_status_id_str", InReplyToStatusIDStr)

	}

	{
		var InReplyToUserID j.Value
		if InReplyToUserID != nil {
			InReplyToUserID = j.NewString(*value.InReplyToUserID)
		} else {
			InReplyToUserID = j.NewNull()
		}
		members[11] = j.NewMember("", "in_reply_to_user_id", InReplyToUserID)

	}

	{
		var InReplyToUserIDStr j.Value
		if InReplyToUserIDStr != nil {
			InReplyToUserIDStr = j.NewString(*value.InReplyToUserIDStr)
		} else {
			InReplyToUserIDStr = j.NewNull()
		}
		members[12] = j.NewMember("", "in_reply_to_user_id_str", InReplyToUserIDStr)

	}

	{
		Metadata, err := value.Metadata.ToJSON()
		if err != nil {
			return nil, err
		}
		members[13] = j.NewMember("", "metadata", Metadata)
	}

	{
		var Place j.Value
		if Place != nil {
			Place = j.NewString(*value.Place)
		} else {
			Place = j.NewNull()
		}
		members[14] = j.NewMember("", "place", Place)

	}

	{
		RetweetCount := j.NewNumberInt64(int64(value.RetweetCount))
		members[15] = j.NewMember("", "retweet_count", RetweetCount)

	}

	{
		Retweeted := j.NewBool(value.Retweeted)
		members[16] = j.NewMember("", "retweeted", Retweeted)

	}

	{
		Source := j.NewString(value.Source)
		members[17] = j.NewMember("", "source", Source)

	}

	{
		Text := j.NewString(value.Text)
		members[18] = j.NewMember("", "text", Text)

	}

	{
		Truncated := j.NewBool(value.Truncated)
		members[19] = j.NewMember("", "truncated", Truncated)

	}

	{
		User, err := value.User.ToJSON()
		if err != nil {
			return nil, err
		}
		members[20] = j.NewMember("", "user", User)
	}

	return j.NewObject(members), nil
}

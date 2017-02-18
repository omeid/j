package benchmark

import (
	"github.com/omeid/j"
	"github.com/pkg/errors"
)

// FromJSON makes the lords of json happy.
func (value *User) FromJSON(input j.Value) error {

	if input.Type() != j.ObjectType {
		return errors.New("Unexpected type.")
	}

	if field := input.Member("contributors_enabled"); field != nil {
		if field.Type() != j.BoolType {
			return errors.New("Unexpected type.")
		}

		value.ContributorsEnabled = field.Bool()
	}

	if field := input.Member("created_at"); field != nil {
		if field.Type() != j.StringType {
			return errors.New("Unexpected type.")
		}

		value.CreatedAt = string(field.String())
	}

	if field := input.Member("default_profile"); field != nil {
		if field.Type() != j.BoolType {
			return errors.New("Unexpected type.")
		}

		value.DefaultProfile = field.Bool()
	}

	if field := input.Member("default_profile_image"); field != nil {
		if field.Type() != j.BoolType {
			return errors.New("Unexpected type.")
		}

		value.DefaultProfileImage = field.Bool()
	}

	if field := input.Member("description"); field != nil {
		if field.Type() != j.StringType {
			return errors.New("Unexpected type")
		}

		value.Description = string(field.String())
	}

	if field := input.Member("entities"); field != nil {
		if field.Type() != j.ObjectType {
			return errors.New("Unexpected type")
		}

		UserEntities := UserEntities{}
		err := UserEntities.FromJSON(field)
		if err != nil {
			return err
		}
		value.Entities = UserEntities
	}

	if field := input.Member("favourites_count"); field != nil {
		if field.Type() != j.NumberType {
			return errors.New("Unexpected type")
		}

		n, err := field.Int64()
		if err != nil {
			return err
		}

		value.FavouritesCount = int(n)
	}

	if field := input.Member("follow_request_sent"); field != nil {

		switch field.Type() {
		case j.NullType:
			value.FollowRequestSent = nil
		case j.StringType:
			str := string(field.String())
			value.FollowRequestSent = &str
		default:
			return errors.New("Unexpected type")
		}
	}

	if field := input.Member("followers_count"); field != nil {
		if field.Type() != j.NumberType {
			return errors.New("Unexpected type")
		}

		n, err := field.Int64()
		if err != nil {
			return err
		}

		value.FollowersCount = int(n)
	}

	if field := input.Member("following"); field != nil {
		switch field.Type() {
		case j.NullType:
			value.Following = nil
		case j.StringType:
			str := string(field.String())
			value.Following = &str
		default:
			return errors.New("Unexpected type")
		}
	}

	if field := input.Member("friends_count"); field != nil {
		if field.Type() != j.NumberType {
			return errors.New("Unexpected type")
		}

		n, err := field.Int64()
		if err != nil {
			return err
		}

		value.FriendsCount = int(n)
	}

	if field := input.Member("geo_enabled"); field != nil {
		if field.Type() != j.BoolType {
			return errors.New("Unexpected type.")
		}

		value.GeoEnabled = field.Bool()
	}

	if field := input.Member("id"); field != nil {
		if field.Type() != j.NumberType {
			return errors.New("Unexpected type")
		}

		n, err := field.Int64()
		if err != nil {
			return err
		}

		value.ID = int(n)
	}

	if field := input.Member("id_str"); field != nil {
		if field.Type() != j.StringType {
			return errors.New("Unexpected type.")
		}

		value.IDStr = string(field.String())
	}

	if field := input.Member("is_translator"); field != nil {
		if field.Type() != j.BoolType {
			return errors.New("Unexpected type.")
		}

		value.IsTranslator = field.Bool()
	}

	if field := input.Member("lang"); field != nil {
		if field.Type() != j.StringType {
			return errors.New("Unexpected type")
		}

		value.Lang = string(field.String())
	}

	if field := input.Member("listed_count"); field != nil {
		if field.Type() != j.NumberType {
			return errors.New("Unexpected type")
		}

		n, err := field.Int64()
		if err != nil {
			return err
		}

		value.ListedCount = int(n)
	}

	if field := input.Member("location"); field != nil {
		if field.Type() != j.StringType {
			return errors.New("Unexpected type")
		}

		value.Location = string(field.String())
	}

	if field := input.Member("name"); field != nil {
		if field.Type() != j.StringType {
			return errors.New("Unexpected type")
		}

		value.Name = string(field.String())
	}

	if field := input.Member("notifications"); field != nil {
		switch field.Type() {
		case j.NullType:
			value.Notifications = nil
		case j.StringType:
			str := string(field.String())
			value.Notifications = &str
		default:
			return errors.New("Unexpected type")
		}
	}

	if field := input.Member("profile_background_color"); field != nil {
		if field.Type() != j.StringType {
			return errors.New("Unexpected type")
		}

		value.ProfileBackgroundColor = string(field.String())
	}

	if field := input.Member("profile_background_image_url"); field != nil {
		if field.Type() != j.StringType {
			return errors.New("Unexpected type")
		}

		value.ProfileBackgroundImageURL = string(field.String())
	}

	if field := input.Member("profile_background_image_url_https"); field != nil {
		if field.Type() != j.StringType {
			return errors.New("Unexpected type")
		}

		value.ProfileBackgroundImageURLHTTPS = string(field.String())
	}

	if field := input.Member("profile_background_tile"); field != nil {
		if field.Type() != j.BoolType {
			return errors.New("Unexpected type.")
		}

		value.ProfileBackgroundTile = field.Bool()
	}

	if field := input.Member("profile_image_url"); field != nil {
		if field.Type() != j.StringType {
			return errors.New("Unexpected type")
		}

		value.ProfileImageURL = string(field.String())
	}

	if field := input.Member("profile_image_url_https"); field != nil {
		if field.Type() != j.StringType {
			return errors.New("Unexpected type")
		}

		value.ProfileImageURLHTTPS = string(field.String())
	}

	if field := input.Member("profile_link_color"); field != nil {
		if field.Type() != j.StringType {
			return errors.New("Unexpected type")
		}

		value.ProfileLinkColor = string(field.String())
	}

	if field := input.Member("profile_sidebar_border_color"); field != nil {
		if field.Type() != j.StringType {
			return errors.New("Unexpected type")
		}

		value.ProfileSidebarBorderColor = string(field.String())
	}

	if field := input.Member("profile_sidebar_fill_color"); field != nil {
		if field.Type() != j.StringType {
			return errors.New("Unexpected type")
		}

		value.ProfileSidebarFillColor = string(field.String())
	}

	if field := input.Member("profile_text_color"); field != nil {
		if field.Type() != j.StringType {
			return errors.New("Unexpected type")
		}

		value.ProfileTextColor = string(field.String())
	}

	if field := input.Member("profile_use_background_image"); field != nil {
		if field.Type() != j.BoolType {
			return errors.New("Unexpected type.")
		}

		value.ProfileUseBackgroundImage = field.Bool()
	}

	if field := input.Member("protected"); field != nil {
		if field.Type() != j.BoolType {
			return errors.New("Unexpected type.")
		}

		value.Protected = field.Bool()
	}

	if field := input.Member("screen_name"); field != nil {
		if field.Type() != j.StringType {
			return errors.New("Unexpected type")
		}

		value.ScreenName = string(field.String())
	}

	if field := input.Member("show_all_inline_media"); field != nil {
		if field.Type() != j.BoolType {
			return errors.New("Unexpected type.")
		}

		value.ShowAllInlineMedia = field.Bool()
	}

	if field := input.Member("statuses_count"); field != nil {
		if field.Type() != j.NumberType {
			return errors.New("Unexpected type")
		}

		n, err := field.Int64()
		if err != nil {
			return err
		}

		value.StatusesCount = int(n)
	}

	if field := input.Member("time_zone"); field != nil {
		if field.Type() != j.StringType {
			return errors.New("Unexpected type")
		}

		value.TimeZone = string(field.String())
	}

	if field := input.Member("url"); field != nil {
		switch field.Type() {
		case j.NullType:
			value.URL = nil
		case j.StringType:
			str := string(field.String())
			value.URL = &str
		default:
			return errors.New("Unexpected type")
		}
	}

	if field := input.Member("utc_offset"); field != nil {
		if field.Type() != j.NumberType {
			return errors.New("Unexpected type")
		}

		n, err := field.Int64()
		if err != nil {
			return err
		}

		value.UtcOffset = int(n)
	}

	if field := input.Member("verified"); field != nil {
		if field.Type() != j.BoolType {
			return errors.New("Unexpected type.")
		}

		value.Verified = field.Bool()
	}

	return nil
}

// ToJSON makes the lords of J happy.
func (value *User) ToJSON() (j.Value, error) {

	members := make([]j.Member, 39)

	{
		ContributorsEnabled := j.NewBool(value.ContributorsEnabled)
		members[0] = j.NewMember("", "contributors_enabled", ContributorsEnabled)
	}

	{
		CreatedAt := j.NewString(value.CreatedAt)
		members[1] = j.NewMember("", "created_at", CreatedAt)
	}

	{
		DefaultProfile := j.NewBool(value.DefaultProfile)
		members[2] = j.NewMember("", "default_profile", DefaultProfile)
	}

	{
		DefaultProfileImage := j.NewBool(value.DefaultProfileImage)
		members[3] = j.NewMember("", "default_profile_image", DefaultProfileImage)
	}

	{
		Description := j.NewString(value.Description)
		members[4] = j.NewMember("", "description", Description)
	}

	{
		Entities, err := value.Entities.ToJSON()
		if err != nil {
			return nil, err
		}
		members[5] = j.NewMember("", "entities", Entities)
	}

	{
		FavouritesCount := j.NewNumberInt64(int64(value.FavouritesCount))
		members[6] = j.NewMember("", "favourites_count", FavouritesCount)
	}

	{
		var FollowRequestSent j.Value
		if value.FollowRequestSent != nil {
			FollowRequestSent = j.NewString(*value.FollowRequestSent)
		} else {
			FollowRequestSent = j.NewNull()
		}
		members[7] = j.NewMember("", "follow_request_sent", FollowRequestSent)
	}

	{
		FollowersCount := j.NewNumberInt64(int64(value.FollowersCount))
		members[8] = j.NewMember("", "followers_count", FollowersCount)
	}

	{
		var Following j.Value
		if value.Following != nil {
			Following = j.NewString(*value.Following)
		} else {
			Following = j.NewNull()
		}
		members[9] = j.NewMember("", "following", Following)
	}

	{
		FriendsCount := j.NewNumberInt64(int64(value.FriendsCount))
		members[10] = j.NewMember("", "friends_count", FriendsCount)
	}

	{
		GeoEnabled := j.NewBool(value.GeoEnabled)
		members[11] = j.NewMember("", "geo_enabled", GeoEnabled)
	}

	{
		ID := j.NewNumberInt64(int64(value.ID))
		members[12] = j.NewMember("", "id", ID)
	}

	{
		IDStr := j.NewString(value.IDStr)
		members[13] = j.NewMember("", "id_str", IDStr)
	}

	{
		IsTranslator := j.NewBool(value.IsTranslator)
		members[14] = j.NewMember("", "is_translator", IsTranslator)
	}

	{
		Lang := j.NewString(value.Lang)
		members[15] = j.NewMember("", "lang", Lang)
	}

	{
		ListedCount := j.NewNumberInt64(int64(value.ListedCount))
		members[16] = j.NewMember("", "listed_count", ListedCount)
	}

	{
		Location := j.NewString(value.Location)
		members[17] = j.NewMember("", "location", Location)
	}

	{
		Name := j.NewString(value.Name)
		members[18] = j.NewMember("", "name", Name)
	}

	{
		var Notifications j.Value
		if value.Notifications != nil {
			Notifications = j.NewString(*value.Notifications)
		} else {
			Notifications = j.NewNull()
		}
		members[19] = j.NewMember("", "notifications", Notifications)
	}

	{
		ProfileBackgroundColor := j.NewString(value.ProfileBackgroundColor)
		members[20] = j.NewMember("", "profile_background_color", ProfileBackgroundColor)
	}

	{
		ProfileBackgroundImageURL := j.NewString(value.ProfileBackgroundImageURL)
		members[21] = j.NewMember("", "profile_background_image_url", ProfileBackgroundImageURL)
	}

	{
		ProfileBackgroundImageURLHTTPS := j.NewString(value.ProfileBackgroundImageURLHTTPS)
		members[22] = j.NewMember("", "profile_background_image_url_https", ProfileBackgroundImageURLHTTPS)
	}

	{
		ProfileBackgroundTile := j.NewBool(value.ProfileBackgroundTile)
		members[23] = j.NewMember("", "profile_background_tile", ProfileBackgroundTile)
	}

	{
		ProfileImageURL := j.NewString(value.ProfileImageURL)
		members[24] = j.NewMember("", "profile_image_url", ProfileImageURL)
	}

	{
		ProfileImageURLHTTPS := j.NewString(value.ProfileImageURLHTTPS)
		members[25] = j.NewMember("", "profile_image_url_https", ProfileImageURLHTTPS)
	}

	{
		ProfileLinkColor := j.NewString(value.ProfileLinkColor)
		members[26] = j.NewMember("", "profile_link_color", ProfileLinkColor)
	}

	{
		ProfileSidebarBorderColor := j.NewString(value.ProfileSidebarBorderColor)
		members[27] = j.NewMember("", "profile_sidebar_border_color", ProfileSidebarBorderColor)
	}

	{
		ProfileSidebarFillColor := j.NewString(value.ProfileSidebarFillColor)
		members[28] = j.NewMember("", "profile_sidebar_fill_color", ProfileSidebarFillColor)
	}

	{
		ProfileTextColor := j.NewString(value.ProfileTextColor)
		members[29] = j.NewMember("", "profile_text_color", ProfileTextColor)
	}

	{
		ProfileUseBackgroundImage := j.NewBool(value.ProfileUseBackgroundImage)
		members[30] = j.NewMember("", "profile_use_background_image", ProfileUseBackgroundImage)
	}

	{
		Protected := j.NewBool(value.Protected)
		members[31] = j.NewMember("", "protected", Protected)
	}

	{
		ScreenName := j.NewString(value.ScreenName)
		members[32] = j.NewMember("", "screen_name", ScreenName)
	}

	{
		ShowAllInlineMedia := j.NewBool(value.ShowAllInlineMedia)
		members[33] = j.NewMember("", "show_all_inline_media", ShowAllInlineMedia)
	}

	{
		StatusesCount := j.NewNumberInt64(int64(value.StatusesCount))
		members[34] = j.NewMember("", "statuses_count", StatusesCount)
	}

	{
		TimeZone := j.NewString(value.TimeZone)
		members[35] = j.NewMember("", "time_zone", TimeZone)
	}

	{
		var URL j.Value
		if value.URL == nil {
			URL = j.NewNull()
		} else {
			URL = j.NewString(*value.URL)
		}
		members[36] = j.NewMember("", "url", URL)
	}

	{
		UtcOffset := j.NewNumberInt64(int64(value.UtcOffset))
		members[37] = j.NewMember("", "utc_offset", UtcOffset)
	}

	{
		Verified := j.NewBool(value.Verified)
		members[38] = j.NewMember("", "verified", Verified)
	}

	return j.NewObject(members), nil
}

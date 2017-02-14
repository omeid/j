package benchmark

import (
	"github.com/omeid/j"
	"github.com/omeid/j/mutable"
	"github.com/pkg/errors"
)

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

		value.FollowersCount = int(n)
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

func (value *User) ToJSON() (j.Value, error) {

	obj := mutable.NewObject()

	{
		ContributorsEnabled := mutable.NewBool(value.ContributorsEnabled)
		obj.Add(mutable.NewMember("", "contributors_enabled", ContributorsEnabled.Value()))
	}

	{
		CreatedAt := mutable.NewString(value.CreatedAt)
		obj.Add(mutable.NewMember("", "created_at", CreatedAt.Value()))
	}

	{
		DefaultProfile := mutable.NewBool(value.DefaultProfile)
		obj.Add(mutable.NewMember("", "default_profile", DefaultProfile.Value()))
	}

	{
		DefaultProfileImage := mutable.NewBool(value.DefaultProfileImage)
		obj.Add(mutable.NewMember("", "default_profile_image", DefaultProfileImage.Value()))
	}

	{
		Description := mutable.NewString(value.Description)
		obj.Add(mutable.NewMember("", "description", Description.Value()))
	}

	{
		// Entities                       UserEntities
		// obj.Add(mutable.NewMember("", "entities", nil.Value()))
	}

	{
		FavouritesCount := mutable.NewNumberInt64(int64(value.FavouritesCount))
		obj.Add(mutable.NewMember("", "favourites_count", FavouritesCount.Value()))
	}

	{
		var FollowRequestSent j.Value
		if value.FollowRequestSent != nil {
			FollowRequestSent = mutable.NewString(*value.FollowRequestSent).Value()
		} else {
			FollowRequestSent = mutable.NewNull().Value()
		}
		obj.Add(mutable.NewMember("", "follow_request_sent", FollowRequestSent))
	}

	{
		FollowersCount := mutable.NewNumberInt64(int64(value.FollowersCount))
		obj.Add(mutable.NewMember("", "followers_count", FollowersCount.Value()))
	}

	{
		var Following j.Value
		if value.Following != nil {
			Following = mutable.NewString(*value.Following).Value()
		} else {
			Following = mutable.NewNull().Value()
		}
		obj.Add(mutable.NewMember("", "following", Following))
	}

	{
		FriendsCount := mutable.NewNumberInt64(int64(value.FriendsCount))
		obj.Add(mutable.NewMember("", "friends_count", FriendsCount.Value()))
	}

	{
		GeoEnabled := mutable.NewBool(value.GeoEnabled)
		obj.Add(mutable.NewMember("", "geo_enabled", GeoEnabled.Value()))
	}

	{
		ID := mutable.NewNumberInt64(int64(value.ID))
		obj.Add(mutable.NewMember("", "id", ID.Value()))
	}

	{
		IDStr := mutable.NewString(value.IDStr)
		obj.Add(mutable.NewMember("", "id_str", IDStr.Value()))
	}

	{
		IsTranslator := mutable.NewBool(value.IsTranslator)
		obj.Add(mutable.NewMember("", "is_translator", IsTranslator.Value()))
	}

	{
		Lang := mutable.NewString(value.Lang)
		obj.Add(mutable.NewMember("", "lang", Lang.Value()))
	}

	{
		ListedCount := mutable.NewNumberInt64(int64(value.ListedCount))
		obj.Add(mutable.NewMember("", "listed_count", ListedCount.Value()))
	}

	{
		Location := mutable.NewString(value.Location)
		obj.Add(mutable.NewMember("", "location", Location.Value()))
	}

	{
		Name := mutable.NewString(value.Name)
		obj.Add(mutable.NewMember("", "name", Name.Value()))
	}

	{
		var Notifications j.Value
		if value.Notifications != nil {
			Notifications = mutable.NewString(*value.Notifications).Value()
		} else {
			Notifications = mutable.NewNull().Value()
		}
		obj.Add(mutable.NewMember("", "notifications", Notifications))
	}

	{
		ProfileBackgroundColor := mutable.NewString(value.ProfileBackgroundColor)
		obj.Add(mutable.NewMember("", "profile_background_color", ProfileBackgroundColor.Value()))
	}

	{
		ProfileBackgroundImageURL := mutable.NewString(value.ProfileBackgroundImageURL)
		obj.Add(mutable.NewMember("", "profile_background_image_url", ProfileBackgroundImageURL.Value()))
	}

	{
		ProfileBackgroundImageURLHTTPS := mutable.NewString(value.ProfileBackgroundImageURLHTTPS)
		obj.Add(mutable.NewMember("", "profile_background_image_url_https", ProfileBackgroundImageURLHTTPS.Value()))
	}

	{
		ProfileBackgroundTile := mutable.NewBool(value.ProfileBackgroundTile)
		obj.Add(mutable.NewMember("", "profile_background_tile", ProfileBackgroundTile.Value()))
	}

	{
		ProfileImageURL := mutable.NewString(value.ProfileImageURL)
		obj.Add(mutable.NewMember("", "profile_image_url", ProfileImageURL.Value()))
	}

	{
		ProfileImageURLHTTPS := mutable.NewString(value.ProfileImageURLHTTPS)
		obj.Add(mutable.NewMember("", "profile_image_url_https", ProfileImageURLHTTPS.Value()))
	}

	{
		ProfileLinkColor := mutable.NewString(value.ProfileLinkColor)
		obj.Add(mutable.NewMember("", "profile_link_color", ProfileLinkColor.Value()))
	}

	{
		ProfileSidebarBorderColor := mutable.NewString(value.ProfileSidebarBorderColor)
		obj.Add(mutable.NewMember("", "profile_sidebar_border_color", ProfileSidebarBorderColor.Value()))
	}

	{
		ProfileSidebarFillColor := mutable.NewString(value.ProfileSidebarFillColor)
		obj.Add(mutable.NewMember("", "profile_sidebar_fill_color", ProfileSidebarFillColor.Value()))
	}

	{
		ProfileTextColor := mutable.NewString(value.ProfileTextColor)
		obj.Add(mutable.NewMember("", "profile_text_color", ProfileTextColor.Value()))
	}

	{
		ProfileUseBackgroundImage := mutable.NewBool(value.ProfileUseBackgroundImage)
		obj.Add(mutable.NewMember("", "profile_use_background_image", ProfileUseBackgroundImage.Value()))
	}

	{
		Protected := mutable.NewBool(value.Protected)
		obj.Add(mutable.NewMember("", "protected", Protected.Value()))
	}

	{
		ScreenName := mutable.NewString(value.ScreenName)
		obj.Add(mutable.NewMember("", "screen_name", ScreenName.Value()))
	}

	{
		ShowAllInlineMedia := mutable.NewBool(value.ShowAllInlineMedia)
		obj.Add(mutable.NewMember("", "show_all_inline_media", ShowAllInlineMedia.Value()))
	}

	{
		StatusesCount := mutable.NewNumberInt64(int64(value.StatusesCount))
		obj.Add(mutable.NewMember("", "statuses_count", StatusesCount.Value()))
	}

	{
		TimeZone := mutable.NewString(value.TimeZone)
		obj.Add(mutable.NewMember("", "time_zone", TimeZone.Value()))
	}

	{
		var URL j.Value
		if value.Notifications != nil {
			URL = mutable.NewString(*value.URL).Value()
		} else {
			URL = mutable.NewNull().Value()
		}
		obj.Add(mutable.NewMember("", "url", URL))
	}

	{
		UtcOffset := mutable.NewNumberInt64(int64(value.UtcOffset))
		obj.Add(mutable.NewMember("", "utc_offset", UtcOffset.Value()))
	}

	{
		Verified := mutable.NewBool(value.Verified)
		obj.Add(mutable.NewMember("", "verified", Verified.Value()))
	}

	return obj.Value(), nil
}

package api

import (
	"gvb_server/api/advert_api"
	"gvb_server/api/article_api"
	"gvb_server/api/image_api"
	"gvb_server/api/menu_api"
	"gvb_server/api/message_api"
	"gvb_server/api/settings_api"
	"gvb_server/api/tag_api"
	"gvb_server/api/user_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImagesApi   image_api.ImagesApi
	AdvertApi   advert_api.Advert_Api
	MenuApi     menu_api.MenuApi
	UserApi     user_api.UserApi
	TagApi      tag_api.Tag_Api
	MessageApi  message_api.MessageApi
	ArticleApi  article_api.ArticleApi
}

var ApiGroupApp = new(ApiGroup)

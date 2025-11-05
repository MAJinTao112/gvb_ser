package routers

import (
	"gvb_server/api"
)

func (router RouterGroup) SettingsRouter() {
	settings_api := api.ApiGroupApp.SettingsApi
	router.GET("/settings/:name", settings_api.SettingsInfoView)
	router.PUT("/settings/:name", settings_api.SettingsInfoUpdateView)

}

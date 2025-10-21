package routers

import "gvb_server/api"

func (router RouterGroup) AdvertsRouter() {
	advert_api := api.ApiGroupApp.AdvertApi
	//router.GET("/settings/:name", settings_api.SettingsInfoView)
	router.POST("/adverts", advert_api.AdvertCreat)
	router.GET(("/adverts"), advert_api.AdvertListView)
	router.PUT(("/adverts/:id"), advert_api.AdvertUpdateView)
	router.DELETE("/adverts", advert_api.AdvertRemoveView)

}

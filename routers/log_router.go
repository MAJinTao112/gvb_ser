package routers

import "gvb_server/api"

func (router RouterGroup) LogRouter() {
	app := api.ApiGroupApp.LogApi
	router.GET("logs", app.LogListView)
	router.DELETE("logs", app.LogRemoveListView)

}

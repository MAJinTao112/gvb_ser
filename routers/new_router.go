package routers

import "gvb_server/api"

func (router RouterGroup) NewRouter() {
	app := api.ApiGroupApp.NewApi
	router.GET("new", app.NewListView)

}

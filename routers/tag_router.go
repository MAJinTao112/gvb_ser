package routers

import "gvb_server/api"

func (router RouterGroup) TagRouter() {
	app := api.ApiGroupApp.TagApi
	router.GET("tags", app.TagListView)
	router.POST("tags", app.TagCreat)
	router.PUT("tags", app.TagUpdateView)
	router.DELETE("tags", app.TagRemoveView)

}

package routers

import "gvb_server/api"

func (router RouterGroup) MenuRouter() {
	menu_api := api.ApiGroupApp.MenuApi
	router.POST("/menus", menu_api.MenuCreatView)
	router.GET("/menus", menu_api.MenuListView)
	router.GET("/menu_names", menu_api.MenuNameListView)
	router.PUT("/menus/:id", menu_api.MenuUpdate)
	router.DELETE("/menus", menu_api.MenuDeleteView)
	router.GET("/menus/:id", menu_api.MenuDetailView)

}

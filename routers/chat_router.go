package routers

import (
	"gvb_server/api"
	//"gvb_server/middleware"
)

func (router RouterGroup) ChatRouter() {
	app := api.ApiGroupApp.ChatApi

	router.GET("chatgroups", app.ChatGroupView)
	router.GET("chat_groups_records", app.ChatListView)

}

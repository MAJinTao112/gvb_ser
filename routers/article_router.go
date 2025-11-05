package routers

import (
	"gvb_server/api"
	"gvb_server/middleware"
)

func (router RouterGroup) ArticleRouter() {
	app := api.ApiGroupApp.ArticleApi
	router.POST("articles", middleware.JwtAuth(), app.ArticleCreateView)
	router.GET("articles", middleware.JwtAuth(), app.ArticleListView)
	router.GET("articles/detail", middleware.JwtAuth(), app.ArticleDetailViewByTitleView)
	router.GET("articles/tags", middleware.JwtAuth(), app.ArticleTagListView)
	router.GET("articles/calendar", middleware.JwtAuth(), app.ArticleCalendarView)
	router.PUT("articles", middleware.JwtAuth(), app.ArticleUpdateView)
	router.DELETE("articles", middleware.JwtAuth(), app.ArticleRemoveView)
	router.POST("articles/coll", middleware.JwtAuth(), app.ArticleCollCreateView)
	router.GET("articles/collects", middleware.JwtAuth(), app.ArticleCollListView)
	router.DELETE("articles/collects", middleware.JwtAuth(), app.ArticleCollBatchRemoveView)
	router.GET("articles/texts", app.FullTextSearchView)
	router.GET("articles/:id", middleware.JwtAuth(), app.ArticleDetailView)
}

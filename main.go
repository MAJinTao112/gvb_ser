package main

import (
	"gvb_server/core"
	_ "gvb_server/docs"
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/routers"
)

// @title gvb_serverAPI文档
// @version 1.0
// @description API文档
// @host 127.0.0.01:8080
// @BasePath /
func main() {
	//读取配置
	core.InitConf()
	//fmt.Println(global.Config)
	//初始化日志
	global.Log = core.InitLogger()
	//global.Log.Info("xixixi")
	//初始化连接数据库
	global.DB = core.InitGorm()
	//连接redis
	global.Redis = core.ConnectRedis()
	//连接es
	global.ESClient = core.EsConnect()
	//解析命令行参数

	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}
	//初始化路由
	router := routers.InitRouter()
	//启动路由

	global.Log.Info("server start at: " + global.Config.System.Addr())
	//global.Log.Info("server start at: " + global.Config.System.Addr())
	router.Run(global.Config.System.Addr())

}

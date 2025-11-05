package flag

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/plugins/log_stash"
)

func Makemigrations() {
	//fmt.Println("Makemigrations")
	//fmt.Println("Makemigrations")
	//初始化数据库
	var err error
	//global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})

	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&models.UserModel{},
		&models.MenuModel{},
		&models.MenuBannerModel{},
		&models.BannerModel{},
		&models.MessageModel{},
		&models.AdvertModel{},
		&models.CommentModel{},
		&models.TagModel{},
		&models.FadeBackModel{},
		&models.LoginDataModel{},
		&models.UserCollectModel{},
		&models.ChatModel{},
		&log_stash.LogStashModel{},
	)
	if err != nil {
		global.Log.Error("error：生成数据库表失败")
		return
	}
	global.Log.Info("数据库表生成成功")
	//return
}

package menu_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

func (MenuApi) MenuUpdate(c *gin.Context) {
	var cr MenuRequset
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}
	id := c.Param("id")
	//先把之前的banner清空
	var menuModel models.MenuModel
	err = global.DB.Take(&menuModel, id).Error
	if err != nil {
		res.OkWithMessage("菜单不存在", c)
		return
	}
	global.DB.Model(&menuModel).Association("Banners").Clear()
	//如果选择了banner,那就添加
	if len(cr.ImageSortList) > 0 {
		var bannerList []models.MenuBannerModel
		for _, banner := range cr.ImageSortList {
			bannerList = append(bannerList, models.MenuBannerModel{

				MenuID:   menuModel.ID,
				BannerID: banner.ImageID,
				Sort:     banner.Sort,
			})
		}
		err = global.DB.Create(&bannerList).Error
		if err != nil {
			global.Log.Error(err.Error())
			res.FailWithMessage("创建菜单图片失败", c)
			return
		}
	}
	//普通更新
	maps := structs.Map(&cr)
	err = global.DB.Model(&menuModel).Updates(maps).Error
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("修改菜单失败", c)
		return
		
	}
	res.OkWithMessage("修改菜单成功", c)
}

package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}
type MenuRequset struct {
	Title         string      `json:"title" binding:"required" msg:"请输入菜单名称" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"请输入菜单路径" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"`
	BannerTime    int         `json:"banner_time" structs:"banner_time"`
	Sort          int         `json:"sort" binding:"required" msg:"请输入菜单序号" structs:"sort"`
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`
}

func (MenuApi) MenuCreatView(c *gin.Context) {
	var cr MenuRequset
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
	}
	//重复值判断
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, "title = ? and path = ?", cr.Title, cr.Path).RowsAffected
	fmt.Println(count)
	if count >= 1 {
		res.FailWithMessage("该目录已经存在", c)
		return
	}
	//创建banner数据入库
	menumodel := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	err = global.DB.Create(&menumodel).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("菜单添加失败", c)
		return
	}
	if len(cr.ImageSortList) == 0 {
		res.OkWithMessage("菜单添加成功", c)
		return
	}
	//给第三张表入库
	var menuBannerList []models.MenuBannerModel

	for _, a := range cr.ImageSortList {

		menuBannerList = append(menuBannerList, models.MenuBannerModel{
			MenuID:   menumodel.ID,
			BannerID: a.ImageID,
			Sort:     a.Sort,
		})

	}
	err = global.DB.Create(&menuBannerList).Error
	if err != nil {
		res.FailWithMessage("菜单图片关联失败", c)
		return
	}
	res.OkWithMessage("菜单添加成功,菜单图片关联成功", c)

}

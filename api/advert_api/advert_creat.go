package advert_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type AdvertRequest struct {
	Title  string `json:"title" binding:"required" msg:"请输入标题"`           // 显示的标题
	Href   string `json:"href" binding:"required,url" msg:"请输入一个合法跳转链接"`  // 跳转链接
	Images string `json:"images" binding:"required,url" msg:"请输入☝️合法的图片"` // 图片
	IsShow bool   `json:"is_show"         msg:"请选择是否展示"`                  // 是否展示
}

// AdvertCreateView 添加广告
// @Tags 广告管理
// @Summary 创建广告
// @Description 创建广告
// @Param data body AdvertRequest    true  "表示多个参数"
// @Param token header string  true  "token"
// @Router /api/adverts [post]
// @Produce json
// @Success 200 {object} res.Response{}
func (Advert_Api) AdvertCreat(c *gin.Context) {
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	//重复判断
	var advert models.AdvertModel
	err = global.DB.Take(&advert, "title=?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("该广告已存在", c)
		return
	}
	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加广告失败", c)
		return
	}
	fmt.Println(cr)
	res.OkWithMessage("添加广告成功", c)
}

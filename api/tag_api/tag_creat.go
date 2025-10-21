package tag_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

type TagRequest struct {
	Title string `json:"title" binding:"required" msg:"请输入标题"` // 显示的标题
}

func (Tag_Api) TagCreat(c *gin.Context) {
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}

	//重复判断
	var tag models.TagModel
	err = global.DB.Take(&tag, "title=?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("该标签已存在", c)
		return
	}
	err = global.DB.Create(&models.TagModel{
		Title: cr.Title,
	}).Error

	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加标签失败", c)
		return
	}
	fmt.Println(cr)
	res.OkWithMessage("添加标签成功", c)
}

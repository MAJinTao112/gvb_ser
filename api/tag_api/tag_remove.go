package tag_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// AdvertRemoveView 批量删除广告
// @Tags 广告管理
// @Summary 批量删除广告
// @Description 批量删除广告
// @Param token header string  true  "token"
// @Param data body models.RemoveRequest    true  "广告id列表"
// @Router /api/adverts [delete]
// @Produce json
// @Success 200 {object} res.Response{}
func (Tag_Api) TagRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var tagLIst []models.TagModel
	count := global.DB.Find(&tagLIst, cr.IDList).RowsAffected
	if count < 1 {
		res.FailWithMessage("标签不存在，请检查", c)
		return
	}
	//如果标签下有关联的文章，怎么办？
	global.DB.Delete(&tagLIst)
	res.OkWithMessage(fmt.Sprintf("共删除了%d个标签", count), c)
}

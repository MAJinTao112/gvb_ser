package tag_api

import (
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
)

// AdvertUpdateView 更新广告
// @Tags 广告管理
// @Summary 更新广告
// @Param token header string  true  "token"
// @Description 更新广告
// @Param data body AdvertRequest    true  "广告的一些参数"
// @Param id path int true "id"
// @Router /api/adverts/{id} [put]
// @Produce json
// @Success 200 {object} res.Response{}
func (Tag_Api) TagUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr TagRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		//return
	}
	var tag models.TagModel
	err = global.DB.Take(&tag, id).Error
	if err != nil {
		res.FailWithMessage("该标签不存在", c)
		return
	}

	maps := structs.Map(&cr)
	err = global.DB.Model(&tag).Updates(&maps).Error
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("修改标签成功", c)
}

package advert_api

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
func (Advert_Api) AdvertUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		//return
	}
	var advert models.AdvertModel
	err = global.DB.Take(&advert, id).Error
	if err != nil {
		res.FailWithMessage("该广告不存在", c)
		return
	}

	maps := structs.Map(&cr)
	err = global.DB.Model(&advert).Updates(&maps).Error
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("修改数据成功", c)
}

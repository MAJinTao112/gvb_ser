package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
)

type SettingsUri struct {
	Name string `uri:"name"`
}

func (SettingsApi) SettingsInfoView(c *gin.Context) {

	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		global.Log.Error(err.Error())
		return
	}
	switch cr.Name {
	case "siteinfo":
		res.OkWithData(global.Config.SiteInfo, c)
		return
	case "qiniu":
		res.OkWithData(global.Config.QiNiu, c)
		return
	case "email":
		res.OkWithData(global.Config.Email, c)
		return
	case "jwt":
		res.OkWithData(global.Config.Jwt, c)
		return
	case "qq":
		res.OkWithData(global.Config.QQ, c)
		return
	default:
		res.FailWithMessage("没有对应的参数信息", c)
	}
	//res.OkWithData(global.Config.SiteInfo, c)
}

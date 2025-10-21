package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/models/res"
)

func (SettingsApi) SettingsInfoUpdateView(c *gin.Context) {

	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	switch cr.Name {
	case "siteinfo":
		var info config.SiteInfo
		err = c.ShouldBind(&info)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			return
		}
		global.Config.SiteInfo = info
		core.SetYaml()

		res.OkWith(c)
		return
	case "qiniu":
		var info config.QiNiu
		err = c.ShouldBind(&info)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			return
		}
		global.Config.QiNiu = info
		core.SetYaml()
		res.OkWith(c)
		return
	case "email":
		var info config.Email
		err = c.ShouldBind(&info)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			return
		}
		global.Config.Email = info
		core.SetYaml()

		res.OkWith(c)
		return
	case "jwt":
		var info config.Jwt
		err = c.ShouldBind(&info)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			return
		}
		global.Config.Jwt = info
		core.SetYaml()

		res.OkWith(c)
		return
	case "qq":
		var info config.QQ
		err = c.ShouldBind(&info)
		if err != nil {
			res.FailWithMessage(err.Error(), c)
			return
		}
		global.Config.QQ = info
		core.SetYaml()

		res.OkWith(c)
		return
	default:
		res.FailWithMessage("没有您想要修改的的参数信息", c)
	}
}

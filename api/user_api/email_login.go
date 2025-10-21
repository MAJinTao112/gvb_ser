package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/utils/jwts"
	"gvb_server/utils/pwd"
)

type EmailLoginView struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (UserApi) EmailLoginView(c *gin.Context) {
	var cr EmailLoginView
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.Password).Error
	if err != nil {
		global.Log.Warn("用户名不存在:", err.Error())
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("用户名密码错误")
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	token, err := jwts.GenToken(jwts.JwtPayload{
		//UserName: userModel.UserName,
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   userModel.ID,
	})
	if err != nil {
		global.Log.Error(err)
		return
	}
	res.OkWithData(token, c)
}

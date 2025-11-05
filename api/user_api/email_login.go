package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/plugins/log_stash"
	"gvb_server/utils"
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
	log := log_stash.NewLogByGin(c)
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name = ? or email = ?", cr.UserName, cr.Password).Error
	if err != nil {
		global.Log.Warn("用户名不存在:", err.Error())
		log.Warn(fmt.Sprintf("%s用户名不存在", cr.UserName))
		res.FailWithMessage("用户名或密码错误", c)
		return
	}
	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("用户名密码错误")
		log.Warn(fmt.Sprintf("用户名或密码错误%s%s", cr.UserName, cr.Password))
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
		res.FailWithMessage("token生成失败", c)
		log.Warn(fmt.Sprintf("%stoken生成失败", err.Error()))
		return
	}
	ip, addr := utils.GetAddrByGin(c)
	log = log_stash.New(ip, token)
	log.Info(fmt.Sprintf("%s登陆成功", cr.UserName))
	global.DB.Create(&models.LoginDataModel{
		UserID:    userModel.ID,
		IP:        ip,
		NickName:  userModel.NickName,
		Token:     token,
		Device:    "",
		Addr:      addr,
		LoginType: ctype.SignEmail,
	})
	res.OkWithData(token, c)
}

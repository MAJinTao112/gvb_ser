package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/plugins/qq"
	"gvb_server/utils"
	"gvb_server/utils/jwts"
	"gvb_server/utils/pwd"
	"gvb_server/utils/random"
)

func (UserApi) QQLoginView(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		res.FailWithMessage("没有code", c)
		return
	}
	fmt.Println(code)
	//res.OkWithMessage("成", c)
	qqInfo, err := qq.NewQQLogin(code)
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(qqInfo)
	openID := qqInfo.OpenID
	//根据openID判断用户是否存在
	var user models.UserModel
	ip, addr := utils.GetAddrByGin(c)
	err = global.DB.Take(&user, "token = ?", openID).Error
	if err != nil {
		//res.FailWithMessage("用户不存在", c)
		//注册
		hashpwd := pwd.HashPwd(random.RandStr(16))
		user = models.UserModel{
			NickName:   qqInfo.Nickname,
			UserName:   openID, //直接用openid代替
			Password:   hashpwd,
			Avatar:     qqInfo.Avatar,
			Token:      openID,
			Addr:       addr,
			IP:         ip,
			Role:       ctype.PermissionUser,
			SignStatus: ctype.SignQQ,
		}
		err = global.DB.Create(&user).Error
		if err != nil {
			res.FailWithMessage("注册失败。", c)
			return
		}
	}
	//登录操作
	token, err := jwts.GenToken(jwts.JwtPayload{
		NickName: user.NickName,
		Role:     int(user.Role),
		UserID:   user.ID,
	})
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage("token生成失败", c)
		return
	}

	global.DB.Create(&models.LoginDataModel{
		UserID:    user.ID,
		IP:        ip,
		NickName:  user.NickName,
		Token:     token,
		Device:    "",
		Addr:      addr,
		LoginType: ctype.SignQQ,
	})
	res.OkWithData(token, c)

}

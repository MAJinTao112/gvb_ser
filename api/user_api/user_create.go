package user_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/service/user_ser"
)

type UserCreateRequest struct {
	Nickname string `json:"nickname" binding:"required" msg:"请输入昵称"`
	Username string `json:"username" binding:"required" msg:"请输入用户名"`
	Password string `json:"password" binding:"required" msg:"请输入密码"`
	Role     int    `json:"role" binding:"required" msg:"请选择权限"`
}

func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	err = user_ser.UserService{}.CreatUser(cr.Username, cr.Nickname, cr.Password, ctype.Role(cr.Role), "", c.ClientIP())
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMessage(err.Error(), c)
		return
	}
	global.Log.Infof("用户%s创建成功!", cr.Username)
	res.OkWithMessage(fmt.Sprintf("用户%s创建成功!", cr.Username), c)
	return
}

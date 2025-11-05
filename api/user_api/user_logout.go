package user_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models/res"
	"gvb_server/service"
	"gvb_server/utils/jwts"
)

func (UserApi) LogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	token := c.Request.Header.Get("token")

	err := service.ServiceGroupApp.UserService.Logout(token, claims)

	if err != nil {
		res.FailWithMessage("注销失败", c)
		global.Log.Error(err)
		return
	}
	res.OkWithMessage("注销成功", c)
}

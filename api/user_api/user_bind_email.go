package user_api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/plugins/email"
	"gvb_server/utils/jwts"
	"gvb_server/utils/pwd"
	"gvb_server/utils/random"
)

type BindmailRequest struct {
	Email    string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code     *string `json:"code"`
	Password string  `json:"password" `
}

// ////// binding:"required" msg:"请输入密码"
func (UserApi) UserBindEmailView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var cr BindmailRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	session := sessions.Default(c)

	if cr.Code == nil {
		//第一次，后台发送验证码
		//生成四位数验证码，将生成的验证码存入session
		code := random.Code(4)
		fmt.Println(code)
		session.Set("valid_code", code)
		err := session.Save()
		if err != nil {
			res.FailWithError(err, &cr, c)
			global.Log.Error(err)
			return
		}
		err = email.NewCode().Send(cr.Email, fmt.Sprintf("你的验证码是%s", code))
		if err != nil {
			global.Log.Error(err)
		}
		res.OkWithMessage("验证码已发送，请查收", c)
		return
	}
	code := session.Get("valid_code")
	if code != *cr.Code {
		fmt.Println(code)
		res.FailWithMessage("您输入的验证码有误，请重新输入", c)
		return
	}
	//修改用户的邮箱
	var user models.UserModel
	err = global.DB.Take(&user, claims.UserID).Error
	if err != nil {
		res.FailWithMessage("用户不存在", c)
		return
	}
	if len(cr.Password) < 4 {
		res.FailWithMessage("邮箱密码强度过低", c)
	}
	hashpwd := pwd.HashPwd(cr.Password)

	//第一次的邮箱也要做一致性校验
	err = global.DB.Model(&user).Updates(map[string]any{
		"email":    cr.Email,
		"password": hashpwd,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("绑定邮箱失败", c)
		return
	}
	res.OkWithMessage("邮箱绑定成功", c)
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models/res"
	"gvb_server/service/redis_ser"
	"gvb_server/utils/jwts"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			res.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			res.FailWithMessage("token错误", c)
			c.Abort()
			return
		}
		if redis_ser.CheckLogout(token) {
			res.FailWithMessage("用户已注销，token失效", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

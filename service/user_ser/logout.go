package user_ser

import (
	"gvb_server/service/redis_ser"
	"gvb_server/utils/jwts"
	"time"
)

func (UserService) Logout(token string, claims *jwts.CustomClaims) error {
	exp := claims.ExpiresAt
	now := time.Now()
	diff := exp.Time.Sub(now)
	return redis_ser.Logout(token, diff)

}

package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type JwtPayload struct {
	UserName string `json:"user_name"`
	NickName string `json:"nick_name"`
	Role     int    `json:"role"`
	UserID   uint   `json:"user_id"`
}

type CustomClaims struct {
	JwtPayload
	jwt.StandardClaims
}

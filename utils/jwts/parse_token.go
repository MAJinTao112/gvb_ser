package jwts

import (
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"gvb_server/global"
)

func ParseToken(tokenString string) (*CustomClaims, error) {
	MySecret := []byte(global.Config.Jwt.Secret)
	token, err := jwt.ParseWithClaims(
		tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return MySecret, nil
		})
	if err != nil {
		global.Log.Error("token parse err:", err.Error())
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")

}

package redis_ser

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/utils"
	"time"
)

const prefix = "logout_"

func Logout(token string, diff time.Duration) error {
	err := global.Redis.Set(fmt.Sprintf("%s%s", prefix, token), "", diff).Err()

	return err
}
func CheckLogout(token string) bool {
	keys := global.Redis.Keys("logout_*").Val()
	if utils.InList("logout_"+token, keys) {
		return true
	}
	return false
}

package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5计算出哈希值返回
func Md5(src []byte) string {
	m := md5.New()
	m.Write(src)
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

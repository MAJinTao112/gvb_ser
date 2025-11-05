package random

import (
	"math/rand"
	"time"
)

// Code 生成 n 位数字验证码
func Code(n int) string {
	rand.Seed(time.Now().UnixNano()) // 设置随机种子
	digits := []rune("0123456789")   // 可选字符
	code := make([]rune, n)
	for i := range code {
		code[i] = digits[rand.Intn(len(digits))]
	}
	return string(code)
}

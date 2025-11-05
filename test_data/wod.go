package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	// 获取本机所有网络接口（网卡）
	Core := cron.New(cron.WithSeconds())

	Core.AddFunc("* * * * * *", func() {
		fmt.Println("啦啦啦啦啦啦")
		fmt.Println(time.Now().String())
	})
	Core.Start()
	for {
	}
}

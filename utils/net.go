package utils

import (
	"github.com/sirupsen/logrus" // 用于记录日志（比标准log更灵活）
	"log"                        // 标准日志库
	"net"                        // 网络相关操作
)

// GetIPList 返回本机所有 IPv4 地址
func GetIPList() (ipList []string) {
	// 获取本机所有网络接口（网卡）
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err) // 如果获取失败，直接退出程序
	}

	// 遍历每一个网络接口
	for _, i2 := range interfaces {
		// 获取当前接口绑定的所有地址（可能有 IPv4、IPv6、环回地址等）
		addresses, err := i2.Addrs()
		if err != nil {
			logrus.Error(err) // 打印错误，但继续处理其他接口
			continue
		}

		// 遍历每个地址
		for _, addr := range addresses {
			// addr 类型断言为 *net.IPNet，才可以获取 IP 和掩码
			ipNet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			// 只取 IPv4 地址（排除 IPv6）
			ip4 := ipNet.IP.To4()
			if ip4 == nil {
				continue
			}

			// 转成字符串并添加到结果列表
			ipList = append(ipList, ip4.String())
		}
	}
	return
}

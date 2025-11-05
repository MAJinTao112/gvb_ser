package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"net"
)

// GetAddrByGin 从 gin 上下文获取客户端 IP 并返回 IP 和解析到的地址字符串
func GetAddrByGin(c *gin.Context) (ip, addr string) {
	ip = c.ClientIP()  // 从 gin.Context 获取客户端 IP（可能包含代理头判断，如 X-Forwarded-For）
	addr = GetAddr(ip) // 调用 GetAddr 解析该 IP 对应的地理位置或其他描述
	return ip, addr    // 返回 IP 和地址字符串
}

// GetAddr 根据传入的 IP 字符串返回地址描述（例如 "省-市"）
// 若是内网 IP 或解析出错，返回相应的提示字符串
func GetAddr(ip string) string {
	parseIP := net.ParseIP(ip) // 将字符串形式的 IP 解析为 net.IP 类型
	if IsIntranetIP(parseIP) { // 判断是否为内网 IP（或回环、IPv6 在此实现中也被视作内网/非公网）
		return "内网地址" // 如果是内网，直接返回固定字符串
	}
	record, err := global.AddrDB.City(net.ParseIP(ip)) // 使用全局 AddrDB（GeoIP2 数据库）查询该 IP 的城市信息
	if err != nil {                                    // 若查询返回错误（例如数据库读取失败或 IP 无效）
		return "错误的地址" // 返回错误提示字符串
	}
	var province string               // 用于保存省/州名称（中文）
	if len(record.Subdivisions) > 0 { // Subdivisions 可能为空，访问前必须检查长度
		province = record.Subdivisions[0].Names["zh-CN"] // 取第一个 subdivision（通常为省/州），并获取中文名称
	}
	city := record.City.Names["zh-CN"]          // 从 City 结构体中获取中文城市名称（可能为空）
	return fmt.Sprintf("%s-%s", province, city) // 返回格式化字符串 "省-市"（若为空会出现 "-市" 或 "省-" 或 "-"）
}

// IsIntranetIP 判断给定的 net.IP 是否属于内网或本机等非公网地址
func IsIntranetIP(ip net.IP) bool {
	if ip.IsLoopback() { // 判断是否为回环地址（如 127.0.0.1 或 ::1）
		return true // 回环地址视为内网/本机地址
	}

	ip4 := ip.To4() // 尝试将 IP 转为 IPv4（成功返回 4 字节切片，失败返回 nil）
	if ip4 == nil { // 如果为 nil，说明不是 IPv4（可能为 IPv6）
		return true // 在此实现中把 IPv6 当作内网/非公网处理（注意：这是一种简化处理，不一定适合所有场景）
	}
	// 以下为 RFC1918 规定的私有地址段判断（以及 link-local）
	// 192.168.0.0 - 192.168.255.255
	// 172.16.0.0 - 172.31.255.255
	// 10.0.0.0 - 10.255.255.255
	// 169.254.0.0 - 169.254.255.255 （link-local，本地链路地址）
	return (ip4[0] == 192 && ip4[1] == 168) || // 判断 192.168.x.x
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 32) || // 判断 172.16.x.x ~ 172.31.x.x（注意：<=32 应为 <=31 才完全符合 RFC1918）
		(ip4[0] == 10) || // 判断 10.x.x.x
		(ip4[0] == 169 && ip4[1] == 254) // 判断 169.254.x.x（link-local）
}

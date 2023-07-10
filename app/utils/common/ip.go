package common

import (
	"math/big"
	"net"
	"strings"
)

// ip字符串转int
func IpToInt(ip net.IP) *big.Int {
	if v := ip.To4(); v != nil {
		return big.NewInt(0).SetBytes(v)
	}
	return big.NewInt(0).SetBytes(ip.To16())
}

// int转字符串ip
func IntToIP(i *big.Int) net.IP {
	return net.IP(i.Bytes())
}

// 转字符串转ip
func StringToIP(i string) net.IP {
	return net.ParseIP(i).To4()
}

// 获取Ip和端口
func GetIpAndPort(ip string) (string, string) {
	if strings.Contains(ip, ":") {
		arr := strings.Split(ip, ":")
		return arr[0], arr[1]
	}
	return ip, "22"
}

// GetIp 获取IP
func GetIp() string {
	udpAddr, err := net.ResolveUDPAddr("udp", "8.8.8.8:80")
	if err != nil {
		return ""
	}
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return ""
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

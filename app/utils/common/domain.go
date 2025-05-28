package common

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

var GlobalContext *gin.Context

func SetGlobalContext(c *gin.Context) *gin.Context {
	GlobalContext = c
	return GlobalContext
}

// 获取当前请求源域名
func GetRequestOrigin() string {
	referer := GlobalContext.GetHeader("Referer")
	if referer == "" {
		// 如果Referer为空，使用当前请求的域名
		return GetCurrentRequestDomain()
	}
	return GetHost(referer)
}

// 获取当前请求的完整域名（考虑各种代理情况）
func GetCurrentRequestDomain() string {
	// 获取真实的协议
	scheme := GetRealScheme()
	// 获取真实的主机名
	host := GetRealHost()
	return fmt.Sprintf("%s://%s", scheme, host)
}

// 获取真实的协议（考虑代理情况）
func GetRealScheme() string {
	// 检查各种代理头部来确定真实协议
	if proto := GlobalContext.GetHeader("X-Forwarded-Proto"); proto != "" {
		// 支持多级代理，取第一个值
		if strings.Contains(proto, ",") {
			proto = strings.TrimSpace(strings.Split(proto, ",")[0])
		}
		return proto
	}

	// 检查其他常见的协议头部
	if GlobalContext.GetHeader("X-Forwarded-Ssl") == "on" ||
		GlobalContext.GetHeader("X-Forwarded-Scheme") == "https" ||
		GlobalContext.GetHeader("X-Scheme") == "https" {
		return "https"
	}

	// 检查CloudFlare的头部
	if GlobalContext.GetHeader("CF-Visitor") != "" {
		// CF-Visitor 通常是 JSON 格式，如: {"scheme":"https"}
		if strings.Contains(GlobalContext.GetHeader("CF-Visitor"), `"scheme":"https"`) {
			return "https"
		}
	}

	// 检查TLS连接
	if GlobalContext.Request.TLS != nil {
		return "https"
	}

	// 默认返回http
	return "http"
}

// 获取真实的主机名（考虑代理情况）
func GetRealHost() string {
	// 按优先级检查各种主机头部

	// 1. X-Forwarded-Host (最常用的代理头部)
	if host := GlobalContext.GetHeader("X-Forwarded-Host"); host != "" {
		// 支持多级代理，取第一个值
		if strings.Contains(host, ",") {
			host = strings.TrimSpace(strings.Split(host, ",")[0])
		}
		return host
	}

	// 2. X-Original-Host (某些代理使用)
	if host := GlobalContext.GetHeader("X-Original-Host"); host != "" {
		return host
	}

	// 3. X-Host (某些负载均衡器使用)
	if host := GlobalContext.GetHeader("X-Host"); host != "" {
		return host
	}

	// 4. X-Real-Host (某些代理使用)
	if host := GlobalContext.GetHeader("X-Real-Host"); host != "" {
		return host
	}

	// 5. 回退到标准的Host头部
	return GlobalContext.Request.Host
}

// GetHost 返回指定URL的域名
func GetHost(varURL string) string {
	u, err := url.Parse(varURL)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s://%s", u.Scheme, u.Host)
}

// 替换域名路径
func ReplaceDomainPath(domain string, path string) string {
	u, err := url.Parse(path)
	if err != nil {
		fmt.Println(err)
	}
	u.Host = ""
	u.Scheme = ""
	u.Hostname()
	// 返回替换后的 URL
	if strings.HasSuffix(domain, "/") || strings.HasPrefix(u.String(), "/") {
		return fmt.Sprintf("%s%s", domain, u.String())
	} else {
		return fmt.Sprintf("%s/%s", domain, u.String())
	}
}

// 获取真实的客户端IP（考虑代理情况）
func GetRealClientIP() string {
	// 按优先级检查各种IP头部

	// 1. X-Forwarded-For (最常用，支持多级代理)
	if xForwardedFor := GlobalContext.GetHeader("X-Forwarded-For"); xForwardedFor != "" {
		// X-Forwarded-For 可能包含多个IP，第一个通常是真实客户端IP
		ips := strings.Split(xForwardedFor, ",")
		if len(ips) > 0 {
			clientIP := strings.TrimSpace(ips[0])
			// 验证IP格式
			if clientIP != "" && !isPrivateIP(clientIP) {
				return clientIP
			}
		}
	}

	// 2. X-Real-IP (Nginx常用)
	if xRealIP := GlobalContext.GetHeader("X-Real-IP"); xRealIP != "" {
		if !isPrivateIP(xRealIP) {
			return xRealIP
		}
	}

	// 3. X-Client-IP
	if xClientIP := GlobalContext.GetHeader("X-Client-IP"); xClientIP != "" {
		if !isPrivateIP(xClientIP) {
			return xClientIP
		}
	}

	// 4. CF-Connecting-IP (CloudFlare)
	if cfIP := GlobalContext.GetHeader("CF-Connecting-IP"); cfIP != "" {
		return cfIP
	}

	// 5. True-Client-IP (Akamai, CloudFlare)
	if trueClientIP := GlobalContext.GetHeader("True-Client-IP"); trueClientIP != "" {
		return trueClientIP
	}

	// 6. 回退到RemoteAddr
	if ip := GlobalContext.ClientIP(); ip != "" {
		return ip
	}

	return GlobalContext.Request.RemoteAddr
}

// 检查是否为私有IP地址
func isPrivateIP(ip string) bool {
	// 移除端口号
	if strings.Contains(ip, ":") {
		ip = strings.Split(ip, ":")[0]
	}

	// 简单的私有IP检查
	return strings.HasPrefix(ip, "10.") ||
		strings.HasPrefix(ip, "172.") ||
		strings.HasPrefix(ip, "192.168.") ||
		strings.HasPrefix(ip, "127.") ||
		ip == "localhost"
}

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
		scheme := "http"
		if GlobalContext.Request.TLS != nil || GlobalContext.GetHeader("X-Forwarded-Proto") == "https" {
			scheme = "https"
		}
		return fmt.Sprintf("%s://%s", scheme, GlobalContext.Request.Host)
	}
	return GetHost(referer)
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

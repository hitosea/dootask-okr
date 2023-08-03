package common

import (
	"fmt"
	"net/url"

	"github.com/gin-gonic/gin"
)

var GlobalContext *gin.Context

func SetGlobalContext(c *gin.Context) *gin.Context {
	GlobalContext = c
	return GlobalContext
}

// 获取当前请求源域名
func GetRequestOrigin() string {
	return GetHost(GlobalContext.GetHeader("Referer"))
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
	return fmt.Sprintf("%s%s", domain, u.String())
}

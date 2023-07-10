package common

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 获取当前域名
func GetCurrentDomain(c *gin.Context) string {
	if IsHttps(c) {
		return fmt.Sprintf("https://%s", c.Request.Host)
	} else {
		return fmt.Sprintf("http://%s", c.Request.Host)
	}
}

// 判断当前请求是否https
func IsHttps(c *gin.Context) bool {
	return c.Request.TLS != nil
}

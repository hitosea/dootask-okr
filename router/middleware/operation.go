package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func OperationLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, "search") || c.Request.Method == http.MethodGet {
			c.Next()
			return
		}
		c.Next()
	}
}

package helper

import (
	"dootask-okr/app/constant"
	"dootask-okr/i18n"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// Token 获取Token（Header、Query、Cookie）
func Token(c *gin.Context) string {
	token := c.GetHeader("token")
	if token == "" {
		token = Input(c, "token")
	}
	if token == "" {
		token = Cookie(c, "token")
	}
	return token
}

// Input 获取参数（优先POST、取Query）
func Input(c *gin.Context, key string) string {
	if c.PostForm(key) != "" {
		return strings.TrimSpace(c.PostForm(key))
	}
	return strings.TrimSpace(c.Query(key))
}

// Scheme 获取Scheme
func Scheme(c *gin.Context) string {
	scheme := "http://"
	if c.Request.TLS != nil || c.Request.Header.Get("X-Forwarded-Proto") == "https" {
		scheme = "https://"
	}
	return scheme
}

// Home 获取Home Url
func Home(c *gin.Context) string {
	return fmt.Sprintf("%s%s", Scheme(c), c.Request.Host)
}

// Cookie 获取Cookie
func Cookie(c *gin.Context, name string) string {
	value, _ := c.Cookie(name)
	return value
}

// SetCookie 设置Cookie
func SetCookie(c *gin.Context, name, value string, maxAge int) {
	c.SetCookie(name, value, maxAge, "/", "", false, false)
}

// RemoveCookie 删除Cookie
func RemoveCookie(c *gin.Context, name string) {
	c.SetCookie(name, "", -1, "/", "", false, false)
}

// Response 返回结果
func Response(c *gin.Context, code int, msg string, values ...any) {
	c.Header("Expires", "-1")
	c.Header("Cache-Control", "no-cache")
	c.Header("Pragma", "no-cache")
	var data any
	if len(values) == 1 {
		data = values[0]
	} else if len(values) == 0 {
		data = gin.H{}
	} else {
		data = values
	}

	//
	if strings.Contains(c.GetHeader("Accept"), "application/json") ||
		strings.Contains(c.GetHeader("Content-Type"), "application/json") ||
		strings.Contains(c.GetHeader("X-Requested-With"), "XMLHttpRequest") ||
		strings.Contains(c.Request.URL.Path, "/api/v1/") ||
		!strings.Contains(c.Request.Method, http.MethodGet) {
		// 接口返回
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  msg,
			"data": data,
		})
	} else {
		// 页面返回
		if code == http.StatusMovedPermanently {
			c.Redirect(code, msg)
		} else {
			c.HTML(http.StatusOK, "index", gin.H{
				"CODE": code,
				"MSG":  url.QueryEscape(msg),
			})
		}
	}
}

// Success 成功
func Success(c *gin.Context, values ...any) {
	Response(c, http.StatusOK, "success", values...)
}

// Error 失败
func Error(c *gin.Context, values ...any) {
	Response(c, http.StatusBadRequest, "error", values...)
}

// ErrorWith 失败信息
func ErrorWith(c *gin.Context, msgKey string, err error, values ...any) {
	msgDetail := i18n.GetMsgWithMap(msgKey, map[string]any{"detail": err})
	Response(c, http.StatusBadRequest, msgDetail, values...)
}

// ErrorAuth 身份失败
func ErrorAuth(c *gin.Context, values ...any) {
	msgDetail := i18n.GetMsgWithMap(constant.ErrTypeNotLogin, map[string]any{"detail": nil})
	Response(c, http.StatusUnauthorized, msgDetail, values...)
}

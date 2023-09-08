package router

import (
	v1 "dootask-okr/app/api/v1"
	"dootask-okr/app/api/v1/helper"
	"dootask-okr/app/service"
	"dootask-okr/app/utils/common"
	"dootask-okr/web"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Init(c *gin.Context) {
	common.SetGlobalContext(c)
	//
	c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
	c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Token, X-Xsrf-Token, Language")
	c.Header("Access-Control-Allow-Credentials", "true")
	//
	urlPath := strings.Replace(c.Request.URL.Path, "/microapp/okr/api/v1/", "/api/v1/", -1)
	urlPath = strings.Replace(c.Request.URL.Path, "/manage/apps/okr/api/v1/", "/api/v1/", -1)
	urlPath = strings.Replace(c.Request.URL.Path, "/apps/okr/api/v1/", "/api/v1/", -1)
	// 接口
	if strings.HasPrefix(urlPath, "/api/v1/") {
		// 读取身份
		api := &v1.BaseApi{
			Route:   urlToApiRoute(urlPath[8:]),
			Token:   helper.Token(c), // 判断Token是否有效
			Context: c,
		}
		// 动态路由（不需要登录）
		if callApiMethod(api, false) {
			return
		}
		// 登录验证
		info, err := service.DootaskService.GetUserInfo(api.Token)
		if err != nil {
			helper.ErrorAuth(c, err.Error())
			return
		}
		api.Userinfo = info
		// 动态路由（需要登录）
		if callApiMethod(api, true) {
			return
		}
	}

	// 静态资源
	if strings.HasPrefix(urlPath, "/assets") || strings.HasPrefix(urlPath, "/apps/okr/assets") || strings.HasPrefix(urlPath, "/manage/apps/okr/assets") {
		assets := strings.Replace(urlPath, "/manage/apps/okr/assets", "/assets", -1)
		assets = strings.Replace(assets, "/apps/okr/assets", "/assets", -1)
		c.FileFromFS("dist"+assets, http.FS(web.Assets))
		return
	}
	if strings.HasPrefix(urlPath, "/apps/okr/src/assets") {
		assets := strings.Replace(urlPath, "/apps/okr/src/assets", "/assets", -1)
		c.FileFromFS("src"+assets, http.FS(web.SrcAssets))
		return
	}
	if strings.HasSuffix(urlPath, "/favicon.ico") {
		c.FileFromFS("/favicon.ico", http.FS(web.Favicon))
		return
	}
	// 页面输出
	c.HTML(http.StatusOK, "index", gin.H{
		"CODE": "",
		"MSG":  "",
	})
}

func urlToApiRoute(urlPath string) string {
	caser := cases.Title(language.Und)
	if strings.Contains(urlPath, "/") || strings.Contains(urlPath, "_") || strings.Contains(urlPath, "-") {
		urlPath = strings.ReplaceAll(urlPath, "/", " ")
		urlPath = strings.ReplaceAll(urlPath, "_", " ")
		urlPath = strings.ReplaceAll(urlPath, "-", " ")
		urlPath = strings.ReplaceAll(urlPath, ".", " ")
		urlPath = strings.ReplaceAll(caser.String(urlPath), " ", "")
	} else {
		urlPath = caser.String(urlPath)
	}
	return urlPath
}

func callApiMethod(api *v1.BaseApi, auth bool) bool {
	if api.Route == "" {
		return false
	}
	method := reflect.ValueOf(api).MethodByName(api.Route)
	if !auth {
		method = reflect.ValueOf(&v1.NotAuthBaseApi{BaseApi: *api}).MethodByName(api.Route)
	}
	if method.IsValid() {
		method.Call(nil)
		return true
	} else {
		return false
	}
}

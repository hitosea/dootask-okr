package router

import (
	v1 "dootask-okr/app/api/v1"
	"dootask-okr/app/api/v1/helper"
	"dootask-okr/app/service"
	"dootask-okr/config"
	"dootask-okr/web"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Init(c *gin.Context) {
	urlPath := c.Request.URL.Path
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

	// 开发模式 - 代理web
	if config.CONF.System.Mode == "debug" {
		err := godotenv.Load("./web/.env")
		if err == nil {
			CreatedProxy(c, "http://localhost:"+os.Getenv("DEV_PORT"))
			return
		}
	}

	// 静态资源
	if strings.HasPrefix(urlPath, "/assets") || strings.HasPrefix(urlPath, "/manage/microapp/okr/assets") {
		assets := strings.Replace(urlPath, "/manage/microapp/okr/assets", "/assets", -1)
		c.FileFromFS("dist"+assets, http.FS(web.Assets))
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

package v1

import (
	"dootask-okr/app/interfaces"

	"github.com/gin-gonic/gin"
)

type BaseApi struct {
	Route    string                   `mapstructure:"route"`
	Token    string                   `mapstructure:"token"`
	Userinfo *interfaces.UserInfoResp `mapstructure:"userinfo"`
	Context  *gin.Context             `mapstructure:"context"`
}

type NotAuthBaseApi struct {
	BaseApi BaseApi `mapstructure:"base_api"`
}

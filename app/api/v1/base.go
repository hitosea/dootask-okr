package v1

import (
	"github.com/gin-gonic/gin"
)

type BaseApi struct {
	Route   string       `mapstructure:"route"`
	Token   string       `mapstructure:"token"`
	Context *gin.Context `mapstructure:"context"`
}

type NotAuthBaseApi struct {
	BaseApi BaseApi `mapstructure:"base_api"`
}

package verify

import (
	"dootask-okr/app/api/v1/helper"
	"dootask-okr/app/constant"
	"dootask-okr/app/utils/logger"

	"github.com/gin-gonic/gin"
)

// verifyUtil 参数验证工具类
type verifyUtil struct{}

var VerifyUtil = verifyUtil{}

func (vu verifyUtil) ShouldBindAll(c *gin.Context, obj any) any {
	if err := c.ShouldBind(obj); err != nil {
		logger.Error(err.Error())
		helper.ErrorWith(c, constant.ErrInvalidParameter, nil)
		panic(nil)
	}
	if err := c.ShouldBindJSON(obj); err != nil && err.Error() != "EOF" {
		logger.Error(err.Error())
		helper.ErrorWith(c, constant.ErrInvalidParameter, err)
		panic(nil)
	}
	return obj
}

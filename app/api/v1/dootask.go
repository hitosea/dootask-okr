package v1

import (
	"dootask-okr/app/api/v1/helper"
	"dootask-okr/app/service"
)

// @Tags Dootask
// @Summary 获取用户信息
// @Description 获取用户信息
// @Accept json
// @Param token query string true "token"
// @Success 200 {object} interfaces.Response
// @Router /doo/user/info [get]
func (api *BaseApi) DooUserInfo() {
	token := api.Context.Query("token")
	result, err := service.DootaskService.GetUserInfo(token)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

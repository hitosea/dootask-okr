package v1

import (
	"dootask-okr/app/api/v1/helper"
	"dootask-okr/app/service"
)

// @Tags Dootask
// @Summary 获取用户信息
// @Description 获取用户信息
// @Accept json
// @Success 200 {object} interfaces.Response
// @Router /doo/user/info [get]
func (api *BaseApi) DooUserInfo() {
	result, err := service.DootaskService.GetUserInfo(api.Userinfo.Token)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

// @Tags Dootask
// @Summary 获取用户列表
// @Description 获取用户列表
// @Accept json
// @Success 200 {object} interfaces.Response
// @Router /doo/user/list [get]
func (api *BaseApi) DooUserList() {
	result, err := service.DootaskService.GetUserList(api.Userinfo.Token)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

// @Tags Dootask
// @Summary 获取项目列表
// @Description 获取项目列表
// @Accept json
// @Success 200 {object} interfaces.Response
// @Router /doo/project/list [get]
func (api *BaseApi) DooProjectList() {
	result, err := service.DootaskService.GetProjectList(api.Userinfo.Token)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

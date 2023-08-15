package v1

import (
	"dootask-okr/app/api/v1/helper"
	"dootask-okr/app/interfaces"
	"dootask-okr/app/service"
	"dootask-okr/app/utils/verify"
	"strconv"
)

// @Tags Dootask
// @Summary 获取用户信息
// @Description 获取用户信息
// @Accept json
// @Success 200 {object} interfaces.Response
// @Router /okr/user/info [get]
func (api *BaseApi) OkrUserInfo() {
	result, err := service.DootaskService.GetUserInfo(api.Userinfo.Token)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

// @Tags Dootask
// @Summary 获取指定会员基础信息
// @Description 获取指定会员基础信息
// @Accept json
// @Param userid query number true "用户id"
// @Success 200 {object} interfaces.Response
// @Router /okr/user/basic [get]
func (api *BaseApi) OkrUserBasic() {
	userid, err := strconv.Atoi(api.Context.Query("userid"))
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	result, err := service.DootaskService.GetUserBasic(api.Userinfo.Token, []int{userid})
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result[0])
}

// @Tags Dootask
// @Summary 获取用户列表
// @Description 获取用户列表
// @Accept json
// @Param request query interfaces.OkrUserListReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/user/list [get]
func (api *BaseApi) OkrUserList() {
	var param = interfaces.OkrUserListReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	result, err := service.DootaskService.GetUserList(api.Userinfo, param.Keyword, param.DeptOnly, param.Page, param.PageSize)
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
// @Param request query interfaces.Pages true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/project/list [get]
func (api *BaseApi) OkrProjectList() {
	var param = interfaces.Pages{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	result, err := service.DootaskService.GetProjectList(api.Userinfo.Token, param.Page, param.PageSize)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

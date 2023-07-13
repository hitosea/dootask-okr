package v1

import (
	"dootask-okr/app/api/v1/helper"
	"dootask-okr/app/interfaces"
	"dootask-okr/app/service"
)

// @Tags Okr
// @Summary 创建OKR
// @Description 创建OKR
// @Accept json
// @Param request body interfaces.OkrCreateReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/create [post]
func (api *BaseApi) OkrCreate() {
	var param = interfaces.OkrCreateReq{}
	if err := api.Context.ShouldBindJSON(&param); err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	result, err := service.OkrService.Create(api.Userinfo, param)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

// @Tags Okr
// @Summary 更新OKR
// @Description 更新OKR
// @Accept json
// @Param request body interfaces.OkrUpdateReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/update [post]
func (api *BaseApi) OkrUpdate() {
	var param = interfaces.OkrUpdateReq{}
	if err := api.Context.ShouldBindJSON(&param); err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	result, err := service.OkrService.Update(param)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

// @Tags Okr
// @Summary 我的OKR列表
// @Description 我的OKR列表
// @Accept json
// @Param objective query string true "目标"
// @Success 200 {object} interfaces.Response
// @Router /okr/my/list [get]
func (api *BaseApi) OkrMyList() {
	objective := api.Context.Query("objective")
	result, err := service.OkrService.GetMyList(api.Userinfo, objective)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

// @Tags Okr
// @Summary 参与的OKR列表
// @Description 参与的OKR列表
// @Accept json
// @Param objective query string true "目标"
// @Success 200 {object} interfaces.Response
// @Router /okr/participant/list [get]
func (api *BaseApi) OkrParticipantList() {
	objective := api.Context.Query("objective")
	result, err := service.OkrService.GetParticipantList(api.Userinfo, objective)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

// @Tags Okr
// @Summary 部门OKR列表
// @Description 部门OKR列表
// @Accept json
// @Param request body interfaces.OkrDepartmentListReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/department/list [get]
func (api *BaseApi) OkrDepartmentList() {
	var param = interfaces.OkrDepartmentListReq{}
	if err := api.Context.ShouldBindJSON(&param); err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	result, err := service.OkrService.GetDepartmentList(api.Userinfo, param)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

// @Tags Okr
// @Summary 关注的OKR列表
// @Description 关注的OKR列表
// @Accept json
// @Param objective query string true "目标"
// @Success 200 {object} interfaces.Response
// @Router /okr/follow/list [get]
func (api *BaseApi) OkrFollowList() {
	objective := api.Context.Query("objective")
	result, err := service.OkrService.GetFollowList(api.Userinfo, objective)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

// @Tags Okr
// @Summary OKR复盘列表
// @Description OKR复盘列表
// @Accept json
// @Param objective query string true "目标"
// @Success 200 {object} interfaces.Response
// @Router /okr/replay/list [get]
func (api *BaseApi) OkrReplayList() {
	objective := api.Context.Query("objective")
	result, err := service.OkrService.GetReplayList(api.Userinfo, objective)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

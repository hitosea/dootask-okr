package v1

import (
	"dootask-okr/app/api/v1/helper"
	"dootask-okr/app/interfaces"
	"dootask-okr/app/service"
	"strconv"
)

// @Tags Okr
// @Summary 创建OKR
// @Description 创建OKR
// @Accept json
// @Param request formData interfaces.OkrCreateReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/create [post]
func (api *BaseApi) OkrCreate() {
	var param = interfaces.OkrCreateReq{}
	if err := api.Context.ShouldBindJSON(&param); err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	userid := 1
	result, err := service.OkrService.CreateObjective(userid, param)
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
// @Param request formData interfaces.OkrUpdateReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/update [post]
func (api *BaseApi) OkrUpdate() {
	var param = interfaces.OkrUpdateReq{}
	if err := api.Context.ShouldBindJSON(&param); err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	userid := 1
	result, err := service.OkrService.UpdateObjective(userid, param)
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
// @Param userid query number true "用户id"
// @Param objective query string true "目标"
// @Success 200 {object} interfaces.Response
// @Router /okr/my/list [get]
func (api *BaseApi) OkrMyList() {
	userid, err := strconv.Atoi(api.Context.Query("userid"))
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	objective := api.Context.Query("objective")
	result, err := service.OkrService.GetMyObjectivesWithKeyResults(userid, objective)
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
// @Param userid query number true "用户id"
// @Param objective query string true "目标"
// @Success 200 {object} interfaces.Response
// @Router /okr/participant/list [get]
func (api *BaseApi) OkrParticipantList() {
	userid, err := strconv.Atoi(api.Context.Query("userid"))
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	objective := api.Context.Query("objective")
	result, err := service.OkrService.GetParticipantObjectivesWithKeyResults(userid, objective)
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
// @Param userid query number true "用户id"
// @Param objective query string true "目标"
// @Success 200 {object} interfaces.Response
// @Router /okr/follow/list [get]
func (api *BaseApi) OkrFollowList() {
	userid, err := strconv.Atoi(api.Context.Query("userid"))
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	objective := api.Context.Query("objective")
	result, err := service.OkrService.GetFollowObjectives(userid, objective)
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
// @Param objective_id query number true "目标id"
// @Success 200 {object} interfaces.Response
// @Router /okr/replay/list [get]
func (api *BaseApi) OkrReplayList() {
	objectiveId, err := strconv.Atoi(api.Context.Query("objective_id"))
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	result, err := service.OkrService.GetReplays(objectiveId)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

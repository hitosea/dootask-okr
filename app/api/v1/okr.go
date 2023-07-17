package v1

import (
	"dootask-okr/app/api/v1/helper"
	"dootask-okr/app/interfaces"
	"dootask-okr/app/service"
	"dootask-okr/app/utils/verify"
	"strconv"
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
	result, err := service.OkrService.Update(api.Userinfo, param)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

// @Tags Okr
// @Summary 获取我的OKR列表
// @Description 获取我的OKR列表
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
// @Summary 获取参与的OKR列表
// @Description 获取参与的OKR列表
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
// @Summary 获取部门OKR列表
// @Description 获取部门OKR列表
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
// @Summary 获取关注的OKR列表
// @Description 获取关注的OKR列表
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
// @Summary 获取OKR复盘列表
// @Description 获取OKR复盘列表
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

// @Tags Okr
// @Summary 获取复盘列表by目标id
// @Description 获取复盘列表by目标id
// @Accept json
// @Param id query int true "目标id"
// @Success 200 {object} interfaces.Response
// @Router /okr/replay/okr/list [get]
func (api *BaseApi) OkrReplayOkrList() {
	idStr := api.Context.Query("id")
	if idStr == "" {
		helper.ErrorWith(api.Context, "目标id不能为空", nil)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	result, err := service.OkrService.GetReplayListByOkrId(id)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

// @Tags Okr
// @Summary 获取OKR详情
// @Description 获取OKR详情
// @Accept json
// @Param id query int true "目标id"
// @Success 200 {object} interfaces.Response
// @Router /okr/detail [get]
func (api *BaseApi) OkrDetail() {
	idStr := api.Context.Query("id")
	if idStr == "" {
		helper.ErrorWith(api.Context, "目标id不能为空", nil)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	result, err := service.OkrService.GetOkrDetail(api.Userinfo, id)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

// @Tags Okr
// @Summary 关注或取消关注目标
// @Description 关注或取消关注目标
// @Accept json
// @Param id query int true "目标id"
// @Success 200 {object} interfaces.Response
// @Router /okr/follow [get]
func (api *BaseApi) OkrFollow() {
	objectiveIdStr := api.Context.Query("id")
	if objectiveIdStr == "" {
		helper.ErrorWith(api.Context, "关注目标不能为空", nil)
		return
	}

	objectiveId, err := strconv.Atoi(objectiveIdStr)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	result, err := service.OkrService.FollowObjective(api.Userinfo.Userid, objectiveId)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, result)
}

// @Tags Okr
// @Summary 更新进度和进度状态
// @Description 更新进度和进度状态
// @Accept json
// @Param request formData interfaces.OkrUpdateProgressReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/update/progress [post]
func (api *BaseApi) OkrUpdateProgress() {
	var param = interfaces.OkrUpdateProgressReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	err := service.OkrService.UpdateProgressAndStatus(api.Userinfo, param)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, nil)
}

// @Tags Okr
// @Summary OKR评分
// @Description OKR评分
// @Accept json
// @Param request formData interfaces.OkrScoreReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/score [post]
func (api *BaseApi) OkrScore() {
	var param = interfaces.OkrScoreReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	err := service.OkrService.UpdateScore(api.Userinfo, param)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, nil)
}

// @Tags Okr
// @Summary 结束/重启目标
// @Description 结束/重启目标
// @Accept json
// @Param request formData interfaces.OkrFinishReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/finish [post]
func (api *BaseApi) OkrFinish() {
	var param = interfaces.OkrFinishReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	err := service.OkrService.FinishObjective(param)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, nil)
}

// @Tags Okr
// @Summary 更新参与人
// @Description 更新参与人
// @Accept json
// @Param request formData interfaces.OkrParticipantUpdateReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/participant/update [post]
func (api *BaseApi) OkrParticipantUpdate() {
	var param = interfaces.OkrParticipantUpdateReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	err := service.OkrService.UpdateParticipant(param)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, nil)
}

// @Tags Okr
// @Summary 更新信心指数
// @Description 更新信心指数
// @Accept json
// @Param request formData interfaces.OkrConfidenceUpdateReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/confidence/update [post]
func (api *BaseApi) OkrConfidenceUpdate() {
	var param = interfaces.OkrConfidenceUpdateReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	err := service.OkrService.UpdateConfidence(param)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, nil)
}

// @Tags Okr
// @Summary 添加复盘
// @Description 添加复盘
// @Accept json
// @Param request body interfaces.OkrReplayCreateReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/replay/create [post]
func (api *BaseApi) OkrReplayCreate() {
	var param = interfaces.OkrReplayCreateReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	err := service.OkrService.CreateOkrReplay(api.Userinfo.Userid, param)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, nil)
}

// @Tags Okr
// @Summary 复盘详情
// @Description 复盘详情
// @Accept json
// @Param id query int true "复盘id"
// @Success 200 {object} interfaces.Response
// @Router /okr/replay/detail [get]
func (api *BaseApi) OkrReplayDetail() {
	idStr := api.Context.Query("id")
	if idStr == "" {
		helper.ErrorWith(api.Context, "复盘id不能为空", nil)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	result, err := service.OkrService.GetReplayDetail(id)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, result)
}

// @Tags Okr
// @Summary 取消对齐目标
// @Description 取消对齐目标
// @Accept json
// @Param id query int true "目标id"
// @Success 200 {object} interfaces.Response
// @Router /okr/align/cancel [get]
func (api *BaseApi) OkrAlignCancel() {
	objectiveIdStr := api.Context.Query("id")
	if objectiveIdStr == "" {
		helper.ErrorWith(api.Context, "目标id不能为空", nil)
		return
	}

	objectiveId, err := strconv.Atoi(objectiveIdStr)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	err = service.OkrService.CancelAlignObjective(api.Userinfo.Userid, objectiveId)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, nil)
}

// @Tags Okr
// @Summary 获取对齐目标列表
// @Description 获取对齐目标列表
// @Accept json
// @Success 200 {object} interfaces.Response
// @Router /okr/align/list [get]
func (api *BaseApi) OkrAlignList() {
	result, err := service.OkrService.GetAlignList(api.Userinfo)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, result)
}

// @Tags Okr
// @Summary 获取动态列表by目标id
// @Description 获取动态列表by目标id
// @Accept json
// @Param id query int true "目标id"
// @Success 200 {object} interfaces.Response
// @Router /okr/log/list [get]
func (api *BaseApi) OkrLogList() {
	okrIdStr := api.Context.Query("id")
	if okrIdStr == "" {
		helper.ErrorWith(api.Context, "目标id不能为空", nil)
		return
	}

	okrId, err := strconv.Atoi(okrIdStr)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	result, err := service.OkrService.GetOkrLogList(okrId)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, result)
}

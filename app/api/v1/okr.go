package v1

import (
	"dootask-okr/app/api/v1/helper"
	"dootask-okr/app/interfaces"
	"dootask-okr/app/service"
	"dootask-okr/app/utils/verify"
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
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
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
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
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
// @Param request query interfaces.OkrDepartmentListReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/department/list [get]
func (api *BaseApi) OkrDepartmentList() {
	var param = interfaces.OkrDepartmentListReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
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
// @Param request query interfaces.OkrIdReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/replay/okr/list [get]
func (api *BaseApi) OkrReplayOkrList() {
	var param = interfaces.OkrIdReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	result, err := service.OkrService.GetReplayListByOkrId(param.Id)
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
// @Param request query interfaces.OkrIdReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/detail [get]
func (api *BaseApi) OkrDetail() {
	var param = interfaces.OkrIdReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	result, err := service.OkrService.GetOkrDetail(api.Userinfo, param.Id)
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
// @Param request query interfaces.OkrIdReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/follow [get]
func (api *BaseApi) OkrFollow() {
	var param = interfaces.OkrIdReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	result, err := service.OkrService.FollowObjective(api.Userinfo.Userid, param.Id)
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
// @Summary 取消/重启目标
// @Description 取消/重启目标
// @Accept json
// @Param request formData interfaces.OkrCanceledReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/cancel [post]
func (api *BaseApi) OkrCancel() {
	var param = interfaces.OkrCanceledReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	err := service.OkrService.CancelObjective(param)
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
// @Param request query interfaces.OkrReplayIdReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/replay/detail [get]
func (api *BaseApi) OkrReplayDetail() {
	var param = interfaces.OkrReplayIdReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	result, err := service.OkrService.GetReplayDetail(param.Id)
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
// @Param request query interfaces.OkrAlignCancelReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/align/cancel [get]
func (api *BaseApi) OkrAlignCancel() {
	var param = interfaces.OkrAlignCancelReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	err := service.OkrService.CancelAlignObjective(param.OkrId, param.AlignOkrId)
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
// @Param keywords query string true "关键词"
// @Success 200 {object} interfaces.Response
// @Router /okr/align/list [get]
func (api *BaseApi) OkrAlignList() {
	keywords := api.Context.Query("keywords")
	result, err := service.OkrService.GetAlignList(api.Userinfo, keywords)
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
// @Param request query interfaces.OkrIdReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/log/list [get]
func (api *BaseApi) OkrLogList() {
	var param = interfaces.OkrIdReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	result, err := service.OkrService.GetOkrLogList(param.Id)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, result)
}

// @Tags Okr
// @Summary 获取对齐目标by目标id
// @Description 获取对齐目标by目标id
// @Accept json
// @Param request query interfaces.OkrIdReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/align/detail [get]
func (api *BaseApi) OkrAlignDetail() {
	var param = interfaces.OkrIdReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	result, err := service.OkrService.GetAlignListByOkrId(api.Userinfo, param.Id)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, result)
}

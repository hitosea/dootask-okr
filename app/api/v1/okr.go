package v1

import (
	"dootask-okr/app/api/v1/helper"
	"dootask-okr/app/constant"
	"dootask-okr/app/interfaces"
	"dootask-okr/app/service"
	"dootask-okr/app/utils/common"
	"dootask-okr/app/utils/verify"
)

// @Tags Okr
// @Summary 创建OKR
// @Description 创建OKR
// @Accept json
// @Param request body interfaces.OkrCreateReq true "request"
// @Success 200 {object} interfaces.Response{data=model.Okr}
// @Router /okr/create [post]
func (api *BaseApi) OkrCreate() {
	var param = interfaces.OkrCreateReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	// 类型
	if !common.InArrayInt(param.Type, []int{1, 2}) {
		helper.ErrorWith(api.Context, constant.ErrOkrTypeInvalid, nil)
		return
	}
	// 优先级
	if !common.InArray(param.Priority, []string{"P0", "P1", "P2"}) {
		helper.ErrorWith(api.Context, constant.ErrOkrPriorityInvalid, nil)
		return
	}
	// 归属
	if !common.InArrayInt(param.Ascription, []int{1, 2}) {
		helper.ErrorWith(api.Context, constant.ErrOkrAscriptionInvalid, nil)
		return
	}
	// 可见范围
	if param.VisibleRange > 0 && !common.InArrayInt(param.VisibleRange, []int{1, 2, 3}) {
		helper.ErrorWith(api.Context, constant.ErrOkrVisibleRangeInvalid, nil)
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
// @Success 200 {object} interfaces.Response{data=model.Okr}
// @Router /okr/update [post]
func (api *BaseApi) OkrUpdate() {
	var param = interfaces.OkrUpdateReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	// 类型
	if !common.InArrayInt(param.Type, []int{1, 2}) {
		helper.ErrorWith(api.Context, constant.ErrOkrTypeInvalid, nil)
		return
	}
	// 优先级
	if !common.InArray(param.Priority, []string{"P0", "P1", "P2"}) {
		helper.ErrorWith(api.Context, constant.ErrOkrPriorityInvalid, nil)
		return
	}
	// 可见范围
	if param.VisibleRange > 0 && !common.InArrayInt(param.VisibleRange, []int{1, 2, 3}) {
		helper.ErrorWith(api.Context, constant.ErrOkrVisibleRangeInvalid, nil)
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
// @Param request query interfaces.OkrListBaseReq true "request"
// @Success 200 {object} interfaces.Response{data=interfaces.Pagination{data=[]interfaces.OkrResp}}
// @Router /okr/my/list [get]
func (api *BaseApi) OkrMyList() {
	var param = interfaces.OkrListBaseReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	result, err := service.OkrService.GetMyList(api.Userinfo, param.Objective, param.Page, param.PageSize)
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
// @Param request query interfaces.OkrListBaseReq true "request"
// @Success 200 {object} interfaces.Response{data=interfaces.Pagination{data=[]interfaces.OkrResp}}
// @Router /okr/participant/list [get]
func (api *BaseApi) OkrParticipantList() {
	var param = interfaces.OkrListBaseReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	result, err := service.OkrService.GetParticipantList(api.Userinfo, param.Objective, param.Page, param.PageSize)
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
// @Success 200 {object} interfaces.Response{data=interfaces.Pagination{data=[]interfaces.OkrResp}}
// @Router /okr/department/list [get]
func (api *BaseApi) OkrDepartmentList() {
	var param = interfaces.OkrDepartmentListReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	// 类型
	if param.Type > 0 && !common.InArrayInt(param.Type, []int{1, 2}) {
		helper.ErrorWith(api.Context, constant.ErrOkrTypeInvalid, nil)
		return
	}
	// 是否已完成未评分
	if !common.InArrayInt(param.Completed, []int{0, 1}) {
		helper.ErrorWith(api.Context, constant.ErrOkrIsFinishNotScoreInvalid, nil)
		return
	}
	result, err := service.OkrService.GetDepartmentList(api.Userinfo, param, param.Page, param.PageSize)
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
// @Param request query interfaces.OkrListBaseReq true "request"
// @Success 200 {object} interfaces.Response{data=interfaces.Pagination{data=[]interfaces.OkrResp}}
// @Router /okr/follow/list [get]
func (api *BaseApi) OkrFollowList() {
	var param = interfaces.OkrListBaseReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	result, err := service.OkrService.GetFollowList(api.Userinfo, param.Objective, param.Page, param.PageSize)
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
// @Param request query interfaces.OkrListBaseReq true "request"
// @Success 200 {object} interfaces.Response{data=interfaces.Pagination{data=[]model.OkrReplay}}
// @Router /okr/replay/list [get]
func (api *BaseApi) OkrReplayList() {
	var param = interfaces.OkrListBaseReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	result, err := service.OkrService.GetReplayList(api.Userinfo, param.Objective, param.Page, param.PageSize)
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
// @Param request query interfaces.OkrIdListReq true "request"
// @Success 200 {object} interfaces.Response{data=interfaces.Pagination{data=[]model.OkrReplay}}
// @Router /okr/replay/okr/list [get]
func (api *BaseApi) OkrReplayOkrList() {
	var param = interfaces.OkrIdListReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	result, err := service.OkrService.GetReplayListByOkrId(param.Id, param.Page, param.PageSize)
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
// @Success 200 {object} interfaces.Response{data=interfaces.OkrResp}
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
// @Success 200 {object} interfaces.Response{data=model.OkrFollow}
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
// @Success 200 {object} interfaces.Response{data=model.Okr}
// @Router /okr/update/progress [post]
func (api *BaseApi) OkrUpdateProgress() {
	var param = interfaces.OkrUpdateProgressReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	// 进度
	if param.Progress < 1 || param.Progress > 100 {
		helper.ErrorWith(api.Context, constant.ErrOkrProgressInvalid, nil)
		return
	}
	// 进度状态
	if !common.InArrayInt(param.Status, []int{0, 1, 2, 3}) {
		helper.ErrorWith(api.Context, constant.ErrOkrProgressStatusInvalid, nil)
		return
	}
	result, err := service.OkrService.UpdateProgressAndStatus(api.Userinfo, param)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, result)
}

// @Tags Okr
// @Summary OKR评分
// @Description OKR评分
// @Accept json
// @Param request formData interfaces.OkrScoreReq true "request"
// @Success 200 {object} interfaces.Response{data=model.Okr}
// @Router /okr/score [post]
func (api *BaseApi) OkrScore() {
	var param = interfaces.OkrScoreReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	// 评分
	if param.Score < 0 || param.Score > 10 {
		helper.ErrorWith(api.Context, constant.ErrOkrScoreInvalid, nil)
		return
	}
	result, err := service.OkrService.UpdateScore(api.Userinfo, param)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, result)
}

// @Tags Okr
// @Summary 取消/重启目标
// @Description 取消/重启目标
// @Accept json
// @Param request query interfaces.OkrIdReq true "request"
// @Success 200 {object} interfaces.Response{data=model.Okr}
// @Router /okr/cancel [get]
func (api *BaseApi) OkrCancel() {
	var param = interfaces.OkrIdReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	result, err := service.OkrService.CancelObjective(param.Id)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, result)
}

// @Tags Okr
// @Summary 更新参与人
// @Description 更新参与人
// @Accept json
// @Param request formData interfaces.OkrParticipantUpdateReq true "request"
// @Success 200 {object} interfaces.Response{data=model.Okr}
// @Router /okr/participant/update [post]
func (api *BaseApi) OkrParticipantUpdate() {
	var param = interfaces.OkrParticipantUpdateReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	result, err := service.OkrService.UpdateParticipant(param)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, result)
}

// @Tags Okr
// @Summary 更新信心指数
// @Description 更新信心指数
// @Accept json
// @Param request formData interfaces.OkrConfidenceUpdateReq true "request"
// @Success 200 {object} interfaces.Response{data=model.Okr}
// @Router /okr/confidence/update [post]
func (api *BaseApi) OkrConfidenceUpdate() {
	var param = interfaces.OkrConfidenceUpdateReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	// 信心指数
	if param.Confidence < 0 || param.Confidence > 100 {
		helper.ErrorWith(api.Context, constant.ErrOkrConfidenceInvalid, nil)
		return
	}
	result, err := service.OkrService.UpdateConfidence(param)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, result)
}

// @Tags Okr
// @Summary 添加复盘
// @Description 添加复盘
// @Accept json
// @Param request body interfaces.OkrReplayCreateReq true "request"
// @Success 200 {object} interfaces.Response{data=model.OkrReplay}
// @Router /okr/replay/create [post]
func (api *BaseApi) OkrReplayCreate() {
	var param = interfaces.OkrReplayCreateReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	result, err := service.OkrService.CreateOkrReplay(api.Userinfo.Userid, param)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, result)
}

// @Tags Okr
// @Summary 复盘详情
// @Description 复盘详情
// @Accept json
// @Param request query interfaces.OkrReplayIdReq true "request"
// @Success 200 {object} interfaces.Response{data=model.OkrReplay}
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
// @Param request query interfaces.OkrListBaseReq true "request"
// @Success 200 {object} interfaces.Response{data=interfaces.Pagination{data=[]model.Okr}}
// @Router /okr/align/list [get]
func (api *BaseApi) OkrAlignList() {
	var param = interfaces.OkrListBaseReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	result, err := service.OkrService.GetAlignList(api.Userinfo, param.Objective, param.Page, param.PageSize)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, result)
}

// @Tags Okr
// @Summary 更新对齐目标
// @Description 更新对齐目标
// @Accept json
// @Param request formData interfaces.OkrAlignUpdateReq true "request"
// @Success 200 {object} interfaces.Response{data=model.Okr}
// @Router /okr/align/update [post]
func (api *BaseApi) OkrAlignUpdate() {
	var param = interfaces.OkrAlignUpdateReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	result, err := service.OkrService.UpdateAlignObjective(param.Id, param.AlignObjective)
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
// @Param request query interfaces.OkrIdListReq true "request"
// @Success 200 {object} interfaces.Response{data=interfaces.Pagination{data=[]model.OkrLog}}
// @Router /okr/log/list [get]
func (api *BaseApi) OkrLogList() {
	var param = interfaces.OkrIdListReq{}
	verify.VerifyUtil.ShouldBindAll(api.Context, &param)
	result, err := service.OkrService.GetOkrLogList(param.Id, param.Page, param.PageSize)
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
// @Success 200 {object} interfaces.Response{data=[]interfaces.OkrAlignResp}
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

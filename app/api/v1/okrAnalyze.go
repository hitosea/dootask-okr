package v1

import (
	"dootask-okr/app/api/v1/helper"
	"dootask-okr/app/constant"
	"dootask-okr/app/service"
)

// @Tags OkrAnalyze
// @Summary OKR整体平均完成度
// @Description OKR整体平均完成度
// @Accept json
// @Success 200 {object} interfaces.Response{data=interfaces.OkrAnalyzeOverall}
// @Router /okr/analyze/complete [get]
func (api *BaseApi) OkrAnalyzeComplete() {
	if !api.Userinfo.IsAdmin() {
		helper.ErrorWith(api.Context, constant.ErrNoPermission, nil)
		return
	}
	result, err := service.OkrAnalyzeService.GetOverallCompleteness(api.Userinfo)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

// @Tags OkrAnalyze
// @Summary OKR各部门平均完成度
// @Description OKR各部门平均完成度
// @Accept json
// @Success 200 {object} interfaces.Response{data=[]interfaces.OkrAnalyzeDept}
// @Router /okr/analyze/dept/complete [get]
func (api *BaseApi) OkrAnalyzeDeptComplete() {
	if !api.Userinfo.IsAdmin() {
		helper.ErrorWith(api.Context, constant.ErrNoPermission, nil)
		return
	}
	result, err := service.OkrAnalyzeService.GetDeptCompleteness(api.Userinfo)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

// @Tags OkrAnalyze
// @Summary OKR评分分布
// @Description OKR评分分布
// @Accept json
// @Success 200 {object} interfaces.Response{data=interfaces.OkrAnalyzeScore}
// @Router /okr/analyze/score [get]
func (api *BaseApi) OkrAnalyzeScore() {
	if !api.Userinfo.IsAdmin() {
		helper.ErrorWith(api.Context, constant.ErrNoPermission, nil)
		return
	}
	result, err := service.OkrAnalyzeService.GetScore(api.Userinfo)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

// @Tags OkrAnalyze
// @Summary OKR各部门评分分布
// @Description OKR各部门评分分布
// @Accept json
// @Success 200 {object} interfaces.Response{data=interfaces.OkrAnalyzeScoreDept}
// @Router /okr/analyze/dept/score [get]
func (api *BaseApi) OkrAnalyzeDeptScore() {
	if !api.Userinfo.IsAdmin() {
		helper.ErrorWith(api.Context, constant.ErrNoPermission, nil)
		return
	}
	result, err := service.OkrAnalyzeService.GetDeptScore(api.Userinfo)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

// @Tags OkrAnalyze
// @Summary OKR人员评分率
// @Description OKR人员评分率
// @Accept json
// @Success 200 {object} interfaces.Response{data=interfaces.OkrAnalyzePersonnelScoreRate}
// @Router /okr/analyze/personnel/score/rate [get]
func (api *BaseApi) OkrAnalyzePersonnelScoreRate() {
	if !api.Userinfo.IsAdmin() {
		helper.ErrorWith(api.Context, constant.ErrNoPermission, nil)
		return
	}
	result, err := service.OkrAnalyzeService.GetPersonnelScoreRate(api.Userinfo)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

// @Tags OkrAnalyze
// @Summary OKR各部门评分占比
// @Description OKR部门评分占比
// @Accept json
// @Success 200 {object} interfaces.Response{data=[]interfaces.OkrAnalyzeDeptScoreProportion}
// @Router /okr/analyze/dept/score/proportion [get]
func (api *BaseApi) OkrAnalyzeDeptScoreProportion() {
	if !api.Userinfo.IsAdmin() {
		helper.ErrorWith(api.Context, constant.ErrNoPermission, nil)
		return
	}
	result, err := service.OkrAnalyzeService.GetDeptScoreProportion(api.Userinfo)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

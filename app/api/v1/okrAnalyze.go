package v1

import (
	"dootask-okr/app/api/v1/helper"
	"dootask-okr/app/service"
)

// @Tags OkrAnalyze
// @Summary OKR整体平均完成度
// @Description OKR整体平均完成度
// @Accept json
// @Success 200 {object} interfaces.Response{data=interfaces.OkrAnalyzeOverall}
// @Router /okr/analyze/complete [get]
func (api *BaseApi) OkrAnalyzeComplete() {
	result, err := service.OkrAnalyzeService.GetOverallCompleteness(api.Userinfo.Userid)
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
	result, err := service.OkrAnalyzeService.GetDeptCompleteness(api.Userinfo.Userid)
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
	result, err := service.OkrAnalyzeService.GetScore(api.Userinfo.Userid)
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
	result, err := service.OkrAnalyzeService.GetDeptScore(api.Userinfo.Userid)
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
// @Success 200 {object} interfaces.Response{data=interfaces.OkrAnalyzePersonnel}
// @Router /okr/analyze/personnel [get]
func (api *BaseApi) OkrAnalyzePersonnel() {
	result, err := service.OkrAnalyzeService.GetPersonnel(api.Userinfo.Userid)
	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	helper.Success(api.Context, result)
}

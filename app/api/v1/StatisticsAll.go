package v1

import (
	"dootask-okr/app/api/v1/helper"
	"dootask-okr/app/service"
	"fmt"
)

// @Tags StatisticsAll
// @Summary 获取/未完成、已完成/已取消目标数量
// @Description 获取/未完成、已完成/已取消目标数量
// @Accept json
// @Success 200 {object} interfaces.Response
// @Router /okr/statistics/all [get]

func (api *BaseApi) OkrStatisticsAll() {
	//调用逻辑处理，传当前用户id
	result ,err := service.StatisticsAll(api.Userinfo.Userid)

	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}

	helper.Success(api.Context, result)
}


// @Tags Okr
// @Summary 获取个人OKR整体完成度和与评分和
// @Description 查询个人OKR整体完成度和评分
// @Accept json
// @Param userid query int true "用户id"
// @Success 200 {object} interfaces.Response
// @Router /okr/overall [get]

func (api *BaseApi) OkrOverall() {
	result, err := service.OkrService.GetOkrOverall(api.Userinfo.Userid)

	if err != nil {
		helper.ErrorWith(api.Context, err.Error(), nil)
		return
	}
	fmt.Println(result)
	helper.Success(api.Context, result)
}
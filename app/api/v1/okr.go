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
// @Param request formData interfaces.OkrCreateReq true "request"
// @Success 200 {object} interfaces.Response
// @Router /okr/create [post]
func (api *BaseApi) Create() {
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
func (api *BaseApi) Update() {
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

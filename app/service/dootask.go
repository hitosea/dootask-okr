package service

import (
	"dootask-okr/app/constant"
	"dootask-okr/app/core"
	"dootask-okr/app/interfaces"
	"dootask-okr/app/model"
	"dootask-okr/app/utils/common"
	e "dootask-okr/app/utils/error"
	"dootask-okr/config"
	"encoding/json"

	"fmt"
	"time"
)

var DootaskService = newDootaskService()

type dootaskService struct {
	client *common.HTTPClient
}

func newDootaskService() *dootaskService {
	return &dootaskService{
		client: common.NewHTTPClient(5 * time.Second),
	}
}

// 获取用户的信息
func (s dootaskService) GetUserInfo(token string) (*interfaces.UserInfoResp, error) {
	url := fmt.Sprintf("%s%s?token=%s", config.DooTaskUrl, "/api/users/info", token)
	result, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}
	info, err := s.UnmarshalAndCheckResponse(result)
	if err != nil {
		return nil, err
	}
	userInfo := new(interfaces.UserInfoResp)
	if err := common.MapToStruct(info, userInfo); err != nil {
		return nil, err
	}
	return userInfo, nil
}

// 获取指定会员基础信息
func (s dootaskService) GetUserBasic(token string, userid []int) ([]*interfaces.UserBasicResp, error) {
	domain := common.GetRequestOrigin()
	url := fmt.Sprintf("%s%s?userid=%v&token=%s", config.DooTaskUrl, "/api/users/basic", common.StructToJson(userid), token)
	result, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}
	info, err := s.UnmarshalAndCheckResponse(result)
	if err != nil {
		return nil, err
	}

	list, ok := info["list"].([]interface{})
	if !ok {
		return nil, e.New(constant.ErrDooTaskDataFormat)
	}

	var userBasicArr []*interfaces.UserBasicResp
	if len(list) > 0 {
		for _, v := range list {
			var userBasic interfaces.UserBasicResp
			if err := common.MapToStruct(v.(map[string]interface{}), &userBasic); err != nil {
				return nil, err
			}
			// 替换成当前请求源域名
			userBasic.Userimg = common.ReplaceDomainPath(domain, userBasic.Userimg)
			userBasicArr = append(userBasicArr, &userBasic)
		}
	}
	return userBasicArr, nil
}

// 获取用户列表
func (s dootaskService) GetUserList(user *interfaces.UserInfoResp, deptOnly bool, page, pageSize int) (*interfaces.Pagination, error) {
	var departments string
	if deptOnly {
		departments = common.ArrayImplode(user.Department)
	}
	url := fmt.Sprintf("%s%s?keys[departments]=%s&page=%d&pagesize=%d&token=%s", config.DooTaskUrl, "/api/users/search", departments, page, pageSize, user.Token)
	result, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}

	data, err := s.UnmarshalAndCheckResponse(result)
	if err != nil {
		return nil, err
	}
	count, ok := data["total"].(float64)
	if !ok {
		return nil, e.New(constant.ErrDooTaskDataFormat)
	}

	users, ok := data["data"].([]interface{})
	if !ok {
		return nil, e.New(constant.ErrDooTaskDataFormat)
	}

	return interfaces.PaginationRsp(page, pageSize, int64(count), users), nil
}

// 获取项目列表
func (s dootaskService) GetProjectList(token string, page, pageSize int) (*interfaces.Pagination, error) {
	url := fmt.Sprintf("%s%s?page=%d&pagesize=%d&token=%s", config.DooTaskUrl, "/api/project/lists", page, pageSize, token)
	result, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}

	data, err := s.UnmarshalAndCheckResponse(result)
	if err != nil {
		return nil, err
	}

	count, ok := data["total"].(float64)
	if !ok {
		return nil, e.New(constant.ErrDooTaskDataFormat)
	}

	list, ok := data["data"].([]interface{})
	if !ok {
		return nil, e.New(constant.ErrDooTaskDataFormat)
	}

	return interfaces.PaginationRsp(page, pageSize, int64(count), list), nil
}

// 创建OKR评论会话
func (s dootaskService) DialogOkrAdd(token string, okr *model.Okr) (int, error) {
	url := fmt.Sprintf("%s%s", config.DooTaskUrl, "/api/dialog/okr/add")
	params := make(map[string]interface{})
	params["name"] = okr.Title
	params["link_id"] = okr.Id
	userids := []int{okr.Userid}

	// 当okr有部门时，添加部门负责人
	if okr.DepartmentId != "" {
		departmentIds := common.ExplodeInt(",", okr.DepartmentId, true)
		if len(departmentIds) > 0 {
			// 获取部门负责人userid
			var ownerids []int
			err := core.DB.Model(&model.UserDepartment{}).Where("id in (?)", departmentIds).Pluck("owner_userid", &ownerids).Error
			if err != nil {
				return 0, err
			}
			userids = common.ArrayUniqueInt(append(userids, ownerids...))
		}
	}
	//
	params["userids"] = userids
	result, err := s.client.PostToken(url, params, token)
	if err != nil {
		return 0, err
	}
	data, err := s.UnmarshalAndCheckResponse(result)
	if err != nil {
		return 0, err
	}

	dialogId, ok := data["id"].(float64)
	if !ok {
		return 0, e.New(constant.ErrDooTaskDataFormat)
	}

	return int(dialogId), nil
}

// 推送OKR相关信息
func (s dootaskService) DialogOkrPush(okr *model.Okr, token string, mold int, userids []int) error {
	url := fmt.Sprintf("%s%s", config.DooTaskUrl, "/api/dialog/okr/push")

	oHtml := ""
	krHtml := ""
	if okr.ParentId == 0 {
		oHtml = fmt.Sprintf("<span class=\"mention okr\" data-id=\"%v\">#%v</span>", okr.Id, okr.Title)
	} else {
		oHtml = fmt.Sprintf("<span class=\"mention okr\" data-id=\"%v\">#%v</span>", okr.ParentId, okr.ParentTitle)
		krHtml = fmt.Sprintf("<span class=\"mention okr\" data-id=\"%v\">#%v</span>", okr.ParentId, okr.Title)
	}
	text := ""
	switch mold {
	case 1:
		text = fmt.Sprintf("您有新的OKR %v", oHtml) //创建O时（通知发起/所有KR参与人）
	case 2:
		text = fmt.Sprintf("您参与的OKR %v 目标有变动", oHtml) //O目标变动（通知此O内所有KR参与人）
	case 3:
		text = fmt.Sprintf("您参与的OKR %v 中的KR %v 有变动", oHtml, krHtml) //KR变动（非删除-通知此KR所有参与人）
	case 4:
		text = fmt.Sprintf("您有新的OKR %v 中的KR %v", oHtml, krHtml) //为KR添加新参与人（仅通知新的参与人）
	case 5:
		text = fmt.Sprintf("您参与的OKR %v 周期已修改", oHtml) //O时间变动（通知此O内所有KR参与人）
	case 6:
		text = fmt.Sprintf("您参与的OKR %v 中的KR %v 时间已修改", oHtml, krHtml) //KR时间变动（通知此KR所有参与人）
	}
	// 循环推送
	for _, userid := range userids {
		params := map[string]interface{}{
			"text":   text,
			"userid": userid,
		}
		_, err := s.client.PostToken(url, params, token)
		if err != nil {
			return err
		}
	}
	return nil
}

// 解码并检查返回数据
func (s dootaskService) UnmarshalAndCheckResponse(resp []byte) (map[string]interface{}, error) {
	var ret map[string]interface{}
	if err := json.Unmarshal(resp, &ret); err != nil {
		return nil, e.WithDetail(constant.ErrDooTaskUnmarshalResponse, err, nil)
	}

	retCode, ok := ret["ret"].(float64)
	if !ok {
		return nil, e.New(constant.ErrDooTaskResponseFormat)
	}

	if retCode != 1 {
		msg, ok := ret["msg"].(string)
		if !ok {
			return nil, e.New(constant.ErrDooTaskRequestFailed)
		}
		return nil, e.WithDetail(constant.ErrDooTaskRequestFailedWithErr, msg, nil)
	}

	data, ok := ret["data"].(map[string]interface{})
	if !ok {
		dataList, isList := ret["data"].([]interface{})
		if !isList {
			return nil, e.New(constant.ErrDooTaskDataFormat)
		}

		data = make(map[string]interface{})
		data["list"] = dataList
	}

	return data, nil
}

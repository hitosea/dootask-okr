package service

import (
	"dootask-okr/app/core"
	"dootask-okr/app/interfaces"
	"dootask-okr/app/model"
	"dootask-okr/app/utils/common"
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
	url := fmt.Sprintf("%s%s", config.DooTaskUrl, "/api/users/info?token="+token)
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

// 获取用户列表
func (s dootaskService) GetUserList(token string) ([]interface{}, error) {
	url := fmt.Sprintf("%s%s", config.DooTaskUrl, "/api/users/search?keys[key]=&keys[project_id]=0&keys[no_project_id]=0&keys[dialog_id]=0&keys[bot]=0&keys[disable]=0&take=50&token="+token)
	result, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}

	data, err := s.UnmarshalAndCheckResponse(result)
	if err != nil {
		return nil, err
	}

	users, ok := data["list"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("用户列表格式错误")
	}
	return users, nil
}

// 获取项目列表
func (s dootaskService) GetProjectList(token string) ([]interface{}, error) {
	url := fmt.Sprintf("%s%s", config.DooTaskUrl, "/api/project/lists?type=team&keys[name]=&getuserid=yes&getstatistics=no&token="+token)
	result, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}

	data, err := s.UnmarshalAndCheckResponse(result)
	if err != nil {
		return nil, err
	}

	list, ok := data["data"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("项目列表格式错误")
	}

	return list, nil
}

// 创建OKR评论会话
func (s dootaskService) DialogOkrAdd(token string, okr *model.Okr) (int, error) {
	url := fmt.Sprintf("%s%s", config.DooTaskUrl, "/api/dialog/okr/add")
	params := make(map[string]interface{})
	params["name"] = okr.Title
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
		return 0, fmt.Errorf("dialogId is not an int")
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
		krHtml = fmt.Sprintf("<span class=\"mention okr\" data-id=\"%v\">#%v</span>", okr.Id, okr.Title)
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
		return nil, fmt.Errorf("解析响应失败：%w", err)
	}

	retCode, ok := ret["ret"].(float64)
	if !ok {
		return nil, fmt.Errorf("响应格式错误")
	}

	if retCode != 1 {
		msg, ok := ret["msg"].(string)
		if !ok {
			return nil, fmt.Errorf("请求失败")
		}
		return nil, fmt.Errorf("请求失败：%v", msg)
	}

	data, ok := ret["data"].(map[string]interface{})
	if !ok {
		dataList, isList := ret["data"].([]interface{})
		if !isList {
			return nil, fmt.Errorf("数据格式错误")
		}

		data = make(map[string]interface{})
		data["list"] = dataList
	}

	return data, nil
}

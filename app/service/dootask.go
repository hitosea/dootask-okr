package service

import (
	"dootask-okr/app/interfaces"
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

package service

import (
	"dootask-okr/app/utils/common"
	"encoding/json"
	"fmt"
	"time"
)

const (
	dooUrl = "http://192.168.100.219:2222/api/"
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
func (s dootaskService) GetUserInfo(token string) (interface{}, error) {
	url := fmt.Sprintf("%s%s", dooUrl, "users/info?token="+token)
	result, err := s.client.Get(url)
	if err != nil {
		return nil, err
	}
	return s.UnmarshalAndCheckResponse(result)
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
		return nil, fmt.Errorf("数据格式错误")
	}
	return data, nil
}

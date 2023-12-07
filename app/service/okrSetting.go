package service

import (
	"dootask-okr/app/constant"
	"dootask-okr/app/core"
	"dootask-okr/app/interfaces"
	"dootask-okr/app/model"
	"dootask-okr/app/utils/common"
	e "dootask-okr/app/utils/error"
	"encoding/json"
	"sync"
)

var (
	OkrSettingService = okrSettingService{}
	cache             = make(map[string][]byte) // 缓存
	cacheMu           sync.RWMutex              // 读写锁
)

type okrSettingService struct{}

// OkrSetting okr设置
func (s *okrSettingService) OkrSetting(param interfaces.OkrSettingReq) (map[string]interface{}, error) {
	var ret map[string]interface{}
	//
	if param.Type == "save" {
		// 判断评分权重等于100
		if param.SelfScoreWeight+param.SuperiorScoreWeight != 100 {
			return nil, e.New(constant.ErrOkrScoreWeightInvalid)
		}
		// 保存设置
		settingMap, _ := common.StructToMap(param)
		delete(settingMap, "type")
		data, err := s.SettingOperation(model.SettingOkrKey, settingMap, true)
		if err != nil {
			return nil, err
		}
		ret = data

	} else {
		// 获取设置
		data, _ := s.SettingOperation(model.SettingOkrKey, nil, false)
		ret = data
	}

	// 如果ret为空，初始化数据
	if ret == nil {
		ret = map[string]interface{}{
			"score_department_user": 0,
			"self_score_weight":     model.SelfScoreWeight,
			"superior_score_weight": model.SuperiorScoreWeight,
		}
	}
	return ret, nil
}

// GetSetting 获取设置
func (s *okrSettingService) GetSetting(name string) (map[string]interface{}, error) {
	cacheMu.RLock()
	cached, ok := cache[name]
	cacheMu.RUnlock()

	if ok {
		var setting map[string]interface{}
		err := json.Unmarshal(cached, &setting)
		if err != nil {
			return nil, err
		}
		return setting, nil
	}

	row, err := s.getSettingByName(name)
	if err != nil {
		return nil, err
	}

	var setting map[string]interface{}
	err = json.Unmarshal([]byte(row.Setting), &setting)
	if err != nil {
		return nil, err
	}

	cacheMu.Lock()
	cache[name] = []byte(row.Setting)
	cacheMu.Unlock()

	return setting, nil
}

// Setting 设置操作
func (s *okrSettingService) SettingOperation(name string, newSettings map[string]interface{}, isUpdate bool) (map[string]interface{}, error) {
	if newSettings == nil {
		return s.GetSetting(name)
	}

	err := s.UpdateSetting(name, newSettings, isUpdate)
	if err != nil {
		return nil, err
	}

	return s.GetSetting(name)
}

// UpdateSetting 更新设置
func (s *okrSettingService) UpdateSetting(name string, newSettings map[string]interface{}, isUpdate bool) error {
	serializedSettings, err := json.Marshal(newSettings)
	if err != nil {
		return err
	}

	err = s.updateSettingByName(name, string(serializedSettings))
	if err != nil {
		return err
	}

	cacheMu.Lock()
	cache[name] = serializedSettings
	cacheMu.Unlock()

	return nil
}

// getSettingByName 获取设置
func (s *okrSettingService) getSettingByName(name string) (*model.OkrSetting, error) {
	settingModel := &model.OkrSetting{}
	if err := core.DB.Where("name=?", name).First(settingModel).Error; err != nil {
		return nil, err
	}
	return settingModel, nil
}

// updateSettingByName 更新设置
func (s *okrSettingService) updateSettingByName(name string, setting string) error {
	settingModel := model.OkrSetting{}
	if err := core.DB.Where("name=?", name).First(&settingModel).Error; err != nil {
		// 如果获取失败，则创建一个新的设置
		settingModel = model.OkrSetting{
			Name:    name,
			Setting: setting,
		}
		if err := core.DB.Create(&settingModel).Error; err != nil {
			return err
		}
	} else {
		settingModel.Setting = setting
		if err := core.DB.Save(&settingModel).Error; err != nil {
			return err
		}
	}
	return nil
}

// SettingFind 获取设置
func (s *okrSettingService) SettingFind(name string, keyName string, defaultVal string) (string, error) {
	array, err := s.GetSetting(name)
	if err != nil {
		return defaultVal, err
	}
	if val, ok := array[keyName]; ok {
		return val.(string), nil
	}
	return defaultVal, nil
}

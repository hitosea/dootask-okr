package common

import (
	"encoding/json"
	"strings"
)

// StructToJson 结构体转json
func StructToJson(data interface{}) string {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

// StrToStruct 字符串转对象
func StrToStruct(source string, destination interface{}) error {
	err := json.Unmarshal([]byte(source), destination)
	return err
}

// StrToMap 字符转Map
func StrToMap(source string) (map[string]interface{}, error) {
	res := make(map[string]interface{})
	err := json.Unmarshal([]byte(source), &res)
	return res, err
}

// MapToStruct Map转结构体
func MapToStruct(m map[string]interface{}, s interface{}) error {
	data, err := json.Marshal(m)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, s); err != nil {
		return err
	}
	return nil
}

// StructToStruct 结构体转结构体，并可选择过滤某些字段
func StructToStruct(data interface{}, newdata interface{}, filtrationKeys ...string) error {
	result, err := json.Marshal(data)
	if err != nil {
		return err
	}
	str := string(result)
	// 过滤字段
	if len(filtrationKeys) != 0 {
		keysArr := strings.Split(filtrationKeys[0], ",")
		for _, v := range keysArr {
			str = strings.Replace(string(str), `"`+v+`":`, `"_`+v+`":`, 1)
		}
	}
	// 转新的结构体
	err = json.Unmarshal([]byte(str), newdata)
	if err != nil {
		return err
	}
	//
	return nil
}

// JsonToStr 对象转换成字符串
func JsonToStr(data interface{}) (string, error) {
	result, err := json.Marshal(data)
	return string(result), err
}

// StructToMap 结构体转map
func StructToMap(data interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	result, err := StrToMap(StructToJson(data))
	return result, err
}

// MergeMaps 合并map
func MergeMaps(maps ...map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

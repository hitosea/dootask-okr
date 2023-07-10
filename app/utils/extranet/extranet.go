package extranet

import (
	"crypto/md5"
	"dootask-okr/app/core"
	"dootask-okr/app/utils/common"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	cacheExpiration = 0
)

type IpPoint struct {
	Lat  float64
	Long float64
}

// GetIpGcj02 获取IP地址经纬度
func GetIpGcj02(ip string) (*IpPoint, error) {
	if ip == "" {
		ip = common.GetIp()
	}

	cacheKey := "getIpPoint::" + common.StringMd5(ip)

	// 从缓存获取结果
	result, found := core.Cache.Get(cacheKey)
	if found {
		return result.(*IpPoint), nil
	}

	// 缓存中未找到结果，发送HTTP请求获取IP地址经纬度
	url := "https://www.ifreesite.com/ipaddress/address.php?q=" + ip
	respBody, err := getRespBody(url)
	if err != nil {
		return nil, err
	}

	// 解析结果字符串获取经纬度信息
	lastPos := strings.LastIndex(respBody, ",")
	long, _ := strconv.ParseFloat(common.GetMiddle(respBody[lastPos+1:], "", ")"), 64)
	lat, _ := strconv.ParseFloat(common.GetMiddle(respBody[strings.LastIndex(respBody[:lastPos], ",")+1:], "", ","), 64)

	ipPoint := &IpPoint{Lat: lat, Long: long}

	// 缓存结果
	core.Cache.Set(cacheKey, ipPoint, cacheExpiration)

	return ipPoint, nil
}

// GetIpGcj02ByBaidu 根据IP获取经纬度
func GetIpGcj02ByBaidu(ip string) (*IpPoint, error) {
	if ip == "" {
		ip = common.GetIp()
	}

	cacheKey := "getIpPoint::" + fmt.Sprintf("%x", md5.Sum([]byte(ip)))

	// 从缓存获取结果
	result, found := core.Cache.Get(cacheKey)
	if found {
		return result.(*IpPoint), nil
	}

	// 缓存中未找到结果，发送HTTP请求获取IP地址经纬度
	ak := "app.baidu_app_key"
	url := fmt.Sprintf("http://api.map.baidu.com/location/ip?ak=%s&ip=%s&coor=bd09ll", ak, ip)
	respBody, err := getRespBody(url)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(respBody), &data); err != nil {
		return nil, err
	}

	// x坐标纬度, y坐标经度
	long := data["content"].(map[string]interface{})["point"].(map[string]interface{})["x"].(float64)
	lat := data["content"].(map[string]interface{})["point"].(map[string]interface{})["y"].(float64)

	ipPoint := &IpPoint{Lat: lat, Long: long}

	// 缓存结果
	core.Cache.Set(cacheKey, ipPoint, cacheExpiration)

	return ipPoint, nil
}

// GetIpInfo 获取IP地址详情
func GetIpInfo(ip string) (interface{}, error) {
	if ip == "" {
		ip = common.GetIp()
	}

	cacheKey := "getIpInfo::" + common.StringMd5(ip)

	// 从缓存获取结果
	result, found := core.Cache.Get(cacheKey)
	if found {
		return result, nil
	}

	// 缓存中未找到结果，发送HTTP请求获取IP地址详情
	url := "http://ip.taobao.com/service/getIpInfo.php?accessKey=alibaba-inc&ip=" + ip
	respBody, err := getRespBody(url)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	ret := make(map[string]interface{})
	if err := json.Unmarshal([]byte(respBody), &data); err != nil {
		return nil, err
	}

	if data["code"].(float64) != 0 {
		return nil, fmt.Errorf("error ip: -1")
	}

	data = data["data"].(map[string]interface{})
	ret["text"] = data["country"].(string)
	ret["textSmall"] = data["country"].(string)
	if data["region"].(string) != "" && data["region"].(string) != data["country"].(string) && data["region"].(string) != "XX" {
		ret["text"] = fmt.Sprintf("%v %v", ret["text"], data["region"].(string))
		ret["textSmall"] = data["region"].(string)
	}
	if data["city"].(string) != "" && data["city"].(string) != data["region"].(string) && data["city"].(string) != "XX" {
		ret["text"] = fmt.Sprintf("%v %v", ret["text"], data["city"].(string))
		ret["textSmall"] = fmt.Sprintf("%v %v", ret["textSmall"], data["city"].(string))
	}
	if data["county"].(string) != "" && data["county"].(string) != data["city"].(string) && data["county"].(string) != "XX" {
		ret["text"] = fmt.Sprintf("%v %v", ret["text"], data["county"].(string))
		ret["textSmall"] = fmt.Sprintf("%v %v", ret["textSmall"], data["county"].(string))
	}

	// 缓存结果
	core.Cache.Set(cacheKey, ret, cacheExpiration)

	return ret, nil
}

// IsHoliday 判断是否工作日 年月日（如：20220102）
// 返回值：0: 工作日 1: 非工作日 2: 获取不到远程数据的非工作日（周六、日） 所以可以用>0来判断是否工作日
func IsHoliday(Ymd string) (int, error) {
	t, err := time.Parse("20060102", Ymd)
	if err != nil {
		return 2, nil
	}

	holidayKey := "holiday::" + t.Format("200601")
	// 从缓存获取结果
	result, found := core.Cache.Get(holidayKey)
	if found {
		return result.(int), nil
	}

	// 缓存中未找到结果，发送HTTP请求获取IP地址详情
	url := fmt.Sprintf("https://api.apihubs.cn/holiday/get?field=date&month=%s&workday=2&size=31", t.Format("200601"))
	respBody, err := getRespBody(url)
	if err != nil {
		return 2, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(respBody), &data); err != nil {
		return 2, err
	}

	data = data["data"].(map[string]interface{})
	if data["code"].(float64) != 0 {
		return 2, fmt.Errorf("[holiday] result error")
	}

	holidayData := make([]string, 0)
	for _, v := range data["list"].([]interface{}) {
		holidayData = append(holidayData, v.(map[string]interface{})["date"].(string))
	}

	// 缓存结果
	core.Cache.Set(holidayKey, holidayData, cacheExpiration)

	if len(holidayData) == 0 {
		return 2, nil
	}

	if common.InArray(t.Format("20060102"), holidayData) {
		return 1, nil
	}

	return 0, nil
}

// DrawioIconSearch Drawio图标搜索
func DrawioIconSearch(query string, page, size int) (interface{}, error) {
	url := fmt.Sprintf("https://app.diagrams.net/iconSearch?q=%s&p=%d&c=%d", query, page, size)
	respBody, err := getRespBody(url)
	if err != nil {
		return map[string]interface{}{
			"icons":       []interface{}{},
			"total_count": 0,
		}, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(respBody), &data); err != nil {
		return nil, err
	}

	return data, nil
}

// RandJoke 随机笑话接口
func RandJoke() (interface{}, error) {
	jokeKey := "JUKE_KEY_JOKE"
	if jokeKey != "" {
		data, err := getRespBody("http://v.juhe.cn/joke/randJoke.php?key=" + jokeKey)
		if err != nil {
			return nil, err
		}

		var result map[string]interface{}
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			return "", err
		}

		if result["reason"].(string) == "success" {
			return result["result"].(map[string]interface{})["content"].(string), nil
		}
	}
	return nil, nil
}

// Soups 心灵鸡汤
func Soups() (string, error) {
	soupKey := "JUKE_KEY_SOUP"
	if soupKey != "" {
		data, err := getRespBody("https://apis.juhe.cn/fapig/soup/query?key=" + soupKey)
		if err != nil {
			return "", err
		}

		var result map[string]interface{}
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			return "", err
		}

		if result["reason"].(string) == "success" {
			return result["result"].(map[string]interface{})["text"].(string), nil
		}
	}
	return "", nil
}

// CheckinBotQuickMsg 签到机器人
func CheckinBotQuickMsg(command string) string {
	text := "维护中..."
	switch command {
	case "it":
		data, err := getRespBody("http://vvhan.api.hitosea.com/api/hotlist?type=itNews")
		if err != nil {
			return text
		}

		var result map[string]interface{}
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			return text
		}

		if result["code"].(float64) == 200 {
			i := 1
			array := make([]string, 0)
			for _, v := range result["data"].([]interface{}) {
				item := v.(map[string]interface{})
				if item["title"].(string) != "" && item["desc"].(string) != "" {
					array = append(array, fmt.Sprintf("%d. <strong><a href='%s' target='_blank'>%s</a></strong>\n%s", i, item["mobilUrl"].(string), item["title"].(string), item["desc"].(string)))
					i++
				}
			}
			if len(array) > 0 {
				text = fmt.Sprintf("<strong>%s</strong>（%s）\n\n%s", result["title"].(string), result["update_time"].(string), strings.Join(array, "\n\n"))
			}
		}
		break

	case "36ke":
		data, err := getRespBody("http://vvhan.api.hitosea.com/api/hotlist?type=36Ke")
		if err != nil {
			return text
		}

		var result map[string]interface{}
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			return text
		}

		if result["code"].(float64) == 200 {
			i := 1
			array := make([]string, 0)
			for _, v := range result["data"].([]interface{}) {
				item := v.(map[string]interface{})
				if item["title"].(string) != "" && item["desc"].(string) != "" {
					array = append(array, fmt.Sprintf("%d. <strong><a href='%s' target='_blank'>%s</a></strong>\n%s", i, item["mobilUrl"].(string), item["title"].(string), item["desc"].(string)))
					i++
				}
			}
			if len(array) > 0 {
				text = fmt.Sprintf("<strong>%s</strong>（%s）\n\n%s", result["title"].(string), result["update_time"].(string), strings.Join(array, "\n\n"))
			}
		}
		break

	case "60s":
		data, err := getRespBody("http://vvhan.api.hitosea.com/api/60s?type=json")
		if err != nil {
			return text
		}

		var result map[string]interface{}
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			return text
		}

		if result["code"].(float64) == 200 {
			i := 1
			array := make([]string, 0)
			for _, v := range result["data"].([]interface{}) {
				item := v.(map[string]interface{})
				if item["title"].(string) != "" {
					array = append(array, fmt.Sprintf("%d. %s", i, item["title"].(string)))
					i++
				}
			}
			if len(array) > 0 {
				text = fmt.Sprintf("<strong>%s</strong>（%s）\n\n%s", result["name"].(string), result["time"].([]interface{})[0].(string), strings.Join(array, "\n\n"))
			}
		}
		break

	case "joke":
		text = "笑话被掏空"
		data, err := getRespBody("http://vvhan.api.hitosea.com/api/joke?type=json")
		if err != nil {
			return text
		}

		var result map[string]interface{}
		if err := json.Unmarshal([]byte(data), &result); err != nil {
			return text
		}

		if result["code"].(float64) == 200 {
			if result["joke"].(string) != "" {
				text = fmt.Sprintf("开心笑话：%s", result["joke"].(string))
			}
		}
		break

	case "soup":
		text = "鸡汤分完了"
		data, err := getRespBody("https://api.ayfre.com/jt/?type=bot")
		if err != nil {
			return text
		}

		if data != "" {
			text = fmt.Sprintf("心灵鸡汤：%s", data)
		}
		break
	}
	return text
}

// 发送HTTP请求获取响应体
func getRespBody(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 将结果存入缓存
	return string(body), nil
}

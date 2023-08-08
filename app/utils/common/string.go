package common

import (
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/mozillazg/go-pinyin"
)

// RandString 生成随机字符串
func RandString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// RandNum 生成随机数
func RandNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// GetFirstCharter 获取首字母拼音
func GetFirstCharter(s string) string {
	pin := pinyin.NewArgs()
	pin.Style = pinyin.FirstLetter
	// 判断是否包含中文字符
	for _, c := range s {
		if c > 255 {
			return strings.ToUpper(pinyin.Pinyin(string(c), pin)[0][0])
		}
	}
	return strings.ToUpper(string(s[0]))
}

// 获取中文字符串的拼音，非中文字符返回原字符串
func Cn2Pinyin(s string) string {
	pin := pinyin.NewArgs()
	pin.Style = pinyin.Normal
	runes := []rune(s)
	length := len(runes)
	var result []string
	for i := 0; i < length; i++ {
		r := runes[i]
		if r > 255 {
			// 中文字符
			result = append(result, concat(pinyin.Pinyin(string(r), pin)))
		} else {
			// 非中文字符
			result = append(result, string(r))
		}
	}
	return strings.Join(result, "")
}

// 获取逗号分隔的多个中文字符串拼音，非中文字符返回原字符串
// 多个中文名字之间用逗号分隔，例如："张三, 李四, 王五"
func getMultiNamePinyin(names string) string {
	pin := pinyin.NewArgs()
	pin.Style = pinyin.Normal
	nameList := strings.Split(names, ",")
	for i, _ := range nameList {
		nameList[i] = strings.TrimSpace(nameList[i])
	}
	var result []string
	for _, name := range nameList {
		if name != "" {
			result = append(result, Cn2Pinyin(name))
		}
	}
	return strings.Join(result, ", ")
}

// 将拼音数组合并成一个字符串
func concat(pinyinList [][]string) string {
	var result string
	for _, p := range pinyinList {
		result += p[0]
	}
	return result
}

// 字符串转int
func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

// 检查字符串s是否以prefix开头
func LeftExists(s string, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// 删除字符串s中所有前缀为prefix的子串
func LeftDelete(s string, prefix string) string {
	for LeftExists(s, prefix) {
		s = s[len(prefix):]
	}
	return s
}

// 计算两个字符串切片之间的差集
func StringDiff(slice1, slice2 []string) []string {
	diff := []string{}
	for _, s1 := range slice1 {
		if !StringContains(slice2, s1) {
			diff = append(diff, s1)
		}
	}
	return diff
}

// 检查字符串s是否存在于slice中
func StringContains(slice []string, s string) bool {
	for _, e := range slice {
		if e == s {
			return true
		}
	}
	return false
}

// 检查字符串是否是合法的 MAC 地址格式（例如：00:11:22:33:44:55）
func IsMAC(mac string) bool {
	regex := regexp.MustCompile("^([0-9A-Fa-f]{2}[:]){5}([0-9A-Fa-f]{2})$")
	return regex.MatchString(mac)
}

// 打散字符串，只留为数字的项
func ExplodeInt(delimiter string, str interface{}, reInt bool) []int {
	var array []string
	if str == nil {
		str = delimiter
		delimiter = ","
	}
	// 判断类型
	switch v := str.(type) {
	case string:
		array = strings.Split(v, delimiter)
	case []string:
		array = v
	case []int:
		return v
	}
	return ArrayRetainInt(array, reInt)
}

// 数组只保留数字的
func ArrayRetainInt(arr []string, reInt bool) []int {
	var result []int
	for _, v := range arr {
		// 如果不是数字，则剔除
		if _, err := strconv.Atoi(v); err != nil {
			continue
		}
		// 如果需要格式化，则格式化
		if reInt {
			result = append(result, StringToInt(v))
		}
	}
	return result
}

// GetMiddle 截取指定字符串
func GetMiddle(str, ta, tb string) string {
	if ta != "" && strings.Contains(str, ta) {
		str = str[strings.Index(str, ta)+len(ta):]
	}
	if tb != "" && strings.Contains(str, tb) {
		str = str[:strings.Index(str, tb)]
	}
	return str
}

// IsChineseCharCountValid 检查字符串中的中文字符数量是否超过255个
func IsChineseCharCountValid(str string) bool {
	return utf8.RuneCountInString(str) <= 255
}

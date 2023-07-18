package common

import (
	"strconv"
	"strings"
)

// int 数组去重
func ArrayUniqueInt(result []int) []int {
	j := 0
	seen := make(map[int]bool)
	for _, v := range result {
		if !seen[v] {
			seen[v] = true
			result[j] = v
			j++
		}
	}
	return result[:j]
}

// 交集处理
func ArrayIntersectionProcessing(nums []int, itemIds []int) []int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	var result []int
	for _, v := range itemIds {
		if m[v] {
			result = append(result, v)
		}
	}
	return result
}

// 差集处理
func ArrayDifferenceProcessing(a, b []int) []int {
	m := make(map[int]bool)
	for _, item := range b {
		m[item] = true
	}
	var diff []int
	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return diff
}

// str 数组去重
func ArrayUniqueStr(nums []string) []string {
	result := make([]string, 0, len(nums))
	seen := make(map[string]bool)
	for _, n := range nums {
		if !seen[n] {
			result = append(result, n)
			seen[n] = true
		}
	}
	return result
}

// 数组是否包含
func StringsContains(array []string, val string) (index int) {
	index = -1
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			index = i
			return
		}
	}
	return
}

// 元素是否存在数组中
func InArray(item string, items []string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// 数组是否存在，int
func InArrayInt(needle int, haystack []int) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

// 向数组插入内容
func SliceInsert(s []string, index int, value string) []string {
	rear := append([]string{}, s[index:]...)
	return append(append(s[:index], value), rear...)
}

// 查找数组位置
func FindIndex(tab []string, value string) int {
	for i, v := range tab {
		if v == value {
			return i
		}
	}
	return -1
}

// str 数组Diff
func ArrayDiffStr(arr []string, id int) []int {
	var result []int
	for _, d := range arr {
		i, _ := strconv.Atoi(d)
		if i != id && id != 0 && i != 0 {
			result = append(result, i)
		}
	}
	return result
}

// int 数组Diff
func ArrayDiffInt(arr []int, id int) []int {
	var result []int
	for _, d := range arr {
		if d != id && id != 0 && d != 0 {
			result = append(result, d)
		}
	}
	return result
}

// int 数组转字符串
func ArrayImplode(arr []int) string {
	var strSlice []string
	for _, d := range arr {
		strSlice = append(strSlice, strconv.Itoa(int(d)))
	}
	return strings.Join(strSlice, ",")
}

// str 数组转字符串
func ArrayStringImplode(arr []string) string {
	return strings.Join(arr, ",")
}

// 合并数组
func MergeArray(arr1 []string, arr2 []string) []string {
	result := append([]string{}, arr1...)
	result = append(result, arr2...)
	return result
}

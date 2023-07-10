package common

import (
	e "dootask-okr/app/utils/error"
	"regexp"
)

// PasswordPolicy 检测密码策略是否符合
func PasswordPolicy(password string, complex bool) error {
	if len(password) < 6 {
		return e.New("密码设置不能小于6位数")
	}
	if len(password) > 32 {
		return e.New("密码最多只能设置32位数")
	}
	// 复杂密码
	if complex {
		matched, _ := regexp.MatchString("^[0-9]+$", password)
		if matched {
			return e.New("密码不能全是数字，请包含数字，字母大小写或者特殊字符")
		}
		matched, _ = regexp.MatchString("^[a-zA-Z]+$", password)
		if matched {
			return e.New("密码不能全是字母，请包含数字，字母大小写或者特殊字符")
		}
		matched, _ = regexp.MatchString("^[0-9A-Z]+$", password)
		if matched {
			return e.New("密码不能全是数字+大写字母，密码包含数字，字母大小写或者特殊字符")
		}
		matched, _ = regexp.MatchString("^[0-9a-z]+$/", password)
		if matched {
			return e.New("码不能全是数字+小写字母，密码包含数字，字母大小写或者特殊字符")
		}
	}
	return nil
}

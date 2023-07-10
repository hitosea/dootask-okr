package common

import "net/mail"

// FormatEmail 格式化邮箱地址
func FormatEmail(email string) string {
	o, err := mail.ParseAddress(email)
	if err != nil {
		return ""
	}
	return o.Address
}

// IsEmail 判断是否是邮箱
func IsEmail(email string) bool {
	email = FormatEmail(email)
	if len(email) == 0 {
		return false
	}
	return true
}

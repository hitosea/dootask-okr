package common

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// StringMd5 MD5
func StringMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// StringMd52 MD5
func StringMd52(str, pass string) string {
	text := fmt.Sprintf("%s%s", StringMd5(str), pass)
	return StringMd5(text)
}

// Md5s 返回 $val 的 MD5 哈希值的前 $len 个字符。
func Md5s(val string) string {
	h := md5.New()
	h.Write([]byte(val))
	return hex.EncodeToString(h.Sum(nil))[:16]
}

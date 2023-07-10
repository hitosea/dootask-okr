package common

import (
	"encoding/base64"
	"fmt"
)

// base64加密
func Base64Encode(data string, a ...any) string {
	if len(a) > 0 {
		data = fmt.Sprintf(data, a...)
	}
	sEnc := base64.StdEncoding.EncodeToString([]byte(data))
	return fmt.Sprintf(sEnc)
}

// base64解码
func Base64Decode(data string) string {
	uDec, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return ""
	}
	return string(uDec)
}

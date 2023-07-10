package captcha

import (
	"dootask-okr/app/constant"
	"dootask-okr/app/interfaces"
	"errors"
	"strings"

	"github.com/mojocn/base64Captcha"
)

var store = base64Captcha.DefaultMemStore

func VerifyCode(codeID string, code string) error {
	if codeID == "" {
		return errors.New(constant.ErrCaptchaCode)
	}
	vv := store.Get(codeID, true)
	vv = strings.TrimSpace(vv)
	code = strings.TrimSpace(code)
	if strings.EqualFold(vv, code) {
		return nil
	}
	return errors.New(constant.ErrCaptchaCode)
}

func CreateCaptcha() (*interfaces.CaptchaResponse, error) {
	var driverString base64Captcha.DriverString
	driverString.Source = "1234567890QWERTYUPLKJHGFDSAZXCVBNMqwertyupkjhgfdsazxcvbnm"
	driverString.Width = 140
	driverString.Height = 50
	driverString.NoiseCount = 8
	driverString.Length = 5
	driverString.Fonts = []string{"RitaSmith.ttf", "actionj.ttf", "chromohv.ttf"}
	driver := driverString.ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	if err != nil {
		return nil, err
	}
	return &interfaces.CaptchaResponse{
		CaptchaID: id,
		ImagePath: b64s,
	}, nil
}

package interfaces

type Response struct {
	Code int         `json:"code"` // 状态, [200=成功, 400=失败, 401=未登录, 403=无相关权限, 404=请求接口不存在, 405=请求方法不允许, 500=系统错误]
	Msg  string      `json:"msg"`  // 信息
	Data interface{} `json:"data"` // 数据
}

type CaptchaResponse struct {
	CaptchaID string `json:"captcha_id"` //验证码ID
	ImagePath string `json:"image_path"` //图片路径
}

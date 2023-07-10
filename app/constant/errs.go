package constant

// api
var (
	// 环境
	ErrEnvProhibition    = "ErrEnvProhibition"    //当前环境禁止此操作
	ErrInvalidParameter  = "ErrInvalidParameter"  //参数错误
	ErrCaptchaCode       = "ErrCaptchaCode"       //验证码错误
	ErrTypeNotLogin      = "ErrTypeNotLogin"      //未登录
	ErrMailNotConfig     = "ErrMailNotConfig"     //发送邮箱未配置
	ErrMailToInvalid     = "ErrMailToInvalid"     //请输入正确的收件人地址
	ErrRequestTimeout    = "ErrRequestTimeout"    //请求超时
	ErrMailContentReject = "ErrMailContentReject" //邮件内容被拒绝，请检查邮箱是否开启接收功能

)

// app
var (
	ErrCmdTimeout = "ErrCmdTimeout" //命令执行超时
)

package constant

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
	ErrNoPermission      = "ErrNoPermission"      //权限不足

	// OKR
	ErrOkrKeyResultAtLeastOne     = "ErrOkrKeyResultAtLeastOne"     //至少有一条关键结果
	ErrOkrNoData                  = "ErrOkrNoData"                  //没有查询到相关数据
	ErrOkrInvalidKrScore          = "ErrOkrInvalidKrScore"          //无效的KR分数
	ErrOkrProgressNotEnough       = "ErrOkrProgressNotEnough"       //KR进度尚未达到100%
	ErrOkrOwnerNotScore           = "ErrOkrOwnerNotScore"           //负责人未评分
	ErrOkrKrScoreNotComplete      = "ErrOkrKrScoreNotComplete"      //KR评分未完成
	ErrOkrTypeInvalid             = "ErrOkrTypeInvalid"             //目标类型参数错误
	ErrOkrPriorityInvalid         = "ErrOkrPriorityInvalid"         //目标优先级参数错误
	ErrOkrAscriptionInvalid       = "ErrOkrAscriptionInvalid"       //目标归属参数错误
	ErrOkrVisibleRangeInvalid     = "ErrOkrVisibleRangeInvalid"     //目标可见范围参数错误
	ErrOkrConfidenceInvalid       = "ErrOkrConfidenceInvalid"       //信心指数请输入0-100的整数
	ErrOkrIsFinishNotScoreInvalid = "ErrOkrIsFinishNotScoreInvalid" //是否已完成未评分参数错误
	ErrOkrProgressInvalid         = "ErrOkrProgressInvalid"         //进度请输入0-100的整数
	ErrOkrProgressStatusInvalid   = "ErrOkrProgressStatusInvalid"   //进度状态参数错误
	ErrOkrScoreInvalid            = "ErrOkrScoreInvalid"            //评分请输入0-10的整数
	ErrOkrOwnerScored             = "ErrOkrOwnerScored"             //负责人已评分
	ErrOkrSuperiorScored          = "ErrOkrSuperiorScored"          //上级已评分
	ErrOkrScoredNotUpdateProgress = "ErrOkrScoredNotUpdateProgress" //评分后不允许修改进度
	ErrOkrKeyResultTitleEmpty     = "ErrOkrKeyResultTitleEmpty"     //关键结果标题不能为空
	ErrOkrTimeEmpty               = "ErrOkrTimeEmpty"               //时间不能为空
	ErrOkrCanceled                = "ErrOkrCanceled"                //O目标已取消无法操作
	ErrOkrCompleted               = "ErrOkrCompleted"               //O目标已完成无法操作
	ErrOkrScored                  = "ErrOkrScored"                  //KR已评分无法操作
	ErrOkrTimeFormat              = "ErrOkrTimeFormat"              //时间格式错误
	ErrOkrTitleLengthInvalid      = "ErrOkrTitleLengthInvalid"      //标题内容过长，请限制在255个字符以内
	ErrOkrTimeInvalid             = "ErrOkrTimeInvalid"             //开始时间不能大于结束时间，请重新选择合适的时间段
	ErrOkrOwnerNotCancel          = "ErrOkrOwnerNotCancel"          //仅限负责人操作

	// dootask
	ErrDooTaskDataFormat           = "ErrDooTaskDataFormat"           //数据格式错误
	ErrDooTaskResponseFormat       = "ErrDooTaskResponseFormat"       //响应格式错误
	ErrDooTaskRequestFailed        = "ErrDooTaskRequestFailed"        //请求失败
	ErrDooTaskUnmarshalResponse    = "ErrDooTaskUnmarshalResponse"    //解析响应失败：{{.detail}}
	ErrDooTaskRequestFailedWithErr = "ErrDooTaskRequestFailedWithErr" //请求失败：{{.detail}}
)

var (
	ErrCmdTimeout = "ErrCmdTimeout" //命令执行超时
)

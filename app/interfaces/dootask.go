package interfaces

// UserInfoResp 返回用户信息
type UserInfoResp struct {
	*UserBasicResp
	Token            string   `json:"token"`             // token
	Identity         []string `json:"identity"`          // 身份
	Tel              string   `json:"tel"`               // 手机号
	Changepass       int      `json:"changepass"`        // 是否需要修改密码 1-是 0-否
	CreatedIp        string   `json:"created_ip"`        // 创建ip
	EmailVerity      int      `json:"email_verity"`      // 邮箱是否验证 1-是 0-否
	LastAt           string   `json:"last_at"`           // 最后登录时间
	LastIp           string   `json:"last_ip"`           // 最后登录ip
	LineIp           string   `json:"line_ip"`           // 在线ip
	LoginNum         int      `json:"login_num"`         // 登录次数
	NicknameOriginal string   `json:"nickname_original"` // 昵称原始值
	TaskDialogId     int      `json:"task_dialog_id"`    // 任务对话框id
	CreatedAt        string   `json:"created_at"`        // 创建时间
}

// 用户基础信息
type UserBasicResp struct {
	Az             string `json:"az"`              // 首字母
	Bot            int    `json:"bot"`             // 是否是机器人 1-是 0-否
	Department     []int  `json:"department"`      // 部门
	DepartmentName string `json:"department_name"` // 部门名称
	DisableAt      string `json:"disable_at"`      // 禁用时间
	Userid         int    `json:"userid"`          // 用户id
	Nickname       string `json:"nickname"`        // 昵称
	Userimg        string `json:"userimg"`         // 头像
	Email          string `json:"email"`           // 邮箱
	LineAt         string `json:"line_at"`         // 在线时间
	Pinyin         string `json:"pinyin"`          // 拼音
	Profession     string `json:"profession"`      // 职业
}

// 判断是否是管理员
func (u *UserInfoResp) IsAdmin() bool {
	for _, v := range u.Identity {
		if v == "admin" {
			return true
		}
	}
	return false
}

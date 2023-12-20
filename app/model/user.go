package model

import (
	"dootask-okr/app/core"
	"dootask-okr/app/utils/common"
	"strings"

	"gorm.io/gorm"
)

type User struct {
	Userid           int    `json:"userid"`            // 用户id
	Identity         string `json:"identity"`          // 身份
	Nickname         string `json:"nickname"`          // 昵称
	Email            string `json:"email"`             // 邮箱
	Userimg          string `json:"userimg"`           // 头像
	Department       string `json:"department"`        // 部门
	Profession       string `json:"profession"`        // 职业
	Token            string `json:"token"`             // token
	DepartmentName   string `json:"department_name"`   // 部门名称
	Tel              string `json:"tel"`               // 手机号
	Pinyin           string `json:"pinyin"`            // 拼音
	Az               string `json:"az"`                // 首字母
	Bot              int    `json:"bot"`               // 是否是机器人 1-是 0-否
	Changepass       int    `json:"changepass"`        // 是否需要修改密码 1-是 0-否
	CreatedIp        string `json:"created_ip"`        // 创建ip
	EmailVerity      int    `json:"email_verity"`      // 邮箱是否验证 1-是 0-否
	LastAt           string `json:"last_at"`           // 最后登录时间
	LastIp           string `json:"last_ip"`           // 最后登录ip
	LineAt           string `json:"line_at"`           // 在线时间
	LineIp           string `json:"line_ip"`           // 在线ip
	LoginNum         int    `json:"login_num"`         // 登录次数
	NicknameOriginal string `json:"nickname_original"` // 昵称原始值
	TaskDialogId     int    `json:"task_dialog_id"`    // 任务对话框id
	CreatedAt        string `json:"created_at"`        // 创建时间
	DisableAt        string `json:"disable_at"`        // 禁用时间
}

type UserBasic struct {
	Userid     int    `json:"userid"`     // 用户id
	Identity   string `json:"identity"`   // 身份
	Nickname   string `json:"nickname"`   // 昵称
	Email      string `json:"email"`      // 邮箱
	Userimg    string `json:"userimg"`    // 头像
	Department string `json:"department"` // 部门
	Profession string `json:"profession"` // 职业
	CreatedAt  string `json:"created_at"` // 创建时间
	DisableAt  string `json:"disable_at"` // 禁用时间
}

var UserModel = User{}

// 表名
func (mb *UserBasic) TableName() string {
	return core.DBTableName(&User{})
}

// 是否为管理员
func (m *UserBasic) IsAdmin() bool {
	if m == nil {
		return false
	}
	return strings.Contains(m.Identity, "admin")
}

// 后置钩子函数
func (m *UserBasic) AfterFind(tx *gorm.DB) error {
	if m.Nickname == "" {
		m.Nickname = common.CardFormat(m.Email)
	}
	return nil
}

// 用户昵称
func (m *User) GetNickname() string {
	if m.Nickname == "" {
		return common.CardFormat(m.Email)
	}
	return m.Nickname
}

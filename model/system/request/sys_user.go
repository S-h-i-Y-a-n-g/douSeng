package request

import uuid "github.com/satori/go.uuid"

// User register structure
type Register struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	NickName     string `json:"nickName" gorm:"default:'GSDUser'"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	DeptId       uint   `json:"deptId"`
	AuthorityIds []uint `json:"authorityIds"`
}

// User login structure
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// Modify password structure
type ChangePasswordStruct struct {
	Username    string `json:"username"`    // 用户名
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// Modify  user's auth structure
type SetUserAuth struct {
	AuthorityId uint `json:"authorityId"` // 角色ID
}

// Modify  user's auth structure
type SetUserAuthorities struct {
	ID           uint   `json:"id"`
	AuthorityIds []uint `json:"authorityIds"` // 角色ID
}

// Modify  user's Info structure
type SetUserInfo struct {
	ID           uint      `json:"ID"`           //用户id
	DeptId       uint      `json:"deptId"`       //部门id
	HeadImg      string    `json:"headerImg"`    //头像
	NickName     string    `json:"nickName"`     //昵称
	Phone        string    `json:"phone"`        //手机号
	Email        string    `json:"email"`        //邮箱
	UUID         uuid.UUID `json:"uuid"`         //uuid
	AuthorityIds []uint    `json:"authorityIds"` //角色ID
}

// Modify  user's Info structure
type SetUserDept struct {
	ID     uint      `json:"ID"`     //用户id
	DeptId uint      `json:"deptId"` //部门id
	UUID   uuid.UUID `json:"uuid"`   //uuid
}

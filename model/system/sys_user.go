package system

import (
	"project/global"

	uuid "github.com/satori/go.uuid"
)

type SysUser struct {
	global.GSD_MODEL
	UUID        uuid.UUID      `json:"uuid" gorm:"not null;comment:用户UUID"`
	Username    string         `json:"userName" gorm:"not null;comment:用户登录名"`                               // 用户登录名
	Password    string         `json:"-"  gorm:"not null;comment:用户登录密码"`                                    // 用户登录密码
	NickName    string         `json:"nickName" gorm:"not null;default:系统用户;comment:用户昵称"`                   // 用户昵称
	HeaderImg   string         `json:"headerImg" gorm:"not null;default:uploads/file/head.jpg;comment:用户头像"` // 用户头像
	Email       string         `json:"email" gorm:"not null;comment:用户邮箱"`
	Phone       string         `json:"phone" gorm:"comment:用户手机号"`
	SideMode    string         `json:"sideMode" gorm:"default:dark;comment:用户角色ID"`       // 用户侧边主题
	ActiveColor string         `json:"activeColor" gorm:"default:#1890ff;comment:用户角色ID"` // 活跃颜色
	BaseColor   string         `json:"baseColor" gorm:"default:#fff;comment:用户角色ID"`      // 基础颜色
	AuthorityId uint           `json:"authorityId" gorm:"default:1;comment:用户角色ID"`       // 用户角色ID
	DeptId      uint           `json:"deptId" gorm:"comment:用户所属部门"`
	Dept        SysDept        `json:"dept" gorm:"foreignKey:DeptId;references:ID;"`
	Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
}

type ExcelInfo struct {
	FileName string    `json:"fileName"` // 文件名
	InfoList []SysUser `json:"infoList"`
}

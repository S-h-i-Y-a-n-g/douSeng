package system

type SysAuthorityDept struct {
	SysDeptId               uint   `gorm:"not null;column:sys_dept_id"`
	SysAuthorityAuthorityId string `gorm:"not null;column:sys_authority_authority_id"`
}

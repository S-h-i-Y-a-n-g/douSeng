package system

type SysMenu struct {
	SysBaseMenu
	MenuId      string                 `json:"menuId" gorm:"not null;comment:菜单ID"`
	AuthorityId uint                   `json:"-" gorm:"not null;comment:角色ID"`
	Children    []SysMenu              `json:"children" gorm:"-"`
	Parameters  []SysBaseMenuParameter `json:"parameters" gorm:"foreignKey:SysBaseMenuID;references:MenuId"`
}

func (s SysMenu) TableName() string {
	return "authority_menu"
}

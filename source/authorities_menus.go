package source

import (
	"project/global"

	"github.com/gookit/color"
	"gorm.io/gorm"
)

var AuthoritiesMenus = new(authoritiesMenus)

type authoritiesMenus struct{}

type AuthorityMenus struct {
	AuthorityId uint `gorm:"column:sys_authority_authority_id"`
	BaseMenuId  uint `gorm:"column:sys_base_menu_id"`
}

var authorityMenus = []AuthorityMenus{
	{1, 1},
	{1, 2},
	{1, 3},
	{1, 4},
	{1, 5},
	{1, 6},
	{1, 7},
	{1, 8},
	{1, 9},
	{1, 10},
	{1, 11},
}

// Init @author: [chenguanglan](https://github.com/sFFbLL)
//@description: sys_authority_menus 表数据初始化
func (a *authoritiesMenus) Init() error {
	return global.GSD_DB.Table("sys_authority_menus").Transaction(func(tx *gorm.DB) error {
		if tx.Where("sys_authority_authority_id IN (1)").Find(&[]AuthorityMenus{}).RowsAffected > 10 {
			color.Danger.Println("\n[Mysql] --> sys_authority_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&authorityMenus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_authority_menus 表初始数据成功!")
		return nil
	})
}

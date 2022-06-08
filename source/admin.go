package source

import (
	"project/global"
	"project/model/system"
	"time"

	"github.com/gookit/color"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

var Admin = new(admin)

type admin struct{}

var admins = []system.SysUser{
	{GSD_MODEL: global.GSD_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, UUID: uuid.NewV4(), Username: "admin", Password: "832c5ba281a74faae637f066243f1848", NickName: "超级管理员", HeaderImg: "http://r9qsta3s9.hn-bkt.clouddn.com/head.jpg", Phone: "15083138896", DeptId: 1, AuthorityId: 1},
}

// Init @author: [chenguanglan](https://github.com/sFFbLL)
//@description: sys_users 表数据初始化
func (a *admin) Init() error {
	return global.GSD_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]system.SysUser{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> sys_users 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&admins).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_users 表初始数据成功!")
		return nil
	})
}

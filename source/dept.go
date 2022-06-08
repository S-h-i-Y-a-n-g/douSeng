package source

import (
	"project/global"
	"project/model/system"
	"time"

	"github.com/gookit/color"
	"gorm.io/gorm"
)

var Dept = new(dept)

type dept struct{}

var deptList = []system.SysDept{
	{GSD_MODEL: global.GSD_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, ParentID: "0", DeptName: "顶级部门"},
}

// Init @author: [chenguanglan](https://github.com/sFFbLL)
//@description: sys_users 表数据初始化
func (a *dept) Init() error {
	return global.GSD_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]system.SysDept{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> sys_depts 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&deptList).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_depts 表初始数据成功!")
		return nil
	})
}

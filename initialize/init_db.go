package initialize

import (
	"os"
	"project/global"
	"project/model/system"
	"project/source"

	"go.uber.org/zap"
)

// InitDB InitDB
//@author: [chenguanglan](https://github.com/sFFbLL)
//@function: MysqlTables
//@description: 初始化数据库表专用
//@param: db *gorm.DB
func InitDB() {
	err := initDB(
		source.Admin,
		source.BaseMenu,
		source.Authority,
		source.AuthoritiesMenus,
		source.AuthorityMenu,
		source.Dept,
		source.UserAuthority,
		source.Api,
	)
	if err != nil {
		global.GSD_LOG.Error("init table data failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.GSD_LOG.Info("init table data success")
}

//@author: [chenguanglan](https://github.com/sFFbLL)
//@function: initDB
//@description: 初始化超级管理员数据
//@param: InitDBFunctions system.InitDBFunc 初始化方法
//@return: error
func initDB(InitDBFunctions ...system.InitDBFunc) (err error) {
	for _, v := range InitDBFunctions {
		err = v.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

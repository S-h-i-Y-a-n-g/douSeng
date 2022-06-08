package source

import (
	"project/global"
	"strconv"

	"github.com/gookit/color"
)

var Casbin = new(casbin)

type casbin struct{}

var carbines [][]string

// Init @author: [chenguanglan](https://github.com/sFFbLL)
//@description: casbin_rule 表数据初始化
func (c *casbin) Init() error {
	roles := global.GSD_Casbin.GetFilteredPolicy(0, strconv.Itoa(1))
	if len(roles) != 0 {
		color.Danger.Println("\n[Mysql] --> casbin_rule 表的初始数据已存在!")
		return nil
	}
	for _, sysApi := range apis {
		carbines = append(carbines, []string{
			"1", sysApi.Path, sysApi.Method,
		})
	}
	success, err := global.GSD_Casbin.AddPolicies(carbines)
	if !success {
		return err
	}
	color.Info.Println("\n[Mysql] --> casbin_rule 表初始数据成功!")
	return nil
}

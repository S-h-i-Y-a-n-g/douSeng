package utils

import (
	"fmt"
	"math"
	"project/model/system"
	"project/model/system/request"

	"golang.org/x/tools/container/intsets"
)

type DataScope struct {
}

// GetDataScope @author: [chenguanglan](https://github.com/sFFbLL)
//@function: GetDataScope
//@description: 获取数据权限
//@param: user request.UserCache
//@return: dataScope []uint 部门id, isAll bool 是否为全部
func (DataScope) GetDataScope(user *request.UserCache) (dataScope []uint, isAll bool) {
	keyMap, all := getDataScopeMap(user)
	if all {
		return dataScope, true
	}
	for deptId := range keyMap {
		dataScope = append(dataScope, deptId)
	}
	return dataScope, false
}

// CanDoToTargetUser @author: [chenguanglan](https://github.com/sFFbLL)
//@function: CanDoToTargetUser
//@description: 是否有权操作该数据
//@param: user *request.UserCache, users []system.SysUser 操作用户对象
//@return: bool 是否有权操作对象
func (d DataScope) CanDoToTargetUser(user *request.UserCache, users []*system.SysUser) bool {
	//校验等级
	maxLevel := d.GetMaxLevel(user.Authority)
	for _, user := range users {
		if d.GetMaxLevel(user.Authorities) < maxLevel {
			fmt.Println(d.GetMaxLevel(user.Authorities), maxLevel)
			return false
		}
	}
	//校验dataScope
	keyMap, all := getDataScopeMap(user)
	if !all {
		for _, user := range users {
			if _, ok := keyMap[user.DeptId]; !ok {
				return false
			}
		}
	}
	return true
}

// GetMaxLevel @author: [chenguanglan](https://github.com/sFFbLL)
//@function: GetMaxLevel
//@description: 获取最角色最大等级(实际为最小值)
//@param: roles []system.SysAuthority 角色列表
//@return: int 所有角色中的最高等级
func (DataScope) GetMaxLevel(roles []system.SysAuthority) (maxLevel uint) {
	maxLevel = uint(intsets.MaxInt)
	for _, role := range roles {
		maxLevel = uint(math.Min(float64(role.Level), float64(maxLevel)))
	}
	return maxLevel
}

//@author: [chenguanglan](https://github.com/sFFbLL)
//@function: getDataScopeMap
//@description: 获取数据权限
//@param: user request.UserCache
//@return: dataScope []uint 部门id, isAll bool 是否为全部
func getDataScopeMap(user *request.UserCache) (keyMap map[uint]uint, isAll bool) {
	keyMap = make(map[uint]uint, 0)
	for _, authority := range user.Authority {
		if authority.DataScope == "全部" {
			return keyMap, true
		} else if authority.DataScope == "本级" {
			keyMap[user.DeptId] = 1
		} else {
			for _, dept := range authority.Depts {
				keyMap[dept.ID] = 1
			}
		}
	}
	return keyMap, false
}

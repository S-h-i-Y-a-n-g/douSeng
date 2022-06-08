package system

import (
	v1 "project/api"

	"github.com/gin-gonic/gin"
)

type MenuRouter struct {
}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	menuRouter := Router.Group("menu")
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.AuthorityMenuApi
	{
		menuRouter.POST("getMenuList", authorityMenuApi.GetMenuList)           // 分页获取基础menu列表
		menuRouter.POST("addBaseMenu", authorityMenuApi.AddMenu)               //新增菜单
		menuRouter.POST("deleteBaseMenu", authorityMenuApi.DeleteMenu)         //删除菜单
		menuRouter.POST("updateBaseMenu", authorityMenuApi.UpdateMenu)         //修改菜单
		menuRouter.POST("getMenu", authorityMenuApi.GetUserMenuTree)           //当前用户菜单
		menuRouter.POST("getBaseMenuTree", authorityMenuApi.GetBaseMenuTree)   // 获取用户动态路由
		menuRouter.POST("getMenuAuthority", authorityMenuApi.GetMenuAuthority) // 获取指定角色menu
		menuRouter.POST("getBaseMenuById", authorityMenuApi.GetBaseMenuById)   // 根据id获取菜单
		menuRouter.POST("addMenuAuthority", authorityMenuApi.AddMenuAuthority) //	增加menu和角色关联关系
	}
}

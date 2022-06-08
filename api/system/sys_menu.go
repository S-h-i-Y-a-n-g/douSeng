package system

import (
	"project/global"
	"project/model/common/request"
	"project/model/common/response"
	"project/model/system"
	systemReq "project/model/system/request"
	systemRes "project/model/system/response"
	"project/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityMenuApi struct {
}

// @Tags Menu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取基础menu列表,返回包括列表,总数,页码,每页数量"
// @Router /api/menu/getMenuList [post]
func (a *AuthorityMenuApi) GetMenuList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, list, total := menuService.GetMenuList(); err != nil {
		global.GSD_LOG.Error("获取失败", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// @Tags Menu
// @Summary 菜单管理-新增菜单
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body system.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {object} response.Response{msg=string} "新增菜单"
// @Router /api/menu/addBaseMenu [post]
func (a *AuthorityMenuApi) AddMenu(c *gin.Context) {
	var menu system.SysBaseMenu
	_ = c.ShouldBindJSON(&menu)
	if err := utils.Verify(menu, utils.MenuVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(menu.Meta, utils.MenuMetaVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuService.AddMenu(menu); err != nil {
		global.GSD_LOG.Error("添加菜单失败", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("添加菜单失败", c)
	} else {
		response.OkWithMessage("添加菜单成功", c)
	}
}

// @Tags Menu
// @Summary 菜单管理-删除菜单
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {object} response.Response{msg=string} "删除菜单"
// @Router /api/menu/deleteBaseMenu [post]
func (a *AuthorityMenuApi) DeleteMenu(c *gin.Context) {
	var menu request.GetById
	_ = c.ShouldBindJSON(&menu)
	if err := utils.Verify(menu, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuService.DeleteMenu(menu.ID); err != nil {
		global.GSD_LOG.Error("删除菜单失败", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("菜单删除失败", c)
		return
	} else {
		response.OkWithMessage("删除菜单成功", c)
	}
}

// @Tags Menu
// @Summary 菜单管理-更新菜单
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body system.SysBaseMenu true "路由path, 父菜单ID, 路由name, 对应前端文件路径, 排序标记"
// @Success 200 {object} response.Response{msg=string} "更新菜单"
// @Router /api/menu/updateBaseMenu [post]
func (a *AuthorityMenuApi) UpdateMenu(c *gin.Context) {
	var menu system.SysBaseMenu
	_ = c.ShouldBindJSON(&menu)
	if err := utils.Verify(menu, utils.MenuVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := utils.Verify(menu.Meta, utils.MenuMetaVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuService.UpdateMenu(menu); err != nil {
		global.GSD_LOG.Error("更新菜单失败", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("菜单更新失败", c)
	} else {
		response.OkWithMessage("菜单更新成功", c)
	}
}

// @Tags Menu
// @Summary 获取当前用户菜单
// @Security ApiKeyAuth
// @Produce application/json
// @Param data body request.Empty true "空"
// @Success 200 {object} response.Response{data=systemRes.SysBaseMenusResponse,msg=string} "获取当前用户菜单,返回包括系统菜单列表"
// @Router /api/menu/getMenu [post]
func (a *AuthorityMenuApi) GetUserMenuTree(c *gin.Context) {
	if err, menus := menuService.GetUserMenu(utils.GetUserAuthority(c)); err != nil {
		global.GSD_LOG.Error("获取失败!", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("获取失败", c)
	} else {
		if menus == nil {
			menus = []system.SysMenu{}
		}
		response.OkWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取成功", c)
	}
}

// @Tags Menu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.Empty true "空"
// @Success 200 {object} response.Response{data=systemRes.SysBaseMenusResponse,msg=string} "获取用户动态路由,返回包括系统菜单列表"
// @Router /api/menu/getBaseMenuTree [post]
func (a *AuthorityMenuApi) GetBaseMenuTree(c *gin.Context) {
	if err, menus := menuService.GetBaseMenuTree(); err != nil {
		global.GSD_LOG.Error("获取失败!", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysBaseMenusResponse{Menus: menus}, "获取成功", c)
	}
}

// @Tags Menu
// @Summary 获取指定角色menu
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetAuthorityId true "角色ID"
// @Success 200 {object} response.Response{data=map[string]interface{},msg=string} "获取指定角色menu"
// @Router /api/menu/getMenuAuthority [post]
func (a *AuthorityMenuApi) GetMenuAuthority(c *gin.Context) {
	var param request.GetAuthorityId
	_ = c.ShouldBindJSON(&param)
	if err := utils.Verify(param, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, menus := menuService.GetMenuAuthority(&param); err != nil {
		global.GSD_LOG.Error("获取失败!", zap.Error(err), utils.GetRequestID(c))
		response.FailWithDetailed(systemRes.SysMenusResponse{Menus: menus}, "获取失败", c)
	} else {
		response.OkWithDetailed(gin.H{"menus": menus}, "获取成功", c)
	}
}

// @Tags Menu
// @Summary 根据id获取菜单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "菜单id"
// @Success 200 {object} response.Response{data=systemRes.SysBaseMenuResponse,msg=string} "根据id获取菜单,返回包括系统菜单列表"
// @Router /api/menu/getBaseMenuById [post]
func (a *AuthorityMenuApi) GetBaseMenuById(c *gin.Context) {
	var idInfo request.GetById
	_ = c.ShouldBindJSON(&idInfo)
	if err := utils.Verify(idInfo, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, menu := menuService.GetBaseMenuById(idInfo.ID); err != nil {
		global.GSD_LOG.Error("获取失败!", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysBaseMenuResponse{Menu: menu}, "获取成功", c)
	}
}

// @Tags AuthorityMenu
// @Summary 增加menu和角色关联关系
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.AddMenuAuthorityInfo true "角色ID"
// @Success 200 {object} response.Response{msg=string} "增加menu和角色关联关系"
// @Router /api/menu/addMenuAuthority [post]
func (a *AuthorityMenuApi) AddMenuAuthority(c *gin.Context) {
	var authorityMenu systemReq.AddMenuAuthorityInfo
	_ = c.ShouldBindJSON(&authorityMenu)
	if err := utils.Verify(authorityMenu, utils.AuthorityIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuService.AddMenuAuthority(authorityMenu.Menus, authorityMenu.AuthorityId); err != nil {
		global.GSD_LOG.Error("添加失败!", zap.Error(err), utils.GetRequestID(c))
		response.FailWithMessage("添加失败", c)
	} else {
		response.OkWithMessage("添加成功", c)
	}
}

package source

import (
	"project/global"
	"project/model/system"
	"time"

	"github.com/gookit/color"

	"gorm.io/gorm"
)

var Api = new(api)

type api struct{}

var apis = []system.SysApi{
	{global.GSD_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/api/base/login", "用户登录（必选）", "base", "POST"},
	{global.GSD_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/user/register", "用户注册（必选）", "user", "POST"},
	{global.GSD_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/api/createApi", "创建api", "api", "POST"},
	{global.GSD_MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/api/getApiList", "获取api列表", "api", "POST"},
	{global.GSD_MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/api/getApiById", "获取api详细信息", "api", "POST"},
	{global.GSD_MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/api/deleteApi", "删除Api", "api", "POST"},
	{global.GSD_MODEL{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/api/updateApi", "更新Api", "api", "POST"},
	{global.GSD_MODEL{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/api/getAllApis", "获取所有api", "api", "POST"},
	{global.GSD_MODEL{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/authority/createAuthority", "创建角色", "authority", "POST"},
	{global.GSD_MODEL{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/authority/deleteAuthority", "删除角色", "authority", "POST"},
	{global.GSD_MODEL{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/authority/getAuthorityList", "获取角色列表", "authority", "POST"},
	{global.GSD_MODEL{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/menu/getMenu", "获取菜单树（必选）", "menu", "POST"},
	{global.GSD_MODEL{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/menu/getMenuList", "分页获取基础menu列表", "menu", "POST"},
	{global.GSD_MODEL{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/menu/addBaseMenu", "新增菜单", "menu", "POST"},
	{global.GSD_MODEL{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/menu/getBaseMenuTree", "获取用户动态路由", "menu", "POST"},
	{global.GSD_MODEL{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/menu/addMenuAuthority", "增加menu和角色关联关系", "menu", "POST"},
	{global.GSD_MODEL{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/menu/getMenuAuthority", "获取指定角色menu", "menu", "POST"},
	{global.GSD_MODEL{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/menu/deleteBaseMenu", "删除菜单", "menu", "POST"},
	{global.GSD_MODEL{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/menu/updateBaseMenu", "更新菜单", "menu", "POST"},
	{global.GSD_MODEL{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/menu/getBaseMenuById", "根据id获取菜单", "menu", "POST"},
	{global.GSD_MODEL{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/user/changePassword", "修改密码（建议选择）", "user", "POST"},
	{global.GSD_MODEL{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/user/getUserList", "获取用户列表", "user", "POST"},
	{global.GSD_MODEL{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/user/setUserAuthority", "修改用户角色（必选）", "user", "POST"},
	{global.GSD_MODEL{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/fileUploadAndDownload/upload", "文件上传示例", "fileUploadAndDownload", "POST"},
	{global.GSD_MODEL{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/fileUploadAndDownload/getFileList", "获取上传文件列表", "fileUploadAndDownload", "POST"},
	{global.GSD_MODEL{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/casbin/updateCasbin", "更改角色api权限", "casbin", "POST"},
	{global.GSD_MODEL{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/casbin/getPolicyPathByAuthorityId", "获取权限列表", "casbin", "POST"},
	{global.GSD_MODEL{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/fileUploadAndDownload/deleteFile", "删除文件", "fileUploadAndDownload", "POST"},
	{global.GSD_MODEL{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/jwt/jsonInBlacklist", "jwt加入黑名单(退出，必选)", "jwt", "POST"},
	{global.GSD_MODEL{ID: 30, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/authority/updateAuthority", "更新角色信息", "authority", "PUT"},
	{global.GSD_MODEL{ID: 31, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/user/deleteUser", "删除用户", "user", "DELETE"},
	{global.GSD_MODEL{ID: 32, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/sysOperationRecord/createSysOperationRecord", "新增操作记录", "sysOperationRecord", "POST"},
	{global.GSD_MODEL{ID: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/sysOperationRecord/deleteSysOperationRecord", "删除操作记录", "sysOperationRecord", "DELETE"},
	{global.GSD_MODEL{ID: 34, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/sysOperationRecord/findSysOperationRecord", "根据ID获取操作记录", "sysOperationRecord", "GET"},
	{global.GSD_MODEL{ID: 35, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/sysOperationRecord/getSysOperationRecordList", "获取操作记录列表", "sysOperationRecord", "GET"},
	{global.GSD_MODEL{ID: 36, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/sysOperationRecord/deleteSysOperationRecordByIds", "批量删除操作历史", "sysOperationRecord", "DELETE"},
	{global.GSD_MODEL{ID: 37, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/user/setUserInfo", "设置用户信息（必选）", "user", "PUT"},
	{global.GSD_MODEL{ID: 38, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/user/setSelfInfo", "设置当前用户信息（必选）", "user", "PUT"},
	{global.GSD_MODEL{ID: 39, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/api/deleteApisByIds", "批量删除api", "api", "DELETE"},
	{global.GSD_MODEL{ID: 40, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/user/setUserAuthorities", "设置权限组", "user", "POST"},
	{global.GSD_MODEL{ID: 41, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/user/getUserInfo", "获取自身信息（必选）", "user", "GET"},
	{global.GSD_MODEL{ID: 42, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/department/addDept", "新增部门（必选）", "department", "POST"},
	{global.GSD_MODEL{ID: 43, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/department/deleteDept", "删除部门（必选）", "department", "POST"},
	{global.GSD_MODEL{ID: 44, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/department/updateDept", "更新部门（必选）", "department", "POST"},
	{global.GSD_MODEL{ID: 45, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/department/lists", "查询部门列表（必选）", "department", "POST"},
	{global.GSD_MODEL{ID: 46, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/department/id", "根据parentID查询部门（必选）", "department", "POST"},
	{global.GSD_MODEL{ID: 47, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/user/importExcel", "导出数据（必选）", "user", "POST"},
	{global.GSD_MODEL{ID: 48, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/user/loadExcel", "加载用户数据（必选）", "user", "GET"},
	{global.GSD_MODEL{ID: 49, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/user/exportExcel", "导出用户数据（必选）", "user", "POST"},
	{global.GSD_MODEL{ID: 50, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/user/downloadTemplate", "下载模板数据（必选）", "user", "GET"},
	{global.GSD_MODEL{ID: 51, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/system/getServerInfo", "获取服务器信息", "system", "POST"},
	{global.GSD_MODEL{ID: 52, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/system/getSystemConfig", "获取配置文件内容", "system", "POST"},
	{global.GSD_MODEL{ID: 53, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/system/setSystemConfig", "设置配置文件内容", "system", "POST"},
	{global.GSD_MODEL{ID: 54, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/user/getUserByAuthority", "根据角色id获取用户信息", "user", "POST"},
	{global.GSD_MODEL{ID: 55, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/department/users", "根据部门获取子部门和该部门下的用户", "department", "POST"},
}

// Init @author: [chenguanglan](https://github.com/sFFbLL)
//@description: sys_apis 表数据初始化
func (a *api) Init() error {
	return global.GSD_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 53}).Find(&[]system.SysApi{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_apis 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&apis).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_apis 表初始数据成功!")
		return nil
	})
}

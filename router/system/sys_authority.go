package system

import (
	v1 "project/api"
	"project/middleware"

	"github.com/gin-gonic/gin"
)

type AuthorityRouter struct {
}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("authority").Use(middleware.OperationRecord())
	var authorityApi = v1.ApiGroupApp.SystemApiGroup.AuthorityApi
	{
		authorityRouter.POST("createAuthority", authorityApi.CreateAuthority)   // 创建角色
		authorityRouter.POST("deleteAuthority", authorityApi.DeleteAuthority)   // 删除角色
		authorityRouter.POST("updateAuthority", authorityApi.UpdateAuthority)   // 更新角色
		authorityRouter.POST("getAuthorityList", authorityApi.GetAuthorityList) // 角色列表分页
	}
}

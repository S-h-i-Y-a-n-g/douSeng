package initialize

import (
	"net/http"
	_ "project/docs"
	"project/global"
	"project/middleware"
	"project/router"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.New()
	Router.Use(middleware.GinLogger(), middleware.GinRecovery(true))
	Router.StaticFS("/api/"+global.GSD_CONFIG.Local.Path, http.Dir(global.GSD_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	global.GSD_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.GSD_LOG.Info("use middleware cors")
	Router.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GSD_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	//获取路由组实例
	systemRouter := router.RouterGroupApp.System
	PublicGroup := Router.Group("api")
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	}
	PrivateGroup := Router.Group("api")
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	{
		systemRouter.InitApiRouter(PrivateGroup)                // 注册功能api路由
		systemRouter.InitJwtRouter(PrivateGroup)                // jwt相关路由
		systemRouter.InitUserRouter(PrivateGroup)               // 注册用户路由
		systemRouter.InitMenuRouter(PrivateGroup)               // 注册menu路由
		systemRouter.InitDeptRouter(PrivateGroup)               //注册部门路由
		systemRouter.InitSystemRouter(PrivateGroup)             // system相关路由
		systemRouter.InitCasbinRouter(PrivateGroup)             // 权限相关路由
		systemRouter.InitAuthorityRouter(PrivateGroup)          // 注册角色路由
		systemRouter.InitSysOperationRecordRouter(PrivateGroup) // 操作记录
		//systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)    // 字典详情管理
		systemRouter.InitFileRouter(PrivateGroup) //文件操作

	}
	//注册抖声路由组实例
	douSengRouter := router.RouterGroupApp.DouSeng
	//公共路由组，加一个抖音前缀
	PublicGroup1 := Router.Group("douyin")
	{
		douSengRouter.DouSengPRouter(PublicGroup1) //p接口
		douSengRouter.RelationRouter(PublicGroup1)
		douSengRouter.DouSengXRouter(PublicGroup1)
		douSengRouter.DouShengWRouter(PublicGroup1) //评论接口
	}

	//两个未知接口简单返回一下
	PublicGroup2 := Router.Group("")
	{
		douSengRouter.DouSengPRouter2(PublicGroup2)
	}

	global.GSD_LOG.Info("router register success")
	return Router
}

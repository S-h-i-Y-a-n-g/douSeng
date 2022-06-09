package douseng

import (
	"github.com/gin-gonic/gin"
	v1 "project/api"
)

type DouSengPJHRouter struct{}

func (u *DouSengPJHRouter) DouSengPRouter(Router *gin.RouterGroup) {
	//设置路由组
	douSengPJHRouter := Router.Group("")
	//具体路由
	douSengPJHApi := v1.ApiGroupApp.DouSengApiGroup.DouSengPJHApi
	{
		douSengPJHRouter.GET("feed",douSengPJHApi.Feed)//视频流接口
	}
	//设置路由组
	douSengPJHRouter2 := Router.Group("user")
	douSengPJHApi2 := v1.ApiGroupApp.DouSengApiGroup.DouSengPJHApi
	{
		douSengPJHRouter2.POST("login/",douSengPJHApi2.DouSengLogin)//登陆接口
		douSengPJHRouter2.GET("/",douSengPJHApi2.GetUserInfo)//获取用户信息接口
	}

}
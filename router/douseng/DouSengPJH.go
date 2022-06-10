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
	{
		douSengPJHRouter2.POST("login/",douSengPJHApi.DouSengLogin)//登陆接口
		douSengPJHRouter2.GET("/",douSengPJHApi.GetUserInfo)//获取用户信息接口
		douSengPJHRouter2.POST("register/",douSengPJHApi.DouSengRegister)//注册接口

	}
	douSengPJHRouter3 := Router.Group("publish")
	{
		douSengPJHRouter3.POST("action/",douSengPJHApi.DouSengPublishVideo)//上传视频接口
		douSengPJHRouter3.GET("list/",douSengPJHApi.GetUserFeed)//用户视频列表接口
	}
	douSengPJHRouter4 := Router.Group("favorite")
	{
		douSengPJHRouter4.GET("list/",douSengPJHApi.GetUserFavoriteFeed)//用户点赞视频接口/douyin/favorite/list/
	}
	douSengPJHRouter5 := Router.Group("")
	{
		douSengPJHRouter5.GET("favicon.ico",douSengPJHApi.BZD)//不知道干嘛的接口
		douSengPJHRouter5.GET("/",douSengPJHApi.BZD)//不知道干嘛的接口
	}
}
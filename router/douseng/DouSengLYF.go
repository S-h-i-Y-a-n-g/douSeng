package douseng

import (
	"github.com/gin-gonic/gin"
	v1 "project/api"
)

type DouSengLYFRouter struct{}

func (u *DouSengLYFRouter) RelationRouter(Router *gin.RouterGroup) {
	//设置路由组
	douSengLYFRouter := Router.Group("relation")
	//具体路由
	douSengLYFApi := v1.ApiGroupApp.DouSengApiGroup.DouSengLYFApi
	{
		douSengLYFRouter.GET("follow/list", douSengLYFApi.FollowList)     //关注列表
		douSengLYFRouter.GET("follower/list", douSengLYFApi.FollowerList) //粉丝列表
		douSengLYFRouter.POST("action", douSengLYFApi.Action)             //关注操作
	}
}

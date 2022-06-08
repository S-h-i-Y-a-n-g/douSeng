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
		douSengPJHRouter.GET("feed",douSengPJHApi.Feed)//测试接口

	}
}
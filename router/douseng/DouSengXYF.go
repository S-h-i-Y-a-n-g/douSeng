package douseng

import (
	"github.com/gin-gonic/gin"
	v1 "project/api"
)

type DouSengXYFRouter struct{}

func (u *DouSengXYFRouter) DouSengXRouter(Router *gin.RouterGroup) {
	//设置路由组
	douSengXYFRouter := Router.Group("favorite")

	douSengXYFApi := v1.ApiGroupApp.DouSengApiGroup.DouSengXYFApi
	{
		douSengXYFRouter.POST("action/", douSengXYFApi.Action)	//点赞操作
		//douSengXYFRouter.GET("list/", douSengXYFApi.Action)	//点赞列表
	}

}

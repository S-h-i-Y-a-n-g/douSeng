package douseng

import (
	"github.com/gin-gonic/gin"
	v1 "project/api"
)

type DouSengXYFRouter struct{}

func (u *DouSengXYFRouter) DouSengXYFRouter(Router *gin.RouterGroup) {
	//设置路由组
	douSengXYFRouter := Router.Group("")

	douSengXYFApi := v1.ApiGroupApp.DouSengApiGroup.DouSengXYFApi
	{
		douSengXYFRouter.GET("favorite/action", douSengXYFApi.Action)
	}

}

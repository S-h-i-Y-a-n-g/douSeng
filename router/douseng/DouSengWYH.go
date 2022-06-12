package douseng

import (
	"github.com/gin-gonic/gin"
	v1 "project/api"
)

type DouShengWYHRouter struct{}

func (u *DouShengWYHRouter) DouShengWRouter(Router *gin.RouterGroup) {
	douShengWYHRouter := Router.Group("comment")
	douShengWYHPJHApi := v1.ApiGroupApp.DouSengApiGroup.DouShengWYHApi
	{
		douShengWYHRouter.GET("list/", douShengWYHPJHApi.CommentList)      //评论列表
		douShengWYHRouter.POST("action/", douShengWYHPJHApi.CommentAction) //评论操作
	}

}

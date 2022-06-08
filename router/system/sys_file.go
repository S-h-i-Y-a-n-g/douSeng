package system

import (
	"github.com/gin-gonic/gin"
	v1 "project/api"
	"project/middleware"
)

type FileRouter struct {
}

func (s *FileRouter) InitFileRouter(Router *gin.RouterGroup) {
	fileRouter := Router.Group("file").Use(middleware.OperationRecord())
	var fileApi = v1.ApiGroupApp.SystemApiGroup.SysFileUploadAndDownloadApi
	{
		fileRouter.POST("upload", fileApi.UploadFile)
	}
}

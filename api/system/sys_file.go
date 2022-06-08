package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"project/global"
	"project/model/common/response"
	"project/model/system"
	systemRes "project/model/system/response"
	"project/utils"
)

type SysFileUploadAndDownloadApi struct {
}

// @Tags ExaFileUploadAndDownload
// @Summary 上传文件
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @Router /api/file/upload [post]
func (u *SysFileUploadAndDownloadApi) UploadFile(c *gin.Context) {
	var file system.SysFileUploadAndDownload
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		global.GSD_LOG.Error("接收文件失败!", zap.Any("err", err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	err, file = fileUploadAndDownloadService.UploadFile(header, noSave) // 文件上传后拿到文件路径
	if err != nil {
		global.GSD_LOG.Error("修改数据库链接失败!", zap.Any("err", err), utils.GetRequestID(c))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysFileResponse{File: file}, "上传成功", c)
}

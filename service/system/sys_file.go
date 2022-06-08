package system

import (
	"mime/multipart"
	"project/global"
	"project/model/system"
	"project/utils/upload"
	"strings"
)

type FileUploadAndDownloadService struct {
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: Upload
//@description: 创建文件上传记录
//@param: file model.ExaFileUploadAndDownload
//@return: error

func (e *FileUploadAndDownloadService) Upload(file system.SysFileUploadAndDownload) error {
	return global.GSD_DB.Create(&file).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UploadFile
//@description: 根据配置文件判断是文件上传到本地或者七牛云
//@param: header *multipart.FileHeader, noSave string
//@return: err error, file model.ExaFileUploadAndDownload

func (e *FileUploadAndDownloadService) UploadFile(header *multipart.FileHeader, noSave string) (err error, file system.SysFileUploadAndDownload) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		err = uploadErr
	}
	if noSave == "0" {
		s := strings.Split(header.Filename, ".")
		f := system.SysFileUploadAndDownload{
			Url:  filePath,
			Name: header.Filename,
			Tag:  s[len(s)-1],
			Key:  key,
		}
		return e.Upload(f), f
	}
	return
}

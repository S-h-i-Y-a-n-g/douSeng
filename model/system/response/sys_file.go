package response

import "project/model/system"

type SysFileResponse struct {
	File system.SysFileUploadAndDownload `json:"file"`
}

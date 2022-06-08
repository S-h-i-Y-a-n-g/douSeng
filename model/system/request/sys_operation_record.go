package request

import (
	"project/model/common/request"
	"project/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}

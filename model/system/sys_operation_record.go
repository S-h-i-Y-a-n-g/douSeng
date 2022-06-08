// 自动生成模板SysOperationRecord
package system

import (
	"project/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type SysOperationRecord struct {
	global.GSD_MODEL
	Ip           string        `json:"ip" form:"ip" gorm:"not null;column:ip;comment:请求ip"`                                   // 请求ip
	Method       string        `json:"method" form:"method" gorm:"not null;column:method;comment:请求方法"`                       // 请求方法
	Path         string        `json:"path" form:"path" gorm:"not null;column:path;comment:请求路径"`                             // 请求路径
	Status       int           `json:"status" form:"status" gorm:"not null;column:status;comment:请求状态"`                       // 请求状态
	Latency      time.Duration `json:"latency" form:"latency" gorm:"not null;column:latency;comment:延迟" swaggertype:"string"` // 延迟
	Agent        string        `json:"agent" form:"agent" gorm:"not null;column:agent;comment:代理"`                            // 代理
	ErrorMessage string        `json:"error_message" form:"error_message" gorm:"not null;column:error_message;comment:错误信息"`  // 错误信息
	Body         string        `json:"body" form:"body" gorm:"not null;type:longtext;column:body;comment:请求Body"`             // 请求Body
	Resp         string        `json:"resp" form:"resp" gorm:"not null;type:longtext;column:resp;comment:响应Body"`             // 响应Body
	UserID       int           `json:"user_id" form:"user_id" gorm:"not null;column:user_id;comment:用户id"`                    // 用户id
	User         SysUser       `json:"user"`
}

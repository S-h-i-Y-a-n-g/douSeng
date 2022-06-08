package system

import (
	"project/service"
	"project/utils"
)

type ApiGroup struct {
	BaseApi
	JwtApi
	DeptApi
	AuthorityMenuApi
	AuthorityApi
	SystemApiApi
	OperationRecordApi
	CasbinApi
	SystemApi
	SysFileUploadAndDownloadApi
}

var dataScope = utils.DataScope{}

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService
var authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
var menuService = service.ServiceGroupApp.SystemServiceGroup.MenuService
var apiService = service.ServiceGroupApp.SystemServiceGroup.ApiService
var operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
var casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService
var DeptService = service.ServiceGroupApp.SystemServiceGroup.DeptService
var systemConfigService = service.ServiceGroupApp.SystemServiceGroup.SystemConfigService
var fileUploadAndDownloadService = service.ServiceGroupApp.SystemServiceGroup.FileUploadAndDownloadService

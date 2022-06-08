package system

type SysGroup struct {
	UserService
	JwtService
	OperationRecordService
	CasbinService
	AuthorityService
	MenuService
	ApiService
	DeptService
	SystemConfigService
	FileUploadAndDownloadService
}

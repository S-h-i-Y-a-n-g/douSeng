package system

type RouterGroup struct {
	BaseRouter
	JwtRouter
	UserRouter
	MenuRouter
	DeptRouter
	AuthorityRouter
	ApiRouter
	OperationRecordRouter
	CasbinRouter
	SysRouter
	FileRouter
}

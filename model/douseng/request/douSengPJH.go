package request


//请求入参
type GetFeed struct {
	LatestTime string `json:"latest_time" form:"latest_time"` //可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token 	string `json:"token" form:"token"` // 用户登录状态下设置
}

//DouSeng用户登录入参
type DouSengLogin struct{
	Username          string `form:"username"` //用户名
	Password      string `form:"password"`	//密码
}

//得到用户信息入参 也是用户发布列表入参
type GetUserInfoBo struct {
	UserId int64 `json:"user_id" form:"user_id"`
	Token string `json:"token" form:"token"`
}

//用户注册入参
type UserRegister struct {
	Username      string `form:"username"` //用户名
	Password      string `form:"password"`	//密码
}

//用户上传视频入参
type UploadedFile struct {
	Title   string `form:"title" json:"title"`
	Token 	string `form:"token" json:"token"`
}
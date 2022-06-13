package douseng

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"go.uber.org/zap"
	"mime/multipart"
	"net/http"
	"project/global"
	"project/middleware"
	"project/model/common/response"
	ds "project/model/douseng"
	req "project/model/douseng/request"
	res "project/model/douseng/response"
	"project/model/system"
	systemReq "project/model/system/request"
	"project/utils"
	"time"
)

type DouSengPJHApi struct{}




// @Tags DouSeng
// @Summary 获取视频列表
// @Description Author：PangJiaHao 2022/06/09
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SetUserAuth true "latest_time, token"
// @Success 200 {string} string "{"StatusCode":0,"VideoList":{},"NextTime":"当前时间"}"UserFeedService
// @Router /douyin/feed [get]
func (d *DouSengPJHApi) Feed(c *gin.Context) {
	var GetInfo req.GetFeed
	var userID int //用户id
	//绑定参数
	err := c.ShouldBind(&GetInfo)
	if err != nil {
		global.GSD_LOG.Error("绑定参数失败!", zap.Any("err", err), utils.GetRequestID(c))
	}
	//解析token
	if GetInfo.Token != ""{
		j := &middleware.JWT{SigningKey: []byte(global.GSD_CONFIG.JWT.SigningKey)} // 唯一签名
		userinfo,err:=j.ParseTokenDouSeng(GetInfo.Token)
		if err != nil {
			global.GSD_LOG.Error("token 解析失败!", zap.Any("err", err), utils.GetRequestID(c))
			c.JSON(http.StatusOK,res.DouSengUser{
				DSResponse:res.DSResponse{
					StatusMsg: "token信息错误",
					StatusCode: 1,
				},
			},
			)
		}
		userID = int(userinfo.ID)
	}

	//进入service层处理
	ru:=douSengPJHService.FeedService(userID,GetInfo.LatestTime)
	c.JSON(http.StatusOK, ru)
}

// @Tags DouSeng
// @Summary 获取用户视频列表
// @Description Author：PangJiaHao 2022年6月9日21:40:31
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SetUserAuth true "latest_time, token"
// @Success 200 {string} string "{"StatusCode":0,"VideoList":{},"NextTime":"当前时间"}"UserFeedService
// @Router /douyin/feed [get]
func (d *DouSengPJHApi) GetUserFeed(c *gin.Context) {
	var GetInfo req.GetUserInfoBo
	//绑定参数
	err := c.ShouldBind(&GetInfo)
	if err != nil {
		global.GSD_LOG.Error("绑定参数失败!", zap.Any("err", err), utils.GetRequestID(c))
		c.JSON(http.StatusOK,res.DouSengUser{
			DSResponse:res.DSResponse{
				StatusMsg: "参数错误",
				StatusCode: 1,
			},
		},
		)
	}
	//token验证
	err=middleware.ValidateUser(GetInfo.Token,GetInfo.UserId,c)
	if err != nil {
		return
	}
	//信息无误，进入service层处理
	ru:=douSengPJHService.UserFeedService(int(GetInfo.UserId))
	c.JSON(http.StatusOK, ru)
}


// @Tags DouSeng
// @Summary 获取用户点赞视频列表
// @Description Author：PangJiaHao 2022年6月9日21:40:31
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SetUserAuth true "user_id, token"
// @Success 200 {string} string "{"StatusCode":0,"VideoList":{},"NextTime":"当前时间"}"UserFeedService
// @Router /douyin/publish/list/ [get]
func (d *DouSengPJHApi) GetUserFavoriteFeed(c *gin.Context) {
	var GetInfo req.GetUserInfoBo
	//绑定参数
	err := c.ShouldBind(&GetInfo)
	if err != nil {
		global.GSD_LOG.Error("绑定参数失败!", zap.Any("err", err), utils.GetRequestID(c))
		c.JSON(http.StatusOK,res.DouSengUser{
			DSResponse:res.DSResponse{
				StatusMsg: "参数错误",
				StatusCode: 1,
			},
		},
		)
	}
	//token验证
	err=middleware.ValidateUser(GetInfo.Token,GetInfo.UserId,c)
	if err != nil {
		return
	}
	//信息无误，进入service层处理
	ru:=douSengPJHService.UserFavoriteFeedService(int(GetInfo.UserId))
	c.JSON(http.StatusOK, ru)
}

// @Tags DouSeng
// @Summary DouSeng用户登录
// @Description Author：PangJiaHao 2022/06/09
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SetUserAuth true "user_id, token"
// @Success 200 {string} string "{"StatusCode":0,"user_id":,"token":}"
// @Router /douyin/user/login [post]
func (d *DouSengPJHApi) DouSengLogin(c *gin.Context) {
	var l req.DouSengLogin
	_ = c.ShouldBind(&l)
	if err := utils.Verify(l, utils.DouSengLoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, user := douSengPJHService.DouSengLoginService(l.Password, l.Username); err != nil {
		global.GSD_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err), utils.GetRequestID(c))
		c.JSON(http.StatusOK, res.DSResponse{
			StatusCode: 1,
			StatusMsg:  "用户名不存在或者密码错误",
		})
	} else { //签发token
		d.tokenNext(c, user)
	}
}

// 登录以后签发jwt
func (d *DouSengPJHApi) tokenNext(c *gin.Context, user *ds.UserInfo) {

	j := &middleware.JWT{SigningKey: []byte(global.GSD_CONFIG.JWT.SigningKey)} // 唯一签名

	claims := systemReq.CustomClaimsDouSeng{
		ID:         uint(user.Id),
		Username:   user.Name,
		PassWord:   user.Password,
		BufferTime: global.GSD_CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.GSD_CONFIG.JWT.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    "gsdPlus",                                             // 签名的发行者
		},
	}
	token, err := j.CreateTokenDouSeng(claims)
	if err != nil {
		global.GSD_LOG.Error("获取token失败!", zap.Any("err", err), utils.GetRequestID(c))
		response.FailWithMessage("获取token失败", c)
		return
	}

	userCache := systemReq.DouSengUserCache{
		ID:            uint(user.Id),
		UserName:      user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
	}

	_ = jwtService.SetRedisDouSengUserInfo(userCache)

	//非多点登录则直接返回响应，目前未配置全登录拦截
	if !global.GSD_CONFIG.System.UseMultipoint {
		c.JSON(http.StatusOK, res.DouSengUserLogin{
			DSResponse: res.DSResponse{
				StatusMsg:  "",
				StatusCode: 0,
			},
			Token:  token,
			UserID: user.Id,
		})
		return
	}

	if err, jwtStr := jwtService.GetRedisJWT(user.Name); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Name); err != nil {
			global.GSD_LOG.Error("设置登录状态失败!", zap.Any("err", err), utils.GetRequestID(c))
			c.JSON(http.StatusOK, res.DouSengUserLogin{DSResponse: res.DSResponse{StatusCode: 1, StatusMsg: "设置登录状态失败"}})
			return
		}

		c.JSON(http.StatusOK, res.DouSengUserLogin{
			DSResponse: res.DSResponse{
				StatusMsg:  "登陆成功",
				StatusCode: 0,
			},
			Token:  token,
			UserID: user.Id,
		})
	} else if err != nil {
		global.GSD_LOG.Error("设置登录状态失败!", zap.Any("err", err), utils.GetRequestID(c))
		c.JSON(http.StatusOK, res.DouSengUserLogin{DSResponse: res.DSResponse{StatusCode: 1, StatusMsg: "设置登录状态失败"}})
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JoinInBlacklist(blackJWT); err != nil {
			c.JSON(http.StatusOK, res.DouSengUserLogin{DSResponse: res.DSResponse{StatusCode: 1, StatusMsg: "jwt作废失败"}})
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Name); err != nil {
			c.JSON(http.StatusOK, res.DouSengUserLogin{DSResponse: res.DSResponse{StatusCode: 1, StatusMsg: "设置登录状态失败"}})
			return
		}
		//设置用户缓存
		c.JSON(http.StatusOK, res.DouSengUserLogin{
			DSResponse: res.DSResponse{
				StatusMsg:  "登录成功",
				StatusCode: 0,
			},
			Token:  token,
			UserID: user.Id,
		})
	}
}

// @Tags DouSeng
// @Summary DouSeng得到用户的所有信息
// @Description Author：PangJiaHao 2022/06/09
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SetUserAuth true "user_id, token"
// @Success 200 {string} string "{"StatusCode":0,"user_id":,"token":}"
// @Router /douyin/user/ [get]
func (d *DouSengPJHApi) GetUserInfo(c *gin.Context) {
	var l req.GetUserInfoBo
	_ = c.ShouldBind(&l)
	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		c.JSON(http.StatusOK, res.DouSengUserLogin{
			DSResponse: res.DSResponse{
				StatusMsg:  "参数错误",
				StatusCode: 1,
			},
		})
		return
	}
	//token目前设置为登录时间内均有效，但用户退出时会删除信息
	err:=middleware.ValidateUser(l.Token,l.UserId,c)
	if err != nil {
		return
	}

	//信息同步后查询缓存
	user, err := jwtService.GetRedisDouSengUserInfo(int(l.UserId))
	if err != nil {
		global.GSD_LOG.Error("缓存查询失败!", zap.Any("err", err), utils.GetRequestID(c))
		c.JSON(http.StatusOK, res.DouSengUserLogin{
			DSResponse: res.DSResponse{
				StatusMsg:  "登录失败，请重新登陆",
				StatusCode: 1,
			},
		})
	}

	c.JSON(http.StatusOK, res.DouSengUser{
		DSResponse: res.DSResponse{
			StatusMsg:  "登录成功",
			StatusCode: 0,
		},
		User: res.User{
			Id:            int64(user.ID),
			FollowerCount: user.FollowerCount,
			FollowCount:   user.FollowCount,
			IsFollow:      false,
			Name:          user.UserName,
		},
	})
}

// @Tags DouSeng
// @Summary DouSeng用户注册账号
// @Description Author：PangJiaHao 2022/06/09
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.Register true "用户名, 密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /douyin/user/register [post]
func (d *DouSengPJHApi) DouSengRegister(c *gin.Context) {
	var r req.UserRegister
	_ = c.ShouldBind(&r)
	//参数校验
	if err := utils.Verify(r, utils.DouSengRegisterVerify); err != nil {
		c.JSON(http.StatusOK, res.DSResponse{
			StatusCode: 1,
			StatusMsg:  "参数错误",
		})
		return
	}
	//进Service
	err := douSengPJHService.DouSengRegisterService(r.Username, r.Password)
	if err != nil {
		global.GSD_LOG.Error("注册失败,用户已存在", utils.GetRequestID(c), zap.Error(err))
		c.JSON(http.StatusOK, res.DSResponse{
			StatusCode: 1,
			StatusMsg:  "用户已存在",
		})
		return
	}

	if err, user := douSengPJHService.DouSengLoginService(r.Password,r.Username ); err != nil {
		global.GSD_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Any("err", err), utils.GetRequestID(c))
		c.JSON(http.StatusOK, res.DSResponse{
			StatusCode: 1,
			StatusMsg:  "用户名不存在或者密码错误",
		})
	} else { //签发token
		d.tokenNext(c, user)
	}
}

// @Tags DouSeng
// @Summary DouSeng用户上传视频
// @Description Author：PangJiaHao 2022/06/09
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.Register true "视频数据, token ,title"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /douyin/publish/action/ [post]
func (d *DouSengPJHApi) DouSengPublishVideo(c *gin.Context) {
	bT := time.Now().Second()            // 开始时间
	var successInfo int
	var f req.UploadedFile
	_ = c.ShouldBind(&f)
	_, file, err := c.Request.FormFile("data")
	if err != nil {
		global.GSD_LOG.Error("接收文件失败!", zap.Any("err", err))
		c.JSON(http.StatusOK, res.DSResponse{
			StatusCode: 1,
			StatusMsg:  "接收文件失败",
		})
	}
	//解析token
	j := &middleware.JWT{SigningKey: []byte(global.GSD_CONFIG.JWT.SigningKey)} // 唯一签名

	userinfo, err := j.ParseTokenDouSeng(f.Token)
	if err != nil {
		global.GSD_LOG.Error("token 解析失败!", zap.Any("err", err), utils.GetRequestID(c))
		c.JSON(http.StatusOK, res.DouSengUser{
			DSResponse: res.DSResponse{
				StatusMsg:  "token信息错误",
				StatusCode: 1,
			},
		},
		)
		return
	}

	go Test(file,f,userinfo.ID,&successInfo)


	for {
		if successInfo != 0 || time.Now().Second()-bT>5||bT-time.Now().Second()>5 {
			fmt.Printf("函数内  ： %v",successInfo)
			c.JSON(http.StatusOK,res.DouSengUser{
				DSResponse:res.DSResponse{
					StatusMsg: "上传成功",
					StatusCode: 0,
				},
			},
			)
			successInfo=0
			return
		}
	}

}



func Test(file *multipart.FileHeader,f req.UploadedFile,id uint,successInfo *int){

	//上传视频到七牛云在service，返回路径和名字
	filePath, _, _ :=PostToHealthCode2(file)
	filetp, _, _ :=PostToHealthCode(file)
	//上传到本地
	//n:=up.NewOss()
	//filePath, _, _ :=n.UploadFile(file)

	//将路径存入数据库
	_=douSengPJHService.DouSengUploadService(filePath,f.Title,filetp,int(id))
	*successInfo = 1
	//这里准备开一个协程去检查上传成功的是否是图片，如果是
}

//上传视频到七牛云
func PostToHealthCode(file *multipart.FileHeader) (string, string, error) {
	var accessKey,secretKey,bucket,Furl string
		//这里的配置单独做，目前先链接我的
		accessKey = "FhMyDZi245KmVAntt_MoIl2hqxBAHMik1JooI7Zz"
		secretKey = "SqQ202of1sL-g3iw0d-Jc4kpzxP8wkcBTNaKQ3i5"
		bucket = "bctbsdousheng"
		Furl ="http://rd9jtq8qx.hn-bkt.clouddn.com"

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}

	//通过 *multipart.FileHeader 打开获取
	files, openError := file.Open()
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	defer files.Close()                                              // 创建文件 defer 关闭
	if openError != nil {
		global.GSD_LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
	}
	//put上传到七牛云
	putErr := formUploader.Put(context.Background(), &ret, upToken, fileKey, files, file.Size, &putExtra)

	if putErr != nil {
		global.GSD_LOG.Error("function formUploader.Put() Filed", zap.Any("err", putErr.Error()))
		return "", "", errors.New("function formUploader.Put() Filed, err:" + putErr.Error())
	}
	//这里路径拼接先写死
	return Furl+ "/" + ret.Key, ret.Key, nil
}

//处理一下不晓得在干嘛的接口
func (d *DouSengPJHApi) BZD(c *gin.Context) {

	c.JSON(http.StatusOK,res.DouSengUser{
		DSResponse:res.DSResponse{
			StatusMsg: "成功",
			StatusCode: 0,
		},
	},
	)
}


//
func PostToHealthCode2(file *multipart.FileHeader) (string, string, error) {

	//这里的配置单独做，目前先链接我的
	accessKey := "udGS-HeZnr2aZQC0XJMprzWXnMy2D6AX44OVVklG"
	secretKey := "RvFxcW7T66MN4A5qCyPXl6zl_d8vv34eUALMb0lg"
	bucket := "dousheng1"

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)


	ret := storage.PutRet{}

	putExtra := storage.PutExtra{
	}

	//通过 *multipart.FileHeader 打开获取
	files, openError := file.Open()
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename) // 文件名格式 自己可以改 建议保证唯一性
	defer files.Close()                                                  // 创建文件 defer 关闭
	if openError != nil {
		global.GSD_LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
	}
	//put上传到七牛云
	putErr := formUploader.Put(context.Background(), &ret, upToken, fileKey, files, file.Size, &putExtra)

	if putErr != nil {
		global.GSD_LOG.Error("function formUploader.Put() Filed", zap.Any("err", putErr.Error()))
		return "", "", errors.New("function formUploader.Put() Filed, err:" + putErr.Error())
	}

	//这里路径拼接先写死
	return "http://rd6xoj6dg.hn-bkt.clouddn.com"+ "/" + ret.Key, ret.Key, nil
}

func test(){

	accessKey := "udGS-HeZnr2aZQC0XJMprzWXnMy2D6AX44OVVklG"
	secretKey := "RvFxcW7T66MN4A5qCyPXl6zl_d8vv34eUALMb0lg"
	bucket := "dousheng1"

	mac := qbox.NewMac(accessKey, secretKey)

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	s:=storage.NewBucketManager(mac,&cfg)

	f,err:=s.Stat(bucket,"1655015084luxu.mp4")

	if err != nil {
		fmt.Println("查询失败")
	}
	fmt.Println(f.MimeType)
}

//func (d *DouSengPJHApi)  Play(c *gin.Context)  {
//	var realPath ="D:\\0.ME\\项目\\抖音大作业\\frist\\marchsoft-golangkuangjia-GoSword-develop\\uploads\\file\\"
//	requestUrl :=c.Request.URL.String()
//
//	filePath := requestUrl[len("/uploads/file/"):]
//	fmt.Println("requestUrl =",realPath+filePath)
//	file,err :=os.Open(realPath + filePath)
//	defer file.Close()
//	if err != nil {
//
//	} else {
//		bs,_ := ioutil.ReadAll(file)
//		c.Data(200,"video/mp4",bs)
//	}
//
//}

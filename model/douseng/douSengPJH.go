package douseng

import (
	"errors"
	"go.uber.org/zap"
	"project/global"
	res "project/model/douseng/response"
	"project/utils"
)

//表结构
type Videos struct {
	Id  int64  `json:"id,omitempty"`
	UserId        int64 `json:"user_id"`		//用户id
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`	//视频播放地址
	CoverUrl      string `json:"cover_url,omitempty"`					//视频封面地址
	Title         string `json:"title"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`				//点赞数
	CommentCount  int64  `json:"comment_count,omitempty"`				//评论数
}

type UserFollower struct {
	UserId int64 `json:"user_id"`
	FollowerId int64 `json:"follower_id"`
}

type UserInfo struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
}

const VideosTableName = "ds_video"
const UserTableName = "ds_user"

func (v *Videos) VideosTableName() string {
	return VideosTableName
}

func (v *Videos) UserTableName() string {
	return UserTableName
}


var (
	ErrorUserExist = errors.New("用户已存在")
	ErrorUserIsNotExist = errors.New("用户不存在")
	ErrorUserLogin = errors.New("登陆失败，账号或密码错误")
)



func (v *Videos)GetFeedList() (videoList []res.Video,err error) {
	var Users res.User
	var videosList []Videos
	var video res.Video

	err = global.GSD_DB.Table(v.VideosTableName()).Where("deleted_at = ?",0).Limit(10).Order("id desc").Find(&videosList).Error
	if err != nil {
		return nil, err
	}
	//查询成功后去添加用户信息
	for _,value :=range videosList{
		video.Id=value.Id
		video.CommentCount=value.CommentCount
		video.FavoriteCount=value.FavoriteCount
		video.CoverUrl=value.CoverUrl
		video.PlayUrl=value.PlayUrl
		//查询视频发布者信息
		err = global.GSD_DB.Table(v.UserTableName()).Where("deleted_at = ? AND id = ?",0,value.UserId).Find(&Users).Error
		if err != nil {
			global.GSD_LOG.Error("绑定视频发布者信息失败!", zap.Any("err", err))
		}
		//TODO 查询是否已关注,先写未关注
		Users.IsFollow = false
		video.IsFavorite = false
		video.Author=Users
		//拼接完成一个信息放入列表
		videoList = append(videoList,video)
	}
	return videoList, err
}


//查询follow是否是user的粉丝
func (v *Videos)SelectIsFollow(userId int64,followId int64) (is bool , err error) {
	test:=new(UserFollower)
	err = global.GSD_DB.Table("ds_user_follower").Where("user_id = ? AND follower_id = ?",userId,followId).Find(&test).Error
	if err != nil {
		return false, err
	}else if test.UserId==0 && test.FollowerId ==0{//未关注
		return false, err
	}else {
		return true, err
	}
}

//登录验证
func (v *Videos) DouSengLogin(password,name string)(err error,info *UserInfo)  {
	user:=new(UserInfo)
	Password := utils.MD5V([]byte(password))
	err=global.GSD_DB.Table(v.UserTableName()).Where("name = ? and password = ?",name,Password).Find(&user).Error
	if user.Password != "" {
		user.Password=password
	}else {//为空则登录失败
	return ErrorUserLogin,nil
	}
	return err,user
}

//注册验证
func (v *Videos) DouSengRegister(password,name string)(err error)  {
	user:=new(UserInfo)
	Password := utils.MD5V([]byte(password))
	err=global.GSD_DB.Table(v.UserTableName()).Where("name = ? AND deleted_at = ?",name,0).Find(&user).Error
	if user.Password != "" {
		//已有账号
		return ErrorUserExist
	}
	//没有账号的话去注册
	user.Name=name
	user.Password=Password

	err = global.GSD_DB.Table(v.UserTableName()).Create(&user).Error
	return err
}


//上传视频
func (v *Videos) DouSengUploadVideo (PlayUrl,Title string , userId int)(err error)  {
	user:=new(UserInfo)
	video:=new(Videos)
	err=global.GSD_DB.Table(v.UserTableName()).Where("id = ? AND deleted_at = ?",userId,0).Find(&user).Error
	if user.Password == "" {
		//账号不存在
		return ErrorUserIsNotExist
	}
	//TODO 查询视频是否已经存在 不知道有没有必要，先放着
	err=global.GSD_DB.Table(v.VideosTableName()).Where("play_url = ? AND deleted_at = ?",PlayUrl,0).Find(&video).Error
	if err != nil {

	}

	v.UserId=int64(userId)
	v.PlayUrl = PlayUrl
	v.Title = Title
	v.CoverUrl = "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg" //封面先固定
	err = global.GSD_DB.Table("ds_video").Create(&v).Error
	return err
}

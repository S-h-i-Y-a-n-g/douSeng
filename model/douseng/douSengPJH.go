package douseng

import (
	"errors"
	"go.uber.org/zap"
	"project/global"
	res "project/model/douseng/response"
	"project/utils"
	"time"
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
	CreatedAt     time.Time `json:"created_at"`
}

type UserFollower struct {
	UserId int64 `json:"user_id"`
	FollowerId int64 `json:"follower_id"`
}

type UserVideo struct {
	UserId int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
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


//获取视频列表
func (v *Videos)GetFeedList(userID int) (videoList []res.Video,err error) {
	var Users res.User	//视频作者的用户数据
	var videosList []Videos //视频列表
	var video res.Video	//视频信息

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
		video.Title=value.Title
		//查询视频发布者信息
		Users,err=GetUserInfoById(int(value.UserId))
		if err != nil {
			global.GSD_LOG.Error("绑定视频发布者信息失败!", zap.Any("err", err))
		}
		if userID == 0{ //用户未登录
			Users.IsFollow = false
			video.IsFavorite = false
		}else {	//用户登录状态
			//根据user_id去查是否已经关注
			err, Users.IsFollow= QueryAttentionAuthor(userID,int(Users.Id))
			if err != nil {
				return nil, err
			}
			//再去找是否点赞
			err,video.IsFavorite = SelectFavoriteByUserId(userID,int(video.Id))
			if err != nil {
				return nil, err
			}
		}
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
func (v *Videos) DouSengUploadVideo (PlayUrl,Title,filetp string , userId int)(err error)  {
	user:=new(UserInfo)
	video:=new(Videos)
	videos:=new(Videos)
	err=global.GSD_DB.Table(v.UserTableName()).Where("id = ? AND deleted_at = ?",userId,0).Find(&user).Error
	if user.Password == "" {
		//账号不存在
		return ErrorUserIsNotExist
	}
	//TODO 查询视频是否已经存在 不知道有没有必要，先放着
	err=global.GSD_DB.Table(v.VideosTableName()).Where("play_url = ? AND deleted_at = ?",PlayUrl,0).Find(&video).Error
	if err != nil {

	}
	videos.UserId=int64(userId)
	videos.PlayUrl = PlayUrl
	videos.Title = Title
	videos.CoverUrl = filetp+"?vframe/jpg/offset/0" //封面取第一帧
	err = global.GSD_DB.Table("ds_video").Create(&videos).Error
	return err
}

//发布列表
func (v *Videos)GetUserFeedList(userId int) (videoList []res.Video,err error) {
	var Users res.User
	var videosList []Videos
	var video res.Video

	err = global.GSD_DB.Table(v.VideosTableName()).Where("deleted_at = ? AND user_id = ?",0,userId).Limit(10).Order("id desc").Find(&videosList).Error
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
		video.Title=value.Title

		//查询视频发布者信息
		Users,err=GetUserInfoById(int(value.UserId))
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

//点赞列表
func (v *Videos)GetUserFavoriteFeedList(userId int) (videoList []res.Video,err error) {
	var Users res.User
	var videosList []Videos
	var video res.Video
	var videoID []int

	//先去点赞关系表里查所有用户点赞的视频id
	err=global.GSD_DB.Table("ds_user_video_action").Select("video_id").Where("user_id = ? AND deleted_at = 0",userId).Find(&videoID).Error
	if err != nil {
		return nil, err
	}
	err = global.GSD_DB.Table(v.VideosTableName()).Where("deleted_at = ? AND id in ?",0,videoID).Limit(10).Order("id desc").Find(&videosList).Error
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
		video.Title=value.Title

		//查询视频发布者信息
		Users,err=GetUserInfoById(int(value.UserId))
		if err != nil {
			global.GSD_LOG.Error("绑定视频发布者信息失败!", zap.Any("err", err))
		}
		//TODO 查询是否已关注,先写未关注
		Users.IsFollow = false
		//点赞的视频列表都是true
		video.IsFavorite = true
		video.Author=Users
		//拼接完成一个信息放入列表
		videoList = append(videoList,video)
	}
	return videoList, err
}

//得到用户信息
func GetUserInfoById(userId int) (user res.User,err error) {
	err=global.GSD_DB.Table("ds_user").Where("id = ? AND deleted_at = ?",userId,0).Find(&user).Error
	if err != nil {
		return res.User{}, err
	}
	return user,err
}

//查询用户是否点赞了一个视频
func SelectFavoriteByUserId(userId int,videoId int) (err error, bo bool) {
	var num int64
	err=global.GSD_DB.Table("ds_user_video_action").Where("user_id = ? AND video_id = ? AND deleted_at = 0",userId,videoId).Count(&num).Error
	if err != nil {
		return err,false
	}
	if num ==0{
		return err,false
	}else {
		return err,true
	}
}

//查询用户是否关注了视频作者
func QueryAttentionAuthor(userId int,authorId int) (err error, bo bool)  {
	var num int64
	err=global.GSD_DB.Table("ds_user_follower").Where("user_id = ? AND follower_id = ?",authorId,userId).Count(&num).Error
	if err != nil {
		return err,false
	}
	if num ==0{
		return err,false
	}else {
		return err,true
	}
}
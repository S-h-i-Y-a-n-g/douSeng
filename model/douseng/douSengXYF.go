package douseng

import (
	"errors"
	"project/global"
)

const UserFavoriteTableName = "ds_user_video_action"


//表结构
type Favorite struct{
	UserID int
	VideoID int
	DeletedAt int
}


func (r *Favorite) DZAction(videoID int, userID int,Action int ) (err error) {
	f:=new(Favorite)
	var count int64
	err = global.GSD_DB.Table(UserFavoriteTableName).Where("user_id = ? AND video_id = ? ",userID,videoID).Find(&f).Count(&count).Error
	if err != nil {
		return err
	}

	if Action == 1 && count == 0 {	//点赞操作
		f.UserID=userID
		f.VideoID=videoID
		err = global.GSD_DB.Table(UserFavoriteTableName).Create(&f).Error
	}else if Action == 1 && count == 1 && f.DeletedAt == 1 {
		err = UserVideoAction(videoID, userID ,0)
	} else if Action == 2 || count == 1{ //取消点赞
		err = UserVideoAction(videoID, userID ,1)
	}else {
		return errors.New("参数错误")
	}
	//点赞成功后去维护一下video表，后续可以把数据放到redis里
	if err == nil{
		go WeiHuVideo(videoID,Action)
	}
	return err
}

//维护video表
func WeiHuVideo(videoID int,Action int)  {
	v:=new(Videos)
	err := global.GSD_DB.Table("ds_video").Where("id = ? ",videoID).Find(&v).Error
	if err != nil {
		return
	}
	if Action == 1 {	//点赞
		_ = global.GSD_DB.Table("ds_video").Where("id = ?",videoID).Updates(&Videos{
			FavoriteCount: v.FavoriteCount+1,
		}).Error
	}else {				//取消点赞
		_ = global.GSD_DB.Table("ds_video").Where("id = ?",videoID).Updates(&Videos{
			FavoriteCount: v.FavoriteCount-1,
		}).Error
	}

}

//点赞表？
func UserVideoAction(videoID int, userID int,Action int)(err error)  {
	err = global.GSD_DB.Table(UserFavoriteTableName).Where("user_id = ? AND video_id = ?",userID,videoID).Select("deleted_at").Updates(&Favorite{
		DeletedAt: Action,
	}).Error
	return err
}
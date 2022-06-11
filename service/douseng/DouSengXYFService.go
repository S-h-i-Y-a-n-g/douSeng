package douseng

import (
	"fmt"
	ds "project/model/douseng"
	res "project/model/douseng/response"
	"time"
)

type DouSengXYFService struct{}


var x ds.Favorite

//统计点赞/取消点赞数目
const delfavorite = 0
const addfavorite = 0

type favoriteDB struct {
	userId  int64 `json:"user_id"`
	videoId int64 `json:"video_id"`
}

/*
	点赞操作主要逻辑代码
*/

//点击点赞按钮
func (d *DouSengXYFService) FavoriteService(videoID int, userID int,Action int) (err error) {
	//直接往下扔了，后悔弄三层了
	err=x.DZAction(videoID,userID,Action)
	return err
}

func getDelFavorite() int64 {
	return delfavorite
}

func getAddFavorite() int64 {
	return addfavorite
}

func getAllFavorite() int64 {

	return addfavorite - delfavorite

}

//读取库存
func ReadFavoriteNum(video res.Video, user res.User) int64 {
	//select favorite_count from ds_vedio where id = ${video.id} and user_id = ${user.id}

	//return favorite_count
	return 0
}

//延时操作( 定时对数据库信息进行更新 )
//falg  true表示点赞 ，false表示取消点赞
func delayOnce(video res.Video, user res.User) {
	n := time.Now()
	// delay 1 second
	<-time.After(time.Second)
	//update ds_video set favorite_count = ReadFavoriteNum() - getAllFavorite() where id = ${video.id} and user_id = ${user.id}
	fmt.Println("Cost ", time.Since(n))
}

//取消点赞
func delAction(video res.Video, user res.User) *res.GetFavoriteResponse {
	resData := new(res.GetFavoriteResponse)
	//delete from ds_user_follower where useris = ${user.id} and video_id = ${video.id}
	//err := global.GSD_DB.Table(x.GetUserFavoriteTableName()).Where("user_id = ? AND video_id = ? ", user.Id, video.Id).Delete(&favoriteDB{})
	//if err != nil {
	//	resData.StatusCode = 505
	//	resData.StatusMsg = "取消点赞失败"
	//	return resData
	//}

	resData.StatusCode = 200
	resData.StatusMsg = "取消点赞成功"
	//令 delfavorite++
	return resData
}

//点赞
func clickAction(video res.Video, user res.User) *res.GetFavoriteResponse {
	resData := new(res.GetFavoriteResponse)
	//favorite := favoriteDB{userId: user.Id, videoId: video.Id}
	//insert into ds_user_follower valuse( ${user.id},${video.id})
	//err := global.GSD_DB.Table(x.GetUserFavoriteTableName()).Create(favorite)

	//if err != nil {
	//	resData.StatusCode = 505
	//	resData.StatusMsg = "点赞失败"
	//}

	//更新数据库

	resData.StatusCode = 200
	resData.StatusMsg = "点赞成功"
	//令 addfavorite++
	return resData
}

package douseng

import (
	"project/global"
	ds "project/model/douseng"
	res "project/model/douseng/response"
)

type DouSengXYFService struct{}

var v ds.Videos
var x ds.Favorite

const delfavorite = 0
const addfavorite = 0

type favoriteDB struct {
	userId  int64 `json:"user_id"`
	videoId int64 `json:"video_id"`
}

//点赞操作
func (d *DouSengXYFService) FavoriteService(video res.Video, user res.User) *res.GetFavoriteResponse {
	if user == (res.User{}) {
		return &res.GetFavoriteResponse{
			FavoriteResponse: res.FavoriteResponse{StatusCode: 200, StatusMsg: "请先登录"},
		}
	}
	if video.IsFavorite {
		video.IsFavorite = false
		resData := delAction(video, user)
		return resData
	}

	video.IsFavorite = true
	resData := clickAction(video, user)
	return resData
}

//取消点赞
func delAction(video res.Video, user res.User) *res.GetFavoriteResponse {
	resData := new(res.GetFavoriteResponse)
	err := global.GSD_DB.Table(x.GetUserFavoriteTableName()).Where("user_id = ? AND video_id = ? ", user.Id, video.Id).Delete(&favoriteDB{})
	if err != nil {
		resData.StatusCode = 505
		resData.StatusMsg = "取消点赞失败"
	}

	resData.StatusCode = 200
	resData.StatusMsg = "取消点赞成功"
	return resData
}

//点赞
func clickAction(video res.Video, user res.User) *res.GetFavoriteResponse {
	resData := new(res.GetFavoriteResponse)
	favorite := favoriteDB{userId: user.Id, videoId: video.Id}
	err := global.GSD_DB.Table(x.GetUserFavoriteTableName()).Create(favorite)
	if err != nil {
		resData.StatusCode = 505
		resData.StatusMsg = "点赞失败"
	}
	resData.StatusCode = 200
	resData.StatusMsg = "点赞成功"
	return resData
}

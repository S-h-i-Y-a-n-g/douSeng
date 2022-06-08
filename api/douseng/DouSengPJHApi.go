package douseng

import (
	"github.com/gin-gonic/gin"
	"net/http"
	res "project/model/dousheng/response"
	"time"
)

type DouSengPJHApi struct{}


var DemoVideos = []res.Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoUser = res.User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}

func (d *DouSengPJHApi) Feed(c *gin.Context) {
	c.JSON(http.StatusOK, res.GetFeedResponse{
		DSResponse:  res.DSResponse{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
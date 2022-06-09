package douseng

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"project/global"
	req "project/model/douseng/request"
	res "project/model/douseng/response"
	ser "project/service/douseng"
	"project/utils"
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


// @Tags DouSeng
// @Summary 获取视频列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body systemReq.SetUserAuth true "latest_time, token"
// @Success 200 {string} string "{"StatusCode":0,"VideoList":{},"NextTime":"当前时间"}"
// @Router /douyin/feed [get]
func (d *DouSengPJHApi) Feed(c *gin.Context) {
	var GetInfo req.GetFeed
	s := new(ser.DouSengPJHService)
	//绑定参数
	err := c.ShouldBind(&GetInfo)
	if err != nil {
		global.GSD_LOG.Error("绑定参数失败!", zap.Any("err", err), utils.GetRequestID(c))
	}
	//进入service层处理
	ru:=s.FeedService(GetInfo.Token,GetInfo.LatestTime)
	c.JSON(http.StatusOK, ru)
}
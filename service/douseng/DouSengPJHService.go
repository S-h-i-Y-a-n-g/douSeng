package douseng

import (
	"go.uber.org/zap"
	"project/global"
	res "project/model/douseng/response"
	ds "project/model/douseng"
	"time"
)

type DouSengPJHService struct{}



//返回视频列表
func (d *DouSengPJHService) FeedService (token string,LatestTime string) *res.GetFeedResponse {
	resData := new(res.GetFeedResponse)
	var vi ds.Videos

	//获取视频列表
	videoList,err:=vi.GetFeedList()
	if err != nil {
		global.GSD_LOG.Error("获取视频列表失败!", zap.Any("err", err))
		return &res.GetFeedResponse{
			DSResponse:  res.DSResponse{StatusCode: 200,StatusMsg: "未知错误"},
			NextTime:  time.Now().Unix(),
		}
	}
	resData.VideoList=videoList
	resData.StatusCode=0
	resData.StatusMsg="success"
	if LatestTime != "" {
		//这里应该是赋值LatesTime 测试先不搞
		resData.NextTime=time.Now().Unix()
	}else {
		resData.NextTime=time.Now().Unix()
	}
	return resData
}

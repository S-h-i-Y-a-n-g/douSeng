package douseng

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"project/global"
	"project/middleware"
	request "project/model/douseng/response"
	v1 "project/service/douseng"
	"project/utils"
)

type DouSengXYFApi struct{}

type ActionInfo struct {
	Token string `json:"token" form:"token"`
	VideoId int `json:"video_id" form:"video_id"`
	ActionType int `json:"action_type" form:"action_type"`
}

var v v1.DouSengXYFService


//点赞操作接口，by xyf，pjh
func (d *DouSengXYFApi) Action(c *gin.Context) {
	var info ActionInfo
	_ = c.ShouldBind(&info)

	//解析token
	j := &middleware.JWT{SigningKey: []byte(global.GSD_CONFIG.JWT.SigningKey)} // 唯一签名
	userinfo, err := j.ParseTokenDouSeng(info.Token)
	if err != nil {
		global.GSD_LOG.Error("token 解析失败!", zap.Any("err", err), utils.GetRequestID(c))
		c.JSON(http.StatusOK, request.DouSengUser{
			DSResponse: request.DSResponse{
				StatusMsg:  "token信息错误",
				StatusCode: 1,
			},
		},
		)
	}

	err=v.FavoriteService( info.VideoId,int(userinfo.ID),info.ActionType)

	if err != nil {
		c.JSON(http.StatusOK, request.DouSengUser{
			DSResponse: request.DSResponse{
				StatusMsg:  "妹成功啊",
				StatusCode: 1,
			},
		},
		)
	}
	if info.ActionType == 1 {
		c.JSON(http.StatusOK, request.DouSengUser{
			DSResponse: request.DSResponse{
				StatusMsg:  "给你比心",
				StatusCode: 0,
			},
		},
		)
	}else {
		c.JSON(http.StatusOK, request.DouSengUser{
			DSResponse: request.DSResponse{
				StatusMsg:  "给你鬼脸",
				StatusCode: 0,
			},
		},
		)
	}


}

package douseng

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"project/global"
	"project/middleware"
	"project/model/douseng"
	"project/utils"
)

type RelationActionRequest struct {
	Token      string `json:"token,omitempty" form:"token"`
	ToUserId   int64  `json:"to_user_id,omitempty" form:"to_user_id"`
	ActionType int32  `json:"action_type,omitempty" form:"action_type"`
}
type RelationFollowListRequest struct {
	UserId int64  `json:"user_id,omitempty" form:"user_id"`
	Token  string `json:"token,omitempty" form:"token"`
}
type RelationFollowerListRequest struct {
	UserId int64  `json:"user_id,omitempty" form:"user_id"`
	Token  string `json:"token,omitempty" form:"token"`
}
type RelationBaseResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
type RelationActionResponse struct {
	RelationBaseResponse
}
type RelationFollowListResponse struct {
	RelationBaseResponse
	UserList []*douseng.User `json:"user_list,omitempty"`
}
type RelationFollowerListResponse struct {
	RelationBaseResponse
	UserList []*douseng.User `json:"user_list,omitempty"`
}

type DouSengLYFApi struct{}

func (d *DouSengLYFApi) Action(c *gin.Context) {
	var parameter RelationActionRequest
	//绑定参数
	err := c.ShouldBind(&parameter)
	if err != nil {
		global.GSD_LOG.Error("绑定参数失败!", zap.Any("err", err), utils.GetRequestID(c))
	}
	id, err := tokenValidate(parameter.Token)
	if err != nil || id == parameter.ToUserId {
		var msg string
		if err != nil {
			msg = err.Error()
		} else {
			msg = "不能关注自己捏~"
		}
		c.JSON(http.StatusOK, RelationActionResponse{RelationBaseResponse{
			StatusCode: 1,
			StatusMsg:  msg,
		}})
		return
	}
	//进入service层处理
	res := douSengLYFService.RelationAction(id, parameter.ToUserId, parameter.ActionType, parameter.Token)
	c.JSON(http.StatusOK, RelationActionResponse{RelationBaseResponse: RelationBaseResponse{
		StatusCode: func() int32 {
			if res == nil {
				return 0
			}
			return 1
		}(),
		StatusMsg: func() string {
			if res == nil {
				return ""
			}
			return res.Error()
		}(),
	}})
}
func (d *DouSengLYFApi) FollowList(c *gin.Context) {
	var parameter RelationFollowListRequest
	//绑定参数
	err := c.ShouldBind(&parameter)
	if err != nil {
		global.GSD_LOG.Error("绑定参数失败!", zap.Any("err", err), utils.GetRequestID(c))
	}
	if id, err := tokenValidate(parameter.Token); err != nil || id != parameter.UserId {
		var msg string
		if err != nil {
			msg = err.Error()
		} else {
			msg = "token错误"
		}
		c.JSON(http.StatusOK, RelationActionResponse{RelationBaseResponse{
			StatusCode: 1,
			StatusMsg:  msg,
		}})
		return
	}
	//进入service层处理
	res, err := douSengLYFService.RelationFollowList(parameter.UserId, parameter.Token)
	c.JSON(http.StatusOK, RelationFollowListResponse{RelationBaseResponse: RelationBaseResponse{
		StatusCode: func() int32 {
			if err == nil {
				return 0
			}
			return 1
		}(),
		StatusMsg: func() string {
			if err == nil {
				return ""
			}
			return err.Error()
		}(),
	}, UserList: res})
}
func (d *DouSengLYFApi) FollowerList(c *gin.Context) {
	var parameter RelationFollowerListRequest
	//绑定参数
	err := c.ShouldBind(&parameter)
	if err != nil {
		global.GSD_LOG.Error("绑定参数失败!", zap.Any("err", err), utils.GetRequestID(c))
	}
	if id, err := tokenValidate(parameter.Token); err != nil || id != parameter.UserId {
		var msg string
		if err != nil {
			msg = err.Error()
		} else {
			msg = "token错误"
		}
		c.JSON(http.StatusOK, RelationActionResponse{RelationBaseResponse{
			StatusCode: 1,
			StatusMsg:  msg,
		}})
		return
	}
	//进入service层处理
	res, err := douSengLYFService.RelationFollowerList(parameter.UserId, parameter.Token)
	c.JSON(http.StatusOK, RelationFollowerListResponse{RelationBaseResponse: RelationBaseResponse{
		StatusCode: func() int32 {
			if err == nil {
				return 0
			}
			return 1
		}(),
		StatusMsg: func() string {
			if err == nil {
				return ""
			}
			return err.Error()
		}(),
	}, UserList: res})
}
func tokenValidate(token string) (int64, error) {
	j := &middleware.JWT{SigningKey: []byte(global.GSD_CONFIG.JWT.SigningKey)} // 唯一签名
	res, err := j.ParseTokenDouSeng(token)
	if err != nil {
		return 0, err
	}
	return int64(res.ID), nil
}

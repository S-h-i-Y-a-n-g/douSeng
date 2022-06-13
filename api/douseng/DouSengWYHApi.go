package douseng

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/global"
	"project/middleware"
	"project/model/douseng"
)

type DouShengWYHApi struct{}

func (d *DouShengWYHApi) CommentList(c *gin.Context) {
	videoId := c.Query("video_id")
	if CommentList, err := douShengWYHService.CommentList(videoId); err != nil {
		c.JSON(http.StatusOK, douseng.Response{
			StatusCode: 1,
			StatusMsg:  "Error",
		})
	} else {
		c.JSON(http.StatusOK, douseng.CommentListResponse{
			Response: douseng.Response{
				StatusCode: 0,
				StatusMsg:  "Success"},
			CommentList: CommentList,
		})
	}
}


//评论操作
func (d *DouShengWYHApi) CommentAction(c *gin.Context) {
	var action douseng.CommentAction
	err := c.ShouldBind(&action)

	if err != nil {
		c.JSON(http.StatusOK, douseng.Response{
			StatusCode: 1,
			StatusMsg:  "Data in wrong formatr",
		})
		return
	}

	j := &middleware.JWT{SigningKey: []byte(global.GSD_CONFIG.JWT.SigningKey)}
	if userinfo, err := j.ParseTokenDouSeng(action.Token); err != nil {
		c.JSON(http.StatusOK, douseng.Response{
			StatusMsg:  "Token error",
			StatusCode: 1,
		})
		return
	} else {
		action.UserId = uint64(userinfo.ID)
	}

	if action.ActionType != "1" && action.ActionType != "2" {
		c.JSON(http.StatusOK, douseng.Response{
			StatusCode: 1,
			StatusMsg:  "Action_type error",
		})
		return
	}

	if comment, err := douShengWYHService.CommentAction(action); err != nil {
		c.JSON(http.StatusOK, douseng.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, douseng.CommentResponse{
			Response: douseng.Response{
				StatusCode: 0,
				StatusMsg:  "Success",
			},
			Comment: comment,
		})
	}
}

package douseng

import (
	"errors"
	"gorm.io/gorm"
	"project/global"
	"project/model/douseng"
)

type DouShengWYHService struct{}

func (d DouShengWYHService) CommentList(videoId string) (commentList []douseng.Comment, err error) {
	err = global.GSD_DB.Table("ds_comment").Order("created_at desc").Find(&commentList, "video_id = ?", videoId).Error

	for i, _ := range commentList {
		var user douseng.CommentUser
		global.GSD_DB.Table("ds_user").Take(&user, commentList[i].UserId)
		commentList[i].User = user
		commentList[i].CreateDate = commentList[i].CreatedAt.Format("01-02")
	}

	return commentList, err
}

func (d DouShengWYHService) CommentAction(commentRequest douseng.CommentAction) (comment douseng.Comment, err error) {
	if commentRequest.ActionType == "1" {
		comment.UserId = commentRequest.UserId
		comment.VideoId = commentRequest.VideoId
		comment.Content = commentRequest.CommentText
		if err := global.GSD_DB.Table("ds_comment").Create(&comment).Error; err != nil {
			return comment, errors.New("Comment failed")
		} else {
			var user douseng.CommentUser
			global.GSD_DB.Table("ds_user").Take(&user, commentRequest.UserId)
			comment.User = user
			comment.CreateDate = comment.CreatedAt.Format("01-02")
			var video douseng.CommentVideo
			global.GSD_DB.Table("ds_video").Take(&video, commentRequest.VideoId).Update("comment_count", video.CommentCount+1)
		}

	}

	if commentRequest.ActionType == "2" {
		err := global.GSD_DB.Debug().Table("ds_comment").Take(&comment, "id = ? AND user_id = ?", commentRequest.CommentId, commentRequest.UserId).Delete(&comment).Error
		if err == gorm.ErrRecordNotFound {
			return comment, errors.New("评论不存在")
		}
		var user douseng.CommentUser
		global.GSD_DB.Table("ds_user").Take(&user, commentRequest.UserId)
		comment.User = user
		comment.CreateDate = comment.CreatedAt.Format("01-02")
		var video douseng.CommentVideo
		global.GSD_DB.Table("ds_video").Take(&video, commentRequest.VideoId).Update("comment_count", video.CommentCount-1)
	}

	return comment, nil
}

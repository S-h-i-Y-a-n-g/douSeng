package douseng

import (
	"gorm.io/gorm"
	"time"
)

type CommentAction struct {
	UserId      uint64
	Token       string `form:"token"`
	VideoId     string `form:"video_id"`
	ActionType  string `form:"action_type"`
	CommentText string `form:"comment_text"`
	CommentId   string `form:"comment_id"`
}

type CommentUser struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type CommentVideo struct {
	Id           string
	CommentCount uint64
}

type Comment struct {
	Id         uint64         `json:"id"`
	UserId     uint64         `json:"-"`
	User       CommentUser    `json:"user"`
	VideoId    string         `json:"-"`
	Content    string         `json:"content"`
	CreateDate string         `json:"create_date" gorm:"-"`
	CreatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type CommentResponse struct {
	Response
	Comment Comment `json:"comment"`
}

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list"`
}

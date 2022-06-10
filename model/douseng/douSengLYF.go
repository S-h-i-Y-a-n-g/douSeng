package douseng

import (
	"gorm.io/gorm"
	"project/global"
	"sync"
)

// User 用户信息
type User struct {
	Id            int64  `json:"id" gorm:"id"`
	Name          string `json:"name" gorm:"name"`
	FollowCount   int64  `json:"follow_count,omitempty" gorm:"follow_count"`
	FollowerCount int64  `json:"follower_count,omitempty" gorm:"follower_count"`
	IsFollow      bool   `json:"is_follow" gorm:"-"`
}

func (u User) TableName() string {
	return "ds_user"
}

type UserDao struct {
}

var userDao *UserDao
var userOnce sync.Once

func GetUserDaoInstance() *UserDao {
	userOnce.Do(func() {
		userDao = new(UserDao)
	})
	return userDao
}

// MQueryUserById 根据uid批量查询用户
func (u *UserDao) MQueryUserById(uids []int64) ([]*User, error) {
	if len(uids) == 0 {
		return []*User{}, nil
	}
	var users []*User
	if err := global.GSD_DB.Find(&users, uids).Error; err != nil {
		return nil, err
	}
	return users, nil
}
func (u *UserDao) UpdateFollowCount(uid int64, cnt int) error {
	if err := global.GSD_DB.Model(&User{Id: uid}).UpdateColumn("follow_count", gorm.Expr("follow_count + (?)", cnt)).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserDao) UpdateFollowerCount(uid int64, cnt int) error {
	if err := global.GSD_DB.Model(&User{Id: uid}).UpdateColumn("follower_count", gorm.Expr("follower_count + (?)", cnt)).Error; err != nil {
		return err
	}
	return nil
}

// FollowRelation 关注关系
type FollowRelation struct {
	UserId     int64 `json:"user_id,omitempty" gorm:"user_id"`
	FollowerId int64 `json:"follower_id,omitempty" gorm:"follower_id"`
}

func (f FollowRelation) TableName() string {
	return "ds_user_follower"
}

type FollowDao struct {
}

var followDao *FollowDao
var followOnce sync.Once

func GetFollowDaoInstance() *FollowDao {
	followOnce.Do(func() {
		followDao = new(FollowDao)

	})
	return followDao
}
func (fd *FollowDao) QueryFollow(fid int64, uids []int64) ([]*FollowRelation, error) {
	var followRelations []*FollowRelation
	if err := global.GSD_DB.Where(&FollowRelation{FollowerId: fid}).Where("user_id in (?)", uids).Find(&followRelations).Error; err != nil {
		return nil, err
	}
	return followRelations, nil
}

// QueryFollowById 根据uid查询关注者关系列表
func (fd *FollowDao) QueryFollowById(uid int64) ([]*FollowRelation, error) {
	var followRelations []*FollowRelation
	if err := global.GSD_DB.Where(&FollowRelation{FollowerId: uid}).Find(&followRelations).Error; err != nil {
		return nil, err
	}
	return followRelations, nil
}

// QueryFollowerById 根据uid查询被关注者关系列表
func (fd *FollowDao) QueryFollowerById(uid int64) ([]*FollowRelation, error) {
	var followRelations []*FollowRelation
	if err := global.GSD_DB.Where(&FollowRelation{UserId: uid}).Find(&followRelations).Error; err != nil {
		return nil, err
	}
	return followRelations, nil
}

// InsertFollowRelation  uid对应的用户关注fid对应的用户
func (fd *FollowDao) InsertFollowRelation(uid, fid int64) error {
	if err := global.GSD_DB.Create(&FollowRelation{
		UserId:     uid,
		FollowerId: fid,
	}).Error; err != nil {
		return err
	}
	return nil
}

// DeleteFollowRelation  uid对应的用户取消关注fid对应的用户
func (fd *FollowDao) DeleteFollowRelation(uid, fid int64) error {
	if err := global.GSD_DB.Where(&FollowRelation{
		UserId:     uid,
		FollowerId: fid,
	}).Delete(&FollowRelation{}).Error; err != nil {
		return err
	}
	return nil
}

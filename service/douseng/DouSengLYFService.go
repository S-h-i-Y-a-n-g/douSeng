package douseng

import (
	"errors"
	"project/model/douseng"
	"sync"
)

type DouSengLYFService struct {
}

var userDao = douseng.GetUserDaoInstance()
var followDao = douseng.GetFollowDaoInstance()

// RelationAction 关注操作
func (s *DouSengLYFService) RelationAction(userId, toUserId int64, actionType int32, token string) error {
	userPair, err := userDao.MQueryUserById([]int64{userId, toUserId})
	if err != nil {
		return err
	} else if len(userPair) != 2 {
		return errors.New("用户不存在")
	}
	if actionType == 1 { //1-关注
		if l, err := followDao.QueryFollow(userId, []int64{toUserId}); err != nil {
			return err
		} else if len(l) > 0 {
			return errors.New("用户已关注")
		}
		if err = followDao.InsertFollowRelation(toUserId, userId); err != nil { //插入关注关系
			return err
		}
		if err = userDao.UpdateFollowCount(userId, 1); err != nil { //当前用户关注数加一
			return err
		}
		if err = userDao.UpdateFollowerCount(toUserId, 1); err != nil { //对应用户粉丝数加一
			return err
		}
		return nil
	} else if actionType == 2 { //2-取消关注
		if l, err := followDao.QueryFollow(userId, []int64{toUserId}); err != nil {
			return err
		} else if len(l) <= 0 {
			return errors.New("用户未关注")
		}
		if err = followDao.DeleteFollowRelation(toUserId, userId); err != nil { //移除关注关系
			return err
		}
		if err = userDao.UpdateFollowCount(userId, -1); err != nil { //当前用户关注数减一
			return err
		}
		if err = userDao.UpdateFollowerCount(toUserId, -1); err != nil { //对应用户粉丝数减一
			return err
		}
		return err
	} else {
		return errors.New("操作异常")
	}
}

// RelationFollowList 用户关注列表
func (s *DouSengLYFService) RelationFollowList(userId int64, token string) ([]*douseng.User, error) {
	u, err := userDao.MQueryUserById([]int64{userId}) //查询当前用户是否存在
	if err != nil {
		return nil, err
	} else if len(u) <= 0 {
		return nil, errors.New("用户不存在")
	}
	rs, err := followDao.QueryFollowById(userId)
	if err != nil {
		return nil, err
	}
	var uids = make([]int64, len(rs))
	for i, r := range rs { //提取关注者的uid
		uids[i] = r.UserId
	}
	ul, err := userDao.MQueryUserById(uids) //根据关注者的uid获取关注者用户信息
	for _, u := range ul {
		u.IsFollow = true
	}
	return ul, err
}

// RelationFollowerList 用户粉丝列表
func (s *DouSengLYFService) RelationFollowerList(userId int64, token string) ([]*douseng.User, error) {
	u, err := userDao.MQueryUserById([]int64{userId}) //查询当前用户是否存在
	if err != nil {
		return nil, err
	} else if len(u) <= 0 {
		return nil, errors.New("用户不存在")
	}
	rs, err := followDao.QueryFollowerById(userId) //查询粉丝关系
	if err != nil {
		return nil, err
	}
	if len(rs) == 0 {
		return []*douseng.User{}, nil
	}
	var uids = make([]int64, len(rs))
	for i, r := range rs { //提取粉丝的uid
		uids[i] = r.FollowerId
	}
	var ul []*douseng.User
	var follows = make(map[int64]struct{})

	var wg sync.WaitGroup
	wg.Add(2)
	go func() { //并行根据粉丝的uid批量获取粉丝用户信息
		ul, err = userDao.MQueryUserById(uids)
		wg.Done()
	}()
	go func() { //并行根据粉丝uid查询关注关系
		rl, er := followDao.QueryFollow(userId, uids)
		if er != nil {
			err = er
			return
		}
		for _, r := range rl {
			follows[r.UserId] = struct{}{}
		}
		wg.Done()
	}()
	wg.Wait()
	if err != nil {
		return nil, err
	}
	for _, u := range ul { //对关注的粉丝设置关注关系
		if _, ok := follows[u.Id]; ok {
			u.IsFollow = true
		}
	}
	return ul, err
}

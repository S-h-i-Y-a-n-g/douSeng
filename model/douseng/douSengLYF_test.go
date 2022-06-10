package douseng

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"project/global"
	"reflect"
	"testing"
)

func init() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/dousheng?charset=utf8mb4&parseTime=True&loc=Local"
	global.GSD_DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	global.GSD_DB.Debug()
}
func TestFollowDao_InsertFollowRelation(t *testing.T) {
	type args struct {
		uid int64
		fid int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"测试用户关注", args{
			uid: 1,
			fid: 2,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fd := GetFollowDaoInstance()
			if err := fd.InsertFollowRelation(tt.args.uid, tt.args.fid); (err != nil) != tt.wantErr {
				t.Errorf("Follow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFollowDao_QueryFollowById(t *testing.T) {
	type args struct {
		uid int64
	}
	tests := []struct {
		name    string
		args    args
		want    []*FollowRelation
		wantErr bool
	}{
		// TODO: Add test cases.
		{"根据uid查询关注关系", args{2}, []*FollowRelation{{
			UserId:     1,
			FollowerId: 2,
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fd := &FollowDao{}
			got, err := fd.QueryFollowById(tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryFollowById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryFollowById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFollowDao_QueryFollowerById(t *testing.T) {
	type args struct {
		uid int64
	}
	tests := []struct {
		name    string
		args    args
		want    []*FollowRelation
		wantErr bool
	}{
		// TODO: Add test cases.
		{"根据uid查看粉丝关系", args{1}, []*FollowRelation{{
			UserId:     1,
			FollowerId: 2,
		}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fd := &FollowDao{}
			got, err := fd.QueryFollowerById(tt.args.uid)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryFollowerById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryFollowerById() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDao_MQueryUserById(t *testing.T) {
	type args struct {
		uids []int64
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"正常根据uid批量查询用户", args{uids: []int64{1, 2, 3, 4}}, 3, false},
		{"根据uid批量查询0个用户", args{uids: []int64{}}, 0, false},
		{"根据uid批量查询不存在用户", args{uids: []int64{-1}}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDao{}
			got, err := u.MQueryUserById(tt.args.uids)
			if (err != nil) != tt.wantErr {
				t.Errorf("MQueryUserById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("MQueryUserById() got = %v, want %v", got, tt.want)
			//}
			if len(got) != tt.want {
				t.Errorf("got=%v,want %v", len(got), tt.want)
			}
		})
	}
}
func TestFollowDao_Unfollow(t *testing.T) {
	type args struct {
		uid int64
		fid int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"根据uid以及fid删除关注关系", args{
			uid: 1,
			fid: 2,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fd := &FollowDao{}
			if err := fd.DeleteFollowRelation(tt.args.uid, tt.args.fid); (err != nil) != tt.wantErr {
				t.Errorf("Unfollow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFollowDao_QueryFollow(t *testing.T) {
	type args struct {
		uid  int64
		fids []int64
	}
	tests := []struct {
		name    string
		args    args
		want    []*FollowRelation
		wantErr bool
	}{
		// TODO: Add test cases.
		{"正常测试根据uid和followerId获取关注关系", args{
			uid:  5,
			fids: []int64{6, 7, 8},
		}, []*FollowRelation{{6, 5}, {7, 5}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fd := &FollowDao{}
			got, err := fd.QueryFollow(tt.args.uid, tt.args.fids)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryFollow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryFollow() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserDao_UpdateFollowCount(t *testing.T) {
	type args struct {
		uid int64
		cnt int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"正常测试增加关注数", args{
			uid: 1,
			cnt: 1,
		}, false},
		{"正常测试减少关注数", args{
			uid: 1,
			cnt: -1,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDao{}
			if err := u.UpdateFollowCount(tt.args.uid, tt.args.cnt); (err != nil) != tt.wantErr {
				t.Errorf("UpdateFollowCount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserDao_UpdateFollowerCount(t *testing.T) {
	type args struct {
		uid int64
		cnt int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"正常测试增加粉丝数", args{
			uid: 1,
			cnt: 1,
		}, false},
		{"正常测试减少粉丝数", args{
			uid: 1,
			cnt: -1,
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserDao{}
			if err := u.UpdateFollowerCount(tt.args.uid, tt.args.cnt); (err != nil) != tt.wantErr {
				t.Errorf("UpdateFollowerCount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

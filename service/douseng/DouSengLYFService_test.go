package douseng

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"project/global"
	"testing"
)

func init() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/dousheng?charset=utf8mb4&parseTime=True&loc=Local"
	global.GSD_DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	global.GSD_DB = global.GSD_DB.Debug()
}
func TestDouSengLYFService_relationAction(t *testing.T) {
	type args struct {
		userId     int64
		toUserId   int64
		actionType int32
		token      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"测试关注对象不存在操作", args{
			userId:     1,
			toUserId:   -1,
			actionType: 1,
			token:      "",
		}, true},
		{"测试关注自身不存在操作", args{
			userId:     -1,
			toUserId:   1,
			actionType: 1,
			token:      "",
		}, true},
		{"测试正常关注操作", args{
			userId:     1,
			toUserId:   2,
			actionType: 1,
			token:      "",
		}, false},
		{"测试取消关注对象不存在操作", args{
			userId:     1,
			toUserId:   -1,
			actionType: 2,
			token:      "",
		}, true},
		{"测试取消关注自身不存在操作", args{
			userId:     -1,
			toUserId:   1,
			actionType: 2,
			token:      "",
		}, true},
		{"测试正常取消关注操作", args{
			userId:     1,
			toUserId:   2,
			actionType: 2,
			token:      "",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DouSengLYFService{}
			err := s.RelationAction(tt.args.userId, tt.args.toUserId, tt.args.actionType, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("relationAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestDouSengLYFService_relationFollowList(t *testing.T) {
	type args struct {
		userId int64
		token  string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"根据uid正常查询关注者", args{
			userId: 4,
			token:  "",
		}, 1, false},
		{"根据不存在uid查询关注者", args{
			userId: -1,
			token:  "",
		}, 0, true},
		{"根据uid查询空关注者", args{
			userId: 1,
			token:  "",
		}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DouSengLYFService{}
			got, err := s.RelationFollowList(tt.args.userId, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("relationFollowList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("relationFollowList() got = %v, want %v", got, tt.want)
			//}
			if len(got) != tt.want {
				t.Errorf("relationFollowList() got = %v, want %v", len(got), tt.want)
			}
		})
	}
}

func TestDouSengLYFService_relationFollowerList(t *testing.T) {
	type args struct {
		userId int64
		token  string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"根据uid正常查询粉丝", args{
			userId: 2,
			token:  "",
		}, 1, false},
		{"根据不存在uid查询粉丝", args{
			userId: -1,
			token:  "",
		}, 0, true},
		{"根据uid查询空粉丝", args{
			userId: 5,
			token:  "",
		}, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &DouSengLYFService{}
			got, err := s.RelationFollowerList(tt.args.userId, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("relationFollowerList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("relationFollowerList() got = %v, want %v", got, tt.want)
			//}
			for _, u := range got {
				t.Log(u)
			}
			if len(got) != tt.want {
				t.Errorf("relationFollowList() got = %v, want %v", len(got), tt.want)
			}
		})
	}
}

package request

import (
	"project/model/system"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// Custom claims structure
type CustomClaims struct {
	UUID        uuid.UUID
	ID          uint
	Username    string
	BufferTime  int64
	AuthorityId uint
	jwt.StandardClaims
}

// User cache structure
type UserCache struct {
	UUID        string                `redis:"uuid"`
	ID          uint                  `redis:"id"`
	DeptId      uint                  `redis:"deptId"`
	AuthorityId []uint                `redis:"authorityId"`
	Authority   []system.SysAuthority `redis:"-"`
}

// User cache structure
type UserCacheRedis struct {
	ID          uint   `redis:"id"`
	DeptId      uint   `redis:"deptId"`
	AuthorityId []byte `redis:"authorityId"`
}

/***********************DouSeng***************************************/
// DouSengJWT
type CustomClaimsDouSeng struct {
	ID          uint
	Username    string
	PassWord    string
	BufferTime  int64 //缓存时间
	jwt.StandardClaims
}

// DouSengUser cache structure
type DouSengUserCache struct {
	ID          uint                `redis:"id"`
	UserName    string              `redis:"name"`
	FollowCount   int64  			`redis:"follow_count"`
	FollowerCount int64  			`redis:"follower_count"`
}
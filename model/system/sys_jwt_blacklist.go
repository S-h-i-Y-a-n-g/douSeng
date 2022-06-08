package system

import (
	"project/global"
)

type JwtBlacklist struct {
	global.GSD_MODEL
	Jwt string `gorm:"not null;type:text;comment:jwt"`
}

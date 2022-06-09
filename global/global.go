package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"

	"project/config"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GSD_DB                  *gorm.DB
	GSD_REDIS               *redis.Client
	GSD_CONFIG              config.Server
	GSD_VP                  *viper.Viper
	GSD_LOG                 *zap.Logger
	GSD_Casbin              *casbin.SyncedEnforcer
	GSD_Concurrency_Control = &singleflight.Group{}
	GSD_UP_Token              string
)

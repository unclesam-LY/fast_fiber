package global

import (
	"FastFiber/config"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Version = "1.0.1"

var (
	Config   *config.Config
	DB       *gorm.DB
	Redis    *redis.Client
	Log      *zap.Logger
	MysqlLog logger.Interface
)

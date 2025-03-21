package core

import (
	"FastFiber/global"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func InitRedis() (client *redis.Client) {
	cfg := global.Config.Redis
	if cfg.Addr == "" {
		zap.L().Warn("未配置redis连接")
		return
	}
	client = redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		zap.L().Error("redis连接失败", zap.Error(err))
		return
	}
	zap.L().Info("连接redis成功")
	return
}

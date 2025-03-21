package flags

import (
	"FastFiber/global"
	"FastFiber/models"
	"go.uber.org/zap"
)

func MigrateDB() {
	err := global.DB.AutoMigrate(&models.UserModel{})
	if err != nil {
		zap.L().Error("表结构迁移失败", zap.Error(err))
		return
	}
	zap.L().Info("表结构迁移成功")
}

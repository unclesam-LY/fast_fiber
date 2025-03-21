package core

import (
	"FastFiber/config"
	"FastFiber/flags"
	"FastFiber/global"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func ReadConfig() (cfg *config.Config) {
	cfg = new(config.Config)
	byteData, err := os.ReadFile(flags.Options.File)
	if err != nil {
		zap.L().Error("配置文件读取出错", zap.Error(err))
		return
	}
	err = yaml.Unmarshal(byteData, cfg)
	if err != nil {
		zap.L().Error("配置文件格式错误", zap.Error(err))
		return
	}
	log.Println("配置文件读取成功")
	return
}

func DumpConfig() {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		zap.L().Error("配置文件转换错误", zap.String("path", flags.Options.File), zap.Error(err))
		return
	}
	err = os.WriteFile(flags.Options.File, byteData, 0666)
	if err != nil {
		zap.L().Error("配置文件写入错误", zap.String("path", flags.Options.File), zap.Error(err))
		return
	}
	zap.L().Info("配置文件保存成功",
		zap.String("path", flags.Options.File))
}

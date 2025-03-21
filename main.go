package main

import (
	"FastFiber/core"
	"FastFiber/flags"
	"FastFiber/global"
	"FastFiber/routers"
	"go.uber.org/zap"
)

func main() {
	flags.Parse()

	global.Config = core.ReadConfig()
	global.Log = core.InitZapLogger() // 再初始化日志
	zap.ReplaceGlobals(global.Log)

	global.DB = core.InitGorm()
	global.Redis = core.InitRedis()

	flags.Run()
	routers.Run()

}

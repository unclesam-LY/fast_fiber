package routers

import (
	"FastFiber/global"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go.uber.org/zap"
)

func Run() {
	r := fiber.New()

	// 添加中间件
	r.Use(
		recover.New(), // 错误恢复
		logger.New(logger.Config{ // 访问日志
			Format: "[${time}] | ${status} | ${latency} | ${ip} | ${method} | ${path} | ${error}\n",
		}),
	)
	// 静态文件路由
	r.Static("/uploads", "./uploads")
	g := r.Group("api")

	UserRouter(g)
	// 获取监听地址
	addr := global.Config.System.Addr()
	if global.Config.System.Mode == "release" {
		global.Log.Sugar().Infof("后端服务运行在 %s", addr)
	}

	err := r.Listen(addr)
	if err != nil {
		zap.L().Error("服务启动错误:", zap.Error(err))
		return
	}
}

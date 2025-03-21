package routers

import (
	"FastFiber/api"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(g fiber.Router) {
	app := api.App.UserApi
	g.Get("users", app.UserListView)
}

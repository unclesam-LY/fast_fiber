package user_api

import (
	"github.com/gofiber/fiber/v2"
)

func (UserApi) UserListView(c *fiber.Ctx) error {
	// 示例响应
	return c.JSON(fiber.Map{
		"code":    200,
		"message": "路由绑定成功",
		"data":    []string{"用户1", "用户2"},
	})
}

package user_api

import (
	"FastFiber/global"
	"FastFiber/models"
	"FastFiber/models/res"
	"FastFiber/service/common"
	"FastFiber/utils/desens"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func (UserApi) UserListView(c *fiber.Ctx) error {

	var page models.PageInfo
	err := c.QueryParser(&page)
	if err != nil {
		global.Log.Error("参数绑定失败", zap.Error(err))
		return res.FailWithCode(res.ArgumentError, c)
	}

	var users []models.UserModel
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
	})

	for _, user := range list {
		user.Email = desens.DesensitizationEmail(user.Email)
		users = append(users, user)
	}

	return res.OKWithList(list, count, c)
}

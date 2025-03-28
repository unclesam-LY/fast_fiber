package middleware

import (
	"FastFiber/service/redis_ser"
	"FastFiber/utils/jwts"
	"FastFiber/utils/res"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := c.GetRespHeader("access_token")
		_, err := jwts.CheckAccessToken(accessToken)
		if err != nil {
			return res.FailWithMessage("认证失败", c)
		}
		if redis_ser.HashLogout(accessToken) {
			return res.FailWithMessage("当前登录已经注销", c)
		}
		return c.Next()

	}
}

func AdminMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := c.GetRespHeader("access_token")
		claims, err := jwts.CheckAccessToken(accessToken)
		if err != nil {
			return res.FailWithMessage("认证失败", c)
		}

		if redis_ser.HashLogout(accessToken) {
			return res.FailWithMessage("当前登录已经注销", c)
		}

		if claims.RoleID != 1 {
			return res.FailWithMessage("角色认证失败", c)
		}
		c.Locals("claims", claims)
		return c.Next()

	}
}

func GetAuth(c *fiber.Ctx) (cl *jwts.MyClaims) {
	cl = new(jwts.MyClaims)
	_claims := c.Locals("claims")
	if _claims == nil {
		return nil
	}

	cl, ok := _claims.(*jwts.MyClaims)
	if !ok {
		return nil
	}

	return cl
}

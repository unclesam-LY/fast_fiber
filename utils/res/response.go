package res

import (
	"FastFiber/utils"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

type ListResponse[T any] struct {
	Count int64 `json:"count"`
	List  T     `json:"list"`
}

const (
	Success = 0
	Error   = 7
)

func Result(code int, data any, msg string, c *fiber.Ctx) error {
	return c.JSON(Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

// OK 响应成功
func OK(data any, msg string, c *fiber.Ctx) error {
	return Result(Success, data, msg, c)
}

func OKWithData(data any, c *fiber.Ctx) error {
	return Result(Success, data, "成功", c)
}

func OKWithList(list any, count int64, c *fiber.Ctx) error {
	return OKWithData(ListResponse[any]{
		List:  list,
		Count: count,
	}, c)
}

func OKWithMessage(msg string, c *fiber.Ctx) error {
	return Result(Success, map[string]any{}, msg, c)
}

func OkWith(c *fiber.Ctx) error {
	return Result(Success, map[string]any{}, "成功", c)
}

// Fail 响应失败
func Fail(data any, msg string, c *fiber.Ctx) error {
	return Result(Error, data, msg, c)
}

func FailWithMessage(msg string, c *fiber.Ctx) error {
	return Result(Error, map[string]any{}, msg, c)
}

func FailWithError(err error, obj any, c *fiber.Ctx) error {
	msg := utils.GetValidMsg(err, obj)
	return FailWithMessage(msg, c)
}

func FailWithCode(code ErrorCode, c *fiber.Ctx) error {
	msg, ok := ErrorMap[code]
	if ok {
		return Result(int(code), map[string]any{}, msg, c)

	}
	return Result(Error, map[string]any{}, "未知错误", c)

}

package middleware

import (
	"cufoon.litkeep.service/app/util"
	"github.com/gofiber/fiber/v2"
)

func Validate(s any) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		err := c.BodyParser(s)
		if err != nil {
			return util.ResFailH(c, 400, 1, "请求参数错误")
		}
		return c.Next()
	}
}

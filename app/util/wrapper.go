package util

import "github.com/gofiber/fiber/v2"

type Res struct {
	Code int    `json:"code"`
	Info string `json:"info"`
	Data any    `json:"data"`
}

func ResOK(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusOK).JSON(&Res{
		Code: 0,
		Info: "success",
		Data: data,
	})
}

func ResFail(c *fiber.Ctx, code int, message string) error {
	return c.Status(fiber.StatusOK).JSON(&Res{
		Code: code,
		Info: message,
		Data: nil,
	})
}

func ResFailH(c *fiber.Ctx, status int, code int, message string) error {
	return c.Status(status).JSON(&Res{
		Code: code,
		Info: message,
		Data: nil,
	})
}

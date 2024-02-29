package middleware

import (
	"time"

	"cufoon.litkeep.service/app/util"
	"cufoon.litkeep.service/pkg/flow"
	"cufoon.litkeep.service/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

func (mw *MiddleWare) Auth(c *fiber.Ctx) error {
	token := c.Get("Authorization", "")
	if token == "" {
		return c.SendStatus(401)
	}
	data, err := jwt.Parse(token)
	if err != nil {
		return c.SendStatus(401)
	}
	if data.ExpireTime <= time.Now().UnixMicro() {
		return util.ResFailH(c, 401, 2, "Login Expired")
	}
	err = flow.SetUserID(c, data.UserId)
	if err != nil {
		return c.SendStatus(500)
	}
	return c.Next()
}

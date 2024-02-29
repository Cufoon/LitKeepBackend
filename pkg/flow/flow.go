package flow

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func GetUserID(c *fiber.Ctx) string {
	xUserID := c.Locals(KeyUserID)
	userID, ok := xUserID.(string)
	if ok {
		return userID
	}
	return ""
}

func SetUserID(c *fiber.Ctx, uid string) error {
	c.Locals(KeyUserID, uid)
	id := GetUserID(c)
	if id == uid {
		return nil
	}
	return errors.New("set-error")
}

func GetSessionID(c *fiber.Ctx) string {
	xUserID := c.Locals(KeySessionID)
	userID, ok := xUserID.(string)
	if ok {
		return userID
	}
	return ""
}

func SetSessionID(c *fiber.Ctx, sid string) error {
	c.Locals(KeySessionID, sid)
	id := GetUserID(c)
	if id == sid {
		return nil
	}
	return errors.New("set-error")
}

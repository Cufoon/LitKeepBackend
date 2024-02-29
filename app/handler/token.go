package handler

import (
	"cufoon.litkeep.service/app/util"
	"github.com/gofiber/fiber/v2"
)

func NewTokenHandler() *TokenHandler {
	return &TokenHandler{}
}

type TokenHandler struct {
}

func (th *TokenHandler) Verify(c *fiber.Ctx) error {
	return util.ResOK(c, &map[string]any{"verified": 0})
}

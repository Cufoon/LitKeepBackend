package handler

import (
	"cufoon.litkeep.service/app/util"
	"cufoon.litkeep.service/app/vo"
	"github.com/gofiber/fiber/v2"
)

func NewOtherHandler() *OtherHandler {
	return &OtherHandler{}
}

type OtherHandler struct{}

func (o *OtherHandler) CheckAndroidAppUpdate(c *fiber.Ctx) error {
	data := new(vo.AndroidAppUpdateCheckReq)
	err := c.BodyParser(data)
	if err != nil {
		return util.ResFailH(c, 400, 1, "请求参数错误")
	}
	if data.Now < 30000001 {
		return util.ResOK(c, &vo.AndroidAppUpdateCheckRes{
			Update: true,
			URL:    "https://cufoon.com",
		})
	}
	return util.ResOK(c, &vo.AndroidAppUpdateCheckRes{
		Update: false,
	})
}

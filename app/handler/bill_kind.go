package handler

import (
	"cufoon.litkeep.service/app/schema"
	"cufoon.litkeep.service/app/service"
	"cufoon.litkeep.service/app/util"
	"cufoon.litkeep.service/app/vo"
	"cufoon.litkeep.service/pkg/flow"
	"github.com/gofiber/fiber/v2"
)

func NewBillKindHandler(billKindService *service.BillKindService) *BillKindHandler {
	return &BillKindHandler{billKindService: billKindService}
}

type BillKindHandler struct {
	billKindService *service.BillKindService
}

func (bkh *BillKindHandler) Create(c *fiber.Ctx) error {
	data := new(vo.BillKind)
	err := c.BodyParser(data)
	if err != nil {
		return util.ResFailH(c, 400, 1, "请求参数错误")
	}
	err = bkh.billKindService.Create(&schema.BillKind{
		UserID:      data.UserID,
		KindID:      data.KindID,
		Name:        data.Name,
		Description: data.Description,
	})
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &map[string]any{"created": true})
}

func (bkh *BillKindHandler) Query(c *fiber.Ctx) error {
	userID := flow.GetUserID(c)
	if userID == "" {
		return util.ResFailH(c, 400, 1, "请求参数错误")
	}
	kinds, err := bkh.billKindService.Query(userID)
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &map[string]any{"kind": kinds})
}

func (bkh *BillKindHandler) Modify(c *fiber.Ctx) error {
	data := new(vo.BillKind)
	err := c.BodyParser(data)
	if err != nil || data.UserID == "" || data.KindID == "" {
		return util.ResFailH(c, 400, 1, "请求参数错误")
	}
	userID := flow.GetUserID(c)
	if userID == "" {
		return util.ResFailH(c, 400, 1, "请求参数错误")
	}
	err = bkh.billKindService.Modify(userID, &schema.BillKind{
		UserID:      data.UserID,
		KindID:      data.KindID,
		Name:        data.Name,
		Description: data.Description,
	})
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &map[string]any{"modified": true})
}

func (bkh *BillKindHandler) Delete(c *fiber.Ctx) error {
	data := new(vo.BillKind)
	err := c.BodyParser(data)
	if err != nil {
		return util.ResFailH(c, 400, 1, "请求参数错误")
	}
	userID := flow.GetUserID(c)
	err = bkh.billKindService.Delete(userID, data.KindID)
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &map[string]any{"deleted": true})
}

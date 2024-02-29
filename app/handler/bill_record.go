package handler

import (
	"cufoon.litkeep.service/app/schema"
	"cufoon.litkeep.service/app/service"
	"cufoon.litkeep.service/app/util"
	"cufoon.litkeep.service/app/vo"
	"cufoon.litkeep.service/pkg/flow"
	"github.com/gofiber/fiber/v2"
)

func NewBillRecordHandler(billRecordService *service.BillRecordService, billKindService *service.BillKindService) *BillRecordHandler {
	return &BillRecordHandler{billRecordService: billRecordService, billKindService: billKindService}
}

type BillRecordHandler struct {
	billRecordService *service.BillRecordService
	billKindService   *service.BillKindService
}

func (brh *BillRecordHandler) Create(c *fiber.Ctx) error {
	data := new(vo.BillRecordCreateReq)
	err := c.BodyParser(data)
	if err != nil {
		return util.ResFailH(c, 400, 1, "请求参数错误")
	}
	userID := flow.GetUserID(c)
	err = brh.billRecordService.Create(&schema.BillRecord{
		UserID: userID,
		KindID: data.KindID,
		Type:   data.Type,
		Value:  data.Value,
		Time:   data.Time,
		Mark:   data.Mark,
	})
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &map[string]any{"created": true})
}

func (brh *BillRecordHandler) Query(c *fiber.Ctx) error {
	data := new(vo.BillRecordQueryReq)
	err := c.BodyParser(data)
	if err != nil {
		return util.ResFailH(c, 400, 1, "请求参数错误")
	}
	userID := flow.GetUserID(c)
	record, err := brh.billRecordService.Query(&schema.BillRecordQuery{
		UserID:    userID,
		KindID:    data.KindID,
		StartTime: data.StartTime,
		EndTime:   data.EndTime,
	})
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &map[string]any{"record": record})
}

func (brh *BillRecordHandler) QueryPage(c *fiber.Ctx) error {
	data := new(vo.BillRecordPageQueryReq)
	err := c.BodyParser(data)
	if err != nil {
		return util.ResFailH(c, 400, 1, "请求参数错误")
	}
	userID := flow.GetUserID(c)
	record, err := brh.billRecordService.QueryPage(&schema.BillRecordPageQuery{
		UserID: userID,
		Page:   data.Page,
	})
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &map[string]any{"record": record})
}

func (brh *BillRecordHandler) QueryPageData(c *fiber.Ctx) error {
	userID := flow.GetUserID(c)
	count, err := brh.billRecordService.QueryRecordCount(userID)
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	kinds, err := brh.billKindService.QueryLinear(userID)
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	// 80
	return util.ResOK(c, &map[string]any{
		"kinds": kinds,
		"pageData": &map[string]any{
			"total":      count,
			"totalPages": (count-1)/20 + 1,
			"pageSize":   20,
		}})
}

func (brh *BillRecordHandler) QueryWithKind(c *fiber.Ctx) error {
	data := new(vo.BillRecordQueryReq)
	err := c.BodyParser(data)
	if err != nil {
		return util.ResFailH(c, 400, 1, "请求参数错误")
	}
	userID := flow.GetUserID(c)
	record, err := brh.billRecordService.Query(&schema.BillRecordQuery{
		UserID:    userID,
		KindID:    data.KindID,
		StartTime: data.StartTime,
		EndTime:   data.EndTime,
	})
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	kinds, err := brh.billKindService.QueryLinear(userID)
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &map[string]any{"kinds": kinds, "record": record})
}

func (brh *BillRecordHandler) Modify(c *fiber.Ctx) error {
	data := new(vo.BillRecord)
	err := c.BodyParser(data)
	if err != nil || data.ID == 0 {
		return util.ResFailH(c, 400, 1, "请求参数错误")
	}
	err = brh.billRecordService.Modify(&schema.BillRecord{
		ID:     data.ID,
		UserID: data.UserID,
		KindID: data.KindID,
		Type:   data.Type,
		Value:  data.Value,
		Time:   data.Time,
		Mark:   data.Mark,
	})
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &map[string]any{"modified": true})
}

func (brh *BillRecordHandler) Delete(c *fiber.Ctx) error {
	data := new(vo.BillRecordDeleteReq)
	err := c.BodyParser(data)
	if err != nil {
		return util.ResFailH(c, 400, 1, "请求参数错误")
	}
	userID := flow.GetUserID(c)
	notDeleted, err := brh.billRecordService.Delete(data.Ids, userID)
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &map[string]any{"notDeleted": notDeleted})
}

func (brh *BillRecordHandler) QueryStatisticsDay(c *fiber.Ctx) error {
	data := new(vo.BillRecordStatisticsDayQueryReq)
	err := c.BodyParser(data)
	if err != nil {
		return util.ResFailH(c, 400, 1, "请求参数错误")
	}
	userID := flow.GetUserID(c)
	result, err := brh.billRecordService.QueryStatisticsDay(userID, data)
	if err != nil {
		return util.ResFail(c, 1, err.Error())
	}
	return util.ResOK(c, &map[string]any{"statistic": result})
}

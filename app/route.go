package app

import (
	"cufoon.litkeep.service/app/middleware"
	"github.com/gofiber/fiber/v2"
)

func InitRoute(app *fiber.App, c *Handler, m *middleware.MiddleWare) {
	v1 := app.Group("/v1")

	v1.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	v1.Get("/TokenVerify", m.Auth, c.TokenHandler.Verify)

	v1.Put("/User", c.UserHandler.Register)
	v1.Post("/User", c.UserHandler.Login)
	v1.Get("/UserInfo", m.Auth, c.UserHandler.GetInfo)
	v1.Get("/UserIcon/:id", c.UserHandler.GetIcon)
	v1.Post("/UserInfoNickNameChange", m.Auth, c.UserHandler.ChangeNickName)
	v1.Post("/UserInfoIconChange", m.Auth, c.UserHandler.ChangeIcon)

	v1.Put("/Kind", m.Auth, c.BillKindHandler.Create)
	v1.Post("/Kind", m.Auth, c.BillKindHandler.Query)
	v1.Patch("/Kind", m.Auth, c.BillKindHandler.Modify)
	v1.Post("/KindDelete", m.Auth, c.BillKindHandler.Delete)

	v1.Put("/Bill", m.Auth, c.BillRecordHandler.Create)
	v1.Post("/Bill", m.Auth, c.BillRecordHandler.Query)
	v1.Patch("/Bill", m.Auth, c.BillRecordHandler.Modify)
	v1.Post("/BillDelete", m.Auth, c.BillRecordHandler.Delete)
	v1.Post("/BillAndKind", m.Auth, c.BillRecordHandler.QueryWithKind)
	v1.Post("/BillPageData", m.Auth, c.BillRecordHandler.QueryPageData)
	v1.Post("/BillPage", m.Auth, c.BillRecordHandler.QueryPage)
	v1.Post("/BillStatisticDay", m.Auth, c.BillRecordHandler.QueryStatisticsDay)

	v1.Post("/AndroidApp", c.OtherHandler.CheckAndroidAppUpdate)
}

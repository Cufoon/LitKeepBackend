package app

import (
	"fmt"

	dao2 "cufoon.litkeep.service/app/db/dao"
	entity2 "cufoon.litkeep.service/app/db/entity"

	"cufoon.litkeep.service/app/handler"
	"cufoon.litkeep.service/app/middleware"
	"cufoon.litkeep.service/app/service"
	"cufoon.litkeep.service/conf"
	"cufoon.litkeep.service/pkg/db"
	"cufoon.litkeep.service/pkg/jwt"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

const aesKey = ".xW@~uH^H>u)*!Vhii.q:6uR,5o%ZA}C"
const edPrivateKey = "WZkixDXuIVSiHU0Fn8ire/bUTNOxFW11pWMe1/caDd6l1YgKSC6y3itj4H20zzccsIh7/dSJPzhJFdROuTQqVA=="
const edPublicKey = "pdWICkgust4rY+B9tM83HLCIe/3UiT84SRXUTrk0KlQ="

func StartAPP(path string) {
	err := jwt.Init("LT", aesKey, edPrivateKey, edPublicKey)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	app := fiber.New(fiber.Config{
		Prefork:       false,
		ServerHeader:  "Cufoon.LitKeep.Server",
		CaseSensitive: true,
		AppName:       "LitKeep",
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
		//ReduceMemoryUsage: false,
		//Immutable:         false,
	})

	c, err := conf.NewConf(path)
	if err != nil {
		println(err.Error())
		return
	}

	mariaDB, err := db.NewDB(c)
	if err != nil {
		println(err.Error())
		return
	}

	userDAO := dao2.NewUserDAO(mariaDB, &entity2.User{})
	billKindDAO := dao2.NewBillKindDAO(mariaDB, &entity2.BillKind{})
	billRecordDAO := dao2.NewBillRecordDAO(mariaDB, &entity2.BillRecord{})

	userService := service.NewUserService(userDAO)
	billKindService := service.NewBillKindService(billKindDAO)
	billRecordService := service.NewBillRecordService(billRecordDAO)

	userHandler := handler.NewUserHandler(userService)
	billKindHandler := handler.NewBillKindHandler(billKindService)
	billRecordHandler := handler.NewBillRecordHandler(billRecordService, billKindService)

	tokenHandler := handler.NewTokenHandler()
	otherHandler := handler.NewOtherHandler()

	middleWare := middleware.NewMiddleWare()

	InitRoute(app, &Handler{
		UserHandler:       userHandler,
		BillKindHandler:   billKindHandler,
		BillRecordHandler: billRecordHandler,
		TokenHandler:      tokenHandler,
		OtherHandler:      otherHandler,
	}, middleWare)

	err = app.Listen(c.Server.Port)
	if err != nil {
		println(err.Error())
	}
}

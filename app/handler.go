package app

import "cufoon.litkeep.service/app/handler"

type Handler struct {
	UserHandler       *handler.UserHandler
	BillKindHandler   *handler.BillKindHandler
	BillRecordHandler *handler.BillRecordHandler
	TokenHandler      *handler.TokenHandler
	OtherHandler      *handler.OtherHandler
}

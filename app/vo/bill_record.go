package vo

import (
	"time"
)

type BillRecord struct {
	ID     uint      `json:"id"`
	UserID string    `json:"userID" validate:"required"`
	KindID string    `json:"kindID" validate:"required"`
	Type   int       `json:"type" validate:"required"`
	Value  float64   `json:"value" validate:"required"`
	Time   time.Time `json:"time"`
	Mark   string    `json:"mark"`
}

type BillRecordQueryReq struct {
	KindID    string    `json:"kindID"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

type BillRecordPageQueryReq struct {
	Page int `json:"page"`
}

type BillRecordDeleteReq struct {
	Ids []uint `json:"ids" validate:"required"`
}

type BillRecordCreateReq struct {
	KindID string    `json:"kindID" validate:"required"`
	Type   int       `json:"type" validate:"required"`
	Value  float64   `json:"value" validate:"required"`
	Time   time.Time `json:"time"`
	Mark   string    `json:"mark"`
}

type BillRecordStatisticsDayQueryReq struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}

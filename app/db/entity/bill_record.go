package entity

import (
	"time"

	"gorm.io/gorm"
)

type BillRecord struct {
	gorm.Model
	UserID string `gorm:"column:id_user"`
	Type   int
	Kind   string
	Value  float64
	Time   time.Time
	Mark   string
}

func (BillRecord) TableName() string {
	return "bill_record"
}

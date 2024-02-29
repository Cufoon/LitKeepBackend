package entity

import "gorm.io/gorm"

type BillKind struct {
	gorm.Model
	KindID        string `gorm:"column:id_kind"`
	UserID        string `gorm:"column:id_user"`
	Name          string
	Owner         string
	Description   string
	ReplaceSystem string `gorm:"column:replace_system"`
}

func (BillKind) TableName() string {
	return "bill_kind"
}

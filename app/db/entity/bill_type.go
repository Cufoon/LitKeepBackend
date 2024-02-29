package entity

import (
	"gorm.io/gorm"
)

type BillType struct {
	gorm.Model
	Type        int
	Name        string
	Description string
}

func (BillType) TableName() string {
	return "bill_type"
}

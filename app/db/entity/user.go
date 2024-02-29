package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   string `gorm:"column:id_user"`
	Password string `gorm:"column:passwd"`
	NickName string
	Icon     []byte `gorm:"type:longblob,column:icon"`
	Email    string
	Status   byte
}

func (User) TableName() string {
	return "user"
}

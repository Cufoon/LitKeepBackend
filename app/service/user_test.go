package service

import (
	"testing"

	"cufoon.litkeep.service/app/db/dao"
	"cufoon.litkeep.service/app/db/entity"
	"cufoon.litkeep.service/app/schema"
	"cufoon.litkeep.service/conf"
	"cufoon.litkeep.service/pkg/db"
)

func TestCreateUser(t *testing.T) {
	gc, err := conf.NewConf("../../dev.yaml")
	if err != nil {
		t.Error(err)
		return
	}
	mariaDB, err := db.NewDB(gc)
	if err != nil {
		t.Error(err)
		return
	}
	userDAO := dao.NewUserDAO(mariaDB, &entity.User{})
	userService := NewUserService(userDAO)
	err = userService.Register(&schema.UserRegisterData{
		Email:    "cufoon@gmail.com",
		Password: "123",
	})
	if err != nil {
		t.Error(err)
		return
	}
}

package dao

import (
	"cufoon.litkeep.service/app/db/dto"
	"cufoon.litkeep.service/app/db/entity"
	"gorm.io/gorm"
)

func NewUserDAO(db *gorm.DB, userModel *entity.User) *UserDAO {
	return &UserDAO{db: db, userModel: userModel}
}

type UserDAO struct {
	db        *gorm.DB
	userModel *entity.User
}

func (u *UserDAO) model() *gorm.DB {
	return u.db.Model(u.userModel)
}

func (u *UserDAO) Create(userInfo *dto.UserCreateData) error {
	result := u.model().Select("UserID", "Email", "Password").Create(&entity.User{
		UserID:   userInfo.UserID,
		Email:    userInfo.Email,
		Password: userInfo.Password,
	})
	return result.Error
}

func (u *UserDAO) QueryByUserID(userID string) (*entity.User, error) {
	var user entity.User
	result := u.model().Where("id_user", userID).First(&user)
	return &user, result.Error
}

func (u *UserDAO) QueryByEmail(email string) (*entity.User, error) {
	var user entity.User
	result := u.model().Where("email", email).First(&user)
	return &user, result.Error
}

func (u *UserDAO) ModifyNickNameByUserID(changeInfo *dto.UserChangeNickNameData) error {
	result := u.model().Where("id_user", changeInfo.UserID).Update("nick_name", changeInfo.NickName)
	return result.Error
}

func (u *UserDAO) ModifyIconByUserID(changeInfo *dto.UserChangeIconData) error {
	result := u.model().Where("id_user", changeInfo.UserID).Update("icon", changeInfo.Icon)
	return result.Error
}

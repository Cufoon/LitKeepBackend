package service

import (
	"errors"
	"time"

	"cufoon.litkeep.service/app/db/dao"
	"cufoon.litkeep.service/app/db/dto"

	"cufoon.litkeep.service/app/constant"
	"cufoon.litkeep.service/app/schema"
	"cufoon.litkeep.service/app/util"
	"cufoon.litkeep.service/pkg/jwt"
	"gorm.io/gorm"
)

func NewUserService(userDAO *dao.UserDAO) *UserService {
	return &UserService{
		userDAO: userDAO,
	}
}

type UserService struct {
	userDAO *dao.UserDAO
}

func (us *UserService) Register(userInfo *schema.UserRegisterData) error {
	_, err := us.userDAO.QueryByEmail(userInfo.Email)
	if err == nil {
		return constant.ErrAccountExist
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	newID := util.Generate8Bytes()
	_, err = us.userDAO.QueryByUserID(newID)
	if err == nil {
		return constant.ErrAccountIDExist
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err1 := us.userDAO.Create(&dto.UserCreateData{
			UserID:   newID,
			Email:    userInfo.Email,
			Password: userInfo.Password,
		})
		if err1 != nil {
			return err1
		}
	} else {
		return err
	}
	return nil
}

func (us *UserService) Token(id string, sid uint8, long time.Duration) (string, error) {
	now := time.Now()
	return jwt.Token(&jwt.TokenProperty{
		UserId:     id,
		SessionId:  sid,
		SignedTime: now.UnixMicro(),
		ExpireTime: now.Add(long).UnixMicro(),
	})
}

func (us *UserService) Login(info *schema.UserLoginData) (string, error) {
	user, err := us.userDAO.QueryByEmail(info.Email)
	if err == nil {
		if user.Password == info.Password {
			token, err1 := us.Token(user.UserID, 0, time.Hour*24)
			if err1 != nil {
				return "", err1
			}
			return token, nil
		} else {
			return "", constant.ErrLoginPassWrong
		}
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", constant.ErrLoginEmailNotExist
	}
	return "", err
}

func (us *UserService) GetInfo(userID string) (*schema.UserInfo, error) {
	user, err := us.userDAO.QueryByUserID(userID)
	if err == nil {
		hasIcon := true
		if user.Icon == nil {
			hasIcon = false
		}
		return &schema.UserInfo{
			NickName:   user.NickName,
			UserID:     user.UserID,
			Email:      user.Email,
			HasIcon:    hasIcon,
			UpdateTime: user.UpdatedAt.Local().String(),
		}, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constant.ErrLoginEmailNotExist
	}
	return nil, err
}

func (us *UserService) GetIcon(userID string) ([]byte, error) {
	user, err := us.userDAO.QueryByUserID(userID)
	if err == nil {
		return user.Icon, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, constant.ErrLoginEmailNotExist
	}
	return nil, err
}

func (us *UserService) SetUserNickName(userID string, nickName string) error {
	return us.userDAO.ModifyNickNameByUserID(&dto.UserChangeNickNameData{
		UserID:   userID,
		NickName: nickName,
	})
}

func (us *UserService) SetUserIcon(userID string, icon []byte) error {
	return us.userDAO.ModifyIconByUserID(&dto.UserChangeIconData{
		UserID: userID,
		Icon:   icon,
	})
}

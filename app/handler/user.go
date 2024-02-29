package handler

import (
	"errors"
	"io"

	"cufoon.litkeep.service/app/constant"
	"cufoon.litkeep.service/app/schema"
	"cufoon.litkeep.service/app/service"
	"cufoon.litkeep.service/app/util"
	"cufoon.litkeep.service/app/vo"
	"cufoon.litkeep.service/pkg/flow"
	"github.com/gofiber/fiber/v2"
)

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

type UserHandler struct {
	userService *service.UserService
}

func (uh *UserHandler) Register(c *fiber.Ctx) error {
	data := new(vo.UserRegisterReq)
	err := c.BodyParser(data)
	if err != nil {
		return util.ResFailH(c, 400, 1, "请求参数错误")
	}
	err = uh.userService.Register(&schema.UserRegisterData{
		Email:    data.Email,
		Password: data.Password,
	})
	if err != nil {
		if errors.Is(err, constant.ErrAccountExist) {
			return util.ResFail(c, 1, "该邮箱已经注册过")
		}
		if errors.Is(err, constant.ErrAccountIDExist) {
			return util.ResFail(c, 2, "内部id重复，请稍后重试！")
		}
		return err
	}
	token, err := uh.userService.Login(&schema.UserLoginData{
		Email:    data.Email,
		Password: data.Password,
	})
	if err != nil {
		if errors.Is(err, constant.ErrLoginEmailNotExist) || errors.Is(err, constant.ErrLoginPassWrong) {
			return util.ResFail(c, 2, "账户或者密码错误")
		}
		return err
	}
	return util.ResOK(c, &vo.UserRegisterRes{
		Signed: true,
		Token:  token,
	})
}

func (uh *UserHandler) Login(c *fiber.Ctx) error {
	data := new(vo.UserLoginReq)
	err := c.BodyParser(data)
	if err != nil {
		return util.ResFailH(c, 400, 1, "请求参数错误")
	}
	token, err := uh.userService.Login(&schema.UserLoginData{
		Email:    data.Email,
		Password: data.Password,
	})
	if err != nil {
		if errors.Is(err, constant.ErrLoginEmailNotExist) || errors.Is(err, constant.ErrLoginPassWrong) {
			return util.ResFail(c, 2, "账户或者密码错误")
		}
		return err
	}
	return util.ResOK(c, &vo.UserLoginRes{
		Logined: true,
		Token:   token,
	})
}

func (uh *UserHandler) GetInfo(c *fiber.Ctx) error {
	userID := flow.GetUserID(c)
	if userID == "" {
		return util.ResFailH(c, 400, 1, "请求参数错误")
	}
	info, err := uh.userService.GetInfo(userID)
	if err != nil {
		return err
	}
	iconPath := ""
	if info.HasIcon {
		iconPath = "UserIcon/" + info.UserID + "?time=" + info.UpdateTime
	}
	return util.ResOK(c, &vo.UserGetInfoRes{
		NickName: info.NickName,
		UserID:   info.UserID,
		Email:    info.Email,
		IconPath: iconPath,
	})
}

func (uh *UserHandler) ChangeIcon(c *fiber.Ctx) error {
	userID := flow.GetUserID(c)
	if userID == "" {
		return util.ResFailH(c, 400, 1, "wrong params")
	}
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return err
	}
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()
	fileB, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	err = uh.userService.SetUserIcon(userID, fileB)
	if err != nil {
		return err
	}
	return util.ResOK(c, &vo.UserChangeNickNameRes{
		Changed: true,
	})
}

func (uh *UserHandler) GetIcon(c *fiber.Ctx) error {
	userID := c.Params("id")
	if userID == "" {
		return util.ResFailH(c, 400, 1, "请求参数错误")
	}
	icon, err := uh.userService.GetIcon(userID)
	if err != nil {
		return err
	}
	c.Set("Content-Type", "image/jpeg")
	return c.Send(icon)
}

func (uh *UserHandler) ChangeNickName(c *fiber.Ctx) error {
	userID := flow.GetUserID(c)
	if userID == "" {
		return util.ResFailH(c, 401, 1, "请求参数错误")
	}
	data := new(vo.UserChangeNickNameReq)
	err := c.BodyParser(data)
	if err != nil {
		return util.ResFailH(c, 400, 1, "请求参数错误")
	}
	err = uh.userService.SetUserNickName(userID, data.NickName)
	if err != nil {
		return util.ResFail(c, 1, "修改失败")
	}
	return util.ResOK(c, &vo.UserChangeNickNameRes{
		Changed: true,
	})
}

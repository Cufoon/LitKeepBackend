package vo

type UserInfo struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLoginReq struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserLoginRes struct {
	Logined bool   `json:"logined"`
	Token   string `json:"token"`
}

type UserRegisterReq struct {
	NickName string `json:"nickname"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserRegisterRes struct {
	Signed bool   `json:"signed"`
	Token  string `json:"token"`
}

type UserGetInfoRes struct {
	NickName string `json:"nickname"`
	UserID   string `json:"userID"`
	Email    string `json:"email"`
	IconPath string `json:"iconPath"`
}

type UserChangeNickNameReq struct {
	NickName string `json:"nickname"`
}

type UserChangeNickNameRes struct {
	Changed bool `json:"changed"`
}

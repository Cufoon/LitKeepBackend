package schema

type UserLoginData struct {
	Email    string
	Password string
}

type UserRegisterData struct {
	NickName string
	Email    string
	Password string
}

type UserInfo struct {
	NickName   string
	UserID     string
	Email      string
	HasIcon    bool
	UpdateTime string
}

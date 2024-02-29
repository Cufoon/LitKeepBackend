package dto

type UserCreateData struct {
	UserID   string
	Email    string
	Password string
}

type UserChangeNickNameData struct {
	UserID   string
	NickName string
}

type UserChangeIconData struct {
	UserID string
	Icon   []byte
}

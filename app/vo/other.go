package vo

type AndroidAppUpdateCheckReq struct {
	Now int64 `json:"now"`
}

type AndroidAppUpdateCheckRes struct {
	Update bool   `json:"update"`
	URL    string `json:"url"`
}

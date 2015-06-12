package models

type UserInfo struct {
	Basic
	Data UserInfoData `json:"data"`
}

type UserInfoData struct {
	UserName string `json:"username"`
}

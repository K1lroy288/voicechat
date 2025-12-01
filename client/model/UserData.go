package model

type UserData struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

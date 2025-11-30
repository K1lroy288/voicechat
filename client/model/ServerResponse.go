package model

type ServerResponse struct {
	HttpCode int
	Token    string `json:"token"`
}

package model

type ServerResponse struct {
	HttpCode int
	Error    string `json:"error"`
	Token    string `json:"token"`
	User     User   `json:"user"`
}

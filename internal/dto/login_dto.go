package dto

type ReqLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ResLogin struct {
	Token string `json:"token"`
}
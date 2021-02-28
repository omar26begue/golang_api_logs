package models

type UsuariosSysResponse struct {
	Identificador string `json:"identificador" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid" gorm:"primaryKey"`
	Login         string `json:"login" example:"usuario" format:"string" validate:"required,min=3"`
	IsAdmin       bool   `json:"is_admin" example:"false" format:"bool"`
	token         string `json:"token" example:"jwt token" format:"string"`
}

type UsuariosSysRequest struct {
	Login    string `json:"login" example:"usuario" format:"string"`
	Password string `json:"password" example:"password" format:"string"`
}

type UsuariosSys struct {
	Identificador string `json:"identificador" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid" gorm:"primaryKey"`
	Login         string `json:"login" example:"usuario" format:"string" validate:"required,min=3"`
	Password      string `json:"password" example:"password" format:"string" validate:"required,min=5"`
	IsAdmin       bool   `json:"is_admin" example:"false" format:"bool"`
	token         string `json:"token" example:"jwt token" format:"string"`
}

type UsuariosSysCreate struct {
	Login       string `json:"login" example:"usuario" format:"string" validate:"required,min=3"`
	Password    string `json:"password" example:"password" format:"string" validate:"required,min=5"`
	NewPassword string `json:"new_password" example:"password" format:"string" validate:"required,min=5"`
}

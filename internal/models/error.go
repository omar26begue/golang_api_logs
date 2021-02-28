package models

type HTTPError400 struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type HTTPError401 struct {
	Code    int    `json:"code" example:"401"`
	Message string `json:"message" example:"status bad request"`
}

type HTTPError500 struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"status bad request"`
}

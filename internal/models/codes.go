package models

type Codes struct {
	Identificador string `json:"identificador" gorm:"primary_key"`
	Code          string `json:"code" gorm:"not null"`
	Descipcion    string `json:"descipcion" gorm:"null"`
}

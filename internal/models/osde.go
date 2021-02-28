package models

type OsdeResponse struct {
	Identificador string `json:"identificador" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid" gorm:"primaryKey"`
	NombreOsde    string `json:"nombre_osde" example:"Cubanacan" format:"string" validate:"required,min=3"`
}

type OsdeRequest struct {
	NombreOsde string `json:"nombre_osde" example:"Cubanacan" format:"string" validate:"required,min=3"`
}

type Osde struct {
	Identificador string `json:"identificador" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid" gorm:"primary_key"`
	NombreOsde    string `json:"nombre_osde" example:"Cubanacan" format:"string" validate:"required,min=3"`
}

package interfaces

import (
	"github.com/google/uuid"
	"go-rest-api-logs/internal/database"
	"go-rest-api-logs/internal/models"
	"gopkg.in/go-playground/validator.v9"
)

func GetAll() []models.Osde {
	var db, closeDB = database.GetConnection()
	defer closeDB.Close()

	var osde []models.Osde
	db.Order("nombre_osde asc").Find(&osde)

	return osde
}

func GetFindOneOSDE(identificador string) (models.Osde, int64) {
	var db, closeDB = database.GetConnection()
	defer closeDB.Close()

	var result models.Osde
	var count int64
	db.Model(&models.Osde{}).Where(&models.Osde{Identificador: identificador}).Find(&result).Count(&count)

	return result, count
}

func GetFindOneOSDEName(name string) (models.Osde, int64) {
	var db, closeDB = database.GetConnection()
	defer closeDB.Close()

	var result models.Osde
	var count int64
	db.Model(&models.Osde{}).Where(&models.Osde{NombreOsde: name}).Find(&result).Count(&count)

	return result, count
}

func CreateOsde(osde *models.Osde) (models.Osde, models.HTTPError400, bool) {
	var db, closeDB = database.GetConnection()
	defer closeDB.Close()

	var validate = validator.New()
	err := validate.Struct(osde)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return models.Osde{}, models.HTTPError400{Code: 400, Message: err.Tag()}, true
		}
	}

	// verificando que no existe la osde
	_, count := GetFindOneOSDEName(osde.NombreOsde)

	if count != 0 {
		return models.Osde{}, models.HTTPError400{Code: 400, Message: "La OSDE ya existe en el sistema"}, true
	}

	osde.Identificador = uuid.New().String()
	db.Create(osde)

	return *osde, models.HTTPError400{}, false
}

func UpdateOsde(osde *models.Osde) (models.Osde, models.HTTPError400, bool) {
	var db, closeDB = database.GetConnection()
	defer closeDB.Close()

	_, count := GetFindOneOSDE(osde.Identificador)

	if count == 0 {
		return models.Osde{}, models.HTTPError400{Code: 400, Message: "La OSDE no existe en el sistema"}, true
	}

	db.Model(models.Osde{}).Where("identificador = ?", osde.Identificador).Update("nombre_osde", osde.NombreOsde)

	result, _ := GetFindOneOSDE(osde.Identificador)

	return result, models.HTTPError400{}, false
}

func DeleteOsde(idosde string) (models.Osde, models.HTTPError400, bool) {
	var db, closeDB = database.GetConnection()
	defer closeDB.Close()

	result, count := GetFindOneOSDE(idosde)

	if count == 0 {
		return models.Osde{}, models.HTTPError400{Code: 400, Message: "La OSDE no existe en el sistema."}, true
	}

	db.Model(models.Osde{}).Where("identificador = ?", idosde).Delete(models.Osde{})

	return result, models.HTTPError400{}, false
}

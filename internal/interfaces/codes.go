package interfaces

import (
	"github.com/google/uuid"
	"go-rest-api-logs/internal/database"
	"go-rest-api-logs/internal/models"
	"gopkg.in/go-playground/validator.v9"
)

func GetAllCodesResponse() []models.Codes {
	var db, closeDB = database.GetConnection()
	defer closeDB.Close()

	var codes []models.Codes
	db.Order("code asc").Find(&codes)

	return codes
}

func GetFindOneCodesName(name string) (models.Codes, int64) {
	var db, closeDB = database.GetConnection()
	defer closeDB.Close()

	var result models.Codes
	var count int64
	db.Model(&models.Codes{}).Where(&models.Codes{Code: name}).Find(&result).Count(&count)

	return result, count
}

func CreateCode(code *models.Codes) (models.Codes, models.HTTPError400, bool) {
	var db, closeDB = database.GetConnection()
	defer closeDB.Close()

	_, count := GetFindOneCodesName(code.Code)

	if count != 0 {
		return models.Codes{}, models.HTTPError400{Code: 400, Message: "El código ya existe en el sistema"}, true
	}

	var validate = validator.New()
	err := validate.Struct(code)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return models.Codes{}, models.HTTPError400{Code: 400, Message: err.Tag()}, true
		}
	}

	code.Identificador = uuid.New().String()
	db.Create(code)

	return *code, models.HTTPError400{}, false
}

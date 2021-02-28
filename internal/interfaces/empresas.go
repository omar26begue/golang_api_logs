package interfaces

import (
	"github.com/google/uuid"
	"go-rest-api-logs/internal/database"
	"go-rest-api-logs/internal/models"
	"gopkg.in/go-playground/validator.v9"
)

func GetAllEmpresas() []models.EmpresasResponse {
	var db, closeDB = database.GetConnection()
	defer closeDB.Close()

	var empresas []models.Empresas
	db.Order("nombre_empresa asc").Where(&models.Empresas{Activo: true}).Find(&empresas)

	var tempEmpresa = []models.EmpresasResponse{}

	for _, emp := range empresas {
		var tEmpresa = models.EmpresasResponse{
			Identificador: emp.Identificador,
			NombreEmpresa: emp.NombreEmpresa,
			IdOsde:        emp.IdOsde,
			Activo:        emp.Activo}
		tempEmpresa = append(tempEmpresa, tEmpresa)
	}

	return tempEmpresa
}

func GetFindOneEmpresa(identificador string) (models.EmpresasResponse, int64) {
	var db, closeDB = database.GetConnection()
	defer closeDB.Close()

	var result models.Empresas
	var count int64
	db.Model(&models.Empresas{}).Where(&models.Empresas{Identificador: identificador}).Find(&result).Count(&count)

	var tempEmpresa = models.EmpresasResponse{}

	tempEmpresa.Identificador = result.Identificador
	tempEmpresa.NombreEmpresa = result.NombreEmpresa
	tempEmpresa.IdOsde = result.IdOsde
	tempEmpresa.Activo = result.Activo

	return tempEmpresa, count
}

func GetFindOneEmpresaName(name string) (models.Empresas, int64) {
	var db, closeDB = database.GetConnection()
	defer closeDB.Close()

	var result models.Empresas
	var count int64
	db.Model(&models.Empresas{}).Where(&models.Empresas{NombreEmpresa: name}).Find(&result).Count(&count)

	return result, count
}

func CreateEmpresa(empresa *models.Empresas) (models.Empresas, models.HTTPError400, bool) {
	var db, closeDB = database.GetConnection()
	defer closeDB.Close()

	resultOsde, count := GetFindOneOSDE(empresa.IdOsde)

	if count == 0 {
		return models.Empresas{}, models.HTTPError400{Code: 400, Message: "Ha ocurrido un problema al agregar la empresa."}, true
	}

	empresa.Osde = resultOsde

	var validate = validator.New()
	err := validate.Struct(empresa)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return models.Empresas{}, models.HTTPError400{Code: 400, Message: err.Tag()}, true
		}
	}

	// verificando que no existe la osde
	_, countE := GetFindOneEmpresaName(empresa.NombreEmpresa)

	if countE != 0 {
		return models.Empresas{}, models.HTTPError400{Code: 400, Message: "La empresa ya existe en el sistema"}, true
	}

	empresa.Identificador = uuid.New().String()
	db.Create(empresa)

	return *empresa, models.HTTPError400{}, false
}

func UpdateEmpresa(empresa models.EmpresasRequest, identificador string) (models.EmpresasResponse, models.HTTPError400, bool) {
	var db, closeDB = database.GetConnection()
	defer closeDB.Close()

	resultEmpresa, count := GetFindOneEmpresa(identificador)

	if count == 0 {
		return models.EmpresasResponse{}, models.HTTPError400{Code: 400, Message: "La empresa no existe en el sistema"}, true
	}

	resultEmpresa.NombreEmpresa = empresa.NombreEmpresa
	if _, errUUID := uuid.Parse(empresa.IdOsde); errUUID == nil {
		resultEmpresa.IdOsde = empresa.IdOsde
	}

	db.Model(&models.Empresas{}).Where(models.Empresas{Identificador: identificador}).Update("nombre_empresa", empresa.NombreEmpresa)

	result, _ := GetFindOneEmpresa(identificador)

	return result, models.HTTPError400{}, false
}

func DeleteEmpresa(idempresa string) (models.EmpresasResponse, models.HTTPError400, bool) {
	var db, closeDB = database.GetConnection()
	defer closeDB.Close()

	result, count := GetFindOneEmpresa(idempresa)

	if count == 0 {
		return models.EmpresasResponse{}, models.HTTPError400{Code: 400, Message: "La empresa no existe en el sistema."}, true
	}

	db.Model(models.Empresas{}).Where("identificador = ?", idempresa).Update("activo", false)

	return result, models.HTTPError400{}, false
}

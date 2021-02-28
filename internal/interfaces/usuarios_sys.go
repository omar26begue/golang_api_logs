package interfaces

import (
	"github.com/google/uuid"
	"go-rest-api-logs/internal/database"
	"go-rest-api-logs/internal/models"
	"gopkg.in/go-playground/validator.v9"
)

func GetFindOneUsers(login string) (models.UsuariosSys, int64) {
	var db, closeDB = database.GetConnection()
	defer closeDB.Close()

	var result models.UsuariosSys
	var count int64
	db.Model(&models.UsuariosSys{}).Where(&models.UsuariosSys{Login: login}).Find(&result).Count(&count)

	return result, count
}

func GetFindOneUsersName(name string) (models.UsuariosSys, int64) {
	var db, closeDB = database.GetConnection()
	defer closeDB.Close()

	var result models.UsuariosSys
	var count int64
	db.Model(&models.UsuariosSys{}).Where(&models.UsuariosSys{Login: name}).Find(&result).Count(&count)

	return result, count
}

//CreateUsers Crea un usuario
func CreateUsers(users models.UsuariosSys) (models.UsuariosSys, models.HTTPError400, bool) {
	var db, closeDB = database.GetConnection()
	defer closeDB.Close()

	var validate = validator.New()
	err := validate.Struct(users)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return models.UsuariosSys{}, models.HTTPError400{Code: 400, Message: err.Tag()}, true
		}
	}

	// comprobando la existencia del nombre del usuario
	_, count := GetFindOneUsersName(users.Login)
	if count > 0 {
		return models.UsuariosSys{}, models.HTTPError400{Code: 400, Message: "El usuario ya esta registrado en la base de datos."}, true
	}

	// encriptando la contraseña
	/*bytes, err := bcrypt.GenerateFromPassword([]byte(users.Password), 14)
	if err != nil {
		return models.UsuariosSys{}, models.HTTPError400{Code: 400, Message: err.Error()}, true
	}*/

	users.Identificador = uuid.New().String()
	//users.Password = string(bytes)
	users.Password = users.Password
	db.Create(users)

	return users, models.HTTPError400{}, false
}

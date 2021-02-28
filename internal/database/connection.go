package database

import (
	"database/sql"
	"github.com/spf13/viper"
	"go-rest-api-logs/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//GetConnection Gestiona la conexion con la base de datos de la aplicación
func GetConnection() (*gorm.DB, *sql.DB) {
	dsn := "host=" + viper.GetString("DB_HOST") + " user=" + viper.GetString("DB_USER") + " password=" + viper.GetString("DB_PASSWORD") + " dbname=" + viper.GetString("DB_NAME") + " port=" + viper.GetString("DB_PORT") + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Error de conexion a la base de datos")
	}

	// migrando la base de datos
	if viper.GetBool("DB_INIT") == true {
		db.AutoMigrate(&models.Osde{}, &models.Empresas{}, &models.UsuariosSys{}, &models.Codes{})
	}

	sqlDB, err := db.DB()

	return db, sqlDB
}

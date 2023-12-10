package configs

import (
	"github.com/bshome19/programming-languages-with-db/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

type DatabaseConnection struct{}

var singleton *DatabaseConnection

func DatabaseConnectionInstance() *DatabaseConnection {
	if singleton == nil {
		singleton = &DatabaseConnection{}
	}
	return singleton
}

func (_ *DatabaseConnection) ConnectDB() {
	DB, err = gorm.Open(sqlite.Open("configs/programming-languages.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: "+ err.Error())
	}
	DB.AutoMigrate(&models.ProgrammingLanguage{})
}
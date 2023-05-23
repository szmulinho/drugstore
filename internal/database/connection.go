package database

import (
	"github.com/szmulinho/drugstore/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() *gorm.DB {

	connection, err := gorm.Open(postgres.Open("host=localhost user=postgres password=szmulinho dbname=szmul-med port=5432 sslmode=disable TimeZone=Europe/Warsaw"), &gorm.Config{})

	if err != nil {
		panic("can't connect with database")
	}

	DB = connection

	connection.AutoMigrate(&model.User{})
	connection.AutoMigrate(&model.Drug{})

	return connection
}
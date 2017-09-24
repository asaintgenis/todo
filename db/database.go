package db

import (
	"github.com/asaintgenis/todo/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDB() {
	database, err := gorm.Open("postgres", "postgres://postgres:postgres@127.0.0.1:5432/todo?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer database.Close()

	database.AutoMigrate(model.Todo{})
}

func GetDBConnection() *gorm.DB {
	database, err := gorm.Open("postgres", "postgres://postgres:postgres@127.0.0.1:5432/todo?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return database
}

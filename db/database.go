package db

import (
	"github.com/asaintgenis/todo/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // blank import of postgres dialects
)

// InitDB initialize the DB by Running AutoMigrate at the server launch
func InitDB() {
	database, err := gorm.Open("postgres", "postgres://postgres:postgres@127.0.0.1:5432/todo?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer database.Close()

	database.AutoMigrate(model.Todo{})
}

// GetDBConnection return a new gorm DB connection TODO: manage a connection pool
func GetDBConnection() *gorm.DB {
	database, err := gorm.Open("postgres", "postgres://postgres:postgres@127.0.0.1:5432/todo?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return database
}

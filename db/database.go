package db

import "github.com/jinzhu/gorm"

func GetDBConnection() *gorm.DB {
	db, err := gorm.Open("postgres", "postgres://postgres:postgres@127.0.0.1:5432/todo?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

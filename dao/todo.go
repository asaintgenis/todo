package dao

import (
	"github.com/asaintgenis/todo/model"
	db2 "github.com/asaintgenis/todo/db"
)

func GetTodos() ([]model.Todo, error) {
	todos := []model.Todo{}
	db := db2.GetDBConnection()
	defer db.Close()
	err := db.Find(&todos).Error
	return todos, err
}

func GetTodo(todoID uint) (*model.Todo, error) {
	todo := model.Todo{}
	db := db2.GetDBConnection()
	defer db.Close()
	err := db.Where("id = ?", todoID).First(&todo).Error
	return &todo,err
}

package dao

import (
	db2 "github.com/asaintgenis/todo/db"
	"github.com/asaintgenis/todo/model"
)

// GetTodos return all task in DB
func GetTodos() ([]model.Todo, error) {
	todos := []model.Todo{}
	db := db2.GetDBConnection()
	defer db.Close()
	err := db.Find(&todos).Error
	return todos, err
}

// GetTodo return the todoID task in DB
func GetTodo(todoID uint) (*model.Todo, error) {
	todo := model.Todo{}
	db := db2.GetDBConnection()
	defer db.Close()
	err := db.Where("id = ?", todoID).First(&todo).Error
	return &todo, err
}

// PostTodo create the passed todo in DB
func PostTodo(todo *model.Todo) error {
	todo.ID = 0
	db := db2.GetDBConnection()
	defer db.Close()
	return db.Create(todo).Error
}

// PutTodo update the passed todo in DB
func PutTodo(todo *model.Todo) (*model.Todo, error) {
	db := db2.GetDBConnection()
	defer db.Close()
	err := db.Save(todo).Error
	return todo, err
}

// DeleteTodo delete the todoID task in DB
func DeleteTodo(todoID uint) error {
	db := db2.GetDBConnection()
	defer db.Close()
	todo, err := GetTodo(todoID)
	if err != nil {
		return err
	}
	return db.Delete(&todo).Error
}

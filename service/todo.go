package service

import (
	"github.com/asaintgenis/todo/dao"
	"github.com/asaintgenis/todo/model"
)

// GetTodos return all todos
func GetTodos() ([]model.Todo, error) {
	todo, err := dao.GetTodos()
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// GetTodo return the todoID task
func GetTodo(todoID uint) (*model.Todo, error) {
	todo, err := dao.GetTodo(todoID)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// PostTodo create the passed todo
func PostTodo(todo *model.Todo) (*model.Todo, error) {
	err := dao.PostTodo(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// PutTodo update the passed todo
func PutTodo(todo *model.Todo) (*model.Todo, error) {
	todo, err := dao.PutTodo(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

// DeleteTodo delete the todoID task
func DeleteTodo(todoID uint) error {
	return dao.DeleteTodo(todoID)
}

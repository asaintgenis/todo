package service

import (
	"github.com/asaintgenis/todo/dao"
	"github.com/asaintgenis/todo/model"
)

func GetTodos() ([]model.Todo, error) {
	todo, err := dao.GetTodos()
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func GetTodo(todoID uint) (*model.Todo, error) {
	todo, err := dao.GetTodo(todoID)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func PostTodo(todo *model.Todo) (*model.Todo, error) {
	err := dao.PostTodo(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func PutTodo(todo *model.Todo) (*model.Todo, error) {
	todo, err := dao.PutTodo(todo)
	if err != nil {
		return nil, err
	}
	return todo, nil
}

func DeleteTodo(todoID uint) error {
	return dao.DeleteTodo(todoID)
}

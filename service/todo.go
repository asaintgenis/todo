package service

import (
	"github.com/asaintgenis/todo/model"
	"github.com/asaintgenis/todo/dao"
)

func GetTodos() ([]model.Todo, error) {
	todo, err := dao.GetTodos()
	if err != nil {
		return nil, err
	}
	return todo,nil
}

func GetTodo(todoID uint) (*model.Todo, error) {
	todo, err := dao.GetTodo(todoID)
	if err != nil {
		return nil, err
	}
	return todo,nil
}
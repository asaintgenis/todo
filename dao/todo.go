package dao

import "github.com/asaintgenis/todo/model"

func GetTodos() ([]model.Todo, error) {
	var todos := []model.Todo{}
	err := rs.DB().Where("id = ?", id).First(&bar).Error
	return &bar, err
}

func GetTodo(todoID uint) (*model.Todo, error) {
	todo, err := dao.getTodo(todoID)
	if err != nil {
		return nil, err
	}
	return todo,nil
}

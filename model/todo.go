package model

import "github.com/jinzhu/gorm"

// Todo represente a task
type Todo struct {
	gorm.Model
	Title     string `json:"title"`
	ShortText string `json:"shortText"`
	LongText  string `json:"longText"`
}

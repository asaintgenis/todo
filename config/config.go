package config

import "github.com/jinzhu/gorm"

type AppConfig struct {
	DB gorm.DB
	PORT string
}

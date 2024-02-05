package configs

import (
	"github.com/pamelaborges/goppotunities/schemas"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	db, err := NewDatabase()
	if err != nil {
		return err
	}
	db.AutoMigrate(&schemas.Opening{})
	return nil
}

func GetLogger() *Logger {
	return NewLogger()
}

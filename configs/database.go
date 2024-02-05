package configs

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("./db/db.sql"), &gorm.Config{})

	if err != nil {
		logger.Errorf("Erro ao criar database", err)
		return nil, err
	}
	return db, nil

}

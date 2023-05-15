package storage

import (
	"github.com/adzpm/jumoreski/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewStorage(cfg *Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(cfg.DatabasePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&models.StoragePost{}); err != nil {
		return nil, err
	}

	return db, nil
}

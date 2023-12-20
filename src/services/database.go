package services

import (
	"github.com/psycomentis/psycofolio++/src/models"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateDBInstance(config *Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.Database.ConnectionString), &gorm.Config{})
	if err != nil {
		log.Err(err)
	}
	return db, nil
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Language{})
}

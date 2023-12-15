package config

import (
	"log"

	"github.com/psycomentis/psycofolio++/src/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateDBInstance() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("psycofolio.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Language{})
}

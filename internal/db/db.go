package db

import (
	"fmt"
	"log"

	"github.com/santoadji21/santoadji21-go-fiber-chat/pkg/config"
	"github.com/santoadji21/santoadji21-go-fiber-chat/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg config.Config) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUsername, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// AutoMigrate
	err = DB.AutoMigrate(&models.Message{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	// Drop the unused column if it exists
	// Db migrations messages
	if DB.Migrator().HasColumn(&models.Message{}, "OldColumn") {
		err = DB.Migrator().DropColumn(&models.Message{}, "OldColumn")
		if err != nil {
			log.Fatalf("Failed to drop column: %v", err)
		}
	}

}

func GetDB() *gorm.DB {
	return DB
}

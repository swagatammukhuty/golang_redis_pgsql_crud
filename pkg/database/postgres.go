package database

import (
	"golang_redis_pgsql/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitPgDB() {
	var err error
	dsn := os.Getenv("POSTGRES_DSN")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}
	DB.AutoMigrate(&model.User{})
}

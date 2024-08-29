package database

import (
	"fmt"

	"github.com/ThailanTec/challenger/movies/domain"
	"github.com/ThailanTec/challenger/movies/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func PostgresClient(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUsername, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, domain.ErrDatabaseConnectionFailed
	}

	db.AutoMigrate(&domain.Movie{})
	return db, nil
}

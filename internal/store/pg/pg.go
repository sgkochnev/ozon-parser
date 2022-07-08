package pg

import (
	"fmt"
	"ozon-parser/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func Dial() (*DB, error) {
	cfg := config.Get()
	if cfg.PgDB == "" {
		return nil, nil
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PgAddr, cfg.PgUser, cfg.PgPassword, cfg.PgDB, cfg.PgPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

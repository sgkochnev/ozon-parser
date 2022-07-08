package app

import (
	"errors"
	"fmt"
	"log"
	"ozon-parser/config"
	"time"

	// migrate tools
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const (
	_defaultAttempts = 20
	_defaultTimeout  = time.Second
)

func init() {
	cfg := config.Get()

	if cfg.PgUser == "" || cfg.PgPassword == "" ||
		cfg.PgAddr == "" || cfg.PgPort == "" || cfg.PgDB == "" {
		log.Fatalf("migrate: environment variable not declared: PG_USER or PG_PASSWORD or PG_ADDR or PG_PORT or PG_DB")
	}

	// PG_URL=postgres://user:password@host:port/db
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.PgUser, cfg.PgPassword,
		cfg.PgAddr, cfg.PgPort,
		cfg.PgDB,
	)

	dbURL += "?sslmode=disable"
	fmt.Println(dbURL)

	var (
		attempts = _defaultAttempts
		err      error
		m        *migrate.Migrate
	)

	if cfg.PgMigrationsPath == "" {
		log.Fatalf("migrate: environment variable not declared: PG_MIGRATIONS_PATH")
	}

	for attempts > 0 {
		// m, err = migrate.New("file://migrations", dbURL)
		m, err = migrate.New(cfg.PgMigrationsPath, dbURL)
		if err == nil {
			break
		}

		log.Printf("Migrate: postgres is trying to connect, attempts left: %d", attempts)
		time.Sleep(_defaultTimeout)
		attempts--
	}

	if err != nil {
		log.Fatalf("Migrate: postgres connect error: %s", err)
	}

	err = m.Up()
	defer m.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("Migrate: up error: %s", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("Migrate: no change")
		return
	}

	log.Printf("Migrate: up success")
}

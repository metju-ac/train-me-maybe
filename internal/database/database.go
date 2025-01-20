package database

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/metju-ac/train-me-maybe/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectAndMigrate(config config.DbConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		config.Host, config.Port, config.User, config.Name, config.Password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		return nil, err
	}
	slog.Info("Connected to database", "db", db)

	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.User, config.Password, config.Host, config.Port, config.Name)
	m, err := migrate.New(
		"file://migrations",
		dbURL,
	)
	if err != nil {
		slog.Error("Failed to create migration instance", "error", err)
		return nil, err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		slog.Error("Failed to apply migrations", "error", err)
		return nil, err
	}
	slog.Info("Migrations applied")

	return db, nil
}

package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/monforje/user-service/pkg/config"
	"github.com/pressly/goose/v3"
)

type Postgres struct {
	DB *sqlx.DB
}

func New(cfg config.PostgresConfig) *Postgres {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.SSLMode,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	if err := applyMigrations(db.DB); err != nil {
		log.Fatal(err)
	}

	return &Postgres{
		DB: db,
	}
}

func applyMigrations(db *sql.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	migrationsDir := "migrations"

	currentVersion, err := goose.GetDBVersion(db)
	if err != nil {
		return err
	}

	log.Printf("postgres: current migration version: %d", currentVersion)

	if err := goose.Up(db, migrationsDir); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) Stop() {
	if err := p.DB.Close(); err != nil {
		log.Println(err)
	}
}

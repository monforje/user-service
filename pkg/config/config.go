package config

import (
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App      AppConfig      `yaml:"app"`
	Postgres PostgresConfig `yaml:"postgres"`
}

type AppConfig struct {
	Port string `yaml:"port"`
}

type PostgresConfig struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	DBName          string        `yaml:"dbname"`
	SSLMode         string        `yaml:"sslmode"`
	MaxOpenConns    int           `yaml:"max_open_conns"`
	MaxIdleConns    int           `yaml:"max_idle_conns"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime"`
}

func New() *Config {
	var cfg Config

	if err := cleanenv.ReadConfig("config.yaml", &cfg); err != nil {
		log.Fatal("config: failed to read config:", err)
	}

	log.Println("config: loaded successfully")
	return &cfg
}

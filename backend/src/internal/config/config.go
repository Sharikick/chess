package config

import (
	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	HTTP struct {
		Host string `env:"HTTP_HOST" env-default:"127.0.0.1"`
		Port string `env:"HTTP_PORT" env-default:"8080"`
	}
	Logger struct {
		Level string `env:"LOG_LEVEL" env-default:"INFO"`
	}
	Database struct {
		Host     string `env:"DB_HOST" env-default:"database"`
		Port     string `env:"DB_PORT" env-default:"5432"`
		User     string `env:"DB_USER" env-default:"tsyden"`
		Password string `env:"DB_PASSWORD" env-default:"chess"`
		Name     string `env:"DB_DATABASE" env-default:"chess_db"`
	}
	Config struct {
		Database Database
		Logger   Logger
		HTTP     HTTP
	}
)

var (
	cfg  *Config
	once sync.Once
	err  error
)

func (db *Database) GetConn() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db.User, db.Password, db.Host, db.Port, db.Name)
}

func Load() (*Config, error) {
	once.Do(func() {
		cfg = &Config{}
		err = cleanenv.ReadEnv(cfg)
	})
	return cfg, err
}

package config

import "fmt"

const (
	envHost     = "DB_HOST"
	envPort     = "DB_PORT"
	envUser     = "DB_USER"
	envPassword = "DB_PASSWORD"
	envName     = "DB_DATABASE"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func (db *DatabaseConfig) fill(getEnv func(key, defaultValue string) (string, error)) error {
	host, err := getEnv(envHost, "database")
	if err != nil {
		return err
	}
	db.Host = host

	port, err := getEnv(envPort, "5432")
	if err != nil {
		return err
	}
	db.Port = port

	user, err := getEnv(envUser, "tsyden")
	if err != nil {
		return err
	}
	db.User = user

	password, err := getEnv(envPassword, "chess")
	if err != nil {
		return err
	}
	db.Password = password

	name, err := getEnv(envName, "chess_db")
	if err != nil {
		return err
	}
	db.Name = name

	return nil
}

func (db DatabaseConfig) GetConn() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db.User, db.Password, db.Host, db.Port, db.Name)
}

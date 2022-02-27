package config

import "os"

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "postgres",
			Host:     os.Getenv("DB_HOST"),
			Port:     "5432",
			Username: os.Getenv("DB_USERNAME"),
			Name:     os.Getenv("DB_NAME"),
			Password: os.Getenv("DB_PASSWORD"),
		},
	}
}

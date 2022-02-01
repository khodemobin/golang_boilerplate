package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	Port string
	Env  string
}

type DB struct {
	Host     string
	Database string
	User     string
	Password string
	Port     string
}

type Rabbit struct {
	Host     string
	Port     string
	User     string
	Password string
	VHost    string
}

type Sentry struct {
	Dsn string
}

type Redis struct {
	Address  string
	Password string
	PoolSize string
	Database string
}

type Config struct {
	App    App
	DB     DB
	Rabbit Rabbit
	Sentry Sentry
	Redis  Redis
}

func New() *Config {
	readConfig()

	return &Config{
		App: App{
			Port: os.Getenv("APP_PORT"),
			Env:  os.Getenv("APP_ENV"),
		},
		DB: DB{
			Host:     os.Getenv("DB_HOST"),
			Database: os.Getenv("DB_DATABASE"),
			User:     os.Getenv("DB_USER"),
			Port:     os.Getenv("DB_PORT"),
			Password: os.Getenv("DB_PASSWORD"),
		},
		Rabbit: Rabbit{
			Host:     os.Getenv("RABBIT_HOST"),
			Port:     os.Getenv("RABBIT_PORT"),
			User:     os.Getenv("RABBIT_USER"),
			Password: os.Getenv("RABBIT_PASSWORD"),
			VHost:    os.Getenv("RABBIT_VHOST"),
		},
		Sentry: Sentry{
			Dsn: os.Getenv("SENTRY_DNS"),
		},
		Redis: Redis{
			Address:  os.Getenv("REDIS_ADDRESS"),
			Password: os.Getenv("REDIS_PASSWORD"),
			PoolSize: os.Getenv("REDIS_POOL_SIZE"),
			Database: os.Getenv("REDIS_DATABASE"),
		},
	}
}

func readConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("error load config", err)
	}
}

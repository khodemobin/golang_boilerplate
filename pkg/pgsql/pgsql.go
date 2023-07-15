package pgsql

import (
	"fmt"
	"log"

	"github.com/khodemobin/golang_boilerplate/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	l "gorm.io/gorm/logger"
)

type Pgsql struct {
	DB *gorm.DB
}

func New(cfg *config.Config) (*Pgsql, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  Dsn(cfg),
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		Logger: l.Default.LogMode(l.Silent),
	})
	if err != nil {
		return nil, err
	}

	return &Pgsql{
		DB: db,
	}, nil
}

func (m *Pgsql) Close() error {
	sqlDB, err := m.DB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	return sqlDB.Close()
}

func Dsn(c *config.Config) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Database)
}

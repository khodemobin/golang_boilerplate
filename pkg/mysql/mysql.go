package mysql

import (
	"fmt"
	"log"

	"github.com/khodemobin/golang_boilerplate/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	l "gorm.io/gorm/logger"
)

type Mysql struct {
	DB *gorm.DB
}

func New(cfg *config.Config) (*Mysql, error) {
	db, err := gorm.Open(mysql.Open(Dsn(cfg)), &gorm.Config{
		Logger: l.Default.LogMode(l.Silent),
	})
	if err != nil {
		return nil, err
	}

	return &Mysql{
		DB: db,
	}, nil
}

func (m *Mysql) Close() error {
	sqlDB, err := m.DB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	return sqlDB.Close()
}

func Dsn(c *config.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true&multiStatements=true", c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Database)
}

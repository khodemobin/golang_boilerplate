package mysql

import (
	"fmt"
	"log"

	"github.com/khodemobin/golang_boilerplate/internal/config"
	"github.com/khodemobin/golang_boilerplate/pkg/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	l "gorm.io/gorm/logger"
)

type Mysql struct {
	logger logger.Logger
	DB     *gorm.DB
}

func New(cfg *config.Config, logger logger.Logger) *Mysql {
	db, err := gorm.Open(mysql.Open(Dsn(cfg)), &gorm.Config{
		Logger: l.Default.LogMode(l.Silent),
	})
	if err != nil {
		logger.Fatal(err)
	}

	return &Mysql{
		logger: logger,
		DB:     db,
	}
}

func (m *Mysql) Close() {
	sqlDB, err := m.DB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	err = sqlDB.Close()
	if err != nil {
		m.logger.Fatal(err)
	}
}

func Dsn(c *config.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true&multiStatements=true", c.DB.User, c.DB.Password, c.DB.Host, c.DB.Port, c.DB.Database)
}

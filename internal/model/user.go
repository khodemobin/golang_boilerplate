package model

import (
	"time"

	"gorm.io/gorm"

	"github.com/bxcodec/faker/v3"
)

type User struct {
	ID        uint           `gorm:"primarykey" db:"id" faker:"-"`
	Email     *string        `db:"phone" faker:"phone_number"`
	Username  *string        `db:"phone" faker:"phone_number"`
	Password  *string        `db:"password" faker:"password" `
	CreatedAt time.Time      `faker:"-"`
	UpdatedAt time.Time      `faker:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" faker:"-"`
}

type UserResource struct {
	Phone string `json:"phone"`
	UUID  string `json:"uuid"`
}

func (u User) SeedUser() (*User, error) {
	user := &User{}
	err := faker.FakeData(user)
	return user, err
}

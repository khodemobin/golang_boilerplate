package model

import (
	"github.com/go-faker/faker/v4"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" mysql:"id" faker:"-"`
	Email     *string        `mysql:"email" faker:"email"`
	Username  *string        `mysql:"username" faker:"username"`
	Password  *string        `mysql:"password" faker:"password" `
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

package model

import (
	"time"

	"github.com/go-faker/faker/v4"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primarykey" column:"id" faker:"-"`
	Email     *string        `column:"email" faker:"email"`
	Username  *string        `column:"username" faker:"username"`
	Password  *string        `column:"password" faker:"password" `
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

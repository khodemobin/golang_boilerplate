package cmd

import (
	"github.com/khodemobin/golang_boilerplate/app"
	"github.com/khodemobin/golang_boilerplate/internal/model"
	"github.com/khodemobin/golang_boilerplate/pkg/encrypt"
	"github.com/spf13/cobra"
)

func SeedCommand() *cobra.Command {
	cmdSeed := &cobra.Command{
		Use:   "seed",
		Short: "Insert fake data to db",
		Run: func(cmd *cobra.Command, args []string) {
			pass, _ := encrypt.Hash("123456")
			email := "mobin@gmail.com"
			user, _ := model.User{}.SeedUser()
			user.Email = &email
			user.Password = &pass
			app.DB().Create(user)
		},
	}

	return cmdSeed
}

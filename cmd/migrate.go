package cmd

import (
	"database/sql"

	"github.com/khodemobin/golang_boilerplate/internal/app"
	"github.com/khodemobin/golang_boilerplate/pkg/pgsql"
	"github.com/pressly/goose/v3"
	"github.com/spf13/cobra"
)

func MigrateCommand() *cobra.Command {
	cmdMigrate := &cobra.Command{
		Use:   "migrate [ up & down & create]",
		Short: "Migrate database [ up & down & create]",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			sql, _ := sql.Open("postgres", pgsql.Dsn(app.Config()))
			dir := "migrations"
			err := goose.SetDialect("postgres")
			if err != nil {
				app.Log().Fatal(err)
			}
			if err := goose.Run(args[0], sql, dir, args[1:]...); err != nil {
				app.Log().Fatal(err)
			}
		},
	}

	return cmdMigrate
}

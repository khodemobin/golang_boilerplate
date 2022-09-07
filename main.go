package main

import (
	"github.com/khodemobin/golang_boilerplate/internal/app"
	"log"

	"github.com/khodemobin/golang_boilerplate/cmd"
	"github.com/spf13/cobra"
)

func main() {
	app.New()

	rootCmd := &cobra.Command{
		Use:                "app",
		DisableAutoGenTag:  true,
		DisableSuggestions: true,
		Run: func(c *cobra.Command, args []string) {
			cmd.Execute()
		},
	}
	rootCmd.AddCommand(cmd.ServeCommand(), cmd.MigrateCommand(), cmd.SeedCommand())
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

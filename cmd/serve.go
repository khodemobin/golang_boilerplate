package cmd

import (
	"errors"
	"fmt"
	"github.com/khodemobin/golang_boilerplate/internal/app"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/khodemobin/golang_boilerplate/internal/server"
	"github.com/khodemobin/golang_boilerplate/pkg/helper"

	"github.com/spf13/cobra"
)

func ServeCommand() *cobra.Command {
	cmdServe := &cobra.Command{
		Use:   "serve",
		Short: "Serve application",
		Run: func(cmd *cobra.Command, args []string) {
			Execute()
		},
	}
	return cmdServe
}

func Execute() {
	// start server
	restServer := server.New(helper.IsLocal())
	go func() {
		if err := restServer.Start(helper.IsLocal(), app.Config().App.Port); err != nil {
			msg := fmt.Sprintf("error happen while serving: %v", err)
			app.Log().Error(errors.New(msg))
			log.Println(msg)
		}
	}()

	// wait for close signal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	<-signalChan
	app.Log().Info("Received an interrupt, closing connections...")

	if err := restServer.Shutdown(); err != nil {
		app.Log().Info("Rest server doesn't shutdown in 10 seconds")
	}
}

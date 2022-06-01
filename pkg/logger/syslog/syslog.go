package syslog

import (
	"log"

	"github.com/khodemobin/golang_boilerplate/pkg/logger"
)

type syslog struct{}

func New() logger.Logger {
	return &syslog{}
}

func (s *syslog) Error(err error) {
	log.Println(err.Error())
}

func (s *syslog) Fatal(err error) {
	log.Println(err.Error())
}

func (s *syslog) Warn(msg string) {
	log.Println(msg)
}

func (s *syslog) Info(msg string) {
	log.Println(msg)
}

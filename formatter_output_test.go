package main

import (
	"log"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFormatting(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("hello world")
	logger.Warn("hello world warn")
	logger.Info("log info sending to file application")
	logger.Error("error test")
}

func TestOutput(t *testing.T) {
	logger := logrus.New()

	// initiation create file for logger
	file, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// set output file into logrus
	logger.SetOutput(file)

	logger.Info("hello world")
	logger.Warn("hello world warn")
	logger.Info("log info sending to file application")
	logger.Error("error test")
}

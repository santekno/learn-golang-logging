package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingletone(t *testing.T) {
	// set formatter using JSON
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// initiate log rus
	logrus.Info("info using Logrus singleton")
	logrus.Warn("warn using Logrus singleton")
	logrus.Error("error using Logrus singleton")
}

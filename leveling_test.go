package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLeveling(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("this is using trace level")
	logger.Debug("this is using debug level")
	logger.Info("this is using info level")
	logger.Warn("this is using warn level")
	logger.Error("this is using error level")
}

package main

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()

	logger.Println("Halo this is log using logrus")
	logger.Infoln("hello this is log using logrus")
	fmt.Println("hello this is using print log")
}

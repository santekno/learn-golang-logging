package log

import (
	"errors"
	"os"

	"github.com/santekno/learn-golang-logging/pkg/env"
	"github.com/sirupsen/logrus"
)

var (
	log, _        = NewLogger(&Config{Formatter: &TextFormatter, Level: InfoLevel, LogName: "application.log"})
	JSONFormatter logrus.JSONFormatter
	TextFormatter logrus.TextFormatter
	serviceFields = map[string]interface{}{
		env.EnvironmentName: env.ServiceEnv(),
		env.GoVersionName:   env.GetVersion(),
	}
)

type (
	Level  = logrus.Level
	Logger = *logrus.Logger
)

const (
	// override logrus level
	Paniclevel = logrus.PanicLevel
	FatalLevel = logrus.FatalLevel
	ErrorLevel = logrus.ErrorLevel
	WarnLevel  = logrus.WarnLevel
	InfoLevel  = logrus.InfoLevel
	DebugLevel = logrus.DebugLevel
	TraceLevel = logrus.TraceLevel
)

type Config struct {
	logrus.Formatter
	logrus.Level
	LogName string
}

func NewLogger(cfg *Config) (Logger, error) {
	l := logrus.New()
	if env.IsDevelopment() {
		l.SetFormatter(&logrus.TextFormatter{})
	}
	l.SetFormatter(cfg.Formatter)
	l.SetLevel(cfg.Level)
	return l, nil
}

func SetConfig(cfg *Config) error {
	if cfg.LogName == "" {
		return errors.New("log name is empty")
	}

	if !env.IsDevelopment() {
		// initiation create file for logger
		file, err := os.OpenFile(cfg.LogName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatal(err)
		}

		// set output file into logrus
		log.SetOutput(file)
	}

	log.SetFormatter(cfg.Formatter)
	log.SetLevel(cfg.Level)
	return nil
}

func Debug(args ...interface{}) {
	log.WithFields(serviceFields).Debug(args...)
}

func Info(args ...interface{}) {
	log.WithFields(serviceFields).Info(args...)
}

func Warn(args ...interface{}) {
	log.WithFields(serviceFields).Warn(args...)
}

func Error(args ...interface{}) {
	log.WithFields(serviceFields).Error(args...)
}

func Fatal(args ...interface{}) {
	log.WithFields(serviceFields).Fatal(args...)
}

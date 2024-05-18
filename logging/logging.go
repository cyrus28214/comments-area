package logging

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type LoggingConfig struct {
	Level  string `mapstructure:"level"`
	Fotmat string `mapstructure:"format"`
	Output string `mapstructure:"output"`
}

func GetLogger(config LoggingConfig) *logrus.Logger {
	logger := logrus.New()

	//level
	level, err := logrus.ParseLevel(config.Level)
	if err != nil {
		panic(fmt.Errorf("invalid log level: %s", config.Level))
	}
	logger.SetLevel(level)

	//format
	switch config.Fotmat {
	case "json":
		logger.SetFormatter(&logrus.JSONFormatter{})
	case "text":
		logger.SetFormatter(&logrus.TextFormatter{})
	default:
		panic(fmt.Errorf("invalid log format: %s", config.Fotmat))
	}

	//output
	file, err := os.OpenFile(config.Output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Errorf("error opening log file: %s", err))
	}
	mw := io.MultiWriter(os.Stdout, file)
	logger.SetOutput(mw)

	return logger
}

package logging

import "testing"

// package logging

// import (
// 	"fmt"
// 	"os"

// 	"github.com/sirupsen/logrus"
// )

// type LoggingConfig struct {
// 	Level  string `mapstructure:"level"`
// 	Fotmat string `mapstructure:"format"`
// 	Output string `mapstructure:"output"`
// }

// func GetLogger(config LoggingConfig) *logrus.Logger {
// 	logger := logrus.New()
// 	level, err := logrus.ParseLevel(config.Level)
// 	if err != nil {
// 		panic(fmt.Errorf("invalid log level: %s", config.Level))
// 	}
// 	logger.SetLevel(level)
// 	file, err := os.OpenFile(config.Output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
// 	if err != nil {
// 		panic(fmt.Errorf("error opening log file: %s", err))
// 	}
// 	logger.Out = file
// 	return logger
// }

func TestLogging(t *testing.T) {
	log := GetLogger(LoggingConfig{
		Level:  "debug",
		Fotmat: "text",
		Output: "./test.log",
	})
	log.Info("test log")
}

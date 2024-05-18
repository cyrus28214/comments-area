package logging

import "testing"

func TestLogging(t *testing.T) {
	log := GetLogger(LoggingConfig{
		Level:  "debug",
		Fotmat: "text",
		Output: "./test.log",
	})
	log.Info("test log")
}

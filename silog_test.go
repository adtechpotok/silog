package silog

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"testing"
)

func Test_BaseLogger_Log(t *testing.T) {
	var logger = log.New(os.Stdout, "\r\n", 0)

	func(logger BaseLogger) {
		logger.Print("Test base logger with log")
	}(logger)
}

func Test_BaseLogger_Logrus(t *testing.T) {
	var logger = logrus.New()

	func(logger BaseLogger) {
		logger.Print("Test base logger with logrus")
	}(logger)
}

func Test_Logger_Logrus(t *testing.T) {
	var logger = logrus.New()

	func(logger Logger) {
		logger.Info("Test normal logger with logrus")
	}(logger)
}

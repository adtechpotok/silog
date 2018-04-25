package logrus

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func Test_Logrus(t *testing.T) {
	var log = logrus.New()

	func(logger FieldLogger) {
		logger.WithField("foo", []string{"bar", "baz"}).Info("Test FieldLogger with logrus")
	}(log)
}

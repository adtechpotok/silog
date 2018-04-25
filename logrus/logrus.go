package logrus

import (
	"github.com/adtechpotok/silog"
	"github.com/sirupsen/logrus"
)

// FieldLogger is the interface used by logrus
type FieldLogger interface {
	silog.Logger

	WithField(key string, value interface{}) *logrus.Entry
	WithFields(fields logrus.Fields) *logrus.Entry
	WithError(err error) *logrus.Entry
}

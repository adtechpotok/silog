package silog

import "github.com/sirupsen/logrus"

type Fields map[string]interface{}

// MiniLogger is the interface used by Go's standard library's log package.
type MiniLogger interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})
	Panicln(args ...interface{})

	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
}

// StandardLogger is the interface used by us
type StandardLogger interface {
	MiniLogger

	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Errorln(args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Infoln(args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Warnln(args ...interface{})
}

// AdvancedLogger is the interface used by logrus
type AdvancedLogger interface {
	StandardLogger
	WithField(key string, value interface{}) *logrus.Entry
	WithFields(fields Fields) *logrus.Entry
}
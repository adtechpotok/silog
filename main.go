package silog

import (
	"github.com/sirupsen/logrus"
	"io"
)

// Logrus instance
var logger *logrus.Logger

func init() {
	logger = new(logrus.Logger)
}


func SetOut(writer io.Writer){
	logger.Out = writer
}

// Adds a field to the log entry, note that it doesn't log until you call
// Debug, Print, Info, Warn, Fatal or Panic. It only creates a log entry.
// If you want multiple fields, use `WithFields`.
func WithField(key string, value interface{}) *logrus.Entry {
	return logger.WithField(key, value)
}

// Adds a struct of fields to the log entry. All it does is call `WithField` for
// each `Field`.
func WithFields(fields logrus.Fields) *logrus.Entry {
	return logger.WithFields(fields)
}

// Add an error as single field to the log entry.  All it does is call
// `WithError` for the given `error`.
func WithError(err error) *logrus.Entry {
	return logger.WithError(err)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args)
}

func Printf(format string, args ...interface{}) {
	logger.Printf(format, args)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args)
}

func Warningf(format string, args ...interface{}) {
	logger.Warningf(format, args)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args)
}

func Debug(args ...interface{}) {
	logger.Debug(args)
}

func Info(args ...interface{}) {
	logger.Info(args)
}

func Print(args ...interface{}) {
	logger.Print(args)
}

func Warn(args ...interface{}) {
	logger.Warn(args)
}

func Warning(args ...interface{}) {
	logger.Warning(args)
}

func Error(args ...interface{}) {
	logger.Error(args)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args)
}

func Panic(args ...interface{}) {
	logger.Panic(args)
}

func Debugln(args ...interface{}) {
	logger.Debugln(args)
}

func Infoln(args ...interface{}) {
	logger.Infoln(args)
}

func Println(args ...interface{}) {
	logger.Println(args)
}

func Warnln(args ...interface{}) {
	logger.Warnln(args)
}

func Warningln(args ...interface{}) {
	logger.Warningln(args)
}

func Errorln(args ...interface{}) {
	logger.Errorln(args)
}

func Fatalln(args ...interface{}) {
	logger.Fatalln(args)
}

func Panicln(args ...interface{}) {
	logger.Panicln(args)
}

//When file is opened with appending mode, it's safe to
//write concurrently to a file (within 4k message on Linux).
//In these cases user can choose to disable the lock.
func SetNoLock() {
	logger.SetNoLock()
}

func SetLevel(level logrus.Level) {
	logger.SetLevel(level)
}

func AddHook(hook logrus.Hook) {
	logger.AddHook(hook)
}

package silog

import (
	"github.com/sirupsen/logrus"
	"io"
)

// Logrus instance
var logger *logrus.Logger

// Initialize logger
func init() {
	logger = logrus.New()
}

// GetInstance return our logger instance
func GetInstance() *logrus.Logger {
	return logger
}

// Set logger out
func SetOut(writer io.Writer) {
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

// Print debug format message
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args)
}

// Print info format message
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args)
}

// Print print format message
func Printf(format string, args ...interface{}) {
	logger.Printf(format, args)
}

// Print warn format message
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args)
}

// Print warning format message
func Warningf(format string, args ...interface{}) {
	logger.Warningf(format, args)
}

// Print error format message
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args)
}

// Print fatal format message
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args)
}

// Print panic format message
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args)
}

// Print debug message
func Debug(args ...interface{}) {
	logger.Debug(args)
}

// Print info message
func Info(args ...interface{}) {
	logger.Info(args)
}

//Print print message
func Print(args ...interface{}) {
	logger.Print(args)
}

//Print warn message
func Warn(args ...interface{}) {
	logger.Warn(args)
}

//Print warning message
func Warning(args ...interface{}) {
	logger.Warning(args)
}

// Print error message
func Error(args ...interface{}) {
	logger.Error(args)
}

// Print Fatal message
func Fatal(args ...interface{}) {
	logger.Fatal(args)
}

// Print panic message
func Panic(args ...interface{}) {
	logger.Panic(args)
}

// Print sprintlnn debug message
func Debugln(args ...interface{}) {
	logger.Debugln(args)
}

// Print info sprintlnn message
func Infoln(args ...interface{}) {
	logger.Infoln(args)
}

// Print print sprintlnn message
func Println(args ...interface{}) {
	logger.Println(args)
}

// Print warn sprintlnn message
func Warnln(args ...interface{}) {
	logger.Warnln(args)
}

// Print warning sprintlnn message
func Warningln(args ...interface{}) {
	logger.Warningln(args)
}

// Print error sprintlnn message
func Errorln(args ...interface{}) {
	logger.Errorln(args)
}

// Print fatal sprintlnn message
func Fatalln(args ...interface{}) {
	logger.Fatalln(args)
}

// Print panic sprintlnn message
func Panicln(args ...interface{}) {
	logger.Panicln(args)
}

//When file is opened with appending mode, it's safe to
//write concurrently to a file (within 4k message on Linux).
//In these cases user can choose to disable the lock.
func SetNoLock() {
	logger.SetNoLock()
}

// Set logger level
func SetLevel(level logrus.Level) {
	logger.SetLevel(level)
}

// Add hook to logger
func AddHook(hook logrus.Hook) {
	logger.AddHook(hook)
}

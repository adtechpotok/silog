package silog

// BaseLogger is the interface used by Go's standard library's log package.
type BaseLogger interface {
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

// Logger is the interface used by us
type Logger interface {
	BaseLogger

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

package loggerPkg

type ILogger interface {
	InitLogger()
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	DPanic(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

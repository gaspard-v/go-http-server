package log

type LogConsumerInterface interface {
	Panic(error)
	Fatal(error)
	Error(error)
	Warning(error)
	Notice(error)
	Message(error)
	Debug(error)

	SetModuleName(string)
}

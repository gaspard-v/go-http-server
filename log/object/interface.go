package object

type LogInterface interface {
	Panic(any)
	Fatal(any)
	Error(any)
	Warning(any)
	Notice(any)
	Message(any)
	Debug(any)

	SetModuleName(string)
}

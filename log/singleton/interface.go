package singleon

type LogInterface interface {
	Panic(string, any)
	Fatal(string, any)
	Error(string, any)
	Warning(string, any)
	Notice(string, any)
	Message(string, any)
	Debug(string, any)
}

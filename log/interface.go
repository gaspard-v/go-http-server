package log

type LogConsumerInterface interface {
	OnPanic(error)
	OnFatal(error)
	OnError(error)
	OnWarning(error)
	OnNotice(error)
	OnMessage(error)
	OnDebug(error)

	SetModuleName(string)
}

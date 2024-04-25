package log

import (
	"fmt"
	"log"
)

type ConsoleLog struct {
	moduleName string
}

// OnPanic(error)
// OnFatal(error)
// OnError(error)
// OnWarning(error)
// OnNotice(error)
// OnMessage(error)
// OnDebug(error)
// SetModuleName(string)

func CreateConsoleLog(moduleName string) *ConsoleLog {
	return &ConsoleLog{moduleName}
}

func (logger *ConsoleLog) setLogPrefix(prefix string) {
	fullPrefix := fmt.Sprintf("%s | %s ", logger.moduleName, prefix)
	log.SetPrefix(fullPrefix)
}

func (logger *ConsoleLog) SetModuleName(moduleName string) {
	logger.moduleName = moduleName
}

func (logger *ConsoleLog) OnPanic(err error) {
	logger.setLogPrefix("PANIC")
	log.Panicln(err)
}

func (logger *ConsoleLog) OnFatal(err error) {
	logger.setLogPrefix("FATAL")
	log.Fatalln(err)
}

func (logger *ConsoleLog) OnError(err error) {
	logger.setLogPrefix("ERROR")
	log.Println(err)
}

func (logger *ConsoleLog) OnWarning(err error) {
	logger.setLogPrefix("WARNING")
	log.Println(err)
}

func (logger *ConsoleLog) OnNotice(err error) {
	logger.setLogPrefix("NOTICE")
	log.Println(err)
}

func (logger *ConsoleLog) OnMessage(err error) {
	logger.setLogPrefix("MESSAGE")
	log.Println(err)
}

func (logger *ConsoleLog) OnDebug(err error) {
	logger.setLogPrefix("DEBUG")
	log.Println(err)
}

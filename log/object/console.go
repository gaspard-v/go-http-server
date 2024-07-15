package logObject

import (
	"fmt"
	"log"
)

type ConsoleLog struct {
	moduleName string
}

func CreateConsole(moduleName string) LogInterface {
	return ConsoleLog{moduleName}
}

func (logger ConsoleLog) setLogPrefix(prefix string) {
	fullPrefix := fmt.Sprintf("[%s] module %s", prefix, logger.moduleName)
	log.SetPrefix(fullPrefix)
}

func (logger ConsoleLog) SetModuleName(moduleName string) {
	(&logger).moduleName = moduleName
}

func (logger ConsoleLog) Panic(msg any) {
	logger.setLogPrefix("PANIC")
	log.Panicln(msg)
}

func (logger ConsoleLog) Fatal(msg any) {
	logger.setLogPrefix("FATAL")
	log.Fatalln(msg)
}

func (logger ConsoleLog) Error(msg any) {
	logger.setLogPrefix("ERROR")
	log.Println(msg)
}

func (logger ConsoleLog) Warning(msg any) {
	logger.setLogPrefix("WARNING")
	log.Println(msg)
}

func (logger ConsoleLog) Notice(msg any) {
	logger.setLogPrefix("NOTICE")
	log.Println(msg)
}

func (logger ConsoleLog) Message(msg any) {
	logger.setLogPrefix("MESSAGE")
	log.Println(msg)
}

func (logger ConsoleLog) Debug(msg any) {
	logger.setLogPrefix("DEBUG")
	log.Println(msg)
}

package log

import (
	"fmt"
	"log"
)

type ConsoleLog struct {
	moduleName string
}

// Panic(error)
// Fatal(error)
// Error(error)
// Warning(error)
// Notice(error)
// Message(error)
// Debug(error)
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

func (logger *ConsoleLog) Panic(err error) {
	logger.setLogPrefix("PANIC")
	log.Panicln(err)
}

func (logger *ConsoleLog) Fatal(err error) {
	logger.setLogPrefix("FATAL")
	log.Fatalln(err)
}

func (logger *ConsoleLog) Error(err error) {
	logger.setLogPrefix("ERROR")
	log.Println(err)
}

func (logger *ConsoleLog) Warning(err error) {
	logger.setLogPrefix("WARNING")
	log.Println(err)
}

func (logger *ConsoleLog) Notice(err error) {
	logger.setLogPrefix("NOTICE")
	log.Println(err)
}

func (logger *ConsoleLog) Message(err error) {
	logger.setLogPrefix("MESSAGE")
	log.Println(err)
}

func (logger *ConsoleLog) Debug(err error) {
	logger.setLogPrefix("DEBUG")
	log.Println(err)
}

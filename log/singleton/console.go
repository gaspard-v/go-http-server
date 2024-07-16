package singleon

import (
	"fmt"
	"log"
	"sync"
)

var lock = &sync.Mutex{}

type ConsoleLog struct{}

var consoleLogInstance *ConsoleLog

func GetInstance() LogInterface {
	if consoleLogInstance != nil {
		return consoleLogInstance
	}
	lock.Lock()
	defer lock.Unlock()
	consoleLogInstance = &ConsoleLog{}
	return consoleLogInstance
}

func (logger *ConsoleLog) setLogPrefix(prefix string, module *string) {
	fullPrefix := fmt.Sprintf("[%s] from module [%s]} ", prefix, *module)
	log.SetPrefix(fullPrefix)
}

func (logger *ConsoleLog) Panic(module string, msg any) {
	logger.setLogPrefix("PANIC", &module)
	log.Panicln(msg)
}

func (logger *ConsoleLog) Fatal(module string, msg any) {
	logger.setLogPrefix("FATAL", &module)
	log.Fatalln(msg)
}

func (logger *ConsoleLog) Error(module string, msg any) {
	logger.setLogPrefix("ERROR", &module)
	log.Println(msg)
}

func (logger *ConsoleLog) Warning(module string, msg any) {
	logger.setLogPrefix("WARNING", &module)
	log.Println(msg)
}

func (logger *ConsoleLog) Notice(module string, msg any) {
	logger.setLogPrefix("NOTICE", &module)
	log.Println(msg)
}

func (logger *ConsoleLog) Message(module string, msg any) {
	logger.setLogPrefix("MESSAGE", &module)
	log.Println(msg)
}

func (logger *ConsoleLog) Debug(module string, msg any) {
	logger.setLogPrefix("DEBUG", &module)
	log.Println(msg)
}

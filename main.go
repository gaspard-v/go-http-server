package main

import (
	"sync"

	"github.com/gaspard-v/go-http-server/log"
	rawServer "github.com/gaspard-v/go-http-server/raw"
	rawClient "github.com/gaspard-v/go-http-server/raw/client"
	tcpServer "github.com/gaspard-v/go-http-server/tcp"
	tcpClient "github.com/gaspard-v/go-http-server/tcp/client"
)

func startServer(logger *log.ConsoleLog, wg *sync.WaitGroup, isReady chan bool) {
	defer wg.Done()
	logger.Message("starting server")
	rawConsumer := rawServer.CreateRaw(logger)
	tcp := tcpServer.CreateDefault(rawConsumer, logger)
	isReady <- true
	tcp.Accept()
}

func startClient(logger *log.ConsoleLog, wg *sync.WaitGroup) {
	defer wg.Done()
	client := rawClient.CreateRaw(logger)
	tcp := tcpClient.CreateDefault(client, logger)
	tcp.Connect()
}

func main() {
	logger := log.CreateConsoleLog("main")
	var wg sync.WaitGroup
	isReady := make(chan bool)
	wg.Add(1)
	go startServer(logger, &wg, isReady)
	if <-isReady {
		wg.Add(1)
		go startClient(logger, &wg)
	}
	wg.Wait()
}

package main

import (
	"errors"

	"github.com/gaspard-v/go-http-server/log"
	rawServer "github.com/gaspard-v/go-http-server/raw"
	rawClient "github.com/gaspard-v/go-http-server/raw/client"
	tcpServer "github.com/gaspard-v/go-http-server/tcp"
	tcpClient "github.com/gaspard-v/go-http-server/tcp/client"
)

func startServer(logger *log.ConsoleLog, onReady func(*log.ConsoleLog)) {
	logger.Message(errors.New("starting server"))
	rawConsumer := rawServer.CreateRaw(logger)
	tcp := tcpServer.CreateDefault(rawConsumer, logger)
	tcp.Accept()
	go onReady(logger)
}

func startClient(logger *log.ConsoleLog) {
	client := rawClient.CreateRaw(logger)
	tcp := tcpClient.CreateDefault(client, logger)
	tcp.Connect()
}

func main() {
	logger := log.CreateConsoleLog("main")
	go startServer(logger, startClient)
}

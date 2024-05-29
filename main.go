package main

import (
	"errors"

	"github.com/gaspard-v/go-http-server/log"
	"github.com/gaspard-v/go-http-server/raw"
	"github.com/gaspard-v/go-http-server/tcp"
)

func main() {
	logger := log.CreateConsoleLog("main")
	logger.Message(errors.New("starting server"))
	rawConsumer := raw.CreateRaw(logger)
	tcp := tcp.CreateDefault(rawConsumer, logger)
	tcp.Accept()
}

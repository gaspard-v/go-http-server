package main

import (
	"sync"
	"time"

	"github.com/gaspard-v/go-http-server/raw"
	tcpClient "github.com/gaspard-v/go-http-server/tcp/client"
	tcpServer "github.com/gaspard-v/go-http-server/tcp/server"
)

func startServer(wg *sync.WaitGroup) {
	defer wg.Done()
	r := raw.CreateRawServer()
	t := tcpServer.Default(r)
	t.Accept()
}

func startClient(wg *sync.WaitGroup) {
	defer wg.Done()
	c := raw.CreateRawClient()
	t := tcpClient.Default(c)
	t.Connect().Wait()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go startServer(&wg)
	time.Sleep(10 * time.Second)
	go startClient(&wg)
	wg.Wait()
}

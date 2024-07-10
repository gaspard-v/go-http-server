package tcp

import (
	"fmt"
	"net"
	"sync"

	"github.com/gaspard-v/go-http-server/log"
)

const DEFAULT_ADDRESS string = "127.0.0.1"
const DEFAULT_PORT string = "8080"

type Tcp struct {
	Listener    *net.TCPListener
	tcpConsumer TcpConsumer
	logger      log.LogConsumerInterface
	stopLoop    chan bool
}

func CreateDefault(
	tcpConsumer TcpConsumer,
	logger log.LogConsumerInterface,
) *Tcp {
	address := fmt.Sprintf("%s:%s", DEFAULT_ADDRESS, DEFAULT_PORT)
	stopLoop := make(chan bool)
	return Create(
		address,
		tcpConsumer,
		logger,
		stopLoop,
	)
}

func Create(
	address string,
	tcpConsumer TcpConsumer,
	logger log.LogConsumerInterface,
	stopLoop chan bool,
) *Tcp {
	tcpAddr, error := net.ResolveTCPAddr("tcp", address)
	if error != nil {
		logger.Fatal(error)
	}
	listener, error := net.ListenTCP(tcpAddr.Network(), tcpAddr)
	if error != nil {
		logger.Fatal(error)
	}
	tcp := Tcp{listener, tcpConsumer, logger, stopLoop}
	return &tcp
}

func (tcp *Tcp) Accept() {
	var wg sync.WaitGroup
Exit:
	for {
		select {
		case <-tcp.stopLoop:
			break Exit
		default:
		}
		conn, error := tcp.Listener.AcceptTCP()
		if error != nil {
			tcp.logger.Debug(error)
			continue
		}

		tcp.logger.Debug("Server got a client " + conn.RemoteAddr().String() + ", calling OnAccept callback")
		wg.Add(1)
		go tcp.tcpConsumer.OnAccept(conn, &wg)
	}
	wg.Wait()
}

func (tcp *Tcp) Stop() {
	tcp.stopLoop <- true
}

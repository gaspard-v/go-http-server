package tcp

import (
	"fmt"
	"net"

	"github.com/gaspard-v/go-http-server/log"
)

const DEFAULT_ADDRESS string = "127.0.0.1"
const DEFAULT_PORT string = "8080"

type Tcp struct {
	Listener    net.TCPListener
	tcpConsumer TcpConsumer
	logger      log.LogConsumerInterface
}

func CreateDefault(
	tcpConsumer TcpConsumer,
	logger log.LogConsumerInterface,
) *Tcp {
	address := fmt.Sprintf("%s:%s", DEFAULT_ADDRESS, DEFAULT_PORT)
	return Create(
		address,
		tcpConsumer,
		logger,
	)
}

func Create(
	address string,
	tcpConsumer TcpConsumer,
	logger log.LogConsumerInterface,
) *Tcp {
	tcpAddr, error := net.ResolveTCPAddr("tcp", address)
	if error != nil {
		logger.Fatal(error)
	}
	listener, error := net.ListenTCP(tcpAddr.Network(), tcpAddr)
	if error != nil {
		logger.Fatal(error)
	}
	tcp := Tcp{*listener, tcpConsumer, logger}
	return &tcp
}

func (tcp *Tcp) Accept() {
	for {
		conn, error := tcp.Listener.AcceptTCP()
		if error != nil {
			tcp.logger.Debug(error)
			continue
		}

		go tcp.tcpConsumer.OnAccept(conn)
	}
}

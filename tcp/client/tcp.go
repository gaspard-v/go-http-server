package client

import (
	"fmt"
	"net"

	"github.com/gaspard-v/go-http-server/log"
)

const DEFAULT_ADDRESS string = "127.0.0.1"
const DEFAULT_PORT string = "8080"

type TcpClient struct {
	tcpAddr           *net.TCPAddr
	tcpClientConsumer TcpClientConsumer
	logger            log.LogConsumerInterface
}

func CreateDefault(
	tcpClientConsumer TcpClientConsumer,
	logger log.LogConsumerInterface,
) *TcpClient {
	address := fmt.Sprintf("%s:%s", DEFAULT_ADDRESS, DEFAULT_PORT)
	return Create(
		address,
		tcpClientConsumer,
		logger,
	)
}

func Create(
	address string,
	tcpConsumer TcpClientConsumer,
	logger log.LogConsumerInterface,
) *TcpClient {
	tcpAddr, error := net.ResolveTCPAddr("tcp", address)
	if error != nil {
		logger.Fatal(error)
	}
	tcpClient := TcpClient{tcpAddr, tcpConsumer, logger}
	return &tcpClient
}

func (tcpClient *TcpClient) Connect() {
	conn, error := net.DialTCP("tcp", nil, tcpClient.tcpAddr)
	if error != nil {
		tcpClient.logger.Fatal(error)
	}
	go tcpClient.tcpClientConsumer.OnConnected(conn)
}

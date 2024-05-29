package client

import (
	"net"

	"github.com/gaspard-v/go-http-server/log"
)

type TcpClient struct {
	conn              *net.TCPConn
	tcpClientConsumer TcpClientConsumer
	logger            log.LogConsumerInterface
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
	conn, error := net.DialTCP("tcp", nil, tcpAddr)
	if error != nil {
		logger.Fatal(error)
	}
	tcpClient := TcpClient{conn, tcpConsumer, logger}
	return &tcpClient
}

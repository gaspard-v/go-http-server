package tcpClient

import (
	"fmt"
	"net"

	log "github.com/gaspard-v/go-http-server/log/object"
)

const DEFAULT_ADDRESS string = "127.0.0.1"
const DEFAULT_PORT string = "8080"

type TcpClient struct {
	tcpAddr       *net.TCPAddr
	tcpClientSock TcpClientSock
	log           log.LogInterface
}

func Default(
	tcpClient TcpClientSock,
) *TcpClient {
	address := fmt.Sprintf("%s:%s", DEFAULT_ADDRESS, DEFAULT_PORT)
	return Create(
		address,
		tcpClient,
	)
}

func Create(
	address string,
	tcpConsumer TcpClientSock,
) *TcpClient {
	tcpAddr, error := net.ResolveTCPAddr("tcp", address)
	log := log.Get("tcp.client")
	if error != nil {
		log.Fatal(error)
	}
	tcpClient := TcpClient{tcpAddr, tcpConsumer, log}
	return &tcpClient
}

func (tcpClient *TcpClient) Connect() {
	conn, error := net.DialTCP("tcp", nil, tcpClient.tcpAddr)
	if error != nil {
		tcpClient.log.Fatal(error)
	}
	go tcpClient.tcpClientSock.OnConnected(conn)

	//TODO wait for goroutine!!
}

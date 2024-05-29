package client

import (
	"fmt"
	"net"

	"github.com/gaspard-v/go-http-server/log"
	"github.com/gaspard-v/go-http-server/raw"
)

type RawClientConsumer struct {
	logger log.LogConsumerInterface
}

func CreateRaw(logger log.LogConsumerInterface) *RawClientConsumer {
	return &RawClientConsumer{logger}
}

func (rawClient *RawClientConsumer) readBody(conn *net.TCPConn) uint64 {
	bodySize, error := raw.GetBodySize(conn)
	if error != nil {
		rawClient.logger.Fatal(error)
	}
	data := make([]byte, bodySize)
	readSize, error := conn.Read(data)
	if error != nil {
		rawClient.logger.Fatal(error)
	}
	fmt.Println(data)
	return uint64(readSize)
}

func (rawClient *RawClientConsumer) OnConnected(conn *net.TCPConn) {
	rawClient.readBody(conn)
}

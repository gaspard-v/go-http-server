package raw

import (
	"fmt"
	"net"

	log "github.com/gaspard-v/go-http-server/log/object"
)

type RawClient struct {
	log log.LogInterface
}

func CreateRawClient() *RawClient {
	log := log.Get("raw.client")
	return &RawClient{log}
}

func (rawClient *RawClient) readBody(conn *net.TCPConn) uint64 {
	rawClient.log.Debug("Client is waiting data...")
	bodySize, error := GetBodySize(conn)
	// bodySize := 1
	if error != nil {
		rawClient.log.Fatal(error)
	}
	data := make([]byte, bodySize)
	readSize, error := conn.Read(data)
	if error != nil {
		rawClient.log.Fatal(error)
	}
	fmt.Println(data)
	return uint64(readSize)
}

func (rawClient *RawClient) OnConnected(conn *net.TCPConn) {
	rawClient.readBody(conn)
}

package raw

import (
	"encoding/binary"
	"fmt"
	"net"
	"sync"

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
	bodySize, error := getBodySize(conn)
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

func (rawClient *RawClient) OnConnected(conn *net.TCPConn, wg *sync.WaitGroup) {
	defer wg.Done()
	// rawClient.readBody(conn)
	b := []byte("ABC€")
	h := make([]byte, 8)
	binary.BigEndian.PutUint64(h, uint64(len(b)))
	conn.Write(h)
	conn.Close()
}

package raw

import (
	"encoding/binary"
	"fmt"
	"net"
	"sync"

	log "github.com/gaspard-v/go-http-server/log/object"
	tcpServer "github.com/gaspard-v/go-http-server/tcp/server"
)

const HEADER_SIZE = 4
const CHUNK_SIZE = 4096

type RawSock struct {
	log     log.LogInterface
	tcpConn tcpServer.TcpConnInterface
}

type RawConn struct {
	log  log.LogInterface
	conn *net.TCPConn
}

func CreateRawServer() *RawSock {
	log := log.Get("raw.server")
	return &RawSock{log, nil}
}

func CreateRawConn(conn *net.TCPConn) *RawConn {
	log := log.Get("raw.server.conn")
	return &RawConn{log, conn}
}

func consumeChunk(chunk *[]byte) {
	fmt.Print(chunk)
}

func (raw *RawConn) splitInChunk(
	conn *net.TCPConn,
	body_size uint64,
) uint64 {
	chunk := make([]byte, CHUNK_SIZE)
	bytes_read := uint64(0)
	for bytes_read <= body_size {
		remaining := body_size - bytes_read
		if remaining < CHUNK_SIZE {
			chunk = make([]byte, remaining)
		}
		err := binary.Read(conn, binary.BigEndian, chunk)
		if err != nil {
			raw.log.Fatal(err)
		}
		bytes_read += (uint64(len(chunk)))
		consumeChunk(&chunk)
	}
	return bytes_read
}

func (raw *RawConn) OnReceive() uint64 {
	body_size, err := GetBodySize(raw.conn)
	if err != nil {
		raw.log.Fatal(err)
	}
	return raw.splitInChunk(raw.conn, body_size)
}

func (raw *RawConn) OnSend() uint64 {
	return 0
}

func (raw *RawSock) OnAccept(conn *net.TCPConn, wg *sync.WaitGroup) {
	defer wg.Done()
	raw.tcpConn = CreateRawConn(conn)
	raw.tcpConn.OnReceive()
	conn.Close()
}

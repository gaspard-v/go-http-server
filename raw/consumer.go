package raw

import (
	"encoding/binary"
	"fmt"
	"net"
	"sync"

	"github.com/gaspard-v/go-http-server/log"
	"github.com/gaspard-v/go-http-server/tcp"
)

const HEADER_SIZE = 4
const CHUNK_SIZE = 4096

type RawConsumer struct {
	logger          log.LogConsumerInterface
	tcpConnConsumer tcp.TcpConnConsumer
}

type RawConnConsumer struct {
	logger log.LogConsumerInterface
	conn   *net.TCPConn
}

func CreateRaw(logger log.LogConsumerInterface) *RawConsumer {
	return &RawConsumer{logger, nil}
}

func CreateRawConn(logger log.LogConsumerInterface, conn *net.TCPConn) *RawConnConsumer {
	return &RawConnConsumer{logger, conn}
}

func GetBodySize(conn *net.TCPConn) (uint64, error) {
	var body_size uint64 = 0
	err := binary.Read(conn, binary.BigEndian, body_size)
	if err != nil {
		return 0, err
	}
	return body_size, nil
}

func consumeChunk(chunk *[]byte) {
	fmt.Print(chunk)
}

func (raw *RawConnConsumer) splitInChunk(
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
			raw.logger.Fatal(err)
		}
		bytes_read += (uint64(len(chunk)))
		consumeChunk(&chunk)
	}
	return bytes_read
}

func (raw *RawConnConsumer) OnReceive() uint64 {
	body_size, err := GetBodySize(raw.conn)
	if err != nil {
		raw.logger.Fatal(err)
	}
	return raw.splitInChunk(raw.conn, body_size)
}

func (raw *RawConnConsumer) OnSend() uint64 {
	return 0
}

func (raw *RawConsumer) OnAccept(conn *net.TCPConn, wg *sync.WaitGroup) {
	defer wg.Done()
	raw.tcpConnConsumer = CreateRawConn(raw.logger, conn)
	raw.tcpConnConsumer.OnReceive()
	conn.Close()
}

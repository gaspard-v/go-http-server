package raw

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/gaspard-v/go-http-server/log"
)

const HEADER_SIZE = 4
const CHUNK_SIZE = 4096

type RawConsumer struct {
	logger log.LogConsumerInterface
}

func CreateRaw(logger log.LogConsumerInterface) *RawConsumer {
	return &RawConsumer{logger}
}

func getBodySize(conn *net.TCPConn) (uint64, error) {
	var body_size uint64 = 0
	err := binary.Read(conn, binary.BigEndian, body_size)
	if err != nil {
		return 0, err
	}
	return body_size, nil
}

func onChunck(chunk *[]byte) {
	fmt.Println(chunk)
}

func (raw *RawConsumer) OnAccept(conn *net.TCPConn) uint64 {
	defer conn.Close()
	body_size, err := getBodySize(conn)
	if err != nil {
		raw.logger.OnFatal(err)
	}
	chunk := make([]byte, CHUNK_SIZE)
	bytes_read := uint64(0)
	for bytes_read <= body_size {
		remaining := body_size - bytes_read
		if remaining < CHUNK_SIZE {
			chunk = make([]byte, remaining)
		}
		err = binary.Read(conn, binary.BigEndian, chunk)
		if err != nil {
			raw.logger.OnFatal(err)
		}
		bytes_read += (uint64(len(chunk)))
		onChunck(&chunk)
	}
	return bytes_read
}

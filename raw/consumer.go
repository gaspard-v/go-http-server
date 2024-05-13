package raw

import (
	"errors"
	"fmt"
	"net"

	"github.com/gaspard-v/go-http-server/conv"
	"github.com/gaspard-v/go-http-server/log"
)

type RawConsumer struct {
	logger log.LogConsumerInterface
}

func CreateRaw(logger log.LogConsumerInterface) *RawConsumer {
	return &RawConsumer{logger}
}

func (raw *RawConsumer) OnAccept(conn *net.TCPConn) {
	defer conn.Close()
	header_size := 4
	header_buf := make([]byte, header_size)
	n, err := conn.Read(header_buf)
	if err != nil {
		raw.logger.OnFatal(err)
	}
	if header_size != n {
		raw.logger.OnFatal(errors.New("incorrect read size"))
	}
	adapter := conv.CreateBytesArrayAdapter(&header_buf)
	body_size := adapter.ToUint64()
	buf := make([]byte, body_size)
	conn.Read(buf)
	fmt.Println(buf)
}

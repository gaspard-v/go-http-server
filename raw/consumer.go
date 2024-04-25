package raw

import (
	"errors"
	"fmt"
	"net"

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
	headerSize := 4
	buf := make([]byte, headerSize)
	n, err := conn.Read(buf)
	if err != nil {
		raw.logger.OnFatal(err)
	}
	if headerSize != n {
		raw.logger.OnFatal(errors.New("incorrect read size"))
	}
	fmt.Println(buf)
}

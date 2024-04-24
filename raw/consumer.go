package raw

import "net"

type RawConsumer struct{}

func Create() *RawConsumer {
	return &RawConsumer{}
}

func (raw *RawConsumer) OnAccept(conn *net.TCPConn) {
	defer conn.Close()
	buf := make([]byte, 8192)
	conn.Read(buf)
}

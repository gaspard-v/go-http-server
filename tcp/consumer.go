package tcp

import "net"

type TcpConsumer interface {
	OnAccept(*net.TCPConn)
}

type TcpConnConsumer interface {
	OnReceive() uint64
	OnSend() uint64
}

package tcp

import "net"

type TcpConsumer interface {
	OnAccept(*net.TCPConn) uint64
}

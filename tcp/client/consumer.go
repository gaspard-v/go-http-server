package client

import "net"

type TcpClientConsumer interface {
	Send(*net.TCPConn) uint64
	OnReceive(*net.TCPConn) uint64
}

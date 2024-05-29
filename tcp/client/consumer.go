package client

import "net"

type TcpClientConsumer interface {
	OnConnected(*net.TCPConn)
}

package tcpClient

import "net"

type TcpClientSock interface {
	OnConnected(*net.TCPConn)
}

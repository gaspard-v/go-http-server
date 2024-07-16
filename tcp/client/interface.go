package tcpClient

import (
	"net"
	"sync"
)

type TcpClientSock interface {
	OnConnected(*net.TCPConn, *sync.WaitGroup)
}

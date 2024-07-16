package tcpServer

import (
	"net"
	"sync"
)

type TcpSockInterface interface {
	OnAccept(*net.TCPConn, *sync.WaitGroup)
}

type TcpConnInterface interface {
	OnReceive() uint64
	OnSend() uint64
}

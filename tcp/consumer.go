package tcp

import (
	"net"
	"sync"
)

type TcpConsumer interface {
	OnAccept(*net.TCPConn, *sync.WaitGroup)
}

type TcpConnConsumer interface {
	OnReceive() uint64
	OnSend() uint64
}

package tcp

import "net"

type TcpConsumer interface {
	OnAccept(*net.TCPConn)
}

type ErrorConsumer interface {
	OnFatal(error)
	OnWarning(error)
	OnNotice(error)
}

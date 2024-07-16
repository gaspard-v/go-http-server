package tcpServer

import (
	"fmt"
	"net"
	"sync"

	log "github.com/gaspard-v/go-http-server/log/object"
)

const DEFAULT_ADDRESS string = "127.0.0.1"
const DEFAULT_PORT string = "8080"

type Tcp struct {
	Listener *net.TCPListener
	tcpSock  TcpSockInterface
	log      log.LogInterface
	stopLoop chan bool
}

func Default(
	tcpSock TcpSockInterface,
) *Tcp {
	address := fmt.Sprintf("%s:%s", DEFAULT_ADDRESS, DEFAULT_PORT)
	return Create(
		address,
		tcpSock,
	)
}

func Create(
	address string,
	tcpSock TcpSockInterface,
) *Tcp {
	tcpAddr, error := net.ResolveTCPAddr("tcp", address)
	stopLoop := make(chan bool)
	log := log.Get("tcp.server")
	if error != nil {
		log.Fatal(error)
	}
	listener, error := net.ListenTCP(tcpAddr.Network(), tcpAddr)
	if error != nil {
		log.Fatal(error)
	}
	tcp := Tcp{listener, tcpSock, log, stopLoop}
	return &tcp
}

func (tcp *Tcp) Accept() {
	var wg sync.WaitGroup
Exit:
	for {
		select {
		case <-tcp.stopLoop:
			break Exit
		default:
		}
		conn, error := tcp.Listener.AcceptTCP()
		if error != nil {
			tcp.log.Debug(error)
			continue
		}

		tcp.log.Debug("Server got a client " + conn.RemoteAddr().String() + ", calling OnAccept callback")
		wg.Add(1)
		go tcp.tcpSock.OnAccept(conn, &wg)
	}
	wg.Wait()
}

func (tcp *Tcp) Stop() {
	tcp.stopLoop <- true
}

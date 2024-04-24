package tcp

import (
	"fmt"
	"log"
	"net"
)

const DEFAULT_ADDRESS string = "127.0.0.1"
const DEFAULT_PORT string = "8080"

type Tcp struct {
	Listener    net.TCPListener
	tcpConsumer TcpConsumer
}

func CreateDefault(tcpConsumer TcpConsumer) *Tcp {
	return Create(fmt.Sprintf("%s:%s", DEFAULT_ADDRESS, DEFAULT_PORT), tcpConsumer)
}

func Create(address string, tcpConsumer TcpConsumer) *Tcp {
	tcpAddr, error := net.ResolveTCPAddr("tcp", address)
	if error != nil {
		log.Fatalln(error)
	}
	listener, error := net.ListenTCP(tcpAddr.Network(), tcpAddr)
	if error != nil {
		log.Fatalln(error)
	}
	tcp := Tcp{*listener, tcpConsumer}
	return &tcp
}

func (tcp *Tcp) Accept() {
	for {
		conn, error := tcp.Listener.AcceptTCP()
		if error != nil {
			log.Println(error)
			continue
		}

		go tcp.tcpConsumer.OnAccept(conn)
	}
}

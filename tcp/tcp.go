package tcp

import (
	"fmt"
	"log"
	"net"
)

const DEFAULT_ADDRESS string = "127.0.0.1"
const DEFAULT_PORT string = "8080"

type Tcp struct {
	Listener net.TCPListener
}

func CreateDefault() *Tcp {
	return Create(fmt.Sprintf("%s:%s", DEFAULT_ADDRESS, DEFAULT_PORT))
}

func Create(address string) *Tcp {
	tcpAddr, error := net.ResolveTCPAddr("tcp", address)
	if error != nil {
		log.Fatalln(error)
	}
	listener, error := net.ListenTCP(tcpAddr.Network(), tcpAddr)
	if error != nil {
		log.Fatalln(error)
	}
	tcp := Tcp{*listener}
	return &tcp
}

func (tcp *Tcp) Accept() {
	for {
		conn, error := tcp.Listener.AcceptTCP()
		if error != nil {
			log.Println(error)
			continue
		}

		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// Close the connection when we're done
	defer conn.Close()

	// Read incoming data
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the incoming data
	fmt.Printf("Received: %s", buf)
}

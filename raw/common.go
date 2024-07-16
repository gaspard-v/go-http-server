package raw

import (
	"encoding/binary"
	"net"
)

func getBodySize(conn *net.TCPConn) (uint64, error) {
	var body_size uint64 = 0
	err := binary.Read(conn, binary.BigEndian, body_size)
	if err != nil {
		return 0, err
	}
	return body_size, nil
}

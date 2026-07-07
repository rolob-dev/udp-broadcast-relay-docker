package socket

import (
	"net"

	"golang.org/x/net/ipv4"
)

type Packet struct {
	Data      []byte
	Length    int
	Source    net.Addr
	Control   *ipv4.ControlMessage
	Interface *net.Interface
}
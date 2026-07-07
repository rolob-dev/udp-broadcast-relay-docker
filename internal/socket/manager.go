package socket

import (
	"fmt"
	"net"

	"golang.org/x/net/ipv4"
)

type Manager struct {
	conn   net.PacketConn
	packet *ipv4.PacketConn
}

func New() *Manager {
	return &Manager{}
}

func (m *Manager) Open() error {

	conn, err := net.ListenPacket("udp4", ":1900")
	if err != nil {
		return err
	}

	m.conn = conn
	m.packet = ipv4.NewPacketConn(conn)

	fmt.Println("Opening packet socket")
	fmt.Println()

	fmt.Println("✓ UDP :1900")
	fmt.Println()

	return nil
}

func (m *Manager) Close() error {

	if m.conn != nil {
		return m.conn.Close()
	}

	return nil
}
package socket

import (
	"fmt"
	"net"

	"golang.org/x/net/ipv4"
)

var ssdpGroup = &net.UDPAddr{
	IP:   net.IPv4(239, 255, 255, 250),
	Port: 1900,
}

type Manager struct {
	conn   net.PacketConn
	packet *ipv4.PacketConn
	group  *net.UDPAddr
}

func New() *Manager {
	return &Manager{
		group: ssdpGroup,
	}
}

func (m *Manager) Open() error {

	conn, err := net.ListenPacket("udp4", "0.0.0.0:1900")
	if err != nil {
		return err
	}

	m.conn = conn
	m.packet = ipv4.NewPacketConn(conn)

	// Wir möchten später sowohl das Eingangsinterface
	// als auch die Zieladresse kennen.
	if err := m.packet.SetControlMessage(
		ipv4.FlagInterface|ipv4.FlagDst,
		true,
	); err != nil {
		conn.Close()
		return err
	}

	fmt.Println("Opening packet socket")
	fmt.Println()

	fmt.Println("✓ UDP :1900")
	fmt.Println("✓ Interface control messages enabled")
	fmt.Println()

	return nil
}

func (m *Manager) Close() error {

	if m.packet != nil {
		return m.packet.Close()
	}

	if m.conn != nil {
		return m.conn.Close()
	}

	return nil
}
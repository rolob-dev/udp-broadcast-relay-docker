package socket

import (
	"fmt"
	"net"
)

func (m *Manager) Join(interfaces map[string]*net.Interface) error {

	fmt.Println("Joining SSDP multicast groups")
	fmt.Println()

	for _, iface := range interfaces {

		if err := m.packet.JoinGroup(iface, m.group); err != nil {
			return fmt.Errorf(
				"join group on %s: %w",
				iface.Name,
				err,
			)
		}

		fmt.Printf("✓ %s\n", iface.Name)
	}

	fmt.Println()

	return nil
}
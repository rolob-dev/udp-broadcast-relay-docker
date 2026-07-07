package socket

import (
	"fmt"
	"net"
)

func (m *Manager) Relay(pkt *Packet, interfaces map[string]*net.Interface) error {

	fmt.Println("Relay decision")

	for _, iface := range interfaces {

		// Niemals auf das Interface zurücksenden,
		// von dem das Paket gekommen ist.
		if iface.Index == pkt.Interface.Index {
			continue
		}

		fmt.Printf(
			"✓ %s -> %s\n",
			pkt.Interface.Name,
			iface.Name,
		)
	}

	fmt.Println()

	return nil
}
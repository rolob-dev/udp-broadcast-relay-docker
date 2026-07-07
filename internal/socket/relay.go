package socket

import (
	"fmt"
	"net"

	"golang.org/x/net/ipv4"
)

func (m *Manager) Relay(pkt *Packet, interfaces map[string]*net.Interface) error {

	fmt.Println("Relay decision")

	for _, iface := range interfaces {

		// Niemals auf das Eingangsinterface zurücksenden
		if iface.Index == pkt.Interface.Index {
			continue
		}

		cm := &ipv4.ControlMessage{
			IfIndex: iface.Index,
		}

		fmt.Printf(
			"Would send %d bytes via %s (ifindex=%d) to %s\n",
			pkt.Length,
			iface.Name,
			iface.Index,
			m.group.String(),
		)

		// Im nächsten Schritt wird hier WriteTo() aktiviert.
		_ = cm
	}

	fmt.Println()

	return nil
}
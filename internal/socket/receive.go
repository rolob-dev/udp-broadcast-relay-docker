package socket

import (
	"fmt"
	"net"
)

func (m *Manager) Receive(interfaces map[string]*net.Interface) (*Packet, error) {

	buf := make([]byte, 1500)

	for {

		n, cm, src, err := m.packet.ReadFrom(buf)
		if err != nil {
			return nil, err
		}

		udp, ok := src.(*net.UDPAddr)
		if ok && udp.Port == 1900 {
    		fmt.Println("Ignoring self-generated packet")
    		continue
		}

		// Ohne ControlMessage können wir das Interface
		// nicht bestimmen.
		if cm == nil {
			continue
		}

		// Nur Multicast interessiert uns.
		if !cm.Dst.IsMulticast() {
			continue
		}

		// Nur SSDP (239.255.255.250)
		if !cm.Dst.Equal(m.group.IP) {
			continue
		}

		var iface *net.Interface

		for _, i := range interfaces {
			if i.Index == cm.IfIndex {
				iface = i
				break
			}
		}

		if iface == nil {
			fmt.Printf("Ignoring packet on unknown interface index %d\n", cm.IfIndex)
			continue
		}

		data := make([]byte, n)
		copy(data, buf[:n])

		return &Packet{
			Data:      data,
			Length:    n,
			Source:    src,
			Control:   cm,
			Interface: iface,
		}, nil
	}
}
package socket

import (
	"fmt"
	"net"
)

func (m *Manager) Listen(interfaces map[string]*net.Interface) error {

	fmt.Println("Listening...")
	fmt.Println()

	for {

		pkt, err := m.Receive(interfaces)
		if err != nil {
			return err
		}

		fmt.Println("----------------------------------------")
		fmt.Printf("Interface : %s\n", pkt.Interface.Name)
		fmt.Printf("Source    : %s\n", pkt.Source)
		fmt.Printf("Target    : %s\n", pkt.Control.Dst)
		fmt.Printf("Length    : %d bytes\n", pkt.Length)
		fmt.Println()
		fmt.Println("----------------------------------------")
		fmt.Println()		
		fmt.Println(string(pkt.Data))
		fmt.Println()		
		fmt.Println("----------------------------------------")
		fmt.Println()

		if err := m.Relay(pkt, interfaces); err != nil {
			return err
		}
	}
}
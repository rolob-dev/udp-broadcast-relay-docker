package network

import (
	"fmt"
	"net"

	"github.com/rolob-dev/udp-broadcast-relay-docker/internal/config"
)

func Discover(cfg *config.Config) error {

	interfaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	fmt.Println("Detected interfaces")
	fmt.Println()

	for _, iface := range interfaces {

		fmt.Printf("%s\n", iface.Name)
		fmt.Printf("    Index : %d\n", iface.Index)
		fmt.Printf("    MTU   : %d\n", iface.MTU)
		fmt.Printf("    Flags : %s\n", iface.Flags.String())

		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println()
			continue
		}

		if len(addrs) == 0 {
			fmt.Println("    Address: (none)")
		}

		for _, addr := range addrs {

			ipnet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			fmt.Printf("    Address: %s\n", ipnet.String())

			// Nur IPv4 für unsere Zuordnung verwenden
			if ipnet.IP.To4() == nil {
				continue
			}

			for i := range cfg.Networks {

				if cfg.Networks[i].CIDR.Contains(ipnet.IP) {
					cfg.Networks[i].Interface = &iface
					cfg.Networks[i].Addresses = append(cfg.Networks[i].Addresses, ipnet)
				}
			}
		}

		fmt.Println()
	}

	return nil
}
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

		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		fmt.Printf("%s\n", iface.Name)

		for _, addr := range addrs {

			ipnet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}

			fmt.Printf("    %s\n", ipnet.String())
		}

		fmt.Println()
	}

	return nil
}
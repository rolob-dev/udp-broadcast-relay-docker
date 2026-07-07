package network

import (
	"fmt"
	"net"

	"github.com/rolob-dev/udp-broadcast-relay-docker/internal/config"
)

func Validate(cfg *config.Config) error {

	interfaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	available := make(map[string]net.Interface)

	for _, iface := range interfaces {
		available[iface.Name] = iface
	}

	fmt.Println("Validating interfaces")
	fmt.Println()

	for _, name := range cfg.Interfaces {

		if _, ok := available[name]; !ok {
			return fmt.Errorf("interface %q not found", name)
		}

		fmt.Printf("✓ %s\n", name)
	}

	fmt.Println()

	return nil
}
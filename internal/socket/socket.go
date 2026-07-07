package socket

import (
	"fmt"
	"net"

	"github.com/rolob-dev/udp-broadcast-relay-docker/internal/config"
)

func Open(cfg *config.Config) error {

	fmt.Println("Preparing sockets")
	fmt.Println()

	interfaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	for _, name := range cfg.Interfaces {

		var found *net.Interface

		for _, iface := range interfaces {
			if iface.Name == name {
				found = &iface
				break
			}
		}

		if found == nil {
			return fmt.Errorf("interface %q not found", name)
		}

		fmt.Printf("✓ %s (index=%d mtu=%d flags=%s)\n",
			found.Name,
			found.Index,
			found.MTU,
			found.Flags.String(),
		)
	}

	fmt.Println()

	return nil
}
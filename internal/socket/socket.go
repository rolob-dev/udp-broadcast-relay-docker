package socket

import (
	"fmt"

	"github.com/rolob-dev/udp-broadcast-relay-docker/internal/runtime"
)

func Open(rt *runtime.Runtime) error {

	fmt.Println("Preparing sockets")
	fmt.Println()

	for _, iface := range rt.Interfaces {

		fmt.Printf("✓ %s (index=%d mtu=%d flags=%s)\n",
			iface.Name,
			iface.Index,
			iface.MTU,
			iface.Flags.String(),
		)
	}

	fmt.Println()

	return nil
}
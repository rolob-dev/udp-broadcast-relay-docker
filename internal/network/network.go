package network

import (
	"fmt"
	"net"

	"github.com/rolob-dev/udp-broadcast-relay-docker/internal/runtime"
)

func Validate(rt *runtime.Runtime) error {

	interfaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	rt.Interfaces = make(map[string]*net.Interface)

	available := make(map[string]net.Interface)

	for _, iface := range interfaces {
		available[iface.Name] = iface
	}

	fmt.Println("Validating interfaces")
	fmt.Println()

	for _, name := range rt.Config.Interfaces {

		iface, ok := available[name]
		if !ok {
			return fmt.Errorf("interface %q not found", name)
		}

		rt.Interfaces[name] = &iface

		fmt.Printf("✓ %s\n", name)
	}

	fmt.Println()

	return nil
}
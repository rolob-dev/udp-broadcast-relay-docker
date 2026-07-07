package network

import (
	"net"

	"github.com/rolob-dev/udp-broadcast-relay-docker/internal/config"
)

func Discover(cfg *config.Config) error {

	interfaces, err := net.Interfaces()
	if err != nil {
		return err
	}

	for i := range cfg.Networks {

		for _, iface := range interfaces {

			addrs, err := iface.Addrs()
			if err != nil {
				continue
			}

			for _, addr := range addrs {

				ipnet, ok := addr.(*net.IPNet)
				if !ok {
					continue
				}

				// Nur IPv4
				if ipnet.IP.To4() == nil {
					continue
				}

				// Liegt die IP in unserem konfigurierten Netz?
				if cfg.Networks[i].CIDR.Contains(ipnet.IP) {

					cfg.Networks[i].Interface = &iface
					cfg.Networks[i].Addresses = append(cfg.Networks[i].Addresses, ipnet)

				}
			}
		}
	}

	return nil
}
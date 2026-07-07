package config

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type Network struct {
	CIDR *net.IPNet

	Interface *net.Interface
	Addresses []*net.IPNet
}

type Config struct {
	Networks []Network
}

func Load() (*Config, error) {

	value := os.Getenv("NETWORKS")
	if value == "" {
		return nil, fmt.Errorf("environment variable NETWORKS is not set")
	}

	cfg := &Config{}

	for _, network := range strings.Split(value, ",") {

		network = strings.TrimSpace(network)

		_, cidr, err := net.ParseCIDR(network)
		if err != nil {
			return nil, fmt.Errorf("invalid network '%s': %w", network, err)
		}

		cfg.Networks = append(cfg.Networks, Network{
			CIDR: cidr,
		})
	}

	return cfg, nil
}
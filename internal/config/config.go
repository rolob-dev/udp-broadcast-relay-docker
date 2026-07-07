package config

import (
	"fmt"
	"net"
	"os"
	"strings"
)

type Config struct {
	Interfaces []string
}

func Load() (*Config, error) {

	value := os.Getenv("INTERFACES")
	if value == "" {
		return nil, fmt.Errorf("environment variable INTERFACES is not set")
	}

	cfg := &Config{}

	for _, iface := range strings.Split(value, ",") {

		iface = strings.TrimSpace(iface)

		if iface == "" {
			continue
		}

		cfg.Interfaces = append(cfg.Interfaces, iface)
	}

	if len(cfg.Interfaces) == 0 {
		return nil, fmt.Errorf("no interfaces configured")
	}

	return cfg, nil
}
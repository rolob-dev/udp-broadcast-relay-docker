package main

import (
	"fmt"
	"log"

	"github.com/rolob-dev/udp-broadcast-relay-docker/internal/config"
)

const Name = "SSDP Relay"

var Version = "dev"

func main() {

	fmt.Println(Name)
	fmt.Println("Version:", Version)
	fmt.Println()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Configured networks:")

	for _, network := range cfg.Networks {
		fmt.Printf("  ✓ %s\n", network.CIDR.String())
	}
}
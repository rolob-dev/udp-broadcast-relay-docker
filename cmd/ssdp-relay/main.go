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

	fmt.Println("Configured interfaces")
	fmt.Println()

	for _, iface := range cfg.Interfaces {
		fmt.Printf("✓ %s\n", iface)
	}
}
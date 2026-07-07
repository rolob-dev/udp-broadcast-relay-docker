package main

import (
	"fmt"
	"log"

	"github.com/rolob-dev/udp-broadcast-relay-docker/internal/config"
	"github.com/rolob-dev/udp-broadcast-relay-docker/internal/network"
	"github.com/rolob-dev/udp-broadcast-relay-docker/internal/runtime"
	"github.com/rolob-dev/udp-broadcast-relay-docker/internal/socket"
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

	rt := &runtime.Runtime{
		Config: cfg,
	}

	rt.Socket = socket.New()

	fmt.Println("Configured interfaces")
	fmt.Println()

	for _, iface := range cfg.Interfaces {
		fmt.Printf("✓ %s\n", iface)
	}

	fmt.Println()

	if err := network.Validate(rt); err != nil {
		log.Fatal(err)
	}

	if err := rt.Socket.Open(); err != nil {
		log.Fatal(err)
	}
	defer rt.Socket.Close()

	if err := rt.Socket.Join(rt.Interfaces); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ready.")
}
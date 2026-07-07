package runtime

import (
	"net"

	"github.com/rolob-dev/udp-broadcast-relay-docker/internal/config"
	"github.com/rolob-dev/udp-broadcast-relay-docker/internal/socket"
)

type Runtime struct {
	Config *config.Config

	Interfaces map[string]*net.Interface

	Socket *socket.Manager
}
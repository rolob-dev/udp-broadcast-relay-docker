package socket

import (
	"net"
)

func (m *Manager) AddLocalIP(ip net.IP) {
	if ip == nil {
		return
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return
	}

	m.localIPs[ip4.String()] = struct{}{}
}

func (m *Manager) IsLocalIP(ip net.IP) bool {
	if ip == nil {
		return false
	}

	_, ok := m.localIPs[ip.String()]
	return ok
}
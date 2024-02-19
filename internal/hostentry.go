package internal

import "net"

type HostEntry struct {
	name    string
	ip      net.IP
	aliases []string
}

func NewHostEntry(name string, ip net.IP) *HostEntry {
	return &HostEntry{
		name: name,
		ip:   ip,
	}
}

package main

import (
	"fmt"
)

type tcpScanConfiguration struct {
	ip                 string
	startPort, endPort int
}

type protocol int

const (
	tcp = iota
	udp
	ICMP
)

func (p *protocol) String() string {
	switch *p {
	case tcp:
		return "tcp"
	case udp:
		return "udp"
	case ICMP:
		return "ICMP"
	default:
		return "unknown protocol"
	}
}

func (p *protocol) Set(s string) error {
	switch s {
	case "tcp":
		*p = tcp
	case "udp":
		*p = udp
	case "ICMP":
		*p = ICMP
	default:
		return fmt.Errorf("invalid protocol: %s", s)
	}
	return nil
}

type icmpScanConfiguration struct {
	startIP, endIP string
}

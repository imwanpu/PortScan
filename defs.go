package main

import (
	"fmt"
)

type scanConfiguration struct {
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

// _____________________________________________________________________________________________________________________
type Gender int

const (
	Male Gender = iota
	Female
	Other
)

func (g *Gender) String() string {
	switch *g {
	case Male:
		return "male"
	case Female:
		return "female"
	case Other:
		return "other"
	default:
		return "unknown"
	}
}

func (g *Gender) Set(s string) error {
	switch s {
	case "male":
		*g = Male
	case "female":
		*g = Female
	case "other":
		*g = Other
	default:
		return fmt.Errorf("invalid gender: %q", s)
	}
	return nil
}

package main

import (
	"flag"
)

var (
	pr               protocol = tcp
	concurrencyLevel int
	tcpCfg           = tcpScanConfiguration{
		ip:        "localhost",
		startPort: 0,
		endPort:   0,
	}
	openPorts []int
	icmpCfg   = icmpScanConfiguration{
		startIP: "0.0.0.0",
		endIP:   "0.0.0.0",
	}
	activeHosts []string
)

func parseArgv() {

	flag.IntVar(&concurrencyLevel, "cl", 1, "concurrent level, default 1.\n the maximum concurrency level is 65535 aka 2<<16 - 1.\nthe minimum concurrency level is 1")

	// tcp scan
	flag.Var(&pr, "p", "tcp, udp or ICMP protocol, default tcp")
	flag.StringVar(&tcpCfg.ip, "ip", "127.0.0.1", "ip address, default 127.0.0.1")
	flag.IntVar(&tcpCfg.startPort, "sp", 0, "start port, default 0")
	flag.IntVar(&tcpCfg.endPort, "ep", 0, "end port, default 0")

	// ICMP scan
	flag.StringVar(&icmpCfg.startIP, "si", "0.0.0.0", "start ip, default 0.0.0.0")
	flag.StringVar(&icmpCfg.endIP, "ei", "0.0.0.0", "end ip, default 0.0.0.0")

	flag.Parse()
}

func main() {
	parseArgv()
	if pr == tcp {
		tcpScan(concurrencyLevel)
	} else if pr == ICMP {
		icmpScan(concurrencyLevel)
	}

}

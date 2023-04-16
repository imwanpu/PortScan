package main

var (
	pr               protocol = tcp
	concurrencyLevel int
	scanCfg          = scanConfiguration{
		ip:        "localhost",
		startPort: 0,
		endPort:   0,
	}
	openPorts   []int
	closedPorts []int
)

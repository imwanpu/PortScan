package main

import (
	"flag"
)

func parseArgv() {
	flag.Var(&pr, "p", "tcp, udp or ICMP protocol, default tcp")
	flag.StringVar(&scanCfg.ip, "ip", "localhost", "ip address, default localhost")
	flag.IntVar(&scanCfg.startPort, "sp", 0, "start port, default 0")
	flag.IntVar(&scanCfg.endPort, "ep", 0, "end port, default 0")
	flag.IntVar(&concurrencyLevel, "cl", 1, "concurrent level, default 1.\n the maximum concurrency level is 65535 aka 2<<16 - 1.\nthe minimum concurrency level is 1")
	flag.Parse()
}

// 端口最大值校验, 过大或过小则赋予最大最小值, 并发出 warming 信息
// 端口范围为 0 到 65535 即 2<<16 - 1

func main() {
	parseArgv()
	if pr == tcp {
		tcpScan(concurrencyLevel)
	}

}

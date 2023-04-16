package main

import (
	"fmt"
	"net"
)

func worker(ports chan int, results chan int) {
	for port := range ports {
		address := fmt.Sprintf("%s:%d", scanCfg.ip, port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		_ = conn.Close()
		results <- port
	}
}

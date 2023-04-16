package main

import (
	"errors"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"math"
	"net"
	"sort"
)

func tcpScan(concurrentLevel int) {

	err := tcpArgvCheck()
	if err != nil {
		fmt.Println(err)
	}

	bar := progressbar.NewOptions(
		tcpCfg.endPort-tcpCfg.startPort+1,
		progressbar.OptionSetRenderBlankState(true))

	ports := make(chan int, concurrentLevel)
	results := make(chan int)

	// 根据并发数确定 tcpWorker 数量
	for i := 0; i < concurrentLevel; i++ {
		go tcpWorker(ports, results)
	}

	go func() {
		for i := tcpCfg.startPort; i <= tcpCfg.endPort; i++ {
			ports <- i
		}
	}()

	for i := tcpCfg.startPort; i <= tcpCfg.endPort; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
		_ = bar.Add(1)
	}

	close(ports)
	close(results)

	sort.Ints(openPorts)

	fmt.Println("\n", openPorts)
}

func tcpWorker(ports chan int, results chan int) {
	for port := range ports {
		address := fmt.Sprintf("%s:%d", tcpCfg.ip, port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		_ = conn.Close()
		results <- port
	}
}

func tcpArgvCheck() error {
	// 端口最大值校验, 过大或过小则赋予最大最小值, 并发出 warming 信息
	// 端口范围为 0 到 65535 即 2<<16 - 1
	var errorStr = "Warning: "
	if tcpCfg.startPort > tcpCfg.endPort {
		t := tcpCfg.startPort
		tcpCfg.startPort = tcpCfg.endPort
		tcpCfg.endPort = t
		errorStr += "start port is larger than end port; "
	}
	if tcpCfg.startPort < 0 {
		tcpCfg.startPort = 0
		errorStr += "start port is less than 0; "
	}
	if tcpCfg.endPort > math.MaxUint16 {
		tcpCfg.endPort = math.MaxUint16
		errorStr += fmt.Sprintf("end port is larger than %v", math.MaxUint16)
	}
	if errorStr != "Warning: " {
		return errors.New(errorStr)
	}
	return nil
}

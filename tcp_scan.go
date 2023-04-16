package main

import (
	"fmt"
	"github.com/schollz/progressbar/v3"
	"sort"
)

func tcpScan(concurrentLevel int) {

	bar := progressbar.NewOptions(
		scanCfg.endPort-scanCfg.startPort+1,
		progressbar.OptionSetRenderBlankState(true))

	ports := make(chan int, concurrentLevel)
	results := make(chan int)

	// 根据端口数构建 worker 数量
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := scanCfg.startPort; i <= scanCfg.endPort; i++ {
			ports <- i
		}
	}()

	for i := scanCfg.startPort; i <= scanCfg.endPort; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		} else {
			closedPorts = append(closedPorts, port)
		}
		_ = bar.Add(1)
	}

	close(ports)
	close(results)

	sort.Ints(openPorts)
	sort.Ints(closedPorts)
	fmt.Println("\n")
	fmt.Println(openPorts)

}

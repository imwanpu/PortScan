package main

import (
	"errors"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"net"
	"os/exec"
	"sort"
)

func icmpScan(concurrentLeve int) {

	err := icmpArgvCheck()
	if err != nil {
		println(err)
	}

	bar := progressbar.NewOptions(
		ipSubtract(icmpCfg.startIP, icmpCfg.endIP)+1,
		progressbar.OptionSetRenderBlankState(true))

	ips := make(chan string, concurrentLeve)
	results := make(chan string)

	// 根据并发数确定 icmpWorker 数量
	for i := 0; i < concurrentLeve; i++ {
		go icmpWorker(ips, results)
	}

	go func() {
		for i := icmpCfg.startIP; isFormerIpNotLarger(i, icmpCfg.endIP); increase(i, 1) {
			ips <- i
		}
	}()

	for i := icmpCfg.startIP; isFormerIpNotLarger(i, icmpCfg.endIP); increase(i, 1) {
		ip := <-results
		if ip != "" {
			activeHost = append(activeHost, ip)
		}
		_ = bar.Add(1)
	}

	close(ips)
	close(results)

	sort.Strings(activeHost)
	fmt.Println("\n", activeHost)
}

func icmpWorker(ips, results chan string) {
	for ip := range ips {
		cmd := exec.Command("cmd", "/c", fmt.Sprintf("ping %s", ip))
		if err := cmd.Start(); err != nil {
			results <- ""
			continue
		}
		// cmd.Start() 为耗时操作, 时长可能会超过 10 秒
		if err := cmd.Wait(); err != nil {
			results <- ""
		}
		results <- ip
	}
}

func ping(ip string) error {
	cmd := exec.Command("cmd", "/c", fmt.Sprintf("ping %s", ip))
	if err := cmd.Start(); err != nil {
		return err
	}
	// cmd.Start() 为耗时操作, 时长可能会超过 10 秒
	if err := cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func increase(smallerIp string, step uint32) string {
	uip := ip2uint32(smallerIp)
	uip += step // 执行递增操作
	return uint32toIp(uip)
}

func isFormerIpNotLarger(former, latter string) bool {
	return ip2uint32(former) <= ip2uint32(latter)
}

func ipSubtract(ip1, ip2 string) int {
	uip1 := ip2uint32(ip1)
	uip2 := ip2uint32(ip2)
	if uip1 > uip2 {
		return int(uip1 - uip2)
	} else {
		return int(uip2 - uip1)
	}
}

func ip2uint32(ip string) uint32 {
	parsed := net.ParseIP(ip)
	return (uint32(parsed[12]) << 24) + (uint32(parsed[13]) << 16) + (uint32(parsed[14]) << 8) + uint32(parsed[15])
}

func uint32toIp(u uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", (u>>24)&0xFF, (u>>16)&0xFF, (u>>8)&0xFF, u&0xFF)
}

func icmpArgvCheck() error {
	// 用正则判断输入的是否是合法你的ipv4地址, 并判断大小
	if !isIpv4(icmpCfg.startIP) || !isIpv4(icmpCfg.endIP) {
		panic("format of start ip or end ip is not valid.")
	}
	if isFormerIpNotLarger(icmpCfg.endIP, icmpCfg.startIP) {
		t := icmpCfg.startIP
		icmpCfg.startIP = icmpCfg.endIP
		icmpCfg.endIP = t
		return errors.New("start ip is larger than end ip")
	}
	return nil

}

func isIpv4(ip string) bool {
	parsed := net.ParseIP(ip)
	return parsed != nil && parsed.To4() != nil // 如果解析结果不为空且是 IPv4，则返回 true；否则返回 false。
}

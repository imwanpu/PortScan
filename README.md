## 项目介绍

项目名称: PortScan

用途: 进行 TCP 及 ICMP 端口扫描

背景: 本人在 信息安全综合实验课程 中的一项作业罢了.

## 使用说明

### 编译

环境

- 平台 Windows 11 22H2
- Go 1.20.3 编译环境

编译命令

```shell
# 项目中用到的第三方进度条库
go get github.com/schollz/progressbar/v3 

go build -o PortScan.exe
```

执行上述命令后, 会得到名为 `PortScan` 的可执行文件

### 使用

命令行参数说明

```text
  -cl int
        concurrent level, default 1.
        the maximum concurrency level is 65535 aka 2<<16 - 1.
        the minimum concurrency level is 1 (default 1)
  -ei string
        end ip, default 0.0.0.0 (default "0.0.0.0")
  -ep int
        end port, default 0
  -ip string
        ip address, default 127.0.0.1 (default "127.0.0.1")
  -p value
        tcp, udp or ICMP protocol, default tcp
  -si string
        start ip, default 0.0.0.0 (default "0.0.0.0")
  -sp int
        start port, default 0
```

### 使用示例

```shell
# 获取帮助信息
.\PortScan.exe -help 

# 对本地主机的 1024-2000 端口进行 200 并发量的 TCP 扫描
.\PortScan.exe -p tcp -ip localhost -sp 1024 -ep 2000 -cl 200

# 对 baidu.com 的 79-444 端口进行 100 并发量的 TCP 扫描 ⚠️ 网络不是法外之地
.\PortScan.exe -p tcp -ip baidu.com -sp 79 -ep 444 -cl 100

# 对 192.168.2.1 - 192.168.3.1 进行 并发量为 10 的 ICMP 扫描
.\PortScan.exe -p ICMP -si 192.168.2.1 -ei 192.168.3.1 -cl 10

```

## 交流

- 直接提 issues 就行 

## 必读注意事项
- ⚠️ 网络不是法外之地啊, 不要在未经允许的情况下扫公网. 为了你的人身安全, 请谨慎使用 ⚠️
- 在最大并发, 即 `-cl 65535`, 的情况下, `PortScan.exe` 程序的内存占用会超过 1GiB. 为了你的计算机安全, 请谨慎控制并发量.
- 本项目尽在 windows 平台运行过, 不确定 mac, linux 平台可用

## 鸣谢

- Golang
- [@Zack](https://github.com/schollz)'s [progressbar](https://github.com/schollz/progressbar)

## 版权

MIT
const net = require('net');
const dgram = require('dgram')

TcpStartPort = 1024

// 打开 10 个 tcp 端口
numberOfTcpPort = 10
for (let i = TcpStartPort; i < TcpStartPort + numberOfTcpPort; i++) {
    net.createServer().listen(i, () => {
        console.log('Server is listening on TCP port ' + i + '.');
    })
}



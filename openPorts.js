const net = require('net');


for (let i = 1024; i < 1034; i++) {
    net.createServer().listen(i, () => {
        console.log('Server is listening on port ' + i + '.');
    })
}
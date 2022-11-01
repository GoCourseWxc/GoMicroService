package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

// 命令请求数据
// echo -e '{"method":"HelloService.Hello","params":["world"],"id":1}'| nc 192.168.50.24 1234
func main() {
	//注册rpc服务
	rpc.RegisterName("HelloService", new(HelloService))
	//设置监听
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	for {
		//接收连接
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		//给当前连接提供针对json格式的rpc服务
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

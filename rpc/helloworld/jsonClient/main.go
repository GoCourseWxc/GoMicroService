package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// nc -l 1234 建立服务
func main() {
	//简历tcp连接
	conn, err := net.Dial("tcp", "192.168.50.45:1234")
	if err != nil {
		panic(err)
	}
	//简历基于json编解码的rpc服务
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	//调用rpc服务方法
	err = client.Call("HelloService.Hello", " world", &reply)
	if err != nil {
		panic(err)
	}

	fmt.Println("收到的数据为:", reply)
}

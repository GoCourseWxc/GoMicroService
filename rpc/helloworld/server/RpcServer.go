package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	//rpc注册服务
	//注册rpc服务，维护一个hash表，key值是服务名称，value值是服务的地址
	err := rpc.RegisterName("HelloService", new(HelloService))
	if err != nil {
		fmt.Println("注册rpc服务失败!", err)
		return
	}

	//设置服务监听
	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("net.Listen:err", err)
		return
	}

	//接受传输的数据
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept:err", err)
		return
	}
	defer conn.Close()

	//rpc调用,并返回执行后的数据
	//1.read，获取服务名称和方法名，获取请求数据
	//2.调用对应服务里面的方法，获取传出数据
	//3.write，把数据返回给client
	rpc.ServeConn(conn)

}

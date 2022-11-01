package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//用rpc连接
	client, err := rpc.Dial("tcp", "192.168.50.45:1234")
	if err != nil {
		panic(err)
	}

	var reply string
	//调用服务中的函数
	err = client.Call("HelloService.Hello", "world", &reply)
	if err != nil {
		panic(err)
	}

	fmt.Println("收到的数据为,", reply)
}

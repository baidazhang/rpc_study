package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	//用rpc连接服务器
	conn, err := net.Dial("tcp", "192.168.60.122:8080")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	defer conn.Close()
	var resp string
	client.Call("hello.HelloWorld", "xiao", &resp)
	fmt.Println(resp)
}

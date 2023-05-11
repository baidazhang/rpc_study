package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	//用rpc连接服务器
	conn, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()
	var resp string
	conn.Call("hello.HelloWorld", "xiao", &resp)
	fmt.Println(resp)
}

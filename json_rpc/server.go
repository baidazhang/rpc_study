package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Hello struct {
}

func (h *Hello) HelloWorld(name string, resp *string) (err error) {
	*resp = name
	return err
}

func main() {
	//注册rpc服务， 指定对象和方法
	err := rpc.RegisterName("hello", new(Hello))
	if err != nil {
		fmt.Println("注册 err:", err)
		return
	}
	//设置监听
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("listen err:", err)
		return
	}
	defer listener.Close()
	//建立链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			return
		}
		//绑定服务
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))

	}

}

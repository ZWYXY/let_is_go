package main

import (
	"fmt"
	"net"
)

/**
1. 自定义客户端 基础设置
*/
func main() {

	client := NewClient("127.0.0.1", 8888)
	if client == nil {
		fmt.Println(">>>>>> 连接服务器失败")
		return
	}

	fmt.Println(">>>>>> 连接服务器成功")

	select {}
}

// 定义一个客户端类型（结构体）
type MyClient struct {
	// IP
	ServerIp string
	// 端口
	ServerPort int
	// 名称
	ServerName string
	// 连接
	ServerConn net.Conn
}

func NewClient(ip string, port int) *MyClient {

	client := &MyClient{
		ServerIp:   ip,
		ServerPort: port,
	}

	dial, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		fmt.Println(err)
	}

	client.ServerConn = dial

	return client
}

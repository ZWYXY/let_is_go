package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

var ServerIp string
var ServerPort int

/**
1. 解析命令行参数
*/
func main() {

	// 命令行解析
	flag.Parse()

	client := NewClient(ServerIp, ServerPort)
	if client == nil {
		fmt.Println(">>>>>> 连接服务器失败")
		return
	}

	fmt.Println(">>>>>> 连接服务器成功")

	time.Sleep(10 * 60 * time.Second)
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

// 用来解析命令行参数
func init() {
	flag.StringVar(&ServerIp, "ip", "127.0.0.1", "设置服务器Ip默认是（127.0.0.1）")
	flag.IntVar(&ServerPort, "port", 88888, "设置服务器port默认是（88888）")
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

package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string // ip
	Port int    // 端口
}

// NewServer 创建一个服务类
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

// BusinessHandler 业务处理类
func (this *Server) BusinessHandler(conn net.Conn) {
	fmt.Println("连接已建立")
}

// StartServer Server类的启动方法
func (this *Server) StartServer() {
	// 调用net包下的listen方法开启一个监听器
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("net.Listen err:", err)
			return
		}
		go this.BusinessHandler(conn)
	}
}

package main

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	Ip        string           // ip
	Port      int              // 端口
	OnlineMap map[string]*User // 在线用户表
	mapLock   sync.RWMutex     // 读写锁
	Message   chan string      // 管道
}

// NewServer 创建一个服务对象
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// BusinessHandler 业务处理类
func (this *Server) BusinessHandler(conn net.Conn) {
	//fmt.Println("连接已建立")

	// 当前用户加入到Map中
	this.mapLock.Lock()
	// 创建User
	user := NewUser(conn)
	this.OnlineMap[user.Addr] = user
	this.mapLock.Unlock()

	// 广播消息
	this.doBroadcast(user, "用户上线")
}

func (this *Server) doBroadcast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg

	for key, value := range this.OnlineMap {
		fmt.Println(key)
		value.C <- sendMsg
	}
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

package main

import (
	"fmt" //
	"io"
	"net"  // 网络通信功能相关包
	"sync" // 同步功能相关api包
)

type Server struct {
	Ip        string           // ip
	Port      int              // 端口
	OnlineMap map[string]*User // 在线用户表
	mapLock   sync.RWMutex     // 读写锁
	Message   chan string      // 管道
}

// NewServer 创建Server对象
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// StartServer Server的启动方法
func (this *Server) StartServer() {
	// 调用net包下的listen方法开启一个监听器
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	defer listen.Close()

	// 启动监听Message的goroutine
	go this.ListenMessage()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("net.Listen err:", err)
			continue
		}
		go this.BusinessHandler(conn)
	}
}

// BusinessHandler 业务处理类
func (this *Server) BusinessHandler(conn net.Conn) {
	// 创建User
	user := NewUser(conn)

	// 当前用户加入到Map中
	this.mapLock.Lock()
	this.OnlineMap[user.Addr] = user
	this.mapLock.Unlock()

	// 广播消息
	this.Broadcast(user, "用户上线")

	// 广播用户发送的信息
	go func() {
		//
		bytes := make([]byte, 10)
		n, err := conn.Read(bytes)

		for {
			if n == 0 {
				this.Broadcast(user, "下线")
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println(err)
				return
			}

			// 提取用户消息 去除 '\n'
			msg := string(bytes[:n-1])

			// 广播用户消息
			this.Broadcast(user, msg)
		}
	}()

	// 阻塞处理
	select {}
}

// 回写消息到管道
func (this *Server) Broadcast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg // 回写消息到Server管道
}

// 从管道中取出消息并回写给客户端
func (this *Server) ListenMessage() {
	for {
		// 监听管道中的信息，回写给客户端
		msg := <-this.Message

		this.mapLock.Lock()
		for _, user := range this.OnlineMap {
			// 这里向User的管道里面写入数据，会触发User对象的ListenMsg方法
			user.C <- msg
		}
		this.mapLock.Unlock()
	}
}
